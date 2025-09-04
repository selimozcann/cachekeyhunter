package scan

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	constants "github.com/selimozcann/cachekeyhunter/ck/internal/constant"
	"github.com/selimozcann/cachekeyhunter/ck/internal/types"
)

// GenerateHeaderVariants loads headers from wordlist and builds header-based variants
func GenerateHeaderVariants(wordlistPath string) []types.Variant {
	var variants []types.Variant

	if wordlistPath == "" {
		return variants
	}

	lines, _ := loadLines(wordlistPath)
	for _, line := range lines {
		variants = append(variants, types.Variant{
			Name:    fmt.Sprintf("%s: %s", line, constants.DefaultExampleDomain),
			Headers: map[string]string{line: constants.DefaultExampleDomain},
		})
	}

	// Add hardcoded Forwarded and X-Forwarded-Proto variants
	variants = append(variants,
		types.Variant{
			Name: fmt.Sprintf("%s: host=%s", constants.HeaderForwarded, constants.DefaultExampleDomain),
			Headers: map[string]string{
				constants.HeaderForwarded: constants.DefaultForwardedPrefix + constants.DefaultExampleDomain,
			},
		},
		types.Variant{
			Name: fmt.Sprintf("%s: %s", constants.HeaderXForwardedProto, constants.DefaultProto),
			Headers: map[string]string{
				constants.HeaderXForwardedProto: constants.DefaultProto,
			},
		},
	)

	return variants
}

// GenerateQueryVariants loads query params from wordlist and builds query-based variants
func GenerateQueryVariants(wordlistPath string) []types.Variant {
	var variants []types.Variant

	lines, _ := loadLines(wordlistPath)
	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key, val := parts[0], parts[1]
			variants = append(variants, types.Variant{
				Name:  fmt.Sprintf("Query %s=%s", key, val),
				Query: map[string]string{key: val},
			})
		}
	}
	return variants
}

// loadLines reads a file line-by-line and returns non-empty trimmed lines
func loadLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, scanner.Err()
}
