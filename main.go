package LRparse

import (
	"errors"
	"regexp"
)

func GetFirstMatch(s, left, right string) (string, error) {
	pattern := regexp.MustCompile(regexp.QuoteMeta(left) + "(.*?)" + regexp.QuoteMeta(right))
	match := pattern.FindStringSubmatch(s)

	if len(match) < 2 {
		return "", errors.New("no match found")
	}

	return match[1], nil
}

func MatchRecursive(s, left, right string) ([]string, error) {
	pattern := regexp.MustCompile(regexp.QuoteMeta(left) + "(.*?)" + regexp.QuoteMeta(right))
	matches := pattern.FindAllStringSubmatch(s, -1)

	if len(matches) == 0 {
		return nil, errors.New("no matches found")
	}

	results := make([]string, len(matches))
	for i, match := range matches {
		results[i] = match[1]
	}

	return results, nil
}
