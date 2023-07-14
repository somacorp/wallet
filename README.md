# Wallet

Wallet is a program that enables digital signing and establishment of shared secrets using cryptographic keys. The application runs as a web server and does not leak managed private keys to any clients.

## Configuration and key management

As a user, you are required to create two files in the [`config`](config) directory:

- `conf.json` - Enforces a set of constraints on how signatures can be created.
- `keys.json` - Store private keys, each referenced by a memorable name.

Example files are provided in the [`config`](config) directory.

### About `conf.json`

See [`config/conf-example.json`](config/conf-example.json), which provides two different sets of constraints on how the private keys can be used.

It is recommended that a client-specific prefix constraint (`constraints.validation.prefixes`) is used in order to prevent a malicious client from creating unwanted signatures.

### About `keys.json`

See [`config/keys-example.json`](config/keys-example.json), which provides two different managed private keys under two different aliases.

## Disclaimer

This software aims to securely manage cryptographic private keys. The authors provide no guarantee that the software does not contain security flaws or implementation errors. Use at your own risk; don't trust, verify.
