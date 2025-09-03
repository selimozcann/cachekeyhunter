# cachekeyhunter

CacheKeyHunter is a Go-based HTTP cache key scanner.

## v2 improvements

Version 2 expands the header and query parameter wordlists used for cache key testing. These wordlists help identify how different inputs affect caching behavior and potential cache poisoning vectors.

### Wordlists

- `wordlists/headers.txt` – common headers that may influence cache keys, such as `x-forwarded-proto`, `origin`, and `accept-encoding`.
- `wordlists/params.txt` – query parameters like `cb`, `utm_source`, `session`, and more.

### Example usage

From the `ck` directory:

```
./cachekeyhunter -u https://example.com/page -w wordlists/headers.txt
```

This command scans the target URL using the header wordlist. Use the parameter wordlist to test query parameters as needed.
