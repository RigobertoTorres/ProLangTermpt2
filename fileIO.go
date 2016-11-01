package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Dumps contents of file into dat
	dat, err := ioutil.ReadFile("./goProLangTest.txt")
	check(err)

	strDat := (string(dat))
	n := len(dat)
	prev := strDat[0]
	curr := strDat[0]

	firstWord := true
	wordCount := 0

	// iterates through data, incrementing word count at each word
	for i := 1; i < n; i++ {
		prev = strDat[i-1]
		curr = strDat[i]

		// Keeps track of the beginning of new lines
		if curr == '\u000d' {
			firstWord = true
		}

		if firstWord == true {
			if prev != ' ' && curr == ' ' {
				wordCount = wordCount + 1
				firstWord = false
			}
		}

		if prev == ' ' && curr != ' ' {
			wordCount = wordCount + 1
		}

	}

	// Creates new file to write to
	f, err := os.Create("./outputFile")
	check(err)

	defer f.Close()

	// Converts word count to string and writes it to output file
	d1 := strconv.Itoa(wordCount)
	_, err2 := f.WriteString(d1)
	check(err2)

	fmt.Printf("There are %d words in \"%s\"\n %s\n", wordCount, strDat, "The data was written to \"outputFile\"")

}
