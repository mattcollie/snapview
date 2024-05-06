package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"encoding/csv"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the Snapchat history HTML file as an argument.")
	}

	htmlFilename := os.Args[1]
	data, err := ioutil.ReadFile(htmlFilename)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	if err != nil {
		log.Fatal(err)
	}

	// Open the CSV file
	csvFile, err := os.Create("snapchat_history.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	// Write the CSV header
	csvWriter.Write([]string{"Username", "Timestamp", "Type", "Message"})

	// Replace these selectors with ones that match your HTML file
	doc.Find(".rightpanel").Each(func(i int, s *goquery.Selection) {
		s.Find("div[style=\"background: #f2f2f2; border-radius: 7px; padding: 3px; margin-bottom:4px;\"]").Each(func(i int, s *goquery.Selection) {
			username := s.Find("h4").Text()
			timestamp := s.Find("h6").Text()
			message_type := s.Find("span[style*=\"left: 50%; position: absolute; font-weight: bold; padding-top: 13px;\"]").Text()
			message := s.Find("p").Text()
			csvWriter.Write([]string{username, timestamp, message_type, message})
		})
	})
}
