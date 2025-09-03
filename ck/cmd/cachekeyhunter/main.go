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
	wordlist := flag.String("w", "", "Path to header or parameter wordlist")
	flag.StringVar(wordlist, "wordlist", "", "Path to header or parameter wordlist")

	flag.Parse()

	if *url == "" || *wordlist == "" {
		fmt.Fprintf(os.Stderr, "Usage: %s -u <url> -w <wordlist>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("Starting cache key scan on %s using %s\n", *url, *wordlist)

	base, err := scan.GetBaseLİne(*url)
	if err != nil {
		fmt.Println("Baseline error:", err)
		return
	}

	variants := scan.GenerateVariants(*wordlist)

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
