//// Copyright 2021 The klaytn Authors
//// This file is part of the klaytn library.
////
//// The klaytn library is free software: you can redistribute it and/or modify
//// it under the terms of the GNU Lesser General Public License as published by
//// the Free Software Foundation, either version 3 of the License, or
//// (at your option) any later version.
////
//// The klaytn library is distributed in the hope that it will be useful,
//// but WITHOUT ANY WARRANTY; without even the implied warranty of
//// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
//// GNU Lesser General Public License for more details.
////
//// You should have received a copy of the GNU Lesser General Public License
//// along with the klaytn library. If not, see <http://www.gnu.org/licenses/>.
//
//
package types

//
//import (
//	"crypto/ecdsa"
//	"encoding/json"
//	"fmt"
//	"github.com/klaytn/klaytn/blockchain/types/accountkey"
//	"github.com/klaytn/klaytn/common"
//	"github.com/klaytn/klaytn/common/hexutil"
//	"github.com/klaytn/klaytn/crypto"
//	"github.com/klaytn/klaytn/crypto/sha3"
//	"github.com/klaytn/klaytn/fork"
//	"github.com/klaytn/klaytn/kerrors"
//	"github.com/klaytn/klaytn/params"
//	"github.com/klaytn/klaytn/rlp"
//	"math/big"
//	"reflect"
//)
//
//type AccessList []AccessTuple
//
//type AccessTuple struct {
//	Address      common.Address
//	StorageKeys  []common.Hash
//}
//
//type TxInternalDataAccessList struct {
//	ChainID      *big.Int
//	AccountNonce uint64
//	Price        *big.Int
//	GasLimit	 uint64
//	Recipient    *common.Address // nil means contract creation.
//	Amount       *big.Int
//	Payload      []byte
//	AccessList   AccessList
//
//	// Signature values
//	V            *big.Int
//	R            *big.Int
//	S            *big.Int
//
//	// This is only used when marshaling to JSON.
//	Hash *common.Hash `json:"hash" rlp:"-"`
//}
//
//type TxInternalDataAccessListJSON struct {
//	Type         TxType           `json:"typeInt"`
//	TypeStr      string           `json:"type"`
//	AccountNonce hexutil.Uint64   `json:"nonce"`
//	Price        *hexutil.Big     `json:"gasPrice"`
//	GasLimit     hexutil.Uint64   `json:"gas"`
//	Recipient    *common.Address   `json:"to"`
//	Amount       *hexutil.Big     `json:"value"`
//	Payload      hexutil.Bytes    `json:"input"`
//	TxSignatures TxSignaturesJSON `json:"signatures"`
//	AccessList	 AccessList       `json:"accessList"`
//	Hash         *common.Hash     `json:"hash"`
//}
//
//func newEmptyTxInternalDataAccessList() *TxInternalDataAccessList {
//	return &TxInternalDataAccessList{}
//}
//
//func newTxInternalDataAccessList() *TxInternalDataAccessList {
//	return &TxInternalDataAccessList{
//		AccountNonce: 0,
//		Recipient:    nil,
//		Payload:      []byte{},
//		Amount:       new(big.Int),
//		GasLimit:     0,
//		AccessList:   AccessList{},
//		Price:        new(big.Int),
//		V:            new(big.Int),
//		R:            new(big.Int),
//		S:            new(big.Int),
//	}
//}
//
//func newTxInternalDataAccessListWithValues(nonce uint64, to *common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte, accessList AccessList) *TxInternalDataAccessList{
//	d := newTxInternalDataAccessList()
//
//	d.AccountNonce = nonce
//	d.Recipient = to
//	d.GasLimit = gasLimit
//
//	if len(data) > 0 {
//		d.Payload = common.CopyBytes(data)
//	}
//	if amount != nil {
//		d.Amount.Set(amount)
//	}
//	if gasPrice != nil {
//		d.Price.Set(gasPrice)
//	}
//
//	if accessList != nil {
//		copy(d.AccessList, accessList)
//	}
//
//	return d
//}
//
//func newTxInternalDataAccessListWithMap(values map[TxValueKeyType]interface{}) (*TxInternalDataAccessList, error) {
//	d := newTxInternalDataAccessList()
//
//	if v, ok := values[TxValueKeyNonce].(uint64); ok {
//		d.AccountNonce = v
//		delete(values, TxValueKeyNonce)
//	} else {
//		return nil, errValueKeyNonceMustUint64
//	}
//
//	if v, ok := values[TxValueKeyTo].(common.Address); ok {
//		d.Recipient = &v
//		delete(values, TxValueKeyTo)
//	} else {
//		return nil, errValueKeyToMustAddress
//	}
//
//	if v, ok := values[TxValueKeyAmount].(*big.Int); ok {
//		d.Amount.Set(v)
//		delete(values, TxValueKeyAmount)
//	} else {
//		return nil, errValueKeyAmountMustBigInt
//	}
//
//	if v, ok := values[TxValueKeyData].([]byte); ok {
//		d.Payload = common.CopyBytes(v)
//		delete(values, TxValueKeyData)
//	} else {
//		return nil, errValueKeyDataMustByteSlice
//	}
//
//	if v, ok := values[TxValueKeyGasLimit].(uint64); ok {
//		d.GasLimit = v
//		delete(values, TxValueKeyGasLimit)
//	} else {
//		return nil, errValueKeyGasLimitMustUint64
//	}
//
//	if v, ok := values[TxValueKeyGasPrice].(*big.Int); ok {
//		d.Price.Set(v)
//		delete(values, TxValueKeyGasPrice)
//	} else {
//		return nil, errValueKeyGasPriceMustBigInt
//	}
//
//	if v, ok:= values[TxValueAccessList].(AccessList); ok {
//		copy(d.AccessList, v)
//		delete(values, TxValueAccessList)
//	} else {
//		return nil, errValueKeyAccessListInvalid
//	}
//
//	if len(values) != 0 {
//		for k := range values {
//			logger.Warn("unnecessary key", k.String())
//		}
//		return nil, errUndefinedKeyRemains
//	}
//
//	return d, nil
//}
//
//func (t *TxInternalDataAccessList) Type() TxType {
//	return TxTypeAccessList
//}
//
//func (t *TxInternalDataAccessList) GetRoleTypeForValidation() accountkey.RoleType {
//	return accountkey.RoleTransaction
//}
//
//func (t *TxInternalDataAccessList) ChainId() *big.Int {
//	return deriveChainId(t.V)
//}
//
//func (t *TxInternalDataAccessList) GetAccountNonce() uint64 {
//	return t.AccountNonce
//}
//
//func (t *TxInternalDataAccessList) GetPrice() *big.Int {
//	return new(big.Int).Set(t.Price)
//}
//
//func (t *TxInternalDataAccessList) GetGasLimit() uint64 {
//	return t.GasLimit
//}
//
//func (t *TxInternalDataAccessList) GetRecipient() *common.Address {
//	return t.Recipient
//}
//
//func (t *TxInternalDataAccessList) GetAmount() *big.Int {
//	return new(big.Int).Set(t.Amount)
//}
//
//func (t *TxInternalDataAccessList) GetHash() *common.Hash {
//	return t.Hash
//}
//
//func (t *TxInternalDataAccessList) GetPayload() []byte {
//	return t.Payload
//}
//
//func (t *TxInternalDataAccessList) SetHash(hash *common.Hash) {
//	t.Hash = hash
//}
//
//func (t *TxInternalDataAccessList) SetSignature(signatures TxSignatures) {
//	if len(signatures) != 1 {
//		logger.Crit("LegacyTransaction receives a single signature only!")
//	}
//
//	t.V = signatures[0].V
//	t.R = signatures[0].R
//	t.S = signatures[0].S
//}
//
//func (t *TxInternalDataAccessList) RawSignatureValues() TxSignatures {
//	return TxSignatures{&TxSignature{t.V, t.R, t.S}}
//}
//
//func (t *TxInternalDataAccessList) ValidateSignature() bool {
//	return validateSignature(t.V, t.R, t.S)
//}
//
//func (t *TxInternalDataAccessList) RecoverAddress(txhash common.Hash, homestead bool, vfunc func(*big.Int) *big.Int) (common.Address, error) {
//	V := vfunc(t.V)
//	return recoverPlain(txhash, t.R, t.S, V, homestead)
//}
//
//func (t *TxInternalDataAccessList) RecoverPubkey(txhash common.Hash, homestead bool, vfunc func(*big.Int) *big.Int) ([]*ecdsa.PublicKey, error) {
//	V := vfunc(t.V)
//
//	pk, err := recoverPlainPubkey(txhash, t.R, t.S, V, homestead)
//	if err != nil {
//		return nil, err
//	}
//
//	return []*ecdsa.PublicKey{pk}, nil
//}
//
//func (t *TxInternalDataAccessList) IntrinsicGas(currentBlockNumber uint64) (uint64, error) {
//	return IntrinsicGas(t.Payload, t.Recipient == nil, *fork.Rules(big.NewInt(int64(currentBlockNumber))))
//}
//
//func (t *TxInternalDataAccessList) SerializeForSign() []interface{} {
//	return []interface{}{
//		t.Type(),
//		t.ChainID,
//		t.AccountNonce,
//		t.Price,
//		t.GasLimit,
//		t.Recipient,
//		t.Amount,
//		t.Payload,
//		t.AccessList,
//	}
//}
//
//func (t *TxInternalDataAccessList) SerializeForEthTypedTxSignToBytes() []byte {
//	b, _ := rlp.EncodeToBytes(struct {
//		TxType TxType
//		ChainID      *big.Int
//		AccountNonce uint64
//		Price        *big.Int
//		GasLimit	 uint64
//		Recipient    *common.Address // nil means contract creation.
//		Amount       *big.Int
//		Payload      []byte
//		AccessList   AccessList
//	}{
//		t.Type(),
//		t.ChainID,
//		t.AccountNonce,
//		t.Price,
//		t.GasLimit,
//		t.Recipient,
//		t.Amount,
//		t.Payload,
//		t.AccessList,
//	})
//
//	return b
//}
//
//func (t *TxInternalDataAccessList) SenderTxHash() common.Hash {
//	hw := sha3.NewKeccak256()
//	rlp.Encode(hw, t.Type())
//	rlp.Encode(hw, []interface{}{
//		t.ChainID,
//		t.AccountNonce,
//		t.Price,
//		t.GasLimit,
//		t.Recipient,
//		t.Amount,
//		t.Payload,
//		t.AccessList,
//	})
//
//	h := common.Hash{}
//	hw.Sum(h[:0])
//
//	return h
//}
//
//func (t *TxInternalDataAccessList) IsLegacyTransaction() bool {
//	return true
//}
//
//func (t *TxInternalDataAccessList) Equal(a TxInternalData) bool {
//	ta, ok := a.(*TxInternalDataAccessList)
//	if !ok {
//		return false
//	}
//
//	return t.AccountNonce == ta.AccountNonce &&
//		t.Price.Cmp(ta.Price) == 0 &&
//		t.GasLimit == ta.GasLimit &&
//		equalRecipient(t.Recipient, ta.Recipient) &&
//		t.Amount.Cmp(ta.Amount) == 0 &&
//		reflect.DeepEqual(t.AccessList, ta.AccessList) &&
//		t.V.Cmp(ta.V) == 0 &&
//		t.R.Cmp(ta.R) == 0 &&
//		t.S.Cmp(ta.S) == 0
//}
//
//func (t *TxInternalDataAccessList) String() string {
//	var from, to string
//	tx := &Transaction{data: t}
//
//	v, r, s := t.V, t.R, t.S
//	if v != nil {
//		// make a best guess about the signer and use that to derive
//		// the sender.
//		signer := deriveSigner(v)
//		if f, err := Sender(signer, tx); err != nil { // derive but don't cache
//			from = "[invalid sender: invalid sig]"
//		} else {
//			from = fmt.Sprintf("%x", f[:])
//		}
//	} else {
//		from = "[invalid sender: nil V field]"
//	}
//
//	if t.GetRecipient() == nil {
//		to = "[contract creation]"
//	} else {
//		to = fmt.Sprintf("%x", t.GetRecipient().Bytes())
//	}
//	enc, _ := rlp.EncodeToBytes(t)
//	return fmt.Sprintf(`
//	TX(%x)
//	Contract: %v
//	From:     %s
//	To:       %s
//	Nonce:    %v
//	GasPrice: %#x
//	GasLimit  %#x
//	Value:    %#x
//	Data:     0x%x
//    AccessList: %s
//	V:        %#x
//	R:        %#x
//	S:        %#x
//	Hex:      %x
//`,
//		tx.Hash(),
//		t.GetRecipient() == nil,
//		from,
//		to,
//		t.GetAccountNonce(),
//		t.GetPrice(),
//		t.GetGasLimit(),
//		t.GetAmount(),
//		t.GetPayload(),
//		t.AccessList,
//		v,
//		r,
//		s,
//		enc,
//	)
//}
//
//func (t *TxInternalDataAccessList) Validate(stateDB StateDB, currentBlockNumber uint64) error {
//	if t.Recipient != nil {
//		if common.IsPrecompiledContractAddress(*t.Recipient) {
//			return kerrors.ErrPrecompiledContractAddress
//		}
//	}
//	return t.ValidateMutableValue(stateDB, currentBlockNumber)
//}
//
//func (t *TxInternalDataAccessList) ValidateMutableValue(stateDB StateDB, currentBlockNumber uint64) error {
//	return nil
//}
//
//func (t *TxInternalDataAccessList) FillContractAddress(from common.Address, r *Receipt) {
//	if t.Recipient == nil {
//		r.ContractAddress = crypto.CreateAddress(from, t.AccountNonce)
//	}
//}
//
//func (t *TxInternalDataAccessList) Execute(sender ContractRef, vm VM, stateDB StateDB, currentBlockNumber uint64, gas uint64, value *big.Int) (ret []byte, usedGas uint64, err error) {
//	///////////////////////////////////////////////////////
//	// OpcodeComputationCostLimit: The below code is commented and will be usd for debugging purposes.
//	//start := time.Now()
//	//defer func() {
//	//	elapsed := time.Since(start)
//	//	logger.Debug("[TxInternalDataLegacy] EVM execution done", "elapsed", elapsed)
//	//}()
//	///////////////////////////////////////////////////////
//	if t.Recipient == nil {
//		// Sender's nonce will be increased in '`vm.Create()`
//		ret, _, usedGas, err = vm.Create(sender, t.Payload, gas, value, params.CodeFormatEVM)
//	} else {
//		stateDB.IncNonce(sender.Address())
//		ret, usedGas, err = vm.Call(sender, *t.Recipient, t.Payload, gas, value)
//	}
//	return ret, usedGas, err
//}
//
//func (t *TxInternalDataAccessList) MakeRPCOutput() map[string]interface{} {
//	return map[string]interface{}{
//		"typeInt":    t.Type(),
//		"type":       t.Type().String(),
//		"gas":        hexutil.Uint64(t.GasLimit),
//		"gasPrice":   (*hexutil.Big)(t.Price),
//		"input":      hexutil.Bytes(t.Payload),
//		"nonce":      hexutil.Uint64(t.AccountNonce),
//		"to":         t.Recipient,
//		"value":      (*hexutil.Big)(t.Amount),
//		"accessList": t.AccessList,
//		"signatures": TxSignaturesJSON{&TxSignatureJSON{(*hexutil.Big)(t.V), (*hexutil.Big)(t.R), (*hexutil.Big)(t.S)}},
//	}
//}
//
//func (t *TxInternalDataAccessList) MarshalJSON() ([]byte, error) {
//	return json.Marshal(TxInternalDataAccessListJSON{
//		t.Type(),
//		t.Type().String(),
//		(hexutil.Uint64)(t.AccountNonce),
//		(*hexutil.Big)(t.Price),
//		(hexutil.Uint64)(t.GasLimit),
//		t.Recipient,
//		(*hexutil.Big)(t.Amount),
//		t.Payload,
//		TxSignaturesJSON{&TxSignatureJSON{(*hexutil.Big)(t.V), (*hexutil.Big)(t.R), (*hexutil.Big)(t.S)}},
//		t.AccessList,
//		t.Hash,
//	})
//}
//
//func (t *TxInternalDataAccessList) UnmarshalJSON(bytes []byte) error {
//	js := &TxInternalDataAccessListJSON{}
//	if err := json.Unmarshal(bytes, js); err != nil {
//		return err
//	}
//
//	t.AccountNonce = uint64(js.AccountNonce)
//	t.Price = (*big.Int)(js.Price)
//	t.GasLimit = uint64(js.GasLimit)
//	t.Recipient = js.Recipient
//	t.Amount = (*big.Int)(js.Amount)
//	t.Payload = js.Payload
//	t.AccessList = js.AccessList
//	t.V = (*big.Int)(js.TxSignatures[0].V)
//	t.R = (*big.Int)(js.TxSignatures[0].R)
//	t.S = (*big.Int)(js.TxSignatures[0].S)
//	t.Hash = js.Hash
//
//	return nil
//}
