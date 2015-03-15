package fragments

func GetOptionalParam(conf ConfDict, key string, defaultValue string) string {
	result, ok := conf[key]
	if !ok {
		result = defaultValue
	}
	return result
}
