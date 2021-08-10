package main

import (
	"flag" //命令行参数解析
	"fmt"
	"log" //日志

	"lib/ethernet"
	"lib/water"
)

func main() {

	/***************************************************日志处理**********************************************/

	/***************************************************参数处理**********************************************/
	//定义在命令行中使用的开关参数
	//调用如下 go run vnet.go -plaintext=true -csv=true
	mode := flag.Bool("mode", false, "运行模式")
	p2p := flag.Bool("p2p", false, "是否点对点直连")
	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "婚否")
	delay := flag.Duration("d", 0, "时间间隔")

	// 将命令行解析为定义的标志
	flag.Parse()

	fmt.Println(name, age, married, delay, mode, p2p)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())

	/***************************************************配置文件**********************************************/

	/***************************************************服务相关**********************************************/

	/***************************************************初 始 化**********************************************/
	ifce, err := water.New(water.Config{
		DeviceType: water.TAP,
	})
	if err != nil {
		log.Fatal(err)
	}
	var frame ethernet.Frame

	for {
		frame.Resize(1500)
		n, err := ifce.Read([]byte(frame))
		if err != nil {
			log.Fatal(err)
		}
		frame = frame[:n]
		log.Printf("Dst: %s-> Src: %s[Ethertype: % x]\n", frame.Destination(), frame.Source(), frame.Ethertype())
		//log.Printf("Src: %s", frame.Source())
		//log.Printf("Ethertype: % x\n", frame.Ethertype())
		//log.Printf("Payload: % x\n", frame.Payload())

		ifce.Write([]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01, 0x08, 0x06, 0x00, 0x01, 0x7F, 0x00, 0x00, 0x01, 0x1d, 0x63, 0x71, 0x00, 0x0f, 0x59, 0x17, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67})

	}

	/***************************************************结束处理**********************************************/

}
