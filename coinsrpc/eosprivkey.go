package coinsrpc

import (
	"encoding/hex"
	"net/http"

	hk "github.com/yancaitech/go-hodler-keys"
)

// EOSDumpPrivateKeyArgs struct
type EOSDumpPrivateKeyArgs struct {
	Entropy string `json:"entropy"`
	Seed    string `json:"seed"`
	M1      uint32 `json:"m1"`
	M2      uint32 `json:"m2"`
}

// EOSDumpPrivateKeyReply struct
type EOSDumpPrivateKeyReply struct {
	Pubkey1 string `json:"pubkey1"`
	Pubkey2 string `json:"pubkey2"`
	Wif1    string `json:"wif1"`
	Wif2    string `json:"wif2"`
	Hex1    string `json:"hex1"`
	Hex2    string `json:"hex2"`
}

// DumpPrivateKey entry
func (h *EOS) DumpPrivateKey(r *http.Request, args *EOSDumpPrivateKeyArgs, reply *EOSDumpPrivateKeyReply) (err error) {
	var key hk.Key
	err = key.LoadFromEntropy(args.Entropy, args.Seed, args.M1, args.M2, true)
	if err != nil {
		return err
	}

	pubkey1, err := key.EOSPublicKey()
	if err != nil {
		return err
	}

	wif, err := key.DumpBitcoinWIF(true, true)
	if err != nil {
		return nil
	}

	bs, err := key.DumpBitcoinHex()
	if err != nil {
		return err
	}

	reply.Pubkey1 = pubkey1
	reply.Wif1 = wif
	reply.Hex1 = hex.EncodeToString(bs)

	err = key.LoadFromEntropy(args.Entropy, args.Seed+"active", args.M1, args.M2, true)
	if err != nil {
		return err
	}

	pubkey2, err := key.EOSPublicKey()
	if err != nil {
		return err
	}

	wif, err = key.DumpBitcoinWIF(true, true)
	if err != nil {
		return nil
	}

	bs, err = key.DumpBitcoinHex()
	if err != nil {
		return err
	}

	reply.Pubkey2 = pubkey2
	reply.Wif2 = wif
	reply.Hex2 = hex.EncodeToString(bs)

	return nil
}
