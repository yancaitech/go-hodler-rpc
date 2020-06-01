package coinsrpc

import (
	"errors"
	"net/http"

	hk "github.com/yancaitech/go-hodler-keys"
)

// BCHEntropyToAddressArgs struct
type BCHEntropyToAddressArgs struct {
	Entropy        string `json:"entropy"`
	Seed           string `json:"seed"`
	M1             uint32 `json:"m1"`
	M2             uint32 `json:"m2"`
	Count          uint32 `json:"count"`
	MainNet        bool   `json:"mainnet"`
	CompressPubKey bool   `json:"compresspubkey"`
}

// BCHAddressReply struct
type BCHAddressReply struct {
	Address string `json:"address"`
	M1      uint32 `json:"m1"`
	M2      uint32 `json:"m2"`
}

// BCHEntropyToAddressReply struct
type BCHEntropyToAddressReply struct {
	Addresses []BCHAddressReply `json:"addresses"`
}

// BCHAddressValidateArgs struct
type BCHAddressValidateArgs struct {
	Addr    string `json:"address"`
	MainNet bool   `json:"mainnet"`
}

// BCHAddressValidateReply struct
type BCHAddressValidateReply struct {
	Result string `json:"result"`
}

// AddressValidate entry
func (h *BCH) AddressValidate(r *http.Request, args *BCHAddressValidateArgs, reply *BCHAddressValidateReply) (err error) {
	var key hk.Key
	err = key.BitcoinCashAddressValidate(args.Addr, args.MainNet)
	if err != nil {
		return err
	}
	reply.Result = "ok"
	return nil
}

// EntropyToAddress entry
func (h *BCH) EntropyToAddress(r *http.Request, args *BCHEntropyToAddressArgs, reply *BCHEntropyToAddressReply) (err error) {
	if args.Count > 100 {
		return errors.New("count must less than 100")
	}

	var addrs []BCHAddressReply
	var key hk.Key
	m1 := args.M1
	m2 := args.M2
	for i := uint32(0); i < args.Count; i++ {
		err = key.LoadFromEntropy(args.Entropy, args.Seed, m1, m2, args.CompressPubKey)
		if err != nil {
			return err
		}
		addr, err := key.BitcoinCashAddress(args.MainNet, args.CompressPubKey)
		if err != nil {
			return err
		}

		var adr BCHAddressReply
		adr.Address = addr
		adr.M1 = m1
		adr.M2 = m2
		addrs = append(addrs, adr)

		m2++
	}

	reply.Addresses = addrs

	return nil
}
