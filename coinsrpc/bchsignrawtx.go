package coinsrpc

import (
	"net/http"

	"github.com/gcash/bchd/btcjson"
	hk "github.com/yancaitech/go-hodler-keys"
)

// BCHSignRawTxArgs struct
type BCHSignRawTxArgs struct {
	Entropy        string `json:"entropy"`
	Seed           string `json:"seed"`
	M1             uint32 `json:"m1"`
	M2             uint32 `json:"m2"`
	MainNet        bool   `json:"mainnet"`
	CompressPubKey bool   `json:"compresspubkey"`
	RawTx          string `json:"rawtx"`
}

// BCHSignRawTxReply struct
type BCHSignRawTxReply struct {
	Result string `json:"result"`
}

// BCHDecodeRawTxOutArgs struct
type BCHDecodeRawTxOutArgs struct {
	MainNet      bool   `json:"mainnet"`
	FromAddr     string `json:"fromAddr"`
	ToAddr       string `json:"toAddr"`
	TotalInValue int64  `json:"totalInValue"`
	RawTx        string `json:"rawtx"`
}

// BCHDecodeRawTxOutReply struct
type BCHDecodeRawTxOutReply struct {
	Txid    string `json:"txid"`
	Amount  int64  `json:"amount"`
	Fee     int64  `json:"fee"`
	Change  int64  `json:"change"`
	Raw     string `json:"raw"`
	Spendtx string `json:"spendtx"`
}

// BCHCreateRawTxArgs struct
type BCHCreateRawTxArgs struct {
	MainNet  bool                       `json:"mainnet"`
	Inputs   []btcjson.TransactionInput `json:"inputs"`
	Amounts  map[string]float64         `json:"amounts"` //`jsonrpcusage:"{\"address\":amount,...}"` // In BTC
	Sequence uint32                     `json:"sequence,omitempty"`
}

// BCHCreateRawTxReply struct
type BCHCreateRawTxReply struct {
	RawTx string `json:"rawtx"`
}

// BCHHashTxArgs struct
type BCHHashTxArgs struct {
	RawTx string `json:"rawtx"`
}

// BCHHashTxReply struct
type BCHHashTxReply struct {
	Txid string `json:"txid"`
}

// HashTX entry
func (h *BCH) HashTX(r *http.Request, args *BCHHashTxArgs, reply *BCHHashTxReply) (err error) {
	var key hk.Key
	txid, err := key.BitcoinCashTxid(args.RawTx)
	if err != nil {
		return err
	}
	reply.Txid = txid

	return nil
}

// CreateRawTx entry
func (h *BCH) CreateRawTx(r *http.Request, args *BCHCreateRawTxArgs, reply *BCHCreateRawTxReply) (err error) {
	var key hk.Key
	rawtx, err := key.BitcoinCashCreateRawTransaction(args.MainNet, args.Inputs, args.Amounts, args.Sequence)
	if err != nil {
		return err
	}
	reply.RawTx = rawtx

	return nil
}

// SignRawTx entry
func (h *BCH) SignRawTx(r *http.Request, args *BCHSignRawTxArgs, reply *BCHSignRawTxReply) (err error) {
	var key hk.Key
	wif, err := key.BitcoinWifFromEntropy(args.Entropy, args.Seed, args.M1, args.M2, args.MainNet, args.CompressPubKey)
	if err != nil {
		return err
	}
	reply.Result, err = key.BitcoinCashSignRawTx(wif, args.RawTx)
	if err != nil {
		return err
	}

	return nil
}

// DecodeRawTxOut entry
func (h *BCH) DecodeRawTxOut(r *http.Request, args *BCHDecodeRawTxOutArgs, reply *BCHDecodeRawTxOutReply) (err error) {
	var key hk.Key
	amount, fee, change, raw, spendtx, err := key.BitcoinCashDecodeRawTxOut(args.MainNet, args.FromAddr, args.ToAddr, args.TotalInValue, args.RawTx)
	if err != nil {
		return err
	}
	txid, err := key.BitcoinCashTxid(args.RawTx)
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
