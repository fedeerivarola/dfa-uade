package main

import (
	"regexp"
	"strings"
)

const (
	MustBeEvenRestrictionName           = "Must be even"
	MustBeOddRestrictionName            = "Must be odd"
	ContainsSubstringRestrictionName    = "Contains substring restriction"
	NotContainsSubstringRestrictionName = "Not contains substring restriction"
	SuffixRestrictionName               = "Suffix must be.."
	PrefixRestrictionName               = "Prefix must be.."
	RegexRestrictionName                = "Regex restriction"
)

type MustBeEvenRestriction struct {
	name string
}

type MustBeOddRestriction struct {
	name string
}

type ContainsSubstringRestriction struct {
	name      string
	substring string
}

type NotContainsSubstringRestriction struct {
	name      string
	substring string
}

type PrefixRestriction struct {
	name   string
	prefix string
}

type SuffixRestriction struct {
	name   string
	suffix string
}

type RegexRestriction struct {
	name    string
	pattern string
}

func (ipr MustBeEvenRestriction) apply(stack []Symbol, z Symbol) (bool, error) {
	count := 0

	for i := 0; i < len(stack); i++ {
		y := stack[i].value
		if y == z.value {
			count++
		}
	}

	return count%2 == 0, nil
}

func (r MustBeOddRestriction) apply(stack []Symbol, z Symbol) (bool, error) {
	count := 0

	for i := 0; i < len(stack); i++ {
		y := stack[i].value
		if y == z.value {
			count++
		}
	}

	return count%2 != 0, nil
}

func (r ContainsSubstringRestriction) apply(stack []Symbol, z Symbol) (bool, error) {
	var s string
	for i := 0; i < len(stack); i++ {
		s += stack[i].value
	}
	s += z.value

	return strings.Contains(s, r.substring), nil
}

func (r NotContainsSubstringRestriction) apply(stack []Symbol, z Symbol) (bool, error) {
	var s string
	for i := 0; i < len(stack); i++ {
		s += stack[i].value
	}
	s += z.value

	return !strings.Contains(s, r.substring), nil
}

func (r PrefixRestriction) apply(stack []Symbol, z Symbol) (bool, error) {
	var s string
	for i := 0; i < len(stack); i++ {
		s += stack[i].value
	}
	s += z.value
	//inverted
	return strings.HasSuffix(s, r.prefix), nil
}

func (r SuffixRestriction) apply(stack []Symbol, z Symbol) (bool, error) {
	var s string
	for i := 0; i < len(stack); i++ {
		s += stack[i].value
	}
	s += z.value
	//inverted
	return strings.HasPrefix(s, r.suffix), nil
}

func (r RegexRestriction) apply(stack []Symbol, z Symbol) (bool, error) {
	p, err := regexp.Compile(r.pattern)

	if err != nil {
		var s string
		for i := 0; i < len(stack); i++ {
			s += stack[i].value
		}
		s += z.value
		return p.MatchString(s), nil
	}

	return false, err
}
