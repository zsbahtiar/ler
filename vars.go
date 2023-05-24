package ler

import "regexp"

var re = regexes{
	Condition:   regexp.MustCompile(`(?:condition|cond|c)\s*:\s*\$(\w+)`),
	Element:     regexp.MustCompile(`(?:element|el|e)\s*:\s*\$(\w+)`),
	Expression:  regexp.MustCompile(`{(.+)}`),
	Method:      regexp.MustCompile(`(?:method|m)\s*:\s*\$(\w+)`),
	Pattern:     regexp.MustCompile(`(?:pattern|p)\s*:\s*['"](.+)['"]`),
	Rule:        regexp.MustCompile(`(?:rule|r)\s+['"](.+?)['"]`),
	Rules:       regexp.MustCompile(`\[[\s\S]*?\];`),
	RulesHeader: regexp.MustCompile(`(rule[\s\S]*?);`),
}
