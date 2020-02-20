package utils

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)




func CreateTempFile() string{
	currentTime :=GetCurrentTimeUnix()
	rand.Seed(time.Now().UnixNano())
	randnum := strconv.Itoa(rand.Intn(9223372036854775805))
	currentTime = currentTime+randnum

	err := os.MkdirAll("./data/temp/", os.ModePerm)
	fileName:=currentTime+".txt"
	filePath:="./data/temp/"+fileName
	f,err := os.Create( filePath)

	defer f.Close()

	if err !=nil {

		fmt.Println( err.Error() )

	} else {

		_,err=f.Write([]byte(currentTime))

		if err !=nil{
			fmt.Println(err.Error())
		}else{
			//fmt.Println("err is nil ")
		}

	}


return filePath
}

func DeleteFile( filePath string)  {

	err := os.Remove(filePath)               //删除文件test.txt
	if err != nil {
		//如果删除失败则输出 file remove Error!
		fmt.Printf("file[%v] remove Error=%v \n",filePath,err)

	} else {
		//如果删除成功则输出 file remove OK!
		//fmt.Printf("file[%v] remove OK! \n",filePath)
	}

}