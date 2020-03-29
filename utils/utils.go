package utils

import (
	"fmt"
)

func RemoveDuplicatesFromSlice(s []interface{}) []interface{} {
	m := make(map[interface{}]bool)
	for _, item := range s {
		if _, ok := m[item]; ok {
			// duplicate item
			fmt.Println(item, "is a duplicate")
		} else {
			m[item] = true
		}
	}

	var result []interface{}
	for item, _ := range m {
		result = append(result, item)
	}
	return result
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}