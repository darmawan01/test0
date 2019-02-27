package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(s string) {

	newVal := strings.ToLower(s)
	aHidup, bHidup := strings.Count(newVal, "a"), strings.Count(newVal, "o")
	mati := 0

	for i := 0; i < len(newVal); i++ {
		if string(newVal[i]) != "a" && string(newVal[i]) != "o" {
			mati++
		}
	}

	if aHidup > 1 {
		aHidup = 1
	} else if bHidup > 0 {
		bHidup = 1
	}

	jlmHM := aHidup + bHidup

	fmt.Printf("hasil: huruf mati: %d, huruf hidup: %d \n", mati, jlmHM)
}

func main() {
	var newdata []string

	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}

	data, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		newdata = append(newdata, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for _, val := range newdata {
		check(val)
	}

	defer data.Close()
}
