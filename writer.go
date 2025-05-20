package main

import (
	"encoding/json"
	"os"
	"sync"
)

type PageData struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}

type JLWriter struct {
	file  *os.File
	mutex sync.Mutex
}

type Writer interface {
	Write(PageData) error
	Close()
}

func NewJLWriter(filename string) (*JLWriter, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return &JLWriter{file: f}, nil
}

func (w *JLWriter) Write(data PageData) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = w.file.WriteString(string(jsonBytes) + "\n")
	return err
}

func (w *JLWriter) Close() {
	w.file.Close()
}
