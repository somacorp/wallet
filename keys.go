package main

import (
	"encoding/hex"
	"fmt"
)

type Keys struct {
	privById map[string]string
}

func (k *Keys) getPrivBytes(id string) ([]byte, error) {
	privHex, ok := k.privById[id]
	if !ok {
		return nil, fmt.Errorf("key not found for %v", id)
	}
	privBytes, err := hex.DecodeString(privHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode priv: %v", err)
	}
	return privBytes, nil
}
