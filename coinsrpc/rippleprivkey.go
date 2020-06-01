package coinsrpc

import (
	"encoding/hex"
	"net/http"

	hk "github.com/yancaitech/go-hodler-keys"
)

// XRPDumpPrivateKeyArgs struct
type XRPDumpPrivateKeyArgs struct {
	Entropy string `json:"entropy"`
	Seed    string `json:"seed"`
	M1      uint32 `json:"m1"`
	M2      uint32 `json:"m2"`
}

// XRPDumpPrivateKeyReply struct
type XRPDumpPrivateKeyReply struct {
	Address string `json:"address"`
	Secret  string `json:"secret"`
	Hex     string `json:"hex"`
}

// DumpPrivateKey entry
func (h *XRP) DumpPrivateKey(r *http.Request, args *XRPDumpPrivateKeyArgs, reply *XRPDumpPrivateKeyReply) (err error) {
	var key hk.Key
	err = key.LoadFromEntropy(args.Entropy, args.Seed, args.M1, args.M2, true)
	if err != nil {
		return err
	}
	addr, err := key.RippleAddress()
	if err != nil {
		return err
	}
	secret, err := key.RippleSecret()
	if err != nil {
		return nil
	}
	bs, err := key.DumpBitcoinHex()
	if err != nil {
		return err
	}
	reply.Address = addr
	reply.Secret = secret
	reply.Hex = hex.EncodeToString(bs)

	return nil
}
