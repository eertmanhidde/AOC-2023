package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	re := regexp.MustCompile("[0-9]+")
	total := 0

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// create scanner from the file to read text from
	scanner := bufio.NewScanner(file)
	// scan or read the bytes of text line by line
	for scanner.Scan() {
		text := scanner.Text()
		strings := re.FindAllString(text, -1)

		value := calculateItAll(strings)
		total = total + value
	}

	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func calculateItAll(strings []string) (int) {
	firstItem := strings[0]
	lastItem := strings[len(strings)-1]

	firstCharacter := firstItem[0:1]
	lastCharacter := lastItem[len(lastItem)-1:]

	int1, _ := strconv.Atoi(firstCharacter + lastCharacter)

	return int1
}
