package main

import (
	"log"
	"os"
	"strings"
	"sync"
	"testing"
)

// ------------------------
// Writer interface and Mock Writer for testing
// ------------------------

type PageWriter interface {
	Write(PageData) error
	Close()
}

type MockWriter struct {
	data []PageData
	sync.Mutex
}

func (mw *MockWriter) Write(d PageData) error {
	mw.Lock()
	defer mw.Unlock()
	mw.data = append(mw.data, d)
	return nil
}

func (mw *MockWriter) Close() {}

// ------------------------
// unit test 1: real page scraping and content validation
// ------------------------

func TestScrapePageContent(t *testing.T) {
	var wg sync.WaitGroup
	writer := &MockWriter{}

	testURL := "https://en.wikipedia.org/wiki/Robotics"
	wg.Add(1)
	go ScrapePage(testURL, writer, &wg)
	wg.Wait()

	if len(writer.data) == 0 {
		t.Fatalf("Expected data but got none")
	}

	if !strings.Contains(writer.data[0].Text, "robot") {
		t.Errorf("Expected 'robot' in content, got: %s", writer.data[0].Text[:100])
	}
}

// ------------------------
// unit test 2: synthetic data
// ------------------------

func TestSyntheticDataWrite(t *testing.T) {
	writer := &MockWriter{}
	sample := PageData{
		URL:  "http://example.com",
		Text: "This is synthetic test data.",
	}

	err := writer.Write(sample)
	if err != nil {
		t.Errorf("Failed to write synthetic data: %v", err)
	}

	if len(writer.data) != 1 || writer.data[0].Text != sample.Text {
		t.Errorf("Synthetic data mismatch: got %v", writer.data)
	}
}

// ------------------------
// benchmark test
// ------------------------

func BenchmarkScraper(b *testing.B) {
	logFile, _ := os.Create("logs/benchmark.log")
	log.SetOutput(logFile)
	defer logFile.Close()

	writer := &MockWriter{}

	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(1)
		go ScrapePage("https://en.wikipedia.org/wiki/Robotics", writer, &wg)
		wg.Wait()
	}
}
