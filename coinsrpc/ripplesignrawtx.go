package coinsrpc

import (
	"net/http"

	hk "github.com/yancaitech/go-hodler-keys"
)

// XRPSignRawTxArgs struct
type XRPSignRawTxArgs struct {
	Entropy        string `json:"entropy"`
	Seed           string `json:"seed"`
	M1             uint32 `json:"m1"`
	M2             uint32 `json:"m2"`
	Sequence       uint32 `json:"sequence"`
	LedgerSequence uint32 `json:"ledgerSequence"`
	ToAddr         string `json:"toAddr"`
	Currency       string `json:"currency"`
	Value          string `json:"value"`
	Fee            string `json:"fee"`
	Tag            uint32 `json:"tag"`
}

// XRPSignRawTxReply struct
type XRPSignRawTxReply struct {
	Txid   string `json:"txid"`
	Result string `json:"result"`
}

// XRPDecodeRawTxOutArgs struct
type XRPDecodeRawTxOutArgs struct {
	RawTx string `json:"rawtx"`
}

// XRPDecodeRawTxOutReply struct
type XRPDecodeRawTxOutReply struct {
	hk.RippleTx
}

// DecodeRawTxOut entry
func (h *XRP) DecodeRawTxOut(r *http.Request, args *XRPDecodeRawTxOutArgs, reply *XRPDecodeRawTxOutReply) (err error) {
	var key hk.Key
	rset, err := key.RippleDecodeRawTxOut(args.RawTx)
	if err != nil {
		return err
	}
	reply.Txid = rset.Txid
	reply.Sequence = rset.Sequence
	reply.LedgerSequence = rset.LedgerSequence
	reply.Amount = rset.Amount
	reply.Fee = rset.Fee
	reply.FromAddress = rset.FromAddress
	reply.ToAddress = rset.ToAddress
	reply.Tag = rset.Tag
	return nil
}

// SignRawTx entry
func (h *XRP) SignRawTx(r *http.Request, args *XRPSignRawTxArgs, reply *XRPSignRawTxReply) (err error) {
	var key hk.Key
	reply.Txid, reply.Result, err = key.RippleSignRawTx(args.Entropy, args.Seed, args.M1, args.M2,
		args.Sequence, args.LedgerSequence,
		args.ToAddr, args.Tag, args.Value, args.Currency, args.Fee)
	if err != nil {
		return err
	}

	return nil
}
