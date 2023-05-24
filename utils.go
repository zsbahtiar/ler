package ler

import "github.com/kitabisa/teler-waf/request"

// Count parsed rules.
func (l *Ler) Count() int {
	return len(l.Rules)
}

func toRequestMethod(s string) request.Method {
	var method request.Method

	switch s {
	case "GET":
		method = request.GET
	case "HEAD":
		method = request.HEAD
	case "POST":
		method = request.POST
	case "PUT":
		method = request.PUT
	case "PATCH":
		method = request.PATCH
	case "DELETE":
		method = request.DELETE
	case "CONNECT":
		method = request.CONNECT
	case "OPTIONS":
		method = request.OPTIONS
	case "TRACE":
		method = request.TRACE
	case "ALL":
		method = request.ALL
	default:
		method = request.UNDEFINED
	}

	return method
}

func toRequestElement(s string) request.Element {
	var element request.Element

	switch s {
	case "URI":
		element = request.URI
	case "HEADERS":
		element = request.Headers
	case "BODY":
		element = request.Body
	case "ANY":
		element = request.Any
	}

	return element
}
