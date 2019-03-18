package utils

import (
	"net/http"
	"html"
	"strings"
)

// Tools for General purposes

func RemoveTag(str string) string{
	return html.EscapeString(strings.Replace(str,"</","",-1))
}


func CheckOrInit (key string ,u map[string]string) (string) {
	value, ok := u[key]
	if ok == false {
		u[key] = ""
	}
	return RemoveTag(value)
}

func ProcessPostByDict (r *http.Request, dict []string ) (map[string]string) {
	target 	:= make(map[string]string)
	result  := make(map[string]string)

	for _ , keys := range dict {
		target[keys] = r.PostFormValue(keys)
	}
	// Filter The Empty
	for keys,val := range target {
		if val != "" {
			if value, ok := target[keys]; ok {
				result[keys] = RemoveTag(value)
			}
		}
	}
	return result
}
