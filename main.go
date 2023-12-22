package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"strconv"
	"time"
)

var selection int
var openScanCode string
var deviceName string

type PaySystemConfig struct{
	M_IsEnablePaySystem 	bool 		`xml:"m_IsEnablePaySystem"`
	M_DeviceId 				string 		`xml:"m_DeviceId"`
}

func main() {
	showMenu()
	time.Sleep(100 * time.Second)
}

func readConfigFile (item int, value string){
	currentUser,_ := user.Current()
	xmlPath := currentUser.HomeDir + `/AppData/LocalLow/guet/智慧蜴学车/Configs/PaySystemConfig.xml`
	_,err := os.Stat(currentUser.HomeDir + "/AppData/LocalLow/guet/智慧蜴学车/Configs")

	if err != nil{
		fmt.Println("文件夹不存在")
	}
	
	//读取文件
	file,err := os.Open(xmlPath)

	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Panicln(err)
	}

	v := PaySystemConfig{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		log.Panicln(err)
	}

	//修改文件
	if item == 1{
		v.M_IsEnablePaySystem,_ = strconv.ParseBool(value)
	}else {
		v.M_DeviceId = value
	}
	
	xmlData,err := xml.MarshalIndent(v, "", "	")
	xmlData = append([]byte(xml.Header), xmlData...)
	if err = os.WriteFile(xmlPath, xmlData, 0777); err != nil {
		log.Fatal("覆写文件失败:",err)
	}
}

func showMenu(){
	fmt.Println("1.打开/关闭扫码功能")
	fmt.Println("2.修改模拟器名称")
	fmt.Printf("请输入你的操作:")
	fmt.Scanf("%d \n", &selection)

	if selection == 1 {
		fmt.Printf("输入true/false打开或关闭扫码功能:")
		fmt.Scanf("%s \n", &openScanCode)
		readConfigFile(1,openScanCode)

		if openScanCode == "true" {
			fmt.Println("打开扫码功能")
		} else{
			fmt.Println("关闭扫码功能")
		}
		
		fmt.Println("------------")
	}else if selection == 2{
		fmt.Printf("请输入模拟器名称:")
		fmt.Scanf("%s \n", &deviceName)
		readConfigFile(2,deviceName)
		fmt.Println("名称修改成功")
		fmt.Println("------------")
	}

	showMenu()
}

