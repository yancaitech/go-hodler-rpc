package coinsrpc

import (
	"encoding/hex"
	"net/http"

	hk "github.com/yancaitech/go-hodler-keys"
)

// FILSignRawTxArgs struct
type FILSignRawTxArgs struct {
	Entropy   string `json:"entropy"`
	Seed      string `json:"seed"`
	M1        uint32 `json:"m1"`
	M2        uint32 `json:"m2"`
	Nonce     uint64 `json:"nonce"`
	ToAddress string `json:"toAddress"`
	Value     uint64 `json:"value"`
	GasPrice  uint64 `json:"gasPrice"`
	GasLimit  uint64 `json:"gasLimit"`
	Method    uint64 `json:"method"`
	Params    string `json:"params"`
}

// FILSignRawTxReply struct
type FILSignRawTxReply struct {
	Txid   string `json:"txid"`
	Result string `json:"result"`
	Hextx  string `json:"hex"`
}

// SignRawTx entry
func (h *FIL) SignRawTx(r *http.Request, args *FILSignRawTxArgs, reply *FILSignRawTxReply) (err error) {
	var key hk.Key

	err = key.LoadFromEntropy(args.Entropy, args.Seed, args.M1, args.M2, false)
	if err != nil {
		return err
	}
	addr, err := key.FilecoinAddress()
	if err != nil {
		return err
	}

	var bs []byte
	if len(args.Params) > 0 {
		bs, err = hex.DecodeString(args.Params)
		if err != nil {
			return err
		}
	}

	reply.Result, reply.Hextx, reply.Txid, err = key.FilecoinSignRawTx(args.Entropy, args.Seed, args.M1, args.M2,
		args.Nonce,
		addr,
		args.ToAddress,
		args.Value, args.GasPrice, args.GasLimit, args.Method, bs)
	if err != nil {
		return err
	}

	return nil
}
