package utils

import (
	cryptoRand "crypto/rand"
	"io"
	"math/rand"
	"time"

	"github.com/Archisman-Mridha/blockchain-from-scratch/api/grpc/proto/generated"
)

func RandomBlock() *generated.Block {
	blockHeader := &generated.BlockHeader{
		Version:                        1,
		Height:                         int32(rand.Intn(1000)),
		PreviousBlockHash:              RandomHash(),
		TransactionsMerkleTreeRootHash: RandomHash(),
		Timestamp:                      time.Now().UnixNano(),
	}

	return &generated.Block{
		Header: blockHeader,
	}
}

func RandomHash() []byte {
	hash := make([]byte, 32)
	io.ReadFull(cryptoRand.Reader, hash)
	return hash
}
