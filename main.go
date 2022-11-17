package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

const (
	GAS_LIMIT = 21000
	GAS_PRICE = 500 * 1e9
)

func SignTransaction(tx *types.Transaction, privateKeyStr string) (string, error) {
	privateKey, err := StringToPrivateKey(privateKeyStr)
	if err != nil {
		return "", err
	}
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(5)), privateKey)
	//signTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		return "", nil
	}

	b, err := rlp.EncodeToBytes(signTx)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func StringToPrivateKey(privateKeyStr string) (*ecdsa.PrivateKey, error) {
	privateKeyByte, err := hexutil.Decode(privateKeyStr)
	if err != nil {
		return nil, err
	}
	privateKey, err := crypto.ToECDSA(privateKeyByte)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func main() {
	sign, err := TransferRawSign("0xc2FcF1f82C6F0af412E3e74AeD192EA9Ba62d279", big.NewInt(0.0001*1e18), "0x41f323f198b1bd743f4dfacc6fcae795e51406f4f97761b299145ce138d8bd98")
	if err != nil {
		panic(err)
	}
	fmt.Println(sign)
}

func TransferRawSign(to string, value *big.Int, privateKey string) (string, error) {
	toAddress := common.HexToAddress(to)
	return SignTransaction(types.NewTx(&types.LegacyTx{
		Nonce:    13,
		GasPrice: big.NewInt(GAS_PRICE),
		Gas:      GAS_LIMIT,
		To:       &toAddress,
		Value:    value,
		Data:     nil,
		V:        nil,
		R:        nil,
		S:        nil,
	}), privateKey)
}
