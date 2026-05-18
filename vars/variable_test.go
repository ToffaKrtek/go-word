package vars

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFnPipeline(t *testing.T) {
	assert.Equal(
		t,
		"Hello World",
		fnPipeline(
			map[string]any{
				"pipeline": []any{
					map[string]any{
						"fn":     "echo",
						"fn_map": map[string]any{"text": "Hello World"},
					},
				},
			}, map[string]any{}))
	assert.Equal(
		t,
		"Prev Hello World",
		fnPipeline(
			map[string]any{
				"pipeline": []any{
					map[string]any{
						"fn": "concat",
						"fn_map": map[string]any{
							"concat_keys": []any{"HelloKey", "WorldKey"},
							"separator":   " ",
						},
					},
					map[string]any{
						"fn": "concat",
						"fn_map": map[string]any{
							"concat_keys": []any{"PrevKey", "*"},
							"separator":   " ",
						},
					},
				},
			},
			map[string]any{
				"WorldKey": "World",
				"HelloKey": "Hello",
				"PrevKey":  "Prev",
			},
		))
}

func TestFnEcho(t *testing.T) {
	assert.Equal(t, "Hello World", fnEcho(map[string]any{"text": "Hello World"}))
	assert.Equal(t, "-", fnEcho(map[string]any{}))
}

func TestFnGet(t *testing.T) {
	assert.Equal(t, "World", fnGet(map[string]any{"data_get": "WorldKey"}, map[string]any{"WorldKey": "World"}))
	assert.Equal(t, "-", fnGet(map[string]any{"data_get": "WorldKey"}, map[string]any{}))
}

func TestFnFormat(t *testing.T) {
	assert.Equal(t, "02.01.2006", fnFormat(map[string]any{"data_get": "DateKey", "date_format": "02.01.2006"}, map[string]any{"DateKey": "2006-01-02"}))
	assert.Equal(t, "-", fnFormat(map[string]any{"data_get": "DateKey", "date_format": "02.01.2006"}, map[string]any{}))
	assert.Equal(t, "02.01.2006", fnFormat(map[string]any{"data_get": "DateKey"}, map[string]any{"DateKey": "2006-01-02"}))
	now := time.Now()
	assert.Equal(t, now.Format("02.01.2006"), fnFormat(map[string]any{"data_get": "DateKey"}, map[string]any{"DateKey": "wrong-date"}))
}

func TestFnConcat(t *testing.T) {
	assert.Equal(t, "HelloWorld", fnConcat(map[string]any{"concat_keys": []any{"HelloKey", "WorldKey"}}, map[string]any{"WorldKey": "World", "HelloKey": "Hello"}))
	assert.Equal(t, "Hello World", fnConcat(map[string]any{"concat_keys": []any{"HelloKey", "WorldKey"}, "separator": " "}, map[string]any{"WorldKey": "World", "HelloKey": "Hello"}))
	assert.Equal(t, "Hello World", fnConcat(map[string]any{"concat_keys": []any{"HelloKey", map[string]any{"fn": "echo", "fn_map": map[string]any{"text": "World"}}}, "separator": " "}, map[string]any{"HelloKey": "Hello"}))
}

func TestFnExecFunction(t *testing.T) {
	assert.Equal(t, "Hello World", ExecFunction(map[string]any{"fn": "echo", "fn_map": map[string]any{"text": "Hello World"}}, map[string]any{}))
	assert.Equal(t, "-", ExecFunction(map[string]any{}, map[string]any{}))
	assert.Equal(t, "-", ExecFunction(map[string]any{"fn": ""}, map[string]any{}))
	assert.Equal(t, "-", ExecFunction(map[string]any{"fn": "echo", "fn_map": map[string]any{}}, map[string]any{}))
}

func TestParseFnType(t *testing.T) {
	for i, fType := range FnTypes {
		assert.Equal(t, i, int(ParseFnType(fType)))
	}
	assert.Equal(t, -1, int(ParseFnType("wrong")))
}
