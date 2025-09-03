package report

import (
	"fmt"

	"github.com/fatih/color"
	constants "github.com/selimozcann/cachekeyhunter/ck/internal/constant"
	"github.com/selimozcann/cachekeyhunter/ck/internal/types"
)

func PrintFinding(f types.Finding) {
	switch f.Severity {
	case constants.SeverityHIGH:
		color.Red("[HIGH] %s", f.Detail)
	case constants.SeverityMedium:
		color.Yellow("[MEDIUM] %s", f.Detail)
	case constants.SeverityLow:
		color.Cyan("[LOW] %s", f.Detail)
	default:
		fmt.Printf("[INFO] %s\n", f.Detail)
	}

	fmt.Printf("  URL: %s\n", f.URL)
	if f.Evidence != "" {
		fmt.Printf("  Evidence: %s\n", f.Evidence)
	}
	fmt.Println()
}
