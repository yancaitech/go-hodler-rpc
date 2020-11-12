package coinsrpc

import (
	"errors"
	"net/http"

	hk "github.com/yancaitech/go-hodler-keys"
)

// BTCEntropyToAddressArgs struct
type BTCEntropyToAddressArgs struct {
	Entropy        string `json:"entropy"`
	Seed           string `json:"seed"`
	M1             uint32 `json:"m1"`
	M2             uint32 `json:"m2"`
	Count          uint32 `json:"count"`
	MainNet        bool   `json:"mainnet"`
	CompressPubKey bool   `json:"compresspubkey"`
}

// BTCAddressReply struct
type BTCAddressReply struct {
	Address string `json:"address"`
	Pubkey  string `json:"pubkey"`
	M1      uint32 `json:"m1"`
	M2      uint32 `json:"m2"`
}

// BTCEntropyToAddressReply struct
type BTCEntropyToAddressReply struct {
	Addresses []BTCAddressReply `json:"addresses"`
}

// BTCAddressValidateArgs struct
type BTCAddressValidateArgs struct {
	Addr    string `json:"address"`
	MainNet bool   `json:"mainnet"`
}

// BTCAddressValidateReply struct
type BTCAddressValidateReply struct {
	Result string `json:"result"`
}

// AddressValidate entry
func (h *BTC) AddressValidate(r *http.Request, args *BTCAddressValidateArgs, reply *BTCAddressValidateReply) (err error) {
	var key hk.Key
	err = key.BitcoinAddressValidate(args.Addr, args.MainNet)
	if err != nil {
		return err
	}
	reply.Result = "ok"
	return nil
}

// EntropyToAddress entry
func (h *BTC) EntropyToAddress(r *http.Request, args *BTCEntropyToAddressArgs, reply *BTCEntropyToAddressReply) (err error) {
	if args.Count > 100 {
		return errors.New("count must less than 100")
	}

	var addrs []BTCAddressReply
	var key hk.Key
	m1 := args.M1
	m2 := args.M2
	for i := uint32(0); i < args.Count; i++ {
		err = key.LoadFromEntropy(args.Entropy, args.Seed, m1, m2, args.CompressPubKey)
		if err != nil {
			return err
		}
		addr, err := key.BitcoinAddress(args.MainNet, args.CompressPubKey)
		if err != nil {
			return err
		}
		pubk, err := key.BitcoinPubKeyString(args.CompressPubKey)
		if err != nil {
			return err
		}

		var adr BTCAddressReply
		adr.Address = addr
		adr.Pubkey = pubk
		adr.M1 = m1
		adr.M2 = m2
		addrs = append(addrs, adr)

		m2++
	}

	reply.Addresses = addrs

	return nil
}
