package eth

import (
	"backend/service"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
)

type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

type LogGameStarted struct {
	User      common.Address
	RequestId *big.Int
	Amount    *big.Int
}

type LogGameEnded struct {
	User      common.Address
	RequestId *big.Int
	Win       bool
}

func init() {
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/828ac7e6fdcb46e795a2cf93dca361fa")
	if err != nil {
		panic(err)
	}

	contractAddress := common.HexToAddress(os.Getenv("PRIVATE_KEY"))
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		panic(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	//logGameStartedSig := []byte("GameStarted(address, uint256, uint256)")
	//logGameStartedSigHash := crypto.Keccak256Hash(logGameStartedSig)
	logGameEndedSig := []byte("GameEnded(address, uint256)")
	logGameEndedSigHash := crypto.Keccak256Hash(logGameEndedSig)

	//contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))

	for _, vLog := range logs {
		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			//var transferEvent LogTransfer
			//err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			//if err != nil {
			//	panic(err)
			//}
			service.UpdateUserInfo(
				vLog.Topics[1].Hex(),
				vLog.Topics[2].Hex(),
				vLog.Topics[3].Hex())

		//case logGameStartedSigHash.Hex():
		//
		case logGameEndedSigHash.Hex():
			service.UpdateGameResult(
				vLog.Topics[1].Hex(),
				vLog.Topics[2].Hex(),
			)
		}
	}
}
