package service

import (
	"bfs-rsa/base"
	"fmt"
	"github.com/tidwall/gjson"
)

func Upload2BFS(uploadFilePath string ,hexPulicKey string ) string {

	params1 := make(map[string]string)
	params1["days"] = "7"
	params1["asympubkey"] = hexPulicKey

	api:="/rn/combine/file"

	code,body,err:=base.ExecutePostFile(api,params1,uploadFilePath,"")

	if err !=nil {
		fmt.Println("upload to bfs fail ...")
	}

	fmt.Printf("上传文件 返回的code=%d \n",code)

	fmt.Println("上传结果：body===>")
	fmt.Println(string(body))

	json := string(body)
	afid := gjson.Get(json, "data.afid").String()

	return afid
}



func GetMyPassEN4BFS(afid string )string {
	api:="/rn/rfs/param/%s/%s"
	paramMypass :="sympassen"

	api=fmt.Sprintf(api,afid,paramMypass)

	code,body,err:=base.ExecuteGet(api,"")
	if err !=nil {
		fmt.Println("get mypassen from bfs fail ...")
	}

	fmt.Printf("获取mypassen属性  返回的code=%d \n",code)

	fmt.Println("获取mypassen属性：body===>")
	fmt.Println(string(body))

	json := string(body)
	mypassen := gjson.Get(json, "data.parameter_value").String()


	return mypassen
}