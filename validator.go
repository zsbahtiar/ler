package ler

import "net/http"

// Validation Functions
func isValidCond(c string) bool {
	validConds := []string{"AND", "OR"}
	for _, vc := range validConds {
		if c == vc {
			return true
		}
	}

	return false
}

func isValidMethod(m string) bool {
	validMethods := []string{
		"ALL", "GET", "HEAD", "POST",
		"PUT", "PATCH", "DELETE",
		"CONNECT", "OPTIONS", "TRACE",
	}
	for _, vm := range validMethods {
		if m == vm {
			return true
		}
	}

	return false
}

func isValidElem(e string) bool {
	validElems := []string{
		"URI", "HEADERS",
		"BODY", "ANY",
	}
	for _, ve := range validElems {
		if e == ve {
			return true
		}
	}

	return false
}

var validConditions = map[string]struct{}{
	"AND": {},
	"OR":  {},
}

func isValidCondition(c string) bool {
	_, isValid := validConditions[c]
	return isValid
}

var validHttpMethods = map[string]struct{}{
	"ALL":              {},
	http.MethodGet:     {},
	http.MethodHead:    {},
	http.MethodPost:    {},
	http.MethodPut:     {},
	http.MethodPatch:   {},
	http.MethodDelete:  {},
	http.MethodConnect: {},
	http.MethodOptions: {},
	http.MethodTrace:   {},
}

func isValidHttpMethod(m string) bool {
	_, isValid := validHttpMethods[m]
	return isValid
}

var validElements = map[string]struct{}{
	"URI":     {},
	"HEADERS": {},
	"BODY":    {},
	"ANY":     {},
}

func isValidElement(e string) bool {
	_, isValid := validElements[e]
	return isValid
}
