package coinsrpc

import (
	"encoding/hex"
	"net/http"

	hk "github.com/yancaitech/go-hodler-keys"
)

// DOTDumpPrivateKeyArgs struct
type DOTDumpPrivateKeyArgs struct {
	Entropy string `json:"entropy"`
	Seed    string `json:"seed"`
	M1      uint32 `json:"m1"`
	M2      uint32 `json:"m2"`
}

// DOTDumpPrivateKeyReply struct
type DOTDumpPrivateKeyReply struct {
	Address string `json:"address"`
	Hex     string `json:"hex"`
}

// DumpPrivateKey entry
func (h *DOT) DumpPrivateKey(r *http.Request, args *DOTDumpPrivateKeyArgs, reply *DOTDumpPrivateKeyReply) (err error) {
	var key hk.Key
	err = key.LoadFromEntropy(args.Entropy, args.Seed, args.M1, args.M2, false)
	if err != nil {
		return err
	}
	addr, err := key.PolkadotAddress()
	if err != nil {
		return err
	}
	bs, err := key.DumpBitcoinHex()
	if err != nil {
		return nil
	}
	reply.Address = addr
	reply.Hex = hex.EncodeToString(bs)

	return nil
}
