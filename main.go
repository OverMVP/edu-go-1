package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"l1/internal/constants"
	"l1/internal/entity"
	"l1/internal/validation"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println(constants.TOO_MANY_ARGS)
		os.Exit(1)
	}

	var path string
	flag.StringVar(&path, "path", "", "Path to a file")
	flag.Parse()

	if err := validation.ValidateFilePath(path); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	list := entity.NewVocabulary()

	for scanner.Scan() {
		list.AddOrInc(scanner.Text())
	}

	fmt.Printf(constants.UNIQUE_COUNT_TITLE, list.GetUniqueCount())

	topList := list.GetMostFrequent(10)

	fmt.Printf(constants.MOST_FREQUENT_TITLE, 10)

	for i, word := range topList {
		fmt.Printf(constants.LIST_ITEM, i+1, word.Text, word.Count)
	}
}
