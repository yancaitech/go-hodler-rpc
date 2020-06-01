package coinsrpc

import (
	"net/http"

	"github.com/ltcsuite/ltcd/btcjson"
	hk "github.com/yancaitech/go-hodler-keys"
)

// LTCSignRawTxArgs struct
type LTCSignRawTxArgs struct {
	Entropy        string `json:"entropy"`
	Seed           string `json:"seed"`
	M1             uint32 `json:"m1"`
	M2             uint32 `json:"m2"`
	MainNet        bool   `json:"mainnet"`
	CompressPubKey bool   `json:"compresspubkey"`
	RawTx          string `json:"rawtx"`
}

// LTCSignRawTxReply struct
type LTCSignRawTxReply struct {
	Result string `json:"result"`
}

// LTCDecodeRawTxOutArgs struct
type LTCDecodeRawTxOutArgs struct {
	MainNet      bool   `json:"mainnet"`
	FromAddr     string `json:"fromAddr"`
	ToAddr       string `json:"toAddr"`
	TotalInValue int64  `json:"totalInValue"`
	RawTx        string `json:"rawtx"`
}

// LTCDecodeRawTxOutReply struct
type LTCDecodeRawTxOutReply struct {
	Txid    string `json:"txid"`
	Amount  int64  `json:"amount"`
	Fee     int64  `json:"fee"`
	Change  int64  `json:"change"`
	Raw     string `json:"raw"`
	Spendtx string `json:"spendtx"`
}

// LTCCreateRawTxArgs struct
type LTCCreateRawTxArgs struct {
	MainNet  bool                       `json:"mainnet"`
	Inputs   []btcjson.TransactionInput `json:"inputs"`
	Amounts  map[string]float64         `json:"amounts"` //`jsonrpcusage:"{\"address\":amount,...}"` // In BTC
	Sequence uint32                     `json:"sequence,omitempty"`
}

// LTCCreateRawTxReply struct
type LTCCreateRawTxReply struct {
	RawTx string `json:"rawtx"`
}

// LTCHashTxArgs struct
type LTCHashTxArgs struct {
	RawTx string `json:"rawtx"`
}

// LTCHashTxReply struct
type LTCHashTxReply struct {
	Txid string `json:"txid"`
}

// HashTX entry
func (h *LTC) HashTX(r *http.Request, args *LTCHashTxArgs, reply *LTCHashTxReply) (err error) {
	var key hk.Key
	txid, err := key.LitecoinTxid(args.RawTx)
	if err != nil {
		return err
	}
	reply.Txid = txid

	return nil
}

// CreateRawTx entry
func (h *LTC) CreateRawTx(r *http.Request, args *LTCCreateRawTxArgs, reply *LTCCreateRawTxReply) (err error) {
	var key hk.Key
	rawtx, err := key.LitecoinCreateRawTransaction(args.MainNet, args.Inputs, args.Amounts, args.Sequence)
	if err != nil {
		return err
	}
	reply.RawTx = rawtx

	return nil
}

// SignRawTx entry
func (h *LTC) SignRawTx(r *http.Request, args *LTCSignRawTxArgs, reply *LTCSignRawTxReply) (err error) {
	var key hk.Key
	wif, err := key.BitcoinWifFromEntropy(args.Entropy, args.Seed, args.M1, args.M2, args.MainNet, args.CompressPubKey)
	if err != nil {
		return err
	}
	reply.Result, err = key.LitecoinSignRawTx(wif, args.RawTx)
	if err != nil {
		return err
	}

	return nil
}

// DecodeRawTxOut entry
func (h *LTC) DecodeRawTxOut(r *http.Request, args *LTCDecodeRawTxOutArgs, reply *LTCDecodeRawTxOutReply) (err error) {
	var key hk.Key
	amount, fee, change, raw, spendtx, err := key.LitecoinDecodeRawTxOut(args.MainNet, args.FromAddr, args.ToAddr, args.TotalInValue, args.RawTx)
	if err != nil {
		return err
	}
	txid, err := key.LitecoinTxid(args.RawTx)
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
