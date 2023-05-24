package ler

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
