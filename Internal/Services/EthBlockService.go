package Services

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

func SendEthFun(reciveAddres string, walletKey string, amount string) error {
	ethcli, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/LQf-7XOeuprlZ7AhXdViw8eaFEHAvyXf") //ENV
	if err != nil {
		log.Fatalf("Failed to connect ETH client")
	}
	private, err := crypto.HexToECDSA(walletKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := private.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal(err)
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ethcli.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	ethValue := new(big.Float)
	ethValue.SetString(amount)
	weiFloat := new(big.Float).Mul(ethValue, big.NewFloat(math.Pow10(18)))
	wei := new(big.Int)

	weiFloat.Int(wei)
	gaslimit := uint64(21000)
	gasPrice, err := ethcli.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to get gas price")
	}
	toAddress := common.HexToAddress(reciveAddres)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, wei, gaslimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, private)
	if err != nil {
		log.Fatalf("Failed to sign tx")
	}

	err = ethcli.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send tx")
	}
	return nil
}

// 'walletAddress': sendAddress,
//          'privateKey': walletKey,
//          'amount': amount,
//          'reciveAddress': reciveAddres,
//Flutter code for inspiration
//  void sendTransaction(String receiver, EtherAmount txValue) async {
//    var apiUrl = "Your RPC Url"; // Replace with your API
//    // Replace with your API
//    var httpClient = http.Client();
//    var ethClient = Web3Client(apiUrl, httpClient);
//
//    EthPrivateKey credentials = EthPrivateKey.fromHex('0x' + privateKey);
//
//    EtherAmount etherAmount = await ethClient.getBalance(credentials.address);
//    EtherAmount gasPrice = await ethClient.getGasPrice();
//
//    print(etherAmount);
//
//    await ethClient.sendTransaction(
//      credentials,
//      Transaction(
//        to: EthereumAddress.fromHex(receiver),
//        gasPrice: gasPrice,
//        maxGas: 100000,
//        value: txValue,
//      ),
//      chainId: 11155111,
//    );
//  }
