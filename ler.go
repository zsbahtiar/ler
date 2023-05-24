package ler

import (
	"regexp"

	"github.com/kitabisa/teler-waf/request"
)

type regexes struct {
	Condition, Element *regexp.Regexp
	Method, Pattern    *regexp.Regexp
	Expression, Rule   *regexp.Regexp
	Rules, RulesHeader *regexp.Regexp
}

type Ler struct {
	Rules []Rule
}

type expr struct {
	Method  request.Method
	Element request.Element
	Pattern string
}

type Rule struct {
	Name      string
	Condition string
	Rules     []expr
}
