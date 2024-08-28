package main

import "fmt"

func arrayTutorial() {
	var intArr [6]int32 = [6]int32{1, 2, 3, 4, 5, 6}
	intSlice := intArr[1:4] //[2 3 4]
	intSlice[0] = 10
	fmt.Println(intArr) // [1 10 3 4 5 6]

	sliceInt := []int{10, 20, 30}
	fmt.Println(len(sliceInt), cap(sliceInt)) //3,3

	sliceInt = append(sliceInt, 40)
	fmt.Println(len(sliceInt), cap(sliceInt))//4,6

	for i,v := range sliceInt{
		fmt.Printf("index : %v , value : %v\n" , i, v)
	}
	
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5

	fmt.Println(b)

}

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8,dy)
	for i := range picture{
		picture[i] = make([]uint8,dx)
		for x:= range picture[i]{
			picture[i][x] = uint8((i+x)/2)
		}
	}

	return picture
}