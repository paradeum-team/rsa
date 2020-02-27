package utils

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"
)

func Public2Hex() string {

	dst := "#public"
	mm := dst[1:]
	fmt.Printf("public=%s \n", mm)

	key := getHexPubkey()
	fmt.Printf("hex=%s \n", key)
	return key
}

func Private2Hex() string {

	dst := "#private"
	mm := dst[1:]
	fmt.Printf("private=%s \n", mm)

	key := getHexPrivateKey()
	fmt.Printf("hex=%s \n", key)
	return key
}

func GetFullPublicKey()([]byte ,error){
	cert_bytes, err := ioutil.ReadFile("./rsa_public.key")

	return cert_bytes,err
}

func GetPublicKey4OriginBase64() string {
	cert_bytes, _ := ioutil.ReadFile("./rsa_public.key")
	public := string(cert_bytes)
	p1 := strings.ReplaceAll(public, "-----BEGIN PUBLIC KEY-----\n", "")
	p2 := strings.ReplaceAll(p1, "-----END PUBLIC KEY-----\n", "")

	return p2
}

func getHexPubkey() string {
	p2:=GetPublicKey4OriginBase64()
	return base64ToHex(p2)
}

func GetFullPrivateKey()([]byte ,error){
	cert_bytes, err := ioutil.ReadFile("./rsa_private.key")

	return cert_bytes,err
}

func GetPrivateKey4OriginBase64() string {
	cert_bytes, _ := ioutil.ReadFile("./rsa_private.key")
	public := string(cert_bytes)
	p1 := strings.ReplaceAll(public, "-----BEGIN RSA PRIVATE KEY-----\n", "")
	p2 := strings.ReplaceAll(p1, "-----END RSA PRIVATE KEY-----\n", "")

	return p2
}

func getHexPrivateKey() string {
	p2 :=GetPrivateKey4OriginBase64()
	return base64ToHex(string(p2))
}

func base64ToHex(base string) string {
	p, err := base64.StdEncoding.DecodeString(base)
	if err != nil {
		fmt.Println(err.Error())
	}
	return hex.EncodeToString(p)
}

func Hex2Base64(hexComplainText string ) string {
	body,err:=HexDecode(hexComplainText)
	if err !=nil{
		fmt.Printf("step 1 error ...%v \n",err)
	}
	base64Content:=base64.StdEncoding.EncodeToString(body)

	return base64Content

}
func HexDecode(encodeSrc string) ([]byte, error) {
	b, err := hex.DecodeString(encodeSrc)
	if err != nil {
		fmt.Printf("hex decode error %v \n", err.Error())
	}

	return b, err
}

//base64

func Base64Encode(content [] byte) string {
	base64PublicKey := base64.StdEncoding.EncodeToString(content)
	return base64PublicKey

}

func Base64Decode(encodeBase64PublicKey string) (string, error) {
	decodePublicKey := ""
	sDec, err := base64.StdEncoding.DecodeString(encodeBase64PublicKey)
	if err != nil {
		fmt.Printf("base64 decode failure, error=[%v]\n", err)
		return decodePublicKey, err
	} else {
		decodePublicKey = string(sDec)
	}

	return decodePublicKey, nil
}

//1e100000001047b0315705440e5afda6276d46203095ca27d6bfa91cf4f83e0595f1b63104c08ec6643496f6f90c55f56f8a944db689370308f918d6ee036429
//1e100000001047b0de26237588b6482001eae5ec7b9f6260d0c4a17ce38934671ef92ed9de66d62e676e18ddcdaefbd90acba771fd08231f379dacd81ecfa12c

//1e1000000000002045e104a5d0f25181b4e723cd331c41401d171bdb06a934c3e3a85fa749812f480aed5730e7e1f5dede7be18dd0e2de09a43a30b17676c9bb
//1e10000000000020ee5aac46290f801770a4916d26dfd79c33abf33803cc4c9f3e3d2760d24db99d810291450394197a45c2fa9722e7126ed98580256f3e400f

/**
 * 把src 转成hex(16进制的)
 */
func HexEncode(src string) string {
	return hex.EncodeToString([]byte(src))
}

func HexEncodeWithByte(src []byte) string{
	return hex.EncodeToString(src)
}



