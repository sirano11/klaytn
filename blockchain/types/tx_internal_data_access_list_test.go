package types

//import (
//	"github.com/klaytn/klaytn/common"
//	"math/big"
//	"testing"
//)
//
//var (
//	testAddr = common.HexToAddress("b94f5374fce5edbc8e2a8697c15331677e6ebf0b")
//
//	accessListTx = TxInternalDataAccessList{
//		ChainID: big.NewInt(1),
//		AccountNonce: 3,
//		Price: big.NewInt(1),
//		GasLimit: 25000,
//		Recipient: &testAddr,
//		Amount: big.NewInt(10),
//		Payload: common.FromHex("5544"),
//	}
//
//	emptyEip2718Tx = Transaction{
//		data: &accessListTx,
//	}
//
//	signedEip2718Tx, _ = emptyEip2718Tx.WithSignature(
//		NewEIP155Signer(big.NewInt(1)),
//		common.Hex2Bytes("c9519f4f2b30335884581971573fadf60c6204f59a911df35ee8a540456b266032f1e8e2c5dd761f9e4f88f41c8310aeaba26a8bfcdacfedfa12ec3862d3752101"))
//)
//
//func TestRLPEncode(t *testing.T) {
//	signer := NewEIP155Signer(big.NewInt(1))
//	if signer.Hash(signedEip2718Tx) != common.HexToHash("49b486f0ec0a60dfbbca2d30cb07c9e8ffb2a2ff41f29a1ab6737475f6ff69f3") {
//		t.Errorf("signed EIP-2718 transaction hash mismatch, got %x", signer.Hash(signedEip2718Tx))
//	}
//}
