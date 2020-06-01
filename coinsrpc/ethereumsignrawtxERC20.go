package coinsrpc

import (
	"math/big"
	"net/http"
	"strconv"

	hk "github.com/yancaitech/go-hodler-keys"
)

// ChainID:
// mainnet: 1
// ropsten: 3
// rinkeby: 4
// goerli : 5
// kovan  : 42

// ERC20SignRawTxArgs struct
type ERC20SignRawTxArgs struct {
	Entropy   string `json:"entropy"`
	Seed      string `json:"seed"`
	M1        uint32 `json:"m1"`
	M2        uint32 `json:"m2"`
	Nonce     string `json:"nonce"`
	GasLimit  string `json:"gasLimit"`
	GasPrice  string `json:"gasPrice"`
	Value     string `json:"value"`
	ChainID   string `json:"chainID"`
	Contract  string `json:"contract"`
	ToAddress string `json:"toAddress"`
}

// ERC20SignRawTxReply struct
type ERC20SignRawTxReply struct {
	Result string `json:"result"`
	Txid   string `json:"txid"`
}

// SignRawTxERC20 entry
func (h *ETH) SignRawTxERC20(r *http.Request, args *ERC20SignRawTxArgs, reply *ERC20SignRawTxReply) (err error) {
	nonce, err := strconv.ParseUint(args.Nonce, 10, 64)
	if err != nil {
		return err
	}
	gasLimit, err := strconv.ParseUint(args.GasLimit, 10, 64)
	if err != nil {
		return err
	}
	var gasPrice big.Int
	gasPrice.UnmarshalText([]byte(args.GasPrice))
	var value big.Int
	value.UnmarshalText([]byte(args.Value))
	var chainID big.Int
	chainID.UnmarshalText([]byte(args.ChainID))

	var key hk.Key
	reply.Result, reply.Txid, err = key.EthereumSignRawTxERC20(args.Entropy, args.Seed, args.M1, args.M2, nonce, gasLimit,
		&gasPrice, &value, &chainID, args.Contract, args.ToAddress)
	if err != nil {
		return err
	}

	return nil
}
