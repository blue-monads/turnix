package modules

import (
	"fmt"
	"strings"
)

func ParamaterizedSchema(baseSchema string, pid int64) string {
	return strings.ReplaceAll(baseSchema, "__project__", fmt.Sprintf("z_%d_", pid))
}
