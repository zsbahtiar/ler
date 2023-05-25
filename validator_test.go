package ler

import "testing"

var inputs = []string{
	"ALL", "GET", "HEAD", "POST",
	"PUT", "PATCH", "DELETE",
	"CONNECT", "OPTIONS", "TRACE",
	"AND", "OR",
	"URI", "HEADERS", "BODY", "ANY",
}

func BenchmarkIsValidCond(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, in := range inputs {
			_ = isValidCond(in)
		}
	}
}

func BenchmarkIsValidMethod(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, in := range inputs {
			_ = isValidMethod(in)
		}
	}
}

func BenchmarkIsValidElem(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, in := range inputs {
			_ = isValidElem(in)
		}
	}
}

func BenchmarkIsValidCondition(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, in := range inputs {
			_ = isValidCondition(in)
		}
	}
}

func BenchmarkIsValidHttpMethod(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, in := range inputs {
			_ = isValidHttpMethod(in)
		}
	}
}

func BenchmarkIsValidElement(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, in := range inputs {
			_ = isValidElement(in)
		}
	}
}
