package main

import (
	"bfs-rsa/bfs"
	"bfs-rsa/utils"
	"fmt"
)

func main(){
	bfs.BfsMain()

	fmt.Printf("the local key ... ")
	//---public key
	fmt.Printf("the public key is : \n")
	utils.Public2Hex()

	//---private key
	fmt.Printf("the private key is : \n")
	utils.Private2Hex()
}
