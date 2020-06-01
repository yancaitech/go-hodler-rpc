package coinsrpc

import (
	"errors"
	"net/http"

	hk "github.com/yancaitech/go-hodler-keys"
)

// BSVEntropyToAddressArgs struct
type BSVEntropyToAddressArgs struct {
	Entropy        string `json:"entropy"`
	Seed           string `json:"seed"`
	M1             uint32 `json:"m1"`
	M2             uint32 `json:"m2"`
	Count          uint32 `json:"count"`
	MainNet        bool   `json:"mainnet"`
	CompressPubKey bool   `json:"compresspubkey"`
}

// BSVAddressReply struct
type BSVAddressReply struct {
	Address string `json:"address"`
	M1      uint32 `json:"m1"`
	M2      uint32 `json:"m2"`
}

// BSVEntropyToAddressReply struct
type BSVEntropyToAddressReply struct {
	Addresses []BSVAddressReply `json:"addresses"`
}

// BSVAddressValidateArgs struct
type BSVAddressValidateArgs struct {
	Addr    string `json:"address"`
	MainNet bool   `json:"mainnet"`
}

// BSVAddressValidateReply struct
type BSVAddressValidateReply struct {
	Result string `json:"result"`
}

// AddressValidate entry
func (h *BSV) AddressValidate(r *http.Request, args *BSVAddressValidateArgs, reply *BSVAddressValidateReply) (err error) {
	var key hk.Key
	err = key.BitcoinSVAddressValidate(args.Addr, args.MainNet)
	if err != nil {
		return err
	}
	reply.Result = "ok"
	return nil
}

// EntropyToAddress entry
func (h *BSV) EntropyToAddress(r *http.Request, args *BSVEntropyToAddressArgs, reply *BSVEntropyToAddressReply) (err error) {
	if args.Count > 100 {
		return errors.New("count must less than 100")
	}

	var addrs []BSVAddressReply
	var key hk.Key
	m1 := args.M1
	m2 := args.M2
	for i := uint32(0); i < args.Count; i++ {
		err = key.LoadFromEntropy(args.Entropy, args.Seed, m1, m2, args.CompressPubKey)
		if err != nil {
			return err
		}
		addr, err := key.BitcoinSVAddress(args.MainNet, args.CompressPubKey)
		if err != nil {
			return err
		}

		var adr BSVAddressReply
		adr.Address = addr
		adr.M1 = m1
		adr.M2 = m2
		addrs = append(addrs, adr)

		m2++
	}

	reply.Addresses = addrs

	return nil
}
