package coinsrpc

import (
	"errors"
	"net/http"

	hk "github.com/yancaitech/go-hodler-keys"
)

// EOSEntropyToPubkeyArgs struct
type EOSEntropyToPubkeyArgs struct {
	Entropy string `json:"entropy"`
	Seed    string `json:"seed"`
	M1      uint32 `json:"m1"`
	M2      uint32 `json:"m2"`
	Count   uint32 `json:"count"`
}

// EOSPubkey struct
type EOSPubkey struct {
	Pubkey1 string `json:"pubkey1"`
	Pubkey2 string `json:"pubkey2"`
	M1      uint32 `json:"m1"`
	M2      uint32 `json:"m2"`
}

// EOSEntropyToPubkeyReply struct
type EOSEntropyToPubkeyReply struct {
	Addresses []EOSPubkey `json:"addresses"`
}

// EntropyToPublicKey entry
func (h *EOS) EntropyToPublicKey(r *http.Request, args *EOSEntropyToPubkeyArgs, reply *EOSEntropyToPubkeyReply) (err error) {
	if args.Count > 100 {
		return errors.New("count must less than 100")
	}

	m1 := args.M1
	m2 := args.M2

	var addrs []EOSPubkey
	var key hk.Key
	for i := uint32(0); i < args.Count; i++ {
		err = key.LoadFromEntropy(args.Entropy, args.Seed, m1, m2, true)
		if err != nil {
			return err
		}
		pubkey1, err := key.EOSPublicKey()
		if err != nil {
			return err
		}
		err = key.LoadFromEntropy(args.Entropy, args.Seed+"active", m1, m2, true)
		if err != nil {
			return err
		}
		pubkey2, err := key.EOSPublicKey()
		if err != nil {
			return err
		}

		var adr EOSPubkey
		adr.Pubkey1 = pubkey1
		adr.Pubkey2 = pubkey2
		adr.M1 = m1
		adr.M2 = m2
		addrs = append(addrs, adr)
		m2++
	}

	reply.Addresses = addrs

	return nil
}
