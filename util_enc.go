package main

import (
	"fmt"

	"github.com/somacorp/core/crypto/enc"
)

func useEncoder(encoding string) (enc.Encoder, error) {
	switch encoding {
	case "eth":
		return enc.EthEncoder(), nil
	case "sha256":
		return enc.SHA256Encoder(), nil
	case "trivial":
		return enc.TrivialEncoder(), nil
	default:
		return nil, fmt.Errorf("invalid encoding '%s'", encoding)
	}
}

func encode(encoding string, data []byte) ([]byte, error) {
	e, err := useEncoder(encoding)
	if err != nil {
		return nil, fmt.Errorf("failed to get encoder: %v", err)
	}
	return e.Encode(data)
}
