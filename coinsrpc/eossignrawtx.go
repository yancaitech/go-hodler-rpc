package coinsrpc

import (
	"encoding/json"
	"net/http"

	eos "github.com/yancaitech/go-eos"
	hk "github.com/yancaitech/go-hodler-keys"
)

// mainnet ChainID "aca376f206b8fc25a6ed44dbdc66547c36c6c33e3a119ffbeaef943642f0e906"

// EOSSignRawTxArgs struct
type EOSSignRawTxArgs struct {
	Entropy     string `json:"entropy"`
	Seed        string `json:"seed"`
	M1          uint32 `json:"m1"`
	M2          uint32 `json:"m2"`
	ChainID     string `json:"chainID"`
	HeadBlockID string `json:"headBlockID"`
	FromAccount string `json:"fromAccount"`
	ToAccount   string `json:"toAccount"`
	Quantity    string `json:"quantity"`
	Memo        string `json:"memo"`
}

// EOSSignRawTxReply struct
type EOSSignRawTxReply struct {
	Result eos.PackedTransaction
}

// SignRawTx entry
func (h *EOS) SignRawTx(r *http.Request, args *EOSSignRawTxArgs, reply *EOSSignRawTxReply) (err error) {
	var key hk.Key
	signedtx, err := key.EOSSignRawTx(args.Entropy, args.Seed, args.M1, args.M2,
		args.ChainID, args.HeadBlockID,
		args.FromAccount, args.ToAccount, args.Quantity, args.Memo)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(signedtx), &reply.Result)
	if err != nil {
		return err
	}

	return nil
}
