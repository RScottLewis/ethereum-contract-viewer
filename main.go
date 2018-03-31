package main
import (
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

)

func main() {
	conn, err:=ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal("Failed to connect to the Ethereum net %v",err)
	}
	contract,err :=NewTronToken(common.HexToAddress("0xf230b790E05390FC8295F4d3F60332c93BEd42e2"),conn)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}
	amt, _ := contract.BalanceOf(&bind.CallOpts{},common.HexToAddress("0xf7640d4881c6bbc66d8b73281daa8fc5e172876e2828381d9b7c3a91a7dd7505"))
	fmt.Println(amt)
	}
