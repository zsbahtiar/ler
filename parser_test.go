package ler

import "testing"

func BenchmarkMustParse(b *testing.B) {
	code := `
		rule "Rule 1":
		  condition: $OR
		  [
		    { method: $GET, element: $URI, pattern: "/api/posts" }
		  ];
	`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MustParse(code)
	}
}

func BenchmarkParse(b *testing.B) {
	code := `
		rule "Rule 1":
		  condition: $OR
		  [
		    { method: $GET, element: $URI, pattern: "/api/posts" }
		  ];
	`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Parse(code)
	}
}
