package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/selimozcann/cachekeyhunter/ck/internal/report"
	"github.com/selimozcann/cachekeyhunter/ck/internal/scan"
	"github.com/selimozcann/cachekeyhunter/ck/internal/types"
)

func main() {
	url := flag.String("u", "", "Target URL")
	flag.StringVar(url, "url", "", "Target URL")

	headers := flag.String("w", "", "Path to header wordlist")
	flag.StringVar(headers, "wordlist", "", "Path to header wordlist")

	params := flag.String("q", "", "Path to query param wordlist")
	flag.StringVar(params, "query", "", "Path to query param wordlist")

	flag.Parse()

	if *url == "" || *headers == "" {
		fmt.Fprintf(os.Stderr, "Usage: %s -u <url> -w <headers> [-q <params>]\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *params != "" {
		fmt.Printf("Starting cache key scan on %s using headers %s and params %s\n", *url, *headers, *params)
	} else {
		fmt.Printf("Starting cache key scan on %s using headers %s\n", *url, *headers)
	}

	base, err := scan.GetBaseline(*url)
	if err != nil {
		fmt.Println("Baseline error:", err)
		return
	}

	// Load header variants
	variants := scan.GenerateHeaderVariants(*headers)

	// Load param variants if provided
	if *params != "" {
		variants = append(variants, scan.GenerateQueryVariants(*params)...)
	}

	// Scan with all variants
	for _, v := range variants {
		sig, err := scan.DoVariant(*url, v)
		if err != nil {
			continue
		}
		severity, detail := scan.Compare(base, sig)
		finding := types.Finding{
			URL:      *url,
			Severity: severity,
			Detail:   detail,
			Evidence: fmt.Sprintf("Variant: %s", v.Name),
		}
		report.PrintFinding(finding)
	}
}

