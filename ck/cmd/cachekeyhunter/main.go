package main

import (
	"fmt"
	"os"

	"github.com/selimozcann/cachekeyhunter/ck/internal/report"
	"github.com/selimozcann/cachekeyhunter/ck/internal/scan"
	"github.com/selimozcann/cachekeyhunter/ck/internal/types"
)

func main() {
	url := os.Args[1]
	base, err := scan.GetBaseLİne(url)
	if err != nil {
		fmt.Println("Baseline error:", err)
		return
	}

	variants := scan.GenerateVariants()

	for _, v := range variants {
		sig, err := scan.DoVariant(url, v)
		if err != nil {
			continue
		}
		severity, detail := scan.Compare(base, sig)
		finding := types.Finding{
			URL:      url,
			Severity: severity,
			Detail:   detail,
			Evidence: fmt.Sprintf("Variant: %s", v.Name),
		}
		report.PrintFinding(finding)
	}
}
