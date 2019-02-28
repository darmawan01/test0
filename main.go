package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/test0/ext"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	mati, hidup := ext.Check(text)

	fmt.Printf("hasil: huruf mati: %d, huruf hidup: %d \n", mati, hidup)

}
