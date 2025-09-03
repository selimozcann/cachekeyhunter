package scan

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	constants "github.com/selimozcann/cachekeyhunter/ck/internal/constant"
	"github.com/selimozcann/cachekeyhunter/ck/internal/types"
)

func GenerateVariants() []types.Variant {
	var variants []types.Variant

	headers, _ := loadLines("wordlists/headers.txt")
	for _, h := range headers {
		variants = append(variants, types.Variant{
			Name:    fmt.Sprintf("%s: %s", h, constants.DefaultExampleDomain),
			Headers: map[string]string{h: constants.DefaultExampleDomain},
		})
	}

	params, _ := loadLines("wordlists/params.txt")
	for _, p := range params {
		parts := strings.SplitN(p, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key, val := parts[0], parts[1]
		variants = append(variants, types.Variant{
			Name:  fmt.Sprintf("Query %s=%s", key, val),
			Query: map[string]string{key: val},
		})
	}

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
