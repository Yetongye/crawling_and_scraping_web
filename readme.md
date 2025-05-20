
# Concurrent Go Wikipedia Web Crawler

## Overview

This project implements a concurrent web crawler and scraper written in Go, using the [Colly](https://github.com/gocolly/colly) framework. It is designed to collect information from Wikipedia pages related to intelligent systems and robotics. The output is stored as a JSON Lines (`.jl`) file where each line corresponds to the main content of one crawled page.

The crawler supports:
- Concurrent processing using goroutines and sync.WaitGroup
- Mutex-guarded file writing to prevent race conditions
- Logging of page status and performance
- Configurable input via command-line flags



## How to Build and Run (Application)

### Build the executable:
This repository already includes pre-built executables:

**On macOS/Linux:**
```bash
go build -o crawler.app .
```

**On Windows:**
```bash
GOOS=windows GOARCH=amd64 go build -o crawler.exe .
```



### Run the application:

```bash
./crawler.app
```

#### Available Flags

- `-output`: Path to output `.jl` file. Default is `output/results.jl`.
- `-urls`: (Optional) Comma-separated list of custom URLs. If omitted, defaults from `urls.go` will be used.

#### Examples:
```bash
go run . -output="output/wikipedia_ai.jl"
```

```bash
./crawler.app -output="output/robots.jl" -urls="https://en.wikipedia.org/wiki/Robot,https://en.wikipedia.org/wiki/Chatbot"
```



### Error Handling & User Guidance

| Scenario                       | Behavior / Output                                              |
|-------------------------------|----------------------------------------------------------------|
| Output path not writable      | Prints: `Error: Output directory not writable`                |
| URLs unreachable or 404       | Logs: `[ERROR] Failed to scrape...`                            |
| All successful crawls         | Logs: `[SUCCESS] Scraped ... in 400ms`                         |
| No flags provided             | Uses default URLs and outputs to `output/results.jl`           |



## How to Test & What Has Been Tested

### Run Unit + Benchmark Tests:

```bash
go test -v
go test -bench=. -benchmem
```

### What has been tested:

| Component             | Test Description                                                                 |
|----------------------|----------------------------------------------------------------------------------|
| `ScrapePage`         | Verified it scrapes expected Wikipedia content (`robot`, `intelligent`, etc.)   |
| Synthetic write test | Verified mock writer correctly stores JSON content                              |
| Memory usage test    | Measured before/after RAM allocation via `runtime.MemStats`                     |
| Benchmarking         | Multiple round performance test with `go test -bench`                           |



## Output Format Example (`results.jl`)

Each line is a standalone JSON object:

```json
{"url":"https://en.wikipedia.org/wiki/Robotics","text":"Robotics is an interdisciplinary branch of computer science..."}
{"url":"https://en.wikipedia.org/wiki/Robot","text":"A robot is a machine—especially one programmable by a computer..."}
```



## Performance Observation

All tests passed successfully. The crawler extracted the expected content (Robotics) from Wikipedia and correctly stored it in memory using the MockWriter.

Total test runtime: 0.374 seconds

This confirms both data collection and write-to-storage components are working correctly.




## GenAI Tools

Generative AI tools were used to assist in:

- Designing the modular architecture of the Go crawler
- Drafting unit tests and performance benchmarks
- Helping generate README documentation and interpreting benchmark results

Tool used: **ChatGPT**

All GenAI-assisted code was verified manually and integrated following Go idioms.
