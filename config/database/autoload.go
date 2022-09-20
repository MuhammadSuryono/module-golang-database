package database

import "fmt"

func Init() {
	InitConnectionFromEnvironment().CreateNewConnection()
}

func InitialDefaultValue(envKey interface{}, defaultValue interface{}) interface{} {
	if envKey == "" || envKey == nil {
		return defaultValue
	}
	return envKey
}

func InterfaceToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}
