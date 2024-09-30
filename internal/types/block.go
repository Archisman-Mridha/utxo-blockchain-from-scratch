package types

import (
	"crypto/sha256"
	"log"

	"github.com/Archisman-Mridha/blockchain-from-scratch/api/grpc/proto/generated"
	"github.com/Archisman-Mridha/blockchain-from-scratch/pkg/crypto"
	"google.golang.org/protobuf/proto"
)

// Block hash is the SHA256 hash of the block header.
func GetBlockHash(block *generated.Block) []byte {
	blockAsBytes, err := proto.Marshal(block)
	if err != nil {
		log.Fatalf("Failed marshalling Block : %v", err)
	}

	blockHash := sha256.Sum256(blockAsBytes)
	return blockHash[:]
}

// Before a validator commits a block, it needs to sign the block.
// Everyone else, later, can verify that, that block was committed by that validator.
func SignBlock(privateKey *crypto.PrivateKey, block *generated.Block) *crypto.Signature {
	return privateKey.Sign(GetBlockHash(block))
}
