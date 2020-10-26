package coinsrpc

import (
	"errors"
	"net/http"

	hk "github.com/yancaitech/go-hodler-keys"
)

// DOTEntropyToAddressArgs struct
type DOTEntropyToAddressArgs struct {
	Entropy string `json:"entropy"`
	Seed    string `json:"seed"`
	M1      uint32 `json:"m1"`
	M2      uint32 `json:"m2"`
	Count   uint32 `json:"count"`
}

// DOTAddressReply struct
type DOTAddressReply struct {
	Address string `json:"address"`
	M1      uint32 `json:"m1"`
	M2      uint32 `json:"m2"`
}

// DOTEntropyToAddressReply struct
type DOTEntropyToAddressReply struct {
	Addresses []DOTAddressReply `json:"addresses"`
}

// DOTAddressValidateArgs struct
type DOTAddressValidateArgs struct {
	Addr string `json:"address"`
}

// DOTAddressValidateReply struct
type DOTAddressValidateReply struct {
	Result string `json:"result"`
}

// AddressValidate entry
func (h *DOT) AddressValidate(r *http.Request, args *DOTAddressValidateArgs, reply *DOTAddressValidateReply) (err error) {
	var key hk.Key
	err = key.PolkadotAddressValidate(args.Addr)
	if err != nil {
		return err
	}
	reply.Result = "ok"
	return nil
}

// EntropyToAddress entry
func (h *DOT) EntropyToAddress(r *http.Request, args *DOTEntropyToAddressArgs, reply *DOTEntropyToAddressReply) (err error) {
	if args.Count > 100 {
		return errors.New("count must less than 100")
	}

	m1 := args.M1
	m2 := args.M2

	var addrs []DOTAddressReply
	var key hk.Key
	for i := uint32(0); i < args.Count; i++ {
		err = key.LoadFromEntropy(args.Entropy, args.Seed, m1, m2, false)
		if err != nil {
			return err
		}
		addr, err := key.PolkadotAddress()
		if err != nil {
			return err
		}

		var adr DOTAddressReply
		adr.Address = addr
		adr.M1 = m1
		adr.M2 = m2
		addrs = append(addrs, adr)
		m2++
	}
	reply.Addresses = addrs

	return nil
}
