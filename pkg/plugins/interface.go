package plugins

import (
	"github.com/meshplus/bitxhub-model/pb"
)

// Client defines the interface that interacts with appchain
//go:generate mockgen -destination mock_client/mock_client.go -package mock_client -source interface.go
type Client interface {
	// Initialize initialize plugin client
	Initialize(configPath string, pierID string, extra []byte) error

	// Start starts to listen appchain event
	Start() error

	// Stop stops client
	Stop() error

	// GetIBTP gets an interchain ibtp channel generated by client
	GetIBTP() chan *pb.IBTP

	// SubmitIBTP submits the interchain ibtp to appchain
	SubmitIBTP(*pb.IBTP) (*pb.SubmitIBTPResponse, error)

	// RollbackIBTP rollbacks the interchain ibtp to appchain
	RollbackIBTP(ibtp *pb.IBTP, isSrcChain bool) (*pb.RollbackIBTPResponse, error)

	// increase in meta without actually executing it
	IncreaseInMeta(ibtp *pb.IBTP) (*pb.IBTP, error)

	// GetUpdateMeta gets an to-update meta channel generated by client
	GetUpdateMeta() <-chan *pb.UpdateMeta

	// GetMintEvent gets an mint event from appchain
	GetLockEvent() <-chan *pb.LockEvent

	// Unescrow submits an unescrow event found in bitxhub to appchain
	Unescrow(unlock *pb.UnLock) error

	// GetOutMessage gets interchain ibtp by index and target chain_id from broker contract
	GetOutMessage(to string, idx uint64) (*pb.IBTP, error)

	// GetInMessage gets receipt by index and source chain_id
	GetInMessage(from string, idx uint64) ([][]byte, error)

	// GetOutMeta gets an index map, which implicates the greatest index of
	// ingoing interchain txs for each source chain
	GetInMeta() (map[string]uint64, error)

	// GetOutMeta gets an index map, which implicates the greatest index of
	// outgoing interchain txs for each receiving chain
	GetOutMeta() (map[string]uint64, error)

	// GetReceiptMeta gets an index map, which implicates the greatest index of
	// executed callback txs for each receiving chain
	GetCallbackMeta() (map[string]uint64, error)

	// CommitCallback is a callback function when get receipt from bitxhub success
	CommitCallback(ibtp *pb.IBTP) error

	// GetReceipt gets receipt of an executed IBTP
	GetReceipt(ibtp *pb.IBTP) (*pb.IBTP, error)

	// Name gets name of blockchain from plugin
	Name() string

	// Type gets type of blockchain from plugin
	Type() string

	// QueryFilterLockStart query the lasted lock event height in appchain
	QueryFilterLockStart(appchainIndex uint64) (uint64, error)

	// QueryLockEventByIndex query A lock event by lock index
	QueryLockEventByIndex(index uint64) (*pb.LockEvent, error)

	// QueryAppchainIndex query the lasted lock event index from appchain
	QueryAppchainIndex() (uint64, error)

	// QueryRelayIndex query the lasted burn event index from appchain
	QueryRelayIndex() (uint64, error)
}
