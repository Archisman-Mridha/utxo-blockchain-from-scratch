package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"io"
	"log"
)

const (
	PRIVATE_KEY_LENGTH = 64
	PUBLIC_KEY_LENGTH  = 32
	SIGNATURE_LENGTH   = 64

	SEED_LENGTH = 32
)

type (
	PrivateKey struct {
		key ed25519.PrivateKey
	}

	PublicKey struct {
		key ed25519.PublicKey
	}

	Signature struct {
		value []byte
	}
)

func NewPrivateKey() *PrivateKey {
	seed := make([]byte, SEED_LENGTH)
	if _, err := io.ReadFull(rand.Reader, seed); err != nil {
		log.Fatalf("Failed getting seed from the random number generator : %v", err)
	}

	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(seed),
	}
}

func (p *PrivateKey) Bytes() []byte {
	return p.key
}

func (p *PrivateKey) Sign(payload []byte) *Signature {
	return &Signature{
		value: ed25519.Sign(p.key, payload),
	}
}

func (p *PrivateKey) GetPublicKey() *PublicKey {
	// NOTE : The last 32 bytes of the private key, represents the public key.
	key := make([]byte, PUBLIC_KEY_LENGTH)
	copy(key, p.key[32:])

	return &PublicKey{key}
}

func (p *PublicKey) Bytes() []byte {
	return p.key
}

func PublicKeyFromBytes(bytes []byte) *PublicKey {
	if len(bytes) != PUBLIC_KEY_LENGTH {
		log.Fatalf("Public key byte length isn't %d", PUBLIC_KEY_LENGTH)
	}
	return &PublicKey{
		key: ed25519.PublicKey(bytes),
	}
}

func (s *Signature) Bytes() []byte {
	return s.value
}

func SignatureFromBytes(bytes []byte) *Signature {
	if len(bytes) != SIGNATURE_LENGTH {
		log.Fatalf("Signature byte length isn't %d", SIGNATURE_LENGTH)
	}
	return &Signature{
		value: bytes,
	}
}

func (s *Signature) Verify(publicKey *PublicKey, payload []byte) bool {
	return ed25519.Verify(publicKey.key, payload, s.Bytes())
}