func PrivateKey2Custom() {
	//base :="3082025c02010002818100badd8cd087c0f52f71581bdb51c5d8f347905822443ecf106152f1831a57f7f6ee019f8310609f46bd9e898f3e1fa88991059432e12c1b18be2b470c0b1e5db613defd6be1f5dae0b54b38ee38273740fae066bf4d93c9229c962c6de68609c974cb4720ef5cb0d05416571e23ccf6e546abadf51c43c71329ea81456b56a6c702030100010281800f5c24b60362a604c75151c0de60dfeb67678307160affa43e0bbe546376f5a7f37a68ad324c6eeb36acf06bb8d48b5afb73b4f1d5b67567bc41fc6e0dffe9fde8814940063075a31a5a334486634e00d9071cc6787b0e8b2e4a391abe6585b7ec06cd138695df2cbae3839aad4e96adff6042baaafb4dc12c2e7c41ba05a539024100dacfb42d930b9bf46c1a821d8d01e248f6264726e0b0a47dce8f47b50e2463d3c16507b735a0e91099f4c9626e2827b61e08872c9ba6378f58dbed67ddd1eebb024100da9fe8181470b178aa73f9bd2747e8c12119762706b0b8a9b51b36cb365defa51dabdcbdff6ddb1bf6c1de6a8be83224f05bf5757edbc77914695ef6adb8756502404b6b3e025b65abfa826c5c9ecce472578da841ec0f94330e3ded3add1823c8d8cb1704a2cc744b00f2dfe1adf0a41c93f424225b68c6d0edb2c6133d32c2d2530241008b7bdc63889f01cbbfccb0b8ab1828fb17381c76f7c0c6809828ddc3b75325738e43b9598b5f369d57c90733a941a2c48889d3487c80927ac0b81dbb7ac8a63502406f7a71286588157911fa6fe71045a69338621594128ccae6c6af9a5bedb951060e931d53d1c7c7828e3d02c1c7f3c1a14b1785c2a8da2996c94b1c195f56f897"
	base := "3082025c02010002818100ba39948baabb0bca792d89e94e0202e5cf6a12f3418c0469bd965e355ce76e90311481039a3fa6be89b0ae0a5972da9633b708a63eb788ad026fde0165d880292d177342e5b8ca0b038d7a5a6343b92c7f6ffdbc2cd6800ad945ff51a6bd9f127ffe3749f31b3dbfce6fd416eef75f2c7290a7ae7b0237f2f7573c15acfb3073020301000102818060383c1e661abb1ee4b9b8d6492e14dd34ec33da6875b61a04915b1feb5ed0ebc3d03a50e834172dbfeb0021ecd1c14b76710e1107bac0edd3a2856ce77893ced10c773129f59cfb37460e68f807e29dc1950d861d4bf559f753fd720f18bb3148564f1c8d02d611fce7e5ab992625a5640f51361f8d13b227ae4e5f15b269d1024100dd4d33462f2e27d00d2986e947c54bd3dc53e33edbf4c8b7b7883005a02772c7e5a3d6fed49d190a0fe257bc468bb7c2a78b4da9dd10c4513aed547f12162567024100d76c714444aa5d27aeaa6db6097ab0038de7b800bc05d935ac2aca4e93b2894bae168558b613cb7e9bd66bd4e222dfdc452fccadbee56815f45adc1241ea8915024046baafd294adb9c7c30d2cb34e5efe773e0a09ad437b9ed328f37bf5b0542b593c49fb23032d1a9d9eaa06c483ff8fc1c4eebee9b55ff07ecdc8a0a2e452f3ad024028fd53580010c284e871394ae7e6d652f6cf5e6d95592f894ce71f73701b0a90c1e13b223412f427751389950a2a449a2ef7f7641fec9aaf82bd3f1ea2383439024100ad4c298897def0d3adb9263bcdd5b1db219c80198cbff0b806deda62d7a804b4e4fc2bf6cbdfcf5aad7e44f8fb9d61a904b904c33b4620a5503590d544dc5be7"
	bytes, _ := hex.DecodeString(base)
	s := base64.StdEncoding.EncodeToString(bytes)

	fmt.Println("custome private key=begin")
	fmt.Println(s)
	fmt.Println("custome private key=end ")

}
