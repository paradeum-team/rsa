package rsa

import (
	"fmt"
	"bfs-rsa/utils"
)

 var publicKeyHex ="30819f300d06092a864886f70d010101050003818d0030818902818100ba39948baabb0bca792d89e94e0202e5cf6a12f3418c0469bd965e355ce76e90311481039a3fa6be89b0ae0a5972da9633b708a63eb788ad026fde0165d880292d177342e5b8ca0b038d7a5a6343b92c7f6ffdbc2cd6800ad945ff51a6bd9f127ffe3749f31b3dbfce6fd416eef75f2c7290a7ae7b0237f2f7573c15acfb30730203010001"
 var privateKeyHex ="3082025c02010002818100ba39948baabb0bca792d89e94e0202e5cf6a12f3418c0469bd965e355ce76e90311481039a3fa6be89b0ae0a5972da9633b708a63eb788ad026fde0165d880292d177342e5b8ca0b038d7a5a6343b92c7f6ffdbc2cd6800ad945ff51a6bd9f127ffe3749f31b3dbfce6fd416eef75f2c7290a7ae7b0237f2f7573c15acfb3073020301000102818060383c1e661abb1ee4b9b8d6492e14dd34ec33da6875b61a04915b1feb5ed0ebc3d03a50e834172dbfeb0021ecd1c14b76710e1107bac0edd3a2856ce77893ced10c773129f59cfb37460e68f807e29dc1950d861d4bf559f753fd720f18bb3148564f1c8d02d611fce7e5ab992625a5640f51361f8d13b227ae4e5f15b269d1024100dd4d33462f2e27d00d2986e947c54bd3dc53e33edbf4c8b7b7883005a02772c7e5a3d6fed49d190a0fe257bc468bb7c2a78b4da9dd10c4513aed547f12162567024100d76c714444aa5d27aeaa6db6097ab0038de7b800bc05d935ac2aca4e93b2894bae168558b613cb7e9bd66bd4e222dfdc452fccadbee56815f45adc1241ea8915024046baafd294adb9c7c30d2cb34e5efe773e0a09ad437b9ed328f37bf5b0542b593c49fb23032d1a9d9eaa06c483ff8fc1c4eebee9b55ff07ecdc8a0a2e452f3ad024028fd53580010c284e871394ae7e6d652f6cf5e6d95592f894ce71f73701b0a90c1e13b223412f427751389950a2a449a2ef7f7641fec9aaf82bd3f1ea2383439024100ad4c298897def0d3adb9263bcdd5b1db219c80198cbff0b806deda62d7a804b4e4fc2bf6cbdfcf5aad7e44f8fb9d61a904b904c33b4620a5503590d544dc5be7"
 var mypassenHex = "32b82266380ad87400c54ffbc38a0c292c54b2a324ecea3a4696bd0966ebd8288c201b0faa78c9d121f88c0e68e06096ecda82988a4338598c9c29c694e57c27f1a8f4353c6a1bbaf45171f7a4b2cdf6a6b6693255ceee0538e227a5b470ee7f41c031d2300e24abe818da76190215bda6de59f0982acef755040af8280ee849"



var privateKeyTemplate = `  
-----BEGIN RSA PRIVATE KEY-----
%s
-----END RSA PRIVATE KEY-----  
`

var publicKeyTemplate = `  
-----BEGIN PUBLIC KEY-----
%s
-----END PUBLIC KEY-----  
`



//2.公私钥明文转 base64

func Base64EnPublicKey(  ) string {
	content,err := utils.HexDecode(publicKeyHex)
	if err !=nil {
		fmt.Println("public key to base64")
		return ""
	}else {
		return utils.Base64Encode(content)
	}

}

func Base64EnPrivateKey(  ) string {
	content,err := utils.HexDecode(privateKeyHex)
	if err !=nil {
		fmt.Println("private key to base64")
		return ""
	}else {
		return utils.Base64Encode(content)
	}
}

//3.私钥解密 加密的密码

//3.1 拼装pem公私钥

func GetPemPrivateKey()string {
	private :=Base64EnPrivateKey()
	privateKey:=fmt.Sprintf(privateKeyTemplate,private)
	return privateKey
}

func GetPemPublicKey() string {
	public:=Base64EnPublicKey()
	publicKey:=fmt.Sprintf(publicKeyTemplate,public)
	return publicKey
}

//3.2公钥加密
func RsaEncryptByPublic(cipertext string) string {
	publicKey :=GetPemPublicKey()
	buf :=[]byte(cipertext)
	keys :=[]byte(publicKey)
	enContent,err:=utils.Rsa2Encrypt(buf,keys)
	if err !=nil{
		fmt.Println("公钥加密出错了")
		return ""
	}

	return string(enContent[:])
}

func RsaDecryptByPrivate(enCipertext string) string {
	privateKey :=GetPemPrivateKey()
	keys :=[]byte(privateKey)
	buf :=[]byte(enCipertext)

	deContent,err :=utils.Rsa2Decrypt(buf,keys)
	if err !=nil{
		fmt.Println("解密出错了...")
		return ""
	}
	return string(deContent[:])
}

//----自定义：使用自定义的公私钥
func RsaEncryptWithPublic(cipertext string,publickey []byte )([] byte,error){

	buf :=[]byte(cipertext)
	return utils.Rsa2Encrypt(buf,publickey)
}

func RsaEncryptWithPublic2Hex(cipertext string) string {
	//publicKey:=utils.GetPublicKey4OriginBase64()
	publicKey,err:=utils.GetFullPublicKey()
	if err !=nil{
		fmt.Printf("获取公钥失败"+err.Error())
	}
	enCrypte,err:=RsaEncryptWithPublic(cipertext,publicKey)

	if err !=nil{
		fmt.Printf("公钥加密失败"+err.Error())
	}

	hexMypassen:=utils.HexEncodeWithByte(enCrypte)

	return hexMypassen
}

//---

func RsaDecrypteWithprivate(hexCipertext string) string {
	mypassen,err:=utils.HexDecode(hexCipertext)
	if err !=nil{
		fmt.Println("hex decode mypassen  fail ... ")
	}

	privateKey,err:=utils.GetFullPrivateKey()
	if err !=nil{
		fmt.Println("get private key fail  ... ")
	}

	mypass,err:=utils.Rsa2Decrypt(mypassen,privateKey)

	if err !=nil{
		fmt.Println("private 解密失败 ")
	}

	return string(mypass[:])
}
