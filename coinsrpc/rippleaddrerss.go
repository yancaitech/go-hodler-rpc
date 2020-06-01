package coinsrpc

import (
	"errors"
	"net/http"

	hk "github.com/yancaitech/go-hodler-keys"
)

// XRPEntropyToAddressArgs struct
type XRPEntropyToAddressArgs struct {
	Entropy string `json:"entropy"`
	Seed    string `json:"seed"`
	M1      uint32 `json:"m1"`
	M2      uint32 `json:"m2"`
	Count   uint32 `json:"count"`
}

// XRPAddressReply struct
type XRPAddressReply struct {
	Address string `json:"address"`
	M1      uint32 `json:"m1"`
	M2      uint32 `json:"m2"`
}

// XRPEntropyToAddressReply struct
type XRPEntropyToAddressReply struct {
	Addresses []XRPAddressReply `json:"addresses"`
}

// XRPAddressValidateArgs struct
type XRPAddressValidateArgs struct {
	Addr string `json:"address"`
}

// XRPAddressValidateReply struct
type XRPAddressValidateReply struct {
	Result string `json:"result"`
}

// AddressValidate entry
func (h *XRP) AddressValidate(r *http.Request, args *XRPAddressValidateArgs, reply *XRPAddressValidateReply) (err error) {
	var key hk.Key
	err = key.RippleAddressValidate(args.Addr)
	if err != nil {
		return err
	}
	reply.Result = "ok"
	return nil
}

// EntropyToAddress entry
func (h *XRP) EntropyToAddress(r *http.Request, args *XRPEntropyToAddressArgs, reply *XRPEntropyToAddressReply) (err error) {
	if args.Count > 100 {
		return errors.New("count must less than 100")
	}

	m1 := args.M1
	m2 := args.M2

	var addrs []XRPAddressReply
	var key hk.Key
	for i := uint32(0); i < args.Count; i++ {
		err = key.LoadFromEntropy(args.Entropy, args.Seed, m1, m2, true)
		if err != nil {
			return err
		}
		addr, err := key.RippleAddress()
		if err != nil {
			return err
		}

		var adr XRPAddressReply
		adr.Address = addr
		adr.M1 = m1
		adr.M2 = m2
		addrs = append(addrs, adr)
		m2++
	}
	reply.Addresses = addrs

	return nil
}
