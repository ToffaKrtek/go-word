package vars

import (
	"fmt"
	"strings"
	"time"
)

type Variable struct {
	Name      string `json:"name"`
	Fn        FnType `json:"fn"`
	FnMap     map[string]any
	Condition map[string]string
}

type FnType int

var FnTypes = []string{
	"echo",
	"get",
	"concat",
	"format",
	"pipeline",
}

const (
	FnEcho FnType = iota
	FnGet
	FnConcat
	FnFormat
	FnPipeline
)

func ParseFnType(s string) FnType {
	switch s {
	case "echo":
		return FnEcho
	case "get":
		return FnGet
	case "concat":
		return FnConcat
	case "format":
		return FnFormat
	case "pipeline":
		return FnPipeline
	default:
		return -1
	}
}

func ExecFunction(spec map[string]any, data map[string]any) string {
	fnRaw, ok := spec["fn"]
	if !ok {
		return "-"
	}
	fnStr, ok := fnRaw.(string)
	if !ok {
		return "-"
	}
	fnType := ParseFnType(fnStr)
	if fnType == -1 {
		return "-"
	}
	fnMapRaw, ok := spec["fn_map"]
	if !ok {
		fnMapRaw = map[string]any{}
	}
	fnMap, ok := fnMapRaw.(map[string]any)
	if !ok {
		return "-"
	}
	return fnType.Exec(fnMap, data)
}

func (v Variable) Exec(data map[string]any) string {
	return v.Fn.Exec(v.FnMap, data)
}

func NewVariable(name string, fn FnType, fnMap map[string]any, condition map[string]string) Variable {
	return Variable{
		Name:      name,
		Fn:        fn,
		FnMap:     fnMap,
		Condition: condition,
	}
}

func (f FnType) Exec(fnMap map[string]any, data map[string]any) string {
	switch f {
	case FnEcho:
		return fnEcho(fnMap)
	case FnGet:
		return fnGet(fnMap, data)
	case FnFormat:
		return fnFormat(fnMap, data)
	case FnConcat:
		return fnConcat(fnMap, data)
	case FnPipeline:
		return fnPipeline(fnMap, data)
	default:
		return "-"
	}
}

func fnPipeline(fnMap map[string]any, data map[string]any) string {
	fnsRaw, ok := fnMap["pipeline"]
	if !ok {
		return "-"
	}
	fnsSlice, ok := fnsRaw.([]any)
	if !ok {
		return "-"
	}
	for _, fn := range fnsSlice {
		fn, ok := fn.(map[string]any)
		if !ok {
			break
		}
		res := ExecFunction(fn, data)
		if res == "-" {
			return "-"
		}
		data["*"] = res
	}

	if val, ok := data["*"]; ok {
		return val.(string)
	}
	return "-"
}

func fnConcat(fnMap map[string]any, data map[string]any) string {
	separator := ""
	if sep, ok := fnMap["separator"]; ok {
		if str, ok := sep.(string); ok {
			separator = str
		}
	}
	keysRaw, ok := fnMap["concat_keys"]
	if !ok {
		return "-"
	}
	keysSlice, ok := keysRaw.([]any)
	if !ok {
		return "-"
	}
	var parts []string

	for _, key := range keysSlice {
		switch v := key.(type) {
		case string:
			parts = append(parts, fnGet(map[string]any{"data_get": v}, data))
		case map[string]any:
			parts = append(parts, ExecFunction(v, data))
		default:
			parts = append(parts, fmt.Sprintf("%v", key))
		}
	}
	return strings.Join(parts, separator)
}

func fnFormat(fnMap map[string]any, data map[string]any) string {
	if key, ok := fnMap["data_get"]; ok {
		if val, ok := data[key.(string)]; ok {

			t, err := time.Parse("2006-01-02", val.(string))
			if err != nil {
				t = time.Now()
			}
			if format, ok := fnMap["date_format"]; ok {
				return t.Format(format.(string))
			}
			return t.Format("02.01.2006")
		}
	}
	return "-"
}

func fnGet(fnMap map[string]any, data map[string]any) string {
	if key, ok := fnMap["data_get"]; ok {
		if val, ok := data[key.(string)]; ok {
			return val.(string)
		}
	}
	return "-"
}

func fnEcho(fnMap map[string]any) string {
	if val, ok := fnMap["text"]; ok {
		return val.(string)
	}
	return "-"
}
