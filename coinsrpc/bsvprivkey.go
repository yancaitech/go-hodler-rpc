package coinsrpc

import (
	"encoding/hex"
	"net/http"

	hk "github.com/yancaitech/go-hodler-keys"
)

// BSVDumpPrivateKeyArgs struct
type BSVDumpPrivateKeyArgs struct {
	Entropy        string `json:"entropy"`
	Seed           string `json:"seed"`
	M1             uint32 `json:"m1"`
	M2             uint32 `json:"m2"`
	MainNet        bool   `json:"mainnet"`
	CompressPubKey bool   `json:"compresspubkey"`
}

// BSVDumpPrivateKeyReply struct
type BSVDumpPrivateKeyReply struct {
	Address string `json:"address"`
	Wif     string `json:"wif"`
	Hex     string `json:"hex"`
}

// DumpPrivateKey entry
func (h *BSV) DumpPrivateKey(r *http.Request, args *BSVDumpPrivateKeyArgs, reply *BSVDumpPrivateKeyReply) (err error) {
	var key hk.Key
	err = key.LoadFromEntropy(args.Entropy, args.Seed, args.M1, args.M2, args.CompressPubKey)
	if err != nil {
		return err
	}

	addr, err := key.BitcoinSVAddress(args.MainNet, args.CompressPubKey)
	if err != nil {
		return err
	}

	wif, err := key.DumpBitcoinWIF(args.MainNet, args.CompressPubKey)
	if err != nil {
		return nil
	}

	bs, err := key.DumpBitcoinHex()
	if err != nil {
		return err
	}

	reply.Address = addr
	reply.Wif = wif
	reply.Hex = hex.EncodeToString(bs)

	return nil
}
