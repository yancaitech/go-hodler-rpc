package coinsrpc

import (
	"encoding/hex"
	"net/http"

	hk "github.com/yancaitech/go-hodler-keys"
)

// FILDumpPrivateKeyArgs struct
type FILDumpPrivateKeyArgs struct {
	Entropy string `json:"entropy"`
	Seed    string `json:"seed"`
	M1      uint32 `json:"m1"`
	M2      uint32 `json:"m2"`
}

// FILDumpPrivateKeyReply struct
type FILDumpPrivateKeyReply struct {
	Address string `json:"address"`
	Secret  string `json:"secret"`
	Hex     string `json:"hex"`
}

// DumpPrivateKey entry
func (h *FIL) DumpPrivateKey(r *http.Request, args *FILDumpPrivateKeyArgs, reply *FILDumpPrivateKeyReply) (err error) {
	var key hk.Key
	err = key.LoadFromEntropy(args.Entropy, args.Seed, args.M1, args.M2, false)
	if err != nil {
		return err
	}
	addr, err := key.FilecoinAddress()
	if err != nil {
		return err
	}
	secret, err := key.FilecoinKeyInfo()
	if err != nil {
		return nil
	}
	bs, err := key.DumpBitcoinHex()
	if err != nil {
		return nil
	}
	reply.Address = addr
	reply.Secret = secret
	reply.Hex = hex.EncodeToString(bs)

	return nil
}
