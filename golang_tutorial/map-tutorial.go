package main

import (
	"fmt"
	"strings"
)

func mapTutorial() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println(m["Answer"])//42
	m["Answer"] = 48
	fmt.Println(m["Answer"])//48
	value, ok := m["Answer"]
	fmt.Println(value,ok)//48 true
	delete(m, "Answer")
	fmt.Println(m["Answer"])//0
	value, ok = m["Answer"]
	fmt.Println(value, ok)//0 false
}

func WordCount(s string) map[string]int {
	
	wordSlice := strings.Fields(s)
	wordMap := make(map[string]int)
	for _,v :=range wordSlice {
		value,exist := wordMap[v]
		if exist {
			value += 1
			wordMap[v]=value
		} else{
			wordMap[v]=1
		}
		
	}
	return wordMap
}