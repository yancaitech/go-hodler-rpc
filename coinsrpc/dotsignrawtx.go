package coinsrpc

import (
	"encoding/json"
	"net/http"
	"strconv"

	hk "github.com/yancaitech/go-hodler-keys"
)

// ChainID:  Polkadot, Kusama, Plasm, Acala

// PolkadotTx struct
type PolkadotTx struct {
	Account            string          `json:"account_id"`
	CallCode           string          `json:"call_code"`
	CallModule         string          `json:"call_module"`
	CallModuleFunction string          `json:"call_module_function"`
	Era                string          `json:"era"`
	ExtrinsicHash      string          `json:"extrinsic_hash"`
	ExtrinsicLength    uint64          `json:"extrinsic_length"`
	Nonce              uint64          `json:"nonce"`
	Params             []PolkadotParam `json:"params"`
	Signature          string          `json:"signature"`
	Tip                string          `json:"tip"`
	VersionInfo        string          `json:"version_info"`
}

// PolkadotParam struct
type PolkadotParam struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Value    string `json:"value"`
	ValueRaw string `json:"value_raw"`
}

// DOTSignRawTxArgs struct
type DOTSignRawTxArgs struct {
	Entropy   string `json:"entropy"`
	Seed      string `json:"seed"`
	M1        uint32 `json:"m1"`
	M2        uint32 `json:"m2"`
	ToAddress string `json:"toAddress"`
	Amount    string `json:"amount"`
	Nonce     string `json:"nonce"`
	Fee       string `json:"fee"`
}

// DOTSignRawTxReply struct
type DOTSignRawTxReply struct {
	Result string `json:"result"`
	Txid   string `json:"txid"`
}

// DOTDecodeRawTxOutArgs struct
type DOTDecodeRawTxOutArgs struct {
	RawTx string `json:"rawtx"`
}

// DOTDecodeRawTxOutReply struct
type DOTDecodeRawTxOutReply struct {
	PolkadotTx
}

// DecodeRawTxOut entry
func (h *DOT) DecodeRawTxOut(r *http.Request, args *DOTDecodeRawTxOutArgs, reply *DOTDecodeRawTxOutReply) (err error) {
	var key hk.Key
	_, txjson, err := key.PolkadotDecodeSignedTransaction(args.RawTx)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(txjson), reply)
	if err != nil {
		return err
	}
	return nil
}

// SignRawTx entry
func (h *DOT) SignRawTx(r *http.Request, args *DOTSignRawTxArgs, reply *DOTSignRawTxReply) (err error) {
	amount, err := strconv.ParseUint(args.Amount, 10, 64)
	if err != nil {
		return err
	}
	nonce, err := strconv.ParseUint(args.Nonce, 10, 64)
	if err != nil {
		return err
	}
	fee, err := strconv.ParseUint(args.Fee, 10, 64)
	if err != nil {
		return err
	}
	var key hk.Key
	err = key.LoadFromEntropy(args.Entropy, args.Seed, args.M1, args.M2, false)
	if err != nil {
		return err
	}
	reply.Txid, reply.Result, err = key.PolkadotCreateSignedTransaction(args.ToAddress, amount, nonce, fee)
	if err != nil {
		return err
	}
	return nil
}
