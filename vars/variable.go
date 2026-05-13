package vars

import "time"

type Variable struct {
	Name      string `json:"name"`
	Fn        FnType `json:"fn"`
	FnMap     map[string]interface{}
	Condition map[string]string
}

func NewVariable(name string, fn FnType, fnMap map[string]interface{}, condition map[string]string) Variable {
	return Variable{
		Name:      name,
		Fn:        fn,
		FnMap:     fnMap,
		Condition: condition,
	}
}

type FnType int

const (
	FnEcho FnType = iota
	FnGet
	FnConcat
	FnFormat
)

func (f FnType) Exec(fnMap map[string]interface{}, data map[string]interface{}) string {
	switch f {
	case FnEcho:
		return fnEcho(fnMap)
	case FnGet:
		return fnGet(fnMap, data)
	case FnFormat:
		return fnFormat(fnMap, data)
	default:
		return "-"
	}
}

func fnFormat(fnMap map[string]interface{}, data map[string]interface{}) string {
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

func fnGet(fnMap map[string]interface{}, data map[string]interface{}) string {
	if key, ok := fnMap["data_get"]; ok {
		if val, ok := data[key.(string)]; ok {
			return val.(string)
		}
	}
	return "-"
}

func fnEcho(fnMap map[string]interface{}) string {
	if val, ok := fnMap["text"]; ok {
		return val.(string)
	}
	return "-"
}
