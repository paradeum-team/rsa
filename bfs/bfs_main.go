package bfs

import (
	"bfs-rsa/bfs/service"
	"bfs-rsa/utils"
	"fmt"
	"bfs-rsa/bfs/rsa"
)

func BfsMain() {

	//1.生成一个待上传文件
	//uploadFilePath:=utils.CreateTempFile()
	//publicKeyHex :=utils.Public2Hex()
	utils.Private2Hex()

	//afid:=service.Upload2BFS(uploadFilePath,publicKeyHex)
   afid:="1e10000000000020962c9a29ac520b5d647413a78952f77be52cc97653faeacb5a01fb85149d90b8e7233f5fc9cd75f2342d80da4cd48bc13191e56c4620bc55"

	fmt.Printf("step1:==>上传文件：afid=%s \n",afid)
	mypassen:=service.GetMyPassEN4BFS(afid)

	fmt.Printf("step2:==>加密密码：\n mypassen=%s \n",mypassen)

	mypassDes:=rsa.RsaDecrypteWithprivate(mypassen)
	fmt.Print("step3:==>明文密码：")
	fmt.Println(mypassDes)
	///---- reset 2 value ...
	fmt.Println("reset 2 value ...")
	mypassen="94420bcef89ad29827ef80e777bc1cbb7e5a708f384db7a9f97b44bd817ff442cb25092e7edf4389adb9f508c9f75625f2ce95697cbb0844e75df04ec87cfa22951ccea6cd2758e0de843b91195ba19f8e39de5255292e64655338680cf321a95ce800d46e847321130eb3d11756cd93e0f4ab682876ef8cd16103aa9351a189"
	fmt.Printf("step2:==>加密密码：\n mypassen=%s \n",mypassen)
	mypassDes=rsa.RsaDecrypteWithprivate(mypassen)
	fmt.Print("step3:==>明文密码：")
	fmt.Println(mypassDes)

	fmt.Println("=======-----<<<<<<验证加密密码>>>>>>>>------======begin ")

	cipertext :="49d2ab711504df5491edb0033f7f0054"
	//cipertext :=mypassDes
	fmt.Println("明文mypass="+cipertext)
	hexMypassEn:=rsa.RsaEncryptWithPublic2Hex( cipertext)
	fmt.Println(hexMypassEn)
	fmt.Println("==>2")
	hexMypassEn="33573f7b19e6c788242e3ade6ae60a8902c2a0525d4bfc38a4fb8090cbbef09e6168ffa073bbb9b15a2b353c228130d0ec21dc01f9ebd44a2e305489a3ba7cedfedf56ee1bb3fcf11168949aa8fa6e194c09454c65e0f4282368497b6f11a86fcc49516a9fe1658bee22e888a29a5669c8c174d2c2b487447e976f6f338292a5"//公钥加密转换出来的
	hexMypassEn="98fdfca3daa9eb3d7b344557966a1bfd4900e78ab595fc36c3c2bb63c8ba8f34594c50628ce713c27a4739284082847dc0618f4fa713899d75b0944388a0d2375a49252ea0df417ba48782efa74bdc05419b4fd84d37f45a4b9351f977553ac393d27d8431751b92c2cfd0f7d7245e2d3ac5d91d824f044f66b674dbc3e7b07f"//公钥加密转换出来的
	hexMypassEn="6751f109e66f7a69d1d91e8b6159e45b31b104748508ec830d900ef49bad199500938431d60914fc59ce0a7d33ea1b35e11221a0f178dfc144b9e5360a277eb8fd34421b5e01f716cb571e695bb2cc28b671297b794887e33e4048b02d66184d042334f01330392e5670bee69708e6267c043b1e6af063ccb9aac79553aba230"//bfs上
	fmt.Println(hexMypassEn)
	mypass:=rsa.RsaDecrypteWithprivate(hexMypassEn)
	fmt.Print("明文mypass=")
	fmt.Println(mypass)
	fmt.Println("=======-----<<<<<<>>>>>>>>------======end ")

	convertHex2base64()

}


