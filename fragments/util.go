package fragments

import (
	"strconv"
)

func (c ConfDict) GetOptionalParam(key string, defaultValue string) string {
	result, ok := c[key]
	if !ok {
		result = defaultValue
	}
	return result
}

func (c ConfDict) GetOptionalParamBool(
	key string,
	defaultValue string) (result bool) {
	result, err := strconv.ParseBool(c.GetOptionalParam(key, defaultValue))
	if err != nil {
		result = false
	}
	return
}
