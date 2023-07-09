package licensedb

import (
	"github.com/anton-l/go-license-detector/licensedb/internal"
)

// InvestigateLicenseText takes the license text and returns the most probable reference licenses matched.
// Each match has the confidence assigned, from 0 to 1, 1 means 100% confident.
func InvestigateLicenseText(text []byte) map[string]float32 {
	return internal.InvestigateLicenseText(text)
}
