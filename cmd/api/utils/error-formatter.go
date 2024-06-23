package utils

import (
	"encoding/json"
	"log"
)

func ErrorFormatter(errorMap map[string]error) (string, error) {
	stringMap := make(map[string]string)

	for key, err := range errorMap {
		stringMap[key] = err.Error()
	}

	jsonData, err := json.Marshal(stringMap)

	if err != nil {
		log.Print(err)
		return "" ,err
	}

	return string(jsonData), nil
}