package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	urlDates         = "https://malshare.com/daily"
	urlTxt           = "https://malshare.com/daily/%s/malshare_fileList.%s.all.txt"
	fileNames        = []string{"md5.txt", "sha1.txt", "sha256.txt"}
	folderPath       = "data2"
	filePath         = "/%d/%s/%d"
	findDateRegex    = `<a href="\d{2,4}\-\d{1,2}\-\d{1,2}\/">`
	extractDateRegex = `\d{2,4}\-\d{1,2}\-\d{1,2}`
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func request(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	checkError(err)

	res, err := http.DefaultClient.Do(req) // send req
	checkError(err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body) // parse res
	checkError(err)

	return string(body)
}

func saveToMap(dataMap map[time.Time]string, dateString string) {
	data := request(fmt.Sprintf(urlTxt, dateString, dateString)) // populate map
	date, err := time.Parse(time.DateOnly, dateString)
	checkError(err)
	dataMap[date] = data
}

func saveToFile(path string, dataString string) {

	err := os.MkdirAll(path, os.ModePerm)
	checkError(err)

	data := strings.Fields(dataString)
	count := 0
	for _, v := range data {
		if count > len(fileNames)-1 { //skip every 4th element
			count = 0
			continue
		}
		file, err := os.OpenFile(path+`/`+fileNames[count], os.O_APPEND|os.O_CREATE, 0700)
		checkError(err)
		file.WriteString(v + "\n")
		file.Close()
		count++
	}
}

func main() {
	fmt.Println("Hello em, tat man hinh di em")
	folder, _ := os.Stat(folderPath)
	if folder != nil {
		fmt.Println("folder exist!")
		return
	}

	responseBody := request(urlDates)
	// find HTML tag that contain require dates https://regex101.com/r/VDfst9/2
	regex, err := regexp.Compile(findDateRegex)
	checkError(err)
	match := regex.FindAllString(responseBody, -1)
	dataMap := make(map[time.Time]string)

	for _, v := range match {
		regex, err := regexp.Compile(extractDateRegex) //extract date
		checkError(err)
		dateString := regex.FindString(v)
		saveToMap(dataMap, dateString)
	}

	for date, v := range dataMap {
		currentPath := fmt.Sprintf(folderPath+filePath, date.Year(), date.Month(), date.Day())
		saveToFile(currentPath, v)
	}
}
