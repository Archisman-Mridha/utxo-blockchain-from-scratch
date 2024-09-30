package types

import (
	"testing"

	"github.com/Archisman-Mridha/blockchain-from-scratch/api/grpc/proto/generated"
	"github.com/Archisman-Mridha/blockchain-from-scratch/pkg/crypto"
	"github.com/Archisman-Mridha/blockchain-from-scratch/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T) {
	var (
		privateKeyAlice = crypto.NewPrivateKey()
		publicKeyAlice  = privateKeyAlice.GetPublicKey()

		privateKeyBob = crypto.NewPrivateKey()
		publicKeyBob  = privateKeyBob.GetPublicKey()

		transactionInput = &generated.TransactionInput{
			PreviousTransactionHash:        utils.RandomHash(),
			PreviousTransactionOutputIndex: 0,
			PublicKey:                      privateKeyAlice.GetPublicKey().Bytes(),
		}

		transactionOutput_AliceToBob = &generated.TransactionOutput{
			Amount:    5,
			ToAddress: publicKeyBob.Address().Bytes(),
		}
		transactionOutput_AliceToAlice = &generated.TransactionOutput{
			Amount:    95,
			ToAddress: publicKeyAlice.Address().Bytes(),
		}

		// Represents : Alice sending 5 units to Bob.
		transaction = &generated.Transaction{
			Version: 1,
			Inputs:  []*generated.TransactionInput{transactionInput},
			Outputs: []*generated.TransactionOutput{
				transactionOutput_AliceToBob,
				transactionOutput_AliceToAlice,
			},
		}
	)

	transactionSignature := SignTransaction(privateKeyAlice, transaction)
	transactionInput.Signature = transactionSignature.Bytes()

	t.Run("Transaction verification should be successfull", func(t *testing.T) {
		assert.True(t, VerifyTransaction(transaction))
	})
}
