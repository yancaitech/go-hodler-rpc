package coinsrpc

import (
	"net/http"

	"github.com/bitcoinsv/bsvd/btcjson"
	hk "github.com/yancaitech/go-hodler-keys"
)

// BSVSignRawTxArgs struct
type BSVSignRawTxArgs struct {
	Entropy        string `json:"entropy"`
	Seed           string `json:"seed"`
	M1             uint32 `json:"m1"`
	M2             uint32 `json:"m2"`
	MainNet        bool   `json:"mainnet"`
	CompressPubKey bool   `json:"compresspubkey"`
	RawTx          string `json:"rawtx"`
}

// BSVSignRawTxReply struct
type BSVSignRawTxReply struct {
	Result string `json:"result"`
}

// BSVDecodeRawTxOutArgs struct
type BSVDecodeRawTxOutArgs struct {
	MainNet      bool   `json:"mainnet"`
	FromAddr     string `json:"fromAddr"`
	ToAddr       string `json:"toAddr"`
	TotalInValue int64  `json:"totalInValue"`
	RawTx        string `json:"rawtx"`
}

// BSVDecodeRawTxOutReply struct
type BSVDecodeRawTxOutReply struct {
	Txid    string `json:"txid"`
	Amount  int64  `json:"amount"`
	Fee     int64  `json:"fee"`
	Change  int64  `json:"change"`
	Raw     string `json:"raw"`
	Spendtx string `json:"spendtx"`
}

// BSVCreateRawTxArgs struct
type BSVCreateRawTxArgs struct {
	MainNet  bool                       `json:"mainnet"`
	Inputs   []btcjson.TransactionInput `json:"inputs"`
	Amounts  map[string]float64         `json:"amounts"` //`jsonrpcusage:"{\"address\":amount,...}"` // In BTC
	Sequence uint32                     `json:"sequence,omitempty"`
}

// BSVCreateRawTxReply struct
type BSVCreateRawTxReply struct {
	RawTx string `json:"rawtx"`
}

// BSVHashTxArgs struct
type BSVHashTxArgs struct {
	RawTx string `json:"rawtx"`
}

// BSVHashTxReply struct
type BSVHashTxReply struct {
	Txid string `json:"txid"`
}

// HashTX entry
func (h *BSV) HashTX(r *http.Request, args *BSVHashTxArgs, reply *BSVHashTxReply) (err error) {
	var key hk.Key
	txid, err := key.BitcoinSVTxid(args.RawTx)
	if err != nil {
		return err
	}
	reply.Txid = txid

	return nil
}

// CreateRawTx entry
func (h *BSV) CreateRawTx(r *http.Request, args *BSVCreateRawTxArgs, reply *BSVCreateRawTxReply) (err error) {
	var key hk.Key
	rawtx, err := key.BitcoinSVCreateRawTransaction(args.MainNet, args.Inputs, args.Amounts, args.Sequence)
	if err != nil {
		return err
	}
	reply.RawTx = rawtx

	return nil
}

// SignRawTx entry
func (h *BSV) SignRawTx(r *http.Request, args *BSVSignRawTxArgs, reply *BSVSignRawTxReply) (err error) {
	var key hk.Key
	wif, err := key.BitcoinWifFromEntropy(args.Entropy, args.Seed, args.M1, args.M2, args.MainNet, args.CompressPubKey)
	if err != nil {
		return err
	}
	reply.Result, err = key.BitcoinSVSignRawTx(wif, args.RawTx)
	if err != nil {
		return err
	}

	return nil
}

// DecodeRawTxOut entry
func (h *BSV) DecodeRawTxOut(r *http.Request, args *BSVDecodeRawTxOutArgs, reply *BSVDecodeRawTxOutReply) (err error) {
	var key hk.Key
	amount, fee, change, raw, spendtx, err := key.BitcoinSVDecodeRawTxOut(args.MainNet, args.FromAddr, args.ToAddr, args.TotalInValue, args.RawTx)
	if err != nil {
		return err
	}
	txid, err := key.BitcoinSVTxid(args.RawTx)
	if err != nil {
		return err
	}
	reply.Txid = txid
	reply.Amount = amount
	reply.Fee = fee
	reply.Change = change
	reply.Raw = raw
	reply.Spendtx = spendtx

	return nil
}
