package peermgr

import (
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	peermgr "github.com/meshplus/pier/internal/peermgr/proto"
)

type MessageHandler func(network.Stream, *peermgr.Message)

//go:generate mockgen -destination mock_peermgr/mock_peermgr.go -package mock_peermgr -source peermgr.go
type PeerManager interface {
	// Start
	Start() error

	// Stop
	Stop() error

	// AsyncSend sends message to peer with peer info.
	AsyncSend(uint64, *peermgr.Message) error

	// SendWithStream sends message using existed stream
	SendWithStream(network.Stream, *peermgr.Message) error

	// Send sends message waiting response
	Send(uint64, *peermgr.Message) (*peermgr.Message, error)

	// Peers
	Peers() map[uint64]*peer.AddrInfo

	// RegisterMsgHandler
	RegisterMsgHandler(peermgr.Message_Type, MessageHandler) error
}
