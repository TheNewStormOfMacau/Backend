package eth

import (
	"backend/service"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
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

func Init() {
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/8yLe2tHeoTjkWPFF_lVU5Rw2oHsFa6Jt")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	contractAddress := common.HexToAddress("0xd18321420B9D9AdB69C80cD04e1dDb6A4e785FcF")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logGameEndedSig := []byte("GameEnded(address,uint256)")
	logGameEndedSigHash := crypto.Keccak256Hash(logGameEndedSig)

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to filter logs: %v", err)
	}

	fmt.Println("Listening for events...")

	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Subscription error: %v", err)
		case vLog := <-logs:
			fmt.Printf("Received log: %v\n", vLog)
			switch vLog.Topics[0].Hex() {
			case logTransferSigHash.Hex():
				fmt.Println("Transfer event detected")
				service.UpdateUserInfo(
					vLog.Topics[1].Hex(),
					vLog.Topics[2].Hex(),
					vLog.Topics[3].Hex())
			case logGameEndedSigHash.Hex():
				fmt.Println("GameEnded event detected")
				service.UpdateGameResult(
					vLog.Topics[1].Hex(),
					vLog.Topics[2].Hex(),
				)
			default:
				fmt.Printf("Unknown event: %v\n", vLog.Topics[0].Hex())
			}
		}
	}
}

//func Init() {
//	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/828ac7e6fdcb46e795a2cf93dca361fa")
//	if err != nil {
//		panic(err)
//	}
//
//	contractAddress := common.HexToAddress("0xd18321420B9D9AdB69C80cD04e1dDb6A4e785FcF")
//	query := ethereum.FilterQuery{
//		Addresses: []common.Address{
//			contractAddress,
//		},
//	}
//
//	//logs, err := client.FilterLogs(context.Background(), query)
//	//if err != nil {
//	//	panic(err)
//	//}
//
//	logTransferSig := []byte("Transfer(address,address,uint256)")
//	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
//	//logGameStartedSig := []byte("GameStarted(address, uint256, uint256)")
//	//logGameStartedSigHash := crypto.Keccak256Hash(logGameStartedSig)
//	logGameEndedSig := []byte("GameEnded(address, uint256)")
//	logGameEndedSigHash := crypto.Keccak256Hash(logGameEndedSig)
//
//	logs := make(chan types.Log)
//	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
//	if err != nil {
//		panic(err)
//	}
//	for {
//		select {
//		case err := <-sub.Err():
//			panic(err)
//		case vLog := <-logs:
//			switch vLog.Topics[0].Hex() {
//			case logTransferSigHash.Hex():
//				//var transferEvent LogTransfer
//				//err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
//				//if err != nil {
//				//	panic(err)
//				//}
//				service.UpdateUserInfo(
//					vLog.Topics[1].Hex(),
//					vLog.Topics[2].Hex(),
//					vLog.Topics[3].Hex())
//			case logGameEndedSigHash.Hex():
//				service.UpdateGameResult(
//					vLog.Topics[1].Hex(),
//					vLog.Topics[2].Hex(),
//				)
//			}
//		}
//	}
//
//	//contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
//}
