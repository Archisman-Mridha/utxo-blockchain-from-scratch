package types

import (
	"crypto/sha256"
	"log"

	"github.com/Archisman-Mridha/blockchain-from-scratch/api/grpc/proto/generated"
	"github.com/Archisman-Mridha/blockchain-from-scratch/pkg/crypto"
	"google.golang.org/protobuf/proto"
)

func GetTransactionHash(transaction *generated.Transaction) []byte {
	transactionAsBytes, err := proto.Marshal(transaction)
	if err != nil {
		log.Fatalf("Failed marshalling Block : %v", err)
	}

	transactionHash := sha256.Sum256(transactionAsBytes)
	return transactionHash[:]
}

func SignTransaction(privateKey *crypto.PrivateKey, transaction *generated.Transaction) *crypto.Signature {
	return privateKey.Sign(GetTransactionHash(transaction))
}

func VerifyTransaction(transaction *generated.Transaction) bool {
	for _, transactionInput := range transaction.Inputs {
		transactionInputSignature := crypto.SignatureFromBytes(transactionInput.Signature)

		// When the transaction input was signed, transactionInput.Signature was null.
		transactionInput.Signature = nil
		defer func() {
			transactionInput.Signature = transactionInputSignature.Bytes()
		}()

		transactionSenderPublicKey := crypto.PublicKeyFromBytes(transactionInput.PublicKey)
		if !transactionInputSignature.Verify(transactionSenderPublicKey, GetTransactionHash(transaction)) {
			return false
		}
	}
	return true
}