func convertHex2base64(){
	asymprivatekey:="3082025c02010002818100b2931ed8c5c317a21234b272aaeaac5a4a437fa0f5571106201443f631917b7e803e1cfdc5c45d582c7d4fbbfa6958bbe7fdf9807ce7524b345a09410a4d1221d4269e444b611e9706f3371157f7a59b0af5eb532573af81e848e09c0f157709eb320157135088c91f8e70e02f28fe2fc987cab200f5ba95b2deac3ee954f56d02030100010281810080bdb02f77e3fed5c96a547c76dc59057f24ca8eb051e4e4159c86a2a779cb1e98362f908553dc38055b1270e347afc148afc180f08b94b8c33566168de6a27aac5da51760710464526951ac06fcd430b41e6743ce58646384bf9dabb16342928ccd29cdd59c24446e74fbb6021b19eca4a3ba23c94a230ca276e47db730f5c1024100e09be381276bd1233f6effbc7d26b73b3d156fdd6fe45045b366f00281ef27a26b8be9eb5669888ebd233ffd8e1c44444566c0b4a75f144c3d73df01e19d1d6b024100cb883524b9a10320f3bdee38b843d2d1d85f13ad9847673ba30cecfbf389b94f50a9dec0704cb9234f4baab6c6686173e3bf6e567342aa5323634efd9869d6870240707c018aabe865fb86081bd114c82fd95df4ff69e1607107071a95365c461e0e4c57ec952c587bc828569ff4669827641d71a12e97e798edb994eee92d4916f30240508e69924e9266dd9cfb589544813ba8e8bdc0d91d836d5d924c6463d8617361ec283fcce4e5c9052bb6642c115d70ece0807d1864684767be82e8d9c0806b7302401b894465e13a0ec402dfb4c73302a5fbf1ce2f39586e299c9a221aba3c0cb0814e840bb9fa39ecbdf0d59cc03e676549d44af51cf02867b23c2a6016bf976877"
	asympubkey:="30819f300d06092a864886f70d010101050003818d0030818902818100b2931ed8c5c317a21234b272aaeaac5a4a437fa0f5571106201443f631917b7e803e1cfdc5c45d582c7d4fbbfa6958bbe7fdf9807ce7524b345a09410a4d1221d4269e444b611e9706f3371157f7a59b0af5eb532573af81e848e09c0f157709eb320157135088c91f8e70e02f28fe2fc987cab200f5ba95b2deac3ee954f56d0203010001"

	fmt.Println("=======-----<<<<<<hex 2 base64 >>>>>>>>------======end ")

	privateKey:= utils.Hex2Base64(asymprivatekey)
	publicKey:=utils.Hex2Base64(asympubkey)

	fmt.Println("privateKey:")
	fmt.Println(privateKey)
	fmt.Println("publicKey:")
	fmt.Println(publicKey)

	fmt.Println("equal ...compare...")


	base64PrivateKey:="MIICXAIBAAKBgQCykx7YxcMXohI0snKq6qxaSkN/oPVXEQYgFEP2MZF7foA+HP3FxF1YLH1Pu/ppWLvn/fmAfOdSSzRaCUEKTRIh1CaeREthHpcG8zcRV/elmwr161Mlc6+B6EjgnA8VdwnrMgFXE1CIyR+OcOAvKP4vyYfKsgD1upWy3qw+6VT1bQIDAQABAoGBAIC9sC934/7VyWpUfHbcWQV/JMqOsFHk5BWchqKnecsemDYvkIVT3DgFWxJw40evwUivwYDwi5S4wzVmFo3monqsXaUXYHEEZFJpUawG/NQwtB5nQ85YZGOEv52rsWNCkozNKc3VnCREbnT7tgIbGeyko7ojyUojDKJ25H23MPXBAkEA4JvjgSdr0SM/bv+8fSa3Oz0Vb91v5FBFs2bwAoHvJ6Jri+nrVmmIjr0jP/2OHERERWbAtKdfFEw9c98B4Z0dawJBAMuINSS5oQMg873uOLhD0tHYXxOtmEdnO6MM7PvziblPUKnewHBMuSNPS6q2xmhhc+O/blZzQqpTI2NO/Zhp1ocCQHB8AYqr6GX7hggb0RTIL9ld9P9p4WBxBwcalTZcRh4OTFfslSxYe8goVp/0ZpgnZB1xoS6X55jtuZTu6S1JFvMCQFCOaZJOkmbdnPtYlUSBO6jovcDZHYNtXZJMZGPYYXNh7Cg/zOTlyQUrtmQsEV1w7OCAfRhkaEdnvoLo2cCAa3MCQBuJRGXhOg7EAt+0xzMCpfvxzi85WG4pnJoiGro8DLCBToQLufo57L3w1ZzAPmdlSdRK9RzwKGeyPCpgFr+XaHc="
	base64PublicKey:="MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCykx7YxcMXohI0snKq6qxaSkN/oPVXEQYgFEP2MZF7foA+HP3FxF1YLH1Pu/ppWLvn/fmAfOdSSzRaCUEKTRIh1CaeREthHpcG8zcRV/elmwr161Mlc6+B6EjgnA8VdwnrMgFXE1CIyR+OcOAvKP4vyYfKsgD1upWy3qw+6VT1bQIDAQAB"
	if privateKey==base64PrivateKey{
		fmt.Println("equals private key ")
	}else{
		fmt.Println("not equals private key ")

	}

	if publicKey==base64PublicKey{
		fmt.Println("equals public key ")
	}else{
		fmt.Println("not equals public key ")

	}


	fmt.Println("=======-----<<<<<<hex 2 base64 end >>>>>>>>------======end ")


}