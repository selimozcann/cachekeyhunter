package scan

import (
	constants "github.com/selimozcann/cachekeyhunter/ck/internal/constant"
	"github.com/selimozcann/cachekeyhunter/ck/internal/types"
)

func GenerateVariants() []types.Variant {
	return []types.Variant{
		{
			Name: "X-Forwarded-Host: attacker.example",
			Headers: map[string]string{
				constants.HeaderXForwardedHost: constants.DefaultAttackerDomain,
			},
		},
		{
			Name: "Forwarded: host=attacker.example",
			Headers: map[string]string{
				constants.HeaderForwarded: constants.DefaultForwardedPrefix + constants.DefaultAttackerDomain,
			},
		},
		{
			Name: "X-Forwarded-Proto: http",
			Headers: map[string]string{
				constants.HeaderXForwardedProto: constants.DefaultProto,
			},
		},
		// Instance more smarter >= v2
		// You can easily extend here with more smart variants
		// Example:
		// {
		//     Name: "Via: 1.1 attacker.example",
		//     Headers: map[string]string{
		//         constants.HeaderVia: constants.DefaultViaValue,
		//     },
		// },
	}
}
