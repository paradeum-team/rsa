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
	hexMypassEn="6751f109e66f7a69d1d91e8b6159e45b31b104748508ec830d900ef49bad199500938431d60914fc59ce0a7d33ea1b35e11221a0f178dfc144b9e5360a277eb8fd34421b5e01f716cb571e695bb2cc28b671297b794887e33e4048b02d66184d042334f01330392e5670bee69708e6267c043b1e6af063ccb9aac79553aba230"//bfs上
	hexMypassEn="98fdfca3daa9eb3d7b344557966a1bfd4900e78ab595fc36c3c2bb63c8ba8f34594c50628ce713c27a4739284082847dc0618f4fa713899d75b0944388a0d2375a49252ea0df417ba48782efa74bdc05419b4fd84d37f45a4b9351f977553ac393d27d8431751b92c2cfd0f7d7245e2d3ac5d91d824f044f66b674dbc3e7b07f"//公钥加密转换出来的
	fmt.Println(hexMypassEn)
	mypass:=rsa.RsaDecrypteWithprivate(hexMypassEn)
	fmt.Print("明文mypass=")
	fmt.Println(mypass)
	fmt.Println("=======-----<<<<<<>>>>>>>>------======end ")



}
