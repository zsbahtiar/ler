package main

import (
	"fmt"

	"github.com/kitabisa/ler"
)

func main() {
	lerCode := `                                                                                                      
    rule "Rule name":                                                                                                 
      condition: $OR                                                                                                  
      [                                                                                                               
        { method: $GET, element: $URI, pattern: "\$\{.*://.*\/?\w+?}" }                                               
        { method: $GET, element: $HEADERS, pattern: "curl" }                                                          
      ];                                                                                                           
     `

	ler := ler.MustParse(lerCode)
	for _, rule := range ler.Rules {
		fmt.Printf("Name: %s\nCondition: %s\n", rule.Name, rule.Condition)
		for i, r := range rule.Rules {
			fmt.Printf("Rule %d:\n\tMethod: %+v\n\tElement: %+v\n\tPattern: %s\n", i+1, r.Method, r.Element, r.Pattern)
		}
	}
	println()
	println(ler.Count())
}
