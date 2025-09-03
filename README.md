# CacheKeyHunter

CacheKeyHunter is a CLI tool for exploring how reverse proxies and CDNs build their cache keys. It sends crafted requests to observe how headers and query parameters influence cached responses.

## Features

- **Header poisoning detection** – test individual headers to see if they alter the cache key
- **Query poisoning detection** – append query parameters and watch for cache key differences

## Requirements

- Go 1.22+

## Installation

```bash
git clone https://github.com/CacheKeyHunter/cachekeyhunter.git
cd cachekeyhunter/ck
go build ./cmd/cachekeyhunter
```

## Usage

```bash
$ go run ./ck/cmd/cachekeyhunter -u https://example.com -w ck/wordlists/headers.txt
```
