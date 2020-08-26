package coinsrpc

import (
	"errors"
	"net/http"

	hk "github.com/yancaitech/go-hodler-keys"
)

// FILEntropyToAddressArgs struct
type FILEntropyToAddressArgs struct {
	Entropy string `json:"entropy"`
	Seed    string `json:"seed"`
	M1      uint32 `json:"m1"`
	M2      uint32 `json:"m2"`
	Count   uint32 `json:"count"`
}

// FILAddressReply struct
type FILAddressReply struct {
	Address string `json:"address"`
	M1      uint32 `json:"m1"`
	M2      uint32 `json:"m2"`
}

// FILEntropyToAddressReply struct
type FILEntropyToAddressReply struct {
	Addresses []FILAddressReply `json:"addresses"`
}

// FILAddressValidateArgs struct
type FILAddressValidateArgs struct {
	Addr string `json:"address"`
}

// FILAddressValidateReply struct
type FILAddressValidateReply struct {
	Result string `json:"result"`
}

// AddressValidate entry
func (h *FIL) AddressValidate(r *http.Request, args *FILAddressValidateArgs, reply *FILAddressValidateReply) (err error) {
	var key hk.Key
	err = key.FilecoinAddressValidate(args.Addr)
	if err != nil {
		return err
	}
	reply.Result = "ok"
	return nil
}

// EntropyToAddress entry
func (h *FIL) EntropyToAddress(r *http.Request, args *FILEntropyToAddressArgs, reply *FILEntropyToAddressReply) (err error) {
	if args.Count > 100 {
		return errors.New("count must less than 100")
	}

	m1 := args.M1
	m2 := args.M2

	var addrs []FILAddressReply
	var key hk.Key
	for i := uint32(0); i < args.Count; i++ {
		err = key.LoadFromEntropy(args.Entropy, args.Seed, m1, m2, false)
		if err != nil {
			return err
		}
		addr, err := key.FilecoinAddress()
		if err != nil {
			return err
		}

		var adr FILAddressReply
		adr.Address = addr
		adr.M1 = m1
		adr.M2 = m2
		addrs = append(addrs, adr)
		m2++
	}
	reply.Addresses = addrs

	return nil
}
