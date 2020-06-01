package coinsrpc

import (
	"errors"
	"net/http"

	hk "github.com/yancaitech/go-hodler-keys"
)

// EntropyToETHAddressArgs struct
type EntropyToETHAddressArgs struct {
	Entropy string `json:"entropy"`
	Seed    string `json:"seed"`
	M1      uint32 `json:"m1"`
	M2      uint32 `json:"m2"`
	Count   uint32 `json:"count"`
}

// ETHAddressReply struct
type ETHAddressReply struct {
	Address string `json:"address"`
	M1      uint32 `json:"m1"`
	M2      uint32 `json:"m2"`
}

// EntropyToETHAddressReply struct
type EntropyToETHAddressReply struct {
	Addresses []ETHAddressReply `json:"addresses"`
}

// ETHAddressValidateArgs struct
type ETHAddressValidateArgs struct {
	Addr string `json:"address"`
}

// ETHAddressValidateReply struct
type ETHAddressValidateReply struct {
	Result string `json:"result"`
}

// AddressValidate entry
func (h *ETH) AddressValidate(r *http.Request, args *ETHAddressValidateArgs, reply *ETHAddressValidateReply) (err error) {
	var key hk.Key
	err = key.EthereumAddressValidate(args.Addr)
	if err != nil {
		return err
	}
	reply.Result = "ok"
	return nil
}

// EntropyToAddress entry
func (h *ETH) EntropyToAddress(r *http.Request, args *EntropyToETHAddressArgs, reply *EntropyToETHAddressReply) (err error) {
	if args.Count > 100 {
		return errors.New("count must less than 100")
	}

	var addrs []ETHAddressReply
	var key hk.Key
	m1 := args.M1
	m2 := args.M2
	for i := uint32(0); i < args.Count; i++ {
		err = key.LoadFromEntropy(args.Entropy, args.Seed, m1, m2, false)
		if err != nil {
			return err
		}
		addr, err := key.EthereumAddress()
		if err != nil {
			return err
		}

		var adr ETHAddressReply
		adr.Address = addr
		adr.M1 = m1
		adr.M2 = m2
		addrs = append(addrs, adr)

		m2++
	}

	reply.Addresses = addrs

	return nil
}
