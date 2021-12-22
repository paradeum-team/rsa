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

	//check ip or url
	fmt.Println("check ip ")
	ip:="192.168.0.145"
	flag:=utils.CheckIp(ip)
	fmt.Printf("check ip =%v \n",flag)

	url:=fmt.Sprintf("https://%s:5145",ip)
	flag=utils.CheckPNUrl(url)
	fmt.Printf("check url=%v \n",flag)
	}
