package coinsrpc

import (
	"net/http"

	hk "github.com/yancaitech/go-hodler-keys"
)

// ETHSignMsgArgs struct
type ETHSignMsgArgs struct {
	Entropy string `json:"entropy"`
	Seed    string `json:"seed"`
	M1      uint32 `json:"m1"`
	M2      uint32 `json:"m2"`
	Message string `json:"message"`
}

// ETHSignMsgReply struct
type ETHSignMsgReply struct {
	Address   string `json:"address"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

// ETHVerifyMsgArgs struct
type ETHVerifyMsgArgs struct {
	Address   string `json:"address"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

// ETHVerifyMsgReply struct
type ETHVerifyMsgReply struct {
	Result string `json:"result"`
}

// SignMessage entry
func (h *ETH) SignMessage(r *http.Request, args *ETHSignMsgArgs, reply *ETHSignMsgReply) (err error) {
	var key hk.Key
	err = key.LoadFromEntropy(args.Entropy, args.Seed, args.M1, args.M2, false)
	if err != nil {
		return nil
	}
	reply.Address, err = key.EthereumAddress()
	if err != nil {
		return
	}
	reply.Message = args.Message
	reply.Signature, err = key.EthereumSignMessage(args.Message)
	if err != nil {
		return
	}
	return nil
}

// VerifyMessage entry
func (h *ETH) VerifyMessage(r *http.Request, args *ETHVerifyMsgArgs, reply *ETHVerifyMsgReply) (err error) {
	var key hk.Key
	err = key.EthereumVerifyMessage(args.Message, args.Signature, args.Address)
	if err != nil {
		return
	}
	reply.Result = "Verify OK"
	return nil
}
