module github.com/kattana-io/models

go 1.20

require (
	github.com/ethereum/go-ethereum v1.13.14
	github.com/shopspring/decimal v1.3.1
	google.golang.org/protobuf v1.33.0
)

require (
	github.com/bits-and-blooms/bitset v1.13.0 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.2 // indirect
	github.com/consensys/bavard v0.1.13 // indirect
	github.com/consensys/gnark-crypto v0.12.1 // indirect
	github.com/crate-crypto/go-kzg-4844 v0.7.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0 // indirect
	github.com/ethereum/c-kzg-4844 v0.4.0 // indirect
	github.com/holiman/uint256 v1.2.4 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/supranational/blst v0.3.11 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/exp v0.0.0-20240325151524-a685a6edb6d8 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)

// fallback compability
// see: https://github.com/crate-crypto/go-kzg-4844/issues/64
replace github.com/crate-crypto/go-kzg-4844 v1.0.0 => github.com/crate-crypto/go-kzg-4844 v0.7.0
