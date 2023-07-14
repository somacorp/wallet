package main

import "fmt"

func validate(v Validation, data []byte) error {
	s := string(data)
	if !hasPrefixFrom(s, v.Prefixes) {
		return fmt.Errorf("forbidden prefix '%s' - not in %s", s, v.Prefixes)
	}
	return nil
}
