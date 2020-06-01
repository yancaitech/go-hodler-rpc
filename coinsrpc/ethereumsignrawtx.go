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

// ETHSignRawTxArgs struct
type ETHSignRawTxArgs struct {
	Entropy   string `json:"entropy"`
	Seed      string `json:"seed"`
	M1        uint32 `json:"m1"`
	M2        uint32 `json:"m2"`
	Nonce     string `json:"nonce"`
	GasLimit  string `json:"gasLimit"`
	GasPrice  string `json:"gasPrice"`
	Value     string `json:"value"`
	ChainID   string `json:"chainID"`
	ToAddress string `json:"toAddress"`
}

// ETHSignRawTxReply struct
type ETHSignRawTxReply struct {
	Result string `json:"result"`
	Txid   string `json:"txid"`
}

// ETHHashTxArgs struct
type ETHHashTxArgs struct {
	RawTx string `json:"rawtx"`
}

// ETHHashTxReply struct
type ETHHashTxReply struct {
	Txid string `json:"txid"`
}

// ETHDecodeRawTxOutArgs struct
type ETHDecodeRawTxOutArgs struct {
	ChainID string `json:"chainID"`
	RawTx   string `json:"rawtx"`
}

// ETHDecodeRawTxOutReply struct
type ETHDecodeRawTxOutReply struct {
	hk.EthereumTx
}

// DecodeRawTxOut entry
func (h *ETH) DecodeRawTxOut(r *http.Request, args *ETHDecodeRawTxOutArgs, reply *ETHDecodeRawTxOutReply) (err error) {
	var chainID big.Int
	chainID.UnmarshalText([]byte(args.ChainID))

	var key hk.Key
	eset, err := key.EthereumDecodeRawTxOut(&chainID, args.RawTx)
	if err != nil {
		return err
	}
	reply.ChainID = eset.ChainID
	reply.Txid = eset.Txid
	reply.Nonce = eset.Nonce
	reply.GasLimit = eset.GasLimit
	reply.GasPrice = eset.GasPrice
	reply.FromAddress = eset.FromAddress
	reply.Recipient = eset.Recipient
	reply.Value = eset.Value
	reply.Payload = eset.Payload
	reply.V = eset.V
	reply.R = eset.R
	reply.S = eset.S
	return nil
}

// HashTX entry
func (h *ETH) HashTX(r *http.Request, args *ETHHashTxArgs, reply *ETHHashTxReply) (err error) {
	var key hk.Key
	txid, err := key.EthereumTxid(args.RawTx)
	if err != nil {
		return err
	}
	reply.Txid = txid
	return nil
}

// SignRawTx entry
func (h *ETH) SignRawTx(r *http.Request, args *ETHSignRawTxArgs, reply *ETHSignRawTxReply) (err error) {
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
	reply.Result, reply.Txid, err = key.EthereumSignRawTx(args.Entropy, args.Seed, args.M1, args.M2, nonce, gasLimit,
		&gasPrice, &value, &chainID, args.ToAddress)
	if err != nil {
		return err
	}
	return nil
}
