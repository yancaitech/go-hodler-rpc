package coinsrpc

import (
	"net/http"

	hk "github.com/yancaitech/go-hodler-keys"
)

// BTCSignMsgArgs struct
type BTCSignMsgArgs struct {
	Entropy        string `json:"entropy"`
	Seed           string `json:"seed"`
	M1             uint32 `json:"m1"`
	M2             uint32 `json:"m2"`
	MainNet        bool   `json:"mainnet"`
	CompressPubKey bool   `json:"compresspubkey"`
	Message        string `json:"message"`
}

// BTCSignMsgReply struct
type BTCSignMsgReply struct {
	Address   string `json:"address"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

// BTCVerifyMsgArgs struct
type BTCVerifyMsgArgs struct {
	Address   string `json:"address"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
	MainNet   bool   `json:"mainnet"`
}

// BTCVerifyMsgReply struct
type BTCVerifyMsgReply struct {
	Result string `json:"result"`
}

// SignMessage entry
func (h *BTC) SignMessage(r *http.Request, args *BTCSignMsgArgs, reply *BTCSignMsgReply) (err error) {
	var key hk.Key
	err = key.LoadFromEntropy(args.Entropy, args.Seed, args.M1, args.M2, args.CompressPubKey)
	if err != nil {
		return nil
	}
	reply.Address, err = key.BitcoinAddress(args.MainNet, args.CompressPubKey)
	if err != nil {
		return
	}
	reply.Message = args.Message
	reply.Signature, err = key.BitcoinSignMessage(args.Message, args.CompressPubKey)
	if err != nil {
		return
	}
	return nil
}

// VerifyMessage entry
func (h *BTC) VerifyMessage(r *http.Request, args *BTCVerifyMsgArgs, reply *BTCVerifyMsgReply) (err error) {
	var key hk.Key
	err = key.BitcoinVerifyMessage(args.Message, args.Signature, args.Address, args.MainNet)
	if err != nil {
		return
	}
	reply.Result = "Verify OK"
	return nil
}
