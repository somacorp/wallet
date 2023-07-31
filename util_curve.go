package main

import (
	"fmt"

	"github.com/somacorp/core/crypto/crv"
)

func useCurve(curve string) (crv.Curve, error) {
	switch curve {
	case "Curve25519":
		return crv.Curve25519(), nil
	case "secp256k1":
		return crv.Secp256k1(), nil
	default:
		return nil, fmt.Errorf("invalid curve '%s'", curve)
	}
}

func createPublicKey(curve string, privBytes []byte) ([]byte, error) {
	c, err := useCurve(curve)
	if err != nil {
		return nil, fmt.Errorf("invalid curve '%s'", curve)
	}
	return c.PublicKey(privBytes)
}

func createECDHSharedSecret(curve string, privBytes, pubBytes []byte) ([]byte, error) {
	c, err := useCurve(curve)
	if err != nil {
		return nil, fmt.Errorf("invalid curve '%s'", curve)
	}
	return c.ECDH(privBytes, pubBytes)
}

func createSignature(curve string, data, privBytes []byte) ([]byte, error) {
	c, err := useCurve(curve)
	if err != nil {
		return nil, fmt.Errorf("invalid curve '%s'", curve)
	}
	return c.Sign(data, privBytes)
}

func verifySignature(curve string, pubBytes, data, sig []byte) (bool, error) {
	c, err := useCurve(curve)
	if err != nil {
		return false, fmt.Errorf("invalid curve '%s'", curve)
	}
	return c.Verify(pubBytes, data, sig), nil
}
