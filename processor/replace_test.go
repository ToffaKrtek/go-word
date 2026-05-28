package processor

import (
	"testing"

	"github.com/ToffaKrtek/go-word/vars"
	"github.com/stretchr/testify/assert"
)

func TestReplaceVariables_SinglePlaceholder(t *testing.T) {
	input := "Привет, ${Имя}! Дата: ${Дата}."
	data := map[string]any{
		"name": "Алексей",
		"date": "2023-01-01",
	}
	vars := []vars.Variable{
		vars.NewVariable(
			"Имя",
			vars.FnGet,
			map[string]any{"data_get": "name"},
			map[string]string{},
		),
		vars.NewVariable(
			"Дата",
			vars.FnFormat,
			map[string]any{"date_format": "02.01.2006", "data_get": "date"},
			map[string]string{},
		),
	}
	expected := "Привет, Алексей! Дата: 01.01.2023."
	result := ReplaceVariables(input, data, vars)
	assert.Equal(t, expected, result)
}
