package ler

import (
	"fmt"
	"strings"
)

// MustParse parses the ler code and panics if it can't be parsed.
func MustParse(code string) *Ler {
	l := &Ler{}

	rulesMatches := re.RulesHeader.FindAllString(code, -1)
	for _, rule := range rulesMatches {
		l.Rules = append(l.Rules, pler(rule, true))
	}

	return l
}

// Parse parses the ler code.
func Parse(code string) *Ler {
	l := &Ler{}

	rulesMatches := re.RulesHeader.FindAllString(code, -1)
	for _, rule := range rulesMatches {
		l.Rules = append(l.Rules, pler(rule, false))
	}

	return l
}

// pler stands for "p"arse "ler" code.
func pler(code string, must bool) Rule {
	rule := Rule{}

	// Extract Rule Name
	match := re.Rule.FindStringSubmatch(code)
	if len(match) < 2 && must {
		panic(fmt.Sprintf(errCouldntFind, `"rule" name`))
	} else {
		rule.Name = match[1]
	}

	if rule.Name == "" && must {
		panic(fmt.Sprintf(errInvalid, "rule name", rule.Name, rule.Name))
	}

	// Extract Condition
	match = re.Condition.FindStringSubmatch(code)
	if len(match) < 2 && must {
		panic(fmt.Sprintf(errCouldntFind, `"condition" or "cond" and/or "c"`))
	} else {
		rule.Condition = strings.TrimSpace(match[1])
	}

	if !isValidCond(rule.Condition) && must {
		panic(fmt.Sprintf(errInvalid, "condition", rule.Name, rule.Condition))
	}

	// Extract Rules
	match = re.Rules.FindStringSubmatch(code)
	if len(match) < 1 && must {
		panic(fmt.Sprintf(errCouldntFind, `rules definition`))
	}

	expression := re.Expression.FindAllStringSubmatch(match[0], -1)
	if len(expression) < 1 && must {
		panic(errMustHaveRule)
	}

	for _, match := range expression {
		var method, element string

		if len(match) < 2 && must {
			panic(errMustHaveRule)
		} else {
			ruleStr := match[1]
			methodMatch := re.Method.FindStringSubmatch(ruleStr)
			elementMatch := re.Element.FindStringSubmatch(ruleStr)
			patternMatch := re.Pattern.FindStringSubmatch(ruleStr)

			r := expr{}

			if len(methodMatch) < 1 && must {
				panic(fmt.Sprintf(errCouldntFind, `"method" or "m"`))
			} else {
				method = strings.TrimSpace(methodMatch[1])
			}

			if !isValidMethod(method) && must {
				panic(fmt.Sprintf(errInvalid, "method", rule.Name, method))
			}

			r.Method = toRequestMethod(method)

			if len(elementMatch) < 1 && must {
				panic(fmt.Sprintf(errCouldntFind, `"element" or "el" and/or "e"`))
			} else {
				element = strings.TrimSpace(elementMatch[1])
			}

			if !isValidElem(element) && must {
				panic(fmt.Sprintf(errInvalid, "element", rule.Name, element))
			}

			r.Element = toRequestElement(element)

			if len(patternMatch) < 1 && must {
				panic(fmt.Sprintf(errCouldntFind, `"pattern" or "p"`))
			} else {
				r.Pattern = strings.TrimSpace(patternMatch[1])
			}

			rule.Rules = append(rule.Rules, r)
		}
	}

	return rule
}
