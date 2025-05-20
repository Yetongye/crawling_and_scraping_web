package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	// read command line arguments
	outputPath := flag.String("output", "output/results.jl", "Path to output .jl file")
	urlList := flag.String("urls", "", "Optional: comma-separated list of URLs to crawl")
	flag.Parse()

	// test if the output path is writable
	if _, err := os.Stat(*outputPath); err != nil {
		dir := strings.TrimSuffix(*outputPath, "/"+filepathBase(*outputPath))
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("Error: Output directory not writable: %v\n", err)
			return
		}
	}

	// set up logging
	logFile, _ := os.Create("logs/run.log")
	log.SetOutput(logFile)
	defer logFile.Close()

	// use default URLs or user-provided URLs
	urls := WikipediaURLs
	if *urlList != "" {
		urls = strings.Split(*urlList, ",")
	}

	writer, err := NewJLWriter(*outputPath)
	if err != nil {
		fmt.Printf("Error: Unable to open output file: %v\n", err)
		return
	}
	defer writer.Close()

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go ScrapePage(url, writer, &wg)
	}

	wg.Wait()
	fmt.Println("Crawling complete. Output saved to:", *outputPath)
}

// WikipediaURLs is a list of Wikipedia URLs to scrape
func filepathBase(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}
