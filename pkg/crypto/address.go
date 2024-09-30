package crypto

import "encoding/hex"

const ADDRESS_LENGTH = 20

type Address struct {
	value []byte
}

// Returns the address associated with the public key.
// An address is an alternate representation of someone's public key. The first 20 bytes of the
// public key, will be the address of the public key's owner.
func (p *PublicKey) Address() *Address {
	return &Address{
		value: p.key[len(p.key)-ADDRESS_LENGTH:],
	}
}

func (a *Address) Bytes() []byte {
	return a.value
}

func (a Address) String() string {
	return hex.EncodeToString(a.value)
}
