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
	mati := (len(newVal) - 1) - (aHidup + bHidup)

	if aHidup > 1 {
		aHidup = 1
	} else if bHidup > 0 {
		bHidup = 1
	}

	jlmHM := aHidup + bHidup

	fmt.Printf("hasil: huruf mati: %d, huruf hidup: %d \n", mati, jlmHM)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	check(text)

}
