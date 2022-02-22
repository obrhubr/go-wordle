package util

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Get_words_from_file(filename string) []string {
	words, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return strings.Split(string(words), "\n")
}
