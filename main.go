package main

import (
	"bufio"
	"fmt"
	"os"
)

func resolve(value []string) {
	fmt.Println(value)
	fmt.Printf(`hasil: "huruf mati: %d, huruf hidup: %d"`, 0, 0)
}

func main() {
	fmt.Print("Load Test .....\n\n")
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

	resolve(newdata)

	defer data.Close()
}
