# cachekeyhunter

CacheKeyHunter is a Go-based HTTP cache key scanner used to explore how headers and query parameters influence caching behaviour.

## Getting started

Clone the repository and build the scanner binary:

```
git clone https://github.com/CacheKeyHunter/cachekeyhunter.git
cd cachekeyhunter/ck
go build ./cmd/cachekeyhunter
```

## Wordlists

Version 2 expands the header and query parameter wordlists used for cache key testing:

- `wordlists/headers.txt` – common headers that may affect cache keys, such as `x-forwarded-proto`, `origin`, and `accept-encoding`.
- `wordlists/params.txt` – query parameters like `cb`, `utm_source`, `session`, and others.

These wordlists help reveal how different inputs change caching behaviour and expose potential cache-poisoning vectors.

## Usage

Run the scanner with a wordlist:

```
./cachekeyhunter -u https://example.com/page -w wordlists/headers.txt
```

This command targets the URL and iterates through the header wordlist. Swap in `wordlists/params.txt` to test query parameters instead.

## v2 improvements

Version 2 adds broader header and parameter coverage to aid testing for cache key discrepancies.
