package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddress(t *testing.T) {
	t.Run("Derive address from Public Key", func(t *testing.T) {
		privateKey := NewPrivateKey()
		publicKey := privateKey.GetPublicKey()

		address := publicKey.Address()
		assert.Equal(t, ADDRESS_LENGTH, len(address.Bytes()))
	})
}
