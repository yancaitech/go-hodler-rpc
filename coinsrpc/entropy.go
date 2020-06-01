package coinsrpc

import (
	"errors"
	"net/http"

	hk "github.com/yancaitech/go-hodler-keys"
)

// GenEntropyArgs struct
type GenEntropyArgs struct {
	Bitsize int `json:"bitsize"`
}

// GenEntropyReply struct
type GenEntropyReply struct {
	Entropy string `json:"entropy"`
}

// EntropyToMnemonicArgs struct
type EntropyToMnemonicArgs struct {
	Entropy string `json:"entropy"`
}

// EntropyToMnemonicReply struct
type EntropyToMnemonicReply struct {
	Mnemonic string `json:"mnem"`
}

// EntropyFromMnemonicArgs struct
type EntropyFromMnemonicArgs struct {
	Mnemonic string `json:"mnem"`
}

// EntropyFromMnemonicReply struct
type EntropyFromMnemonicReply struct {
	Entropy string `json:"entropy"`
}

// EntropyGenerate entry
func (h *ENTROPY) EntropyGenerate(r *http.Request, args *GenEntropyArgs, reply *GenEntropyReply) (err error) {
	if args.Bitsize <= 0 {
		return errors.New("bad entropy bitsize, must > 0")
	}
	var key hk.Key
	entropy, err := key.GenerateEntropy(args.Bitsize)
	if err != nil {
		return err
	}
	reply.Entropy = entropy
	return nil
}

// EntropyToMnemonic entry
func (h *ENTROPY) EntropyToMnemonic(r *http.Request, args *EntropyToMnemonicArgs, reply *EntropyToMnemonicReply) (err error) {
	// test:
	// 7dfe828a0d2755bf7c8c2c4f1949aad888e05fbecd726ae4c7ea4b3fdfcc113f
	// "law village penalty bottom inspire text vendor machine excuse ski height rain mix cool will purity helmet change whisper nose worth toward eager today"
	var key hk.Key
	mnem, err := key.EntropyToMnemonic(args.Entropy)
	if err != nil {
		return err
	}
	reply.Mnemonic = mnem
	return nil
}

// EntropyFromMnemonic entry
func (h *ENTROPY) EntropyFromMnemonic(r *http.Request, args *EntropyFromMnemonicArgs, reply *EntropyFromMnemonicReply) (err error) {
	var key hk.Key
	entropy, err := key.EntropyFromMnemonic(args.Mnemonic)
	if err != nil {
		return err
	}
	reply.Entropy = entropy
	return nil
}
