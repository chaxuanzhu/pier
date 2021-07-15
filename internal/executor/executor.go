package executor

import (
	"context"
	"fmt"

	"github.com/meshplus/bitxhub-kit/storage"
	"github.com/meshplus/bitxhub-kit/types"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/meshplus/pier/internal/txcrypto"
	"github.com/meshplus/pier/pkg/plugins"
	"github.com/sirupsen/logrus"
)

var _ Executor = (*ChannelExecutor)(nil)

// ChannelExecutor represents the necessary data for executing interchain txs in appchain
type ChannelExecutor struct {
	client   plugins.Client // the client to interact with appchain
	storage  storage.Storage
	id       string // appchain id
	cryptor  txcrypto.Cryptor
	logger   logrus.FieldLogger
	ctx      context.Context
	cancel   context.CancelFunc
	checkSum bool
}

// New creates new instance of Executor. agent is for interacting with counterpart chain
// client is for interacting with appchain, meta is for recording interchain tx meta information
// and ds is for persisting some runtime messages
func New(client plugins.Client, pierID string, storage storage.Storage, cryptor txcrypto.Cryptor, logger logrus.FieldLogger, checkSum bool) (*ChannelExecutor, error) {
	ctx, cancel := context.WithCancel(context.Background())

	return &ChannelExecutor{
		client:   client,
		ctx:      ctx,
		cancel:   cancel,
		storage:  storage,
		id:       pierID,
		cryptor:  cryptor,
		logger:   logger,
		checkSum: checkSum,
	}, nil
}

// Start implements Executor
func (e *ChannelExecutor) Start() error {
	e.logger.Info("Executor started")
	return nil
}

// Stop implements Executor
func (e *ChannelExecutor) Stop() error {
	e.cancel()

	e.logger.Info("Executor stopped")
	return nil
}

func (e *ChannelExecutor) QueryInterchainMeta() map[string]uint64 {
	execMeta, err := e.client.GetInMeta()
	if err != nil {
		return map[string]uint64{}
	}
	if !e.checkSum {
		return execMeta
	}
	checkSumMeta := make(map[string]uint64, len(execMeta))
	for from, index := range execMeta {
		checkSumMeta[types.NewAddressByStr(from).String()] = index
	}
	return checkSumMeta
}

func (e *ChannelExecutor) QueryCallbackMeta() map[string]uint64 {
	callbackMeta, err := e.client.GetCallbackMeta()
	if err != nil {
		return map[string]uint64{}
	}
	if !e.checkSum {
		return callbackMeta
	}
	checkSumMeta := make(map[string]uint64, len(callbackMeta))
	for from, index := range callbackMeta {
		checkSumMeta[types.NewAddressByStr(from).String()] = index
	}
	return checkSumMeta
}

// getReceipt only generates one receipt given source chain id and interchain tx index
func (e *ChannelExecutor) QueryIBTPReceipt(originalIBTP *pb.IBTP) (*pb.IBTP, error) {
	if originalIBTP == nil {
		return nil, fmt.Errorf("empty original ibtp")
	}
	return e.client.GetReceipt(originalIBTP)
}
