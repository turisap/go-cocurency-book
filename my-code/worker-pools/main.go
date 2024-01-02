package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

type Result struct {
	file       string
	lineNumber int
	text       string
}

type Work struct {
	file    string
	pattern *regexp.Regexp
	result  chan Result
}

func main() {
	jobs := make(chan Work)
	allResults := make(chan chan Result)
	regE, err := regexp.Compile(os.Args[2])
	var wg sync.WaitGroup

	if err != nil {
		fmt.Println("Error compiling regex")
		log.Fatal(err)
	}
	go func() {
		defer close(allResults)
		filepath.Walk(os.Args[1], func(path string, d fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() {
				ch := make(chan Result)
				jobs <- Work{file: path, pattern: regE, result: ch}
				allResults <- ch
			}
			return nil
		})
	}()

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(jobs)
		}()
	}

	for resultCh := range allResults {
		for result := range resultCh {
			fmt.Printf("%s #%d: %s\n", result.file, result.lineNumber, result.text)
		}
	}

	close(jobs)
	wg.Wait()
}

func worker(jobs chan Work) {
	for work := range jobs {
		f, err := os.Open(work.file)
		if err != nil {
			fmt.Println("Error opening file")
			fmt.Println(err)
			continue
		}
		scn := bufio.NewScanner(f)
		lineNumber := 1
		for scn.Scan() {
			result := work.pattern.Find(scn.Bytes())
			if len(result) > 0 {
				work.result <- Result{
					file:       work.file,
					lineNumber: lineNumber,
					text:       string(result),
				}
			}
			lineNumber++
		}

		close(work.result)
		f.Close()
	}
}
