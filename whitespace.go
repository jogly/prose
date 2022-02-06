package prose

import "strings"

func CollapseSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
