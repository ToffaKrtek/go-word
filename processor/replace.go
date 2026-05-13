package processor

import (
	"fmt"
	"strings"

	"github.com/ToffaKrtek/go-word/vars"
)

func ReplaceVariables(input string, data map[string]interface{}, vars []vars.Variable) string {
	for _, v := range vars {
		placeholder := fmt.Sprintf("${%s}", v.Name)
		value := v.Fn.Exec(v.FnMap, data)
		input = strings.ReplaceAll(input, placeholder, value)
	}
	return input
}

func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
