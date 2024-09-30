package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrivateKey(t *testing.T) {
	t.Run("Private key generation", func(t *testing.T) {
		privateKey := NewPrivateKey()
		assert.Equal(t, len(privateKey.Bytes()), PRIVATE_KEY_LENGTH)

		publicKey := privateKey.GetPublicKey()
		assert.Equal(t, len(publicKey.Bytes()), PUBLIC_KEY_LENGTH)
	})

	t.Run("Private key signing", func(t *testing.T) {
		privateKey := NewPrivateKey()

		payload := []byte("HELLO WORLD")
		signature := privateKey.Sign(payload)

		publicKey := privateKey.GetPublicKey()
		signatureValid := signature.Verify(publicKey, payload)
		assert.Equal(t, signatureValid, true)
	})

	t.Run("Signature should be invalid if wrong payload is used", func(t *testing.T) {
		privateKey := NewPrivateKey()

		payload := []byte("HELLO WORLD")
		signature := privateKey.Sign(payload)

		publicKey := privateKey.GetPublicKey()
		signatureValid := signature.Verify(publicKey, []byte("BYE BYE WORLD"))
		assert.Equal(t, signatureValid, false)
	})

	t.Run("Signature should be invalid if wrong public key is used", func(t *testing.T) {
		privateKey := NewPrivateKey()

		payload := []byte("HELLO WORLD")
		signature := privateKey.Sign(payload)

		wrongPublicKey := NewPrivateKey().GetPublicKey()
		signatureValid := signature.Verify(wrongPublicKey, payload)
		assert.Equal(t, signatureValid, false)
	})
}
