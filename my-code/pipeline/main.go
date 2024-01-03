package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	data := getCSV()
	//synchronousPipeline(data)
	//asyncPipeline(data)
	asyncPipelineFan(data)
}

func getCSV() (reader *csv.Reader) {
	f, err := os.Open("./data.csv")

	if err != nil {
		log.Fatal(err)
	}

	reader = csv.NewReader(f)

	return
}
