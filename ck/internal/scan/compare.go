package scan

import (
	Constants "github.com/selimozcann/cachekeyhunter/ck/internal/constant"
	"github.com/selimozcann/cachekeyhunter/ck/internal/types"
)

func Compare(base, variant types.Signals) (string, string) {
	if base.BodyHash != variant.BodyHash && variant.Cache == Constants.Hit {
		return Constants.SeverityHIGH, "Body changed and cached (X-Cache: HIT)"
	}
	if variant.Cache == Constants.Hit {
		return Constants.SeverityMedium, "Cached response with same body"
	}
	if variant.Age > base.Age {
		return Constants.SeverityLow, "Age increased"
	}
	return "INFO", "No significant difference"
}
