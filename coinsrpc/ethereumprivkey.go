package coinsrpc

import (
	"encoding/hex"
	"net/http"

	hk "github.com/yancaitech/go-hodler-keys"
)

// ETHDumpPrivateKeyArgs struct
type ETHDumpPrivateKeyArgs struct {
	Entropy string `json:"entropy"`
	Seed    string `json:"seed"`
	M1      uint32 `json:"m1"`
	M2      uint32 `json:"m2"`
}

// ETHDumpPrivateKeyReply struct
type ETHDumpPrivateKeyReply struct {
	Address string `json:"address"`
	Wif     string `json:"wif"`
	Hex     string `json:"hex"`
}

// DumpPrivateKey entry
func (h *ETH) DumpPrivateKey(r *http.Request, args *ETHDumpPrivateKeyArgs, reply *ETHDumpPrivateKeyReply) (err error) {
	var key hk.Key
	err = key.LoadFromEntropy(args.Entropy, args.Seed, args.M1, args.M2, false)
	if err != nil {
		return err
	}

	addr, err := key.EthereumAddress()
	if err != nil {
		return err
	}

	wif, err := key.DumpBitcoinWIF(true, false)
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
