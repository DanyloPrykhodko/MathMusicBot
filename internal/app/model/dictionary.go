package model

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrorCantParseDictionary = errors.New("can't parse dictionary")
	ErrorCantParseKey        = errors.New("can't parse key")
	ErrorCantParseValue      = errors.New("can't parse value")
)

type Dictionary struct {
	Key   string
	Value string
}

func (d *Dictionary) Parse(s string) (err error) {
	split := regexp.MustCompile(`(\n|\s?[=:\-]\s?)`).Split(s, -1)
	if len(split) < 2 {
		return ErrorCantParseDictionary
	}

	var key, value string

	key = strings.TrimSpace(split[0])
	if key == "" {
		return ErrorCantParseKey
	}

	value = s[strings.Index(s, split[1]):]
	if value == "" {
		return ErrorCantParseValue
	}

	d.Key = key
	d.Value = value

	return
}
