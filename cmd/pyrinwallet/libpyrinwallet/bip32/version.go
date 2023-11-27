package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// PyrinMainnetPrivate is the version that is used for
// pyrin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var PyrinMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// PyrinMainnetPublic is the version that is used for
// pyrin mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var PyrinMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// PyrinTestnetPrivate is the version that is used for
// pyrin testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var PyrinTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// PyrinTestnetPublic is the version that is used for
// pyrin testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var PyrinTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// PyipadevnetPrivate is the version that is used for
// pyrin devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var PyipadevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// PyipadevnetPublic is the version that is used for
// pyrin devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var PyipadevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// PyrinSimnetPrivate is the version that is used for
// pyrin simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var PyrinSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// PyrinSimnetPublic is the version that is used for
// pyrin simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var PyrinSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case PyrinMainnetPrivate:
		return PyrinMainnetPublic, nil
	case PyrinTestnetPrivate:
		return PyrinTestnetPublic, nil
	case PyipadevnetPrivate:
		return PyipadevnetPublic, nil
	case PyrinSimnetPrivate:
		return PyrinSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case PyrinMainnetPrivate:
		return true
	case PyrinTestnetPrivate:
		return true
	case PyipadevnetPrivate:
		return true
	case PyrinSimnetPrivate:
		return true
	}

	return false
}
