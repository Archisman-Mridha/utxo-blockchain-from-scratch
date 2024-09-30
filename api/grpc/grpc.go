package grpc

import (
	"context"
	"log/slog"
	"sync"

	"github.com/Archisman-Mridha/blockchain-from-scratch/api/grpc/proto/generated"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Node struct {
	generated.UnimplementedNodeServer
	peers     map[]bool
	peersLock sync.RWMutex
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) addPeer() {}

func (n *Node) Handshake(ctx context.Context, nodeDetails *generated.NodeDetails) (*generated.NodeDetails, error) {
	peer, _ := peer.FromContext(ctx)
	slog.Info("Received handshake request", slog.String("peer", peer.Addr.String()))

	return &generated.NodeDetails{}, nil
}

func (n *Node) HandleTransaction(ctx context.Context, transaction *generated.Transaction) (*emptypb.Empty, error) {
	peer, _ := peer.FromContext(ctx)
	slog.Info("Received transaction", slog.String("peer", peer.Addr.String()))

	return &emptypb.Empty{}, nil
}
