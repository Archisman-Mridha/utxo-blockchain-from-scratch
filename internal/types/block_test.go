package types

import (
	"testing"

	"github.com/Archisman-Mridha/blockchain-from-scratch/pkg/crypto"
	"github.com/Archisman-Mridha/blockchain-from-scratch/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestBlock(t *testing.T) {
	t.Run("Block hash length should be 32 bytes", func(t *testing.T) {
		block := utils.RandomBlock()

		blockHash := GetBlockHash(block)
		assert.Equal(t, 32, len(blockHash))
	})

	t.Run("Block signature should be verifiable", func(t *testing.T) {
		var (
			block = utils.RandomBlock()

			privateKey = crypto.NewPrivateKey()
			publicKey  = privateKey.GetPublicKey()
		)

		blockSignature := SignBlock(privateKey, block)
		assert.Equal(t, 64, len(blockSignature.Bytes()))
		assert.True(t, blockSignature.Verify(publicKey, GetBlockHash(block)))
	})
}
