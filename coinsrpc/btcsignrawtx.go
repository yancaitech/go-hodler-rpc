package coinsrpc

import (
	"net/http"

	"github.com/btcsuite/btcd/btcjson"
	hk "github.com/yancaitech/go-hodler-keys"
)

// BTCSignRawTxArgs struct
type BTCSignRawTxArgs struct {
	Entropy        string `json:"entropy"`
	Seed           string `json:"seed"`
	M1             uint32 `json:"m1"`
	M2             uint32 `json:"m2"`
	MainNet        bool   `json:"mainnet"`
	CompressPubKey bool   `json:"compresspubkey"`
	RawTx          string `json:"rawtx"`
}

// BTCSignRawTxReply struct
type BTCSignRawTxReply struct {
	Result string `json:"result"`
}

// BTCDecodeRawTxOutArgs struct
type BTCDecodeRawTxOutArgs struct {
	MainNet      bool   `json:"mainnet"`
	FromAddr     string `json:"fromAddr"`
	ToAddr       string `json:"toAddr"`
	TotalInValue int64  `json:"totalInValue"`
	RawTx        string `json:"rawtx"`
}

// BTCDecodeRawTxOutReply struct
type BTCDecodeRawTxOutReply struct {
	Txid    string `json:"txid"`
	Amount  int64  `json:"amount"`
	Fee     int64  `json:"fee"`
	Change  int64  `json:"change"`
	Raw     string `json:"raw"`
	Spendtx string `json:"spendtx"`
}

// BTCCreateRawTxArgs struct
type BTCCreateRawTxArgs struct {
	MainNet  bool                       `json:"mainnet"`
	Inputs   []btcjson.TransactionInput `json:"inputs"`
	Amounts  map[string]float64         `json:"amounts"` //`jsonrpcusage:"{\"address\":amount,...}"` // In BTC
	Sequence uint32                     `json:"sequence,omitempty"`
}

// BTCCreateRawTxReply struct
type BTCCreateRawTxReply struct {
	RawTx string `json:"rawtx"`
}

// BTCHashTxArgs struct
type BTCHashTxArgs struct {
	RawTx string `json:"rawtx"`
}

// BTCHashTxReply struct
type BTCHashTxReply struct {
	Txid string `json:"txid"`
}

// HashTX entry
func (h *BTC) HashTX(r *http.Request, args *BTCHashTxArgs, reply *BTCHashTxReply) (err error) {
	var key hk.Key
	txid, err := key.BitcoinTxid(args.RawTx)
	if err != nil {
		return err
	}
	reply.Txid = txid

	return nil
}

// CreateRawTx entry
func (h *BTC) CreateRawTx(r *http.Request, args *BTCCreateRawTxArgs, reply *BTCCreateRawTxReply) (err error) {
	var key hk.Key
	rawtx, err := key.BitcoinCreateRawTransaction(args.MainNet, args.Inputs, args.Amounts, args.Sequence)
	if err != nil {
		return err
	}
	reply.RawTx = rawtx

	return nil
}

// SignRawTx entry
func (h *BTC) SignRawTx(r *http.Request, args *BTCSignRawTxArgs, reply *BTCSignRawTxReply) (err error) {
	var key hk.Key
	wif, err := key.BitcoinWifFromEntropy(args.Entropy, args.Seed, args.M1, args.M2, args.MainNet, args.CompressPubKey)
	if err != nil {
		return err
	}
	reply.Result, err = key.BitcoinSignRawTx(wif, args.RawTx)
	if err != nil {
		return err
	}

	return nil
}

// DecodeRawTxOut entry
func (h *BTC) DecodeRawTxOut(r *http.Request, args *BTCDecodeRawTxOutArgs, reply *BTCDecodeRawTxOutReply) (err error) {
	var key hk.Key
	amount, fee, change, raw, spendtx, err := key.BitcoinDecodeRawTxOut(args.MainNet, args.FromAddr, args.ToAddr, args.TotalInValue, args.RawTx)
	if err != nil {
		return err
	}
	txid, err := key.BitcoinTxid(args.RawTx)
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
