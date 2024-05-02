package server

import (
	"errors"
	"context"

	bftbytes "github.com/cometbft/cometbft/libs/bytes"
	bftclient "github.com/cometbft/cometbft/rpc/client"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ethereum/go-ethereum/rpc"
	bfttypes "github.com/cometbft/cometbft/types"
)

var errNotImplemented = errors.New("not implemented")

type MonomerClient struct {
	client *rpc.Client
}

var _ client.TendermintRPC = (*MonomerClient)(nil)

func NewMonomerClient(client *rpc.Client) *MonomerClient {
	return &MonomerClient{
		client: client,
	}
}

func (m *MonomerClient) ABCIInfo(ctx context.Context) (*ctypes.ResultABCIInfo, error) {
	result := new(ctypes.ResultABCIInfo)
	if err := m.client.CallContext(ctx, &result, "abci_info"); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *MonomerClient) ABCIQuery(ctx context.Context, path string, data bftbytes.HexBytes) (*ctypes.ResultABCIQuery, error) {
	result := new(ctypes.ResultABCIQuery)
	if err := m.client.CallContext(ctx, &result, "abci_query", path, data); err != nil {
		return nil, err
	}
	return result, nil
}

func (*MonomerClient) ABCIQueryWithOptions(
	ctx context.Context,
	path string,
	data bftbytes.HexBytes,
	opts bftclient.ABCIQueryOptions,
) (*ctypes.ResultABCIQuery, error) {
	return nil, errNotImplemented
}

func (m *MonomerClient) BroadcastTxCommit(ctx context.Context, tx bfttypes.Tx) (*ctypes.ResultBroadcastTxCommit, error) {
	result := new(ctypes.ResultBroadcastTxCommit)
	if err := m.client.CallContext(ctx, &result, "broadcast_tx_commit", tx); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *MonomerClient) BroadcastTxAsync(ctx context.Context, tx bfttypes.Tx) (*ctypes.ResultBroadcastTx, error) {
	result := new(ctypes.ResultBroadcastTx)
	if err := m.client.CallContext(ctx, &result, "broadcast_tx_async", tx); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *MonomerClient) BroadcastTxSync(ctx context.Context, tx bfttypes.Tx) (*ctypes.ResultBroadcastTx, error) {
	result := new(ctypes.ResultBroadcastTx)
	if err := m.client.CallContext(ctx, &result, "broadcast_tx_sync", tx); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *MonomerClient) Validators(ctx context.Context, height *int64, page, perPage *int) (*ctypes.ResultValidators, error) {
	result := new(ctypes.ResultValidators)
	if err := m.client.CallContext(ctx, &result, "validators", height, page, perPage); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *MonomerClient) Status(ctx context.Context) (*ctypes.ResultStatus, error) {
	result := new(ctypes.ResultStatus)
	if err := m.client.CallContext(ctx, &result, "status"); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *MonomerClient) Block(ctx context.Context, height *int64) (*ctypes.ResultBlock, error) {
	result := new(ctypes.ResultBlock)
	if err := m.client.CallContext(ctx, &result, "block", height); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *MonomerClient) BlockchainInfo(ctx context.Context, minHeight, maxHeight int64) (*ctypes.ResultBlockchainInfo, error) {
	result := new(ctypes.ResultBlockchainInfo)
	if err := m.client.CallContext(ctx, &result, "blockchain", minHeight, maxHeight); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *MonomerClient) Commit(ctx context.Context, height *int64) (*ctypes.ResultCommit, error) {
	result := new(ctypes.ResultCommit)
	if err := m.client.CallContext(ctx, &result, "commit", height); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *MonomerClient) Tx(ctx context.Context, hash []byte, prove bool) (*ctypes.ResultTx, error) {
	result := new(ctypes.ResultTx)
	if err := m.client.CallContext(ctx, &result, "tx", hash, prove); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *MonomerClient) TxSearch(ctx context.Context, query string, prove bool, page, perPage *int, orderBy string) (*ctypes.ResultTxSearch, error) {
	result := new(ctypes.ResultTxSearch)
	if err := m.client.CallContext(ctx, &result, "tx_search", query, prove, page, perPage, orderBy); err != nil {
		return nil, err
	}
	return result, nil
}
