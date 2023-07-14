package main

type Conf struct {
	Origins     []string     `json:"origins"`
	Constraints []Constraint `json:"constraints"`
}

type Constraint struct {
	KeyIds     []string   `json:"keyIds"`
	Curves     []string   `json:"curves"`
	Encodings  []string   `json:"encodings"`
	Validation Validation `json:"validation"`
}

type Validation struct {
	Prefixes []string `json:"prefixes"`
}
