package main

import (
	"bufio"
	"flag"
	"log/slog"
	"os"
	"strconv"

	"l1/internal/constants"
	"l1/internal/entity"
	"l1/internal/validation"
)

func main() {
	if len(os.Args) > 2 {
		panic(constants.ErrorTooManyArgs)
	}

	var path string
	flag.StringVar(&path, "path", "", "Path to a file")
	flag.Parse()

	if err := validation.ValidateFilePath(path); err != nil {
		panic(err)
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	list := entity.NewVocabulary()

	for scanner.Scan() {
		list.AddOrInc(scanner.Text())
	}

	slog.Info("Number of unique", "words", list.GetUniqueCount())

	topList := list.GetMostFrequent(10)

	slog.Info("Top 10 most frequent words")

	for i, word := range topList {
		slog.Info("", "index", i, "word", word.Text, "count", strconv.Itoa(word.Count))
	}
}
