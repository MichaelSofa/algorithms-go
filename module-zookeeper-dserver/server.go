package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

// 创建zk连接

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func startServer(port string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	fmt.Println(tcpAddr)
	checkError(err)

	listener, err := net.ListenTCP("tcp4", tcpAddr)
	checkError(err)

	//注册zk节点q
	// 链接zk
	conn, err := GetConnect()
	if err != nil {
		fmt.Printf(" connect zk error: %s ", err)
	}
	defer conn.Close()

	// zk节点注册
	// 注意:需要提前创建好永久节点：go_servers（conn.Create("/go_servers", nil, 0, zk.WorldACL(zk.PermAll)) // zk.WorldACL(zk.PermAll)控制访问权限模式）
	err = RegistServer(conn, port)
	if err != nil {
		fmt.Printf(" regist node error: %s ", err)
	}

	for {
		conn, err := listener.Accept()	// 接收客户端请求
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err)
			continue
		}
		go handleClient(conn, port)
	}

	fmt.Println("aaaaaa")
}

func GetConnect() (conn *zk.Conn, err error) {
	zkList := []string{"localhost:2181"}
	conn, _, err = zk.Connect(zkList, 10*time.Second)
	if err != nil {
		fmt.Println(err)
	}

	return
}

// zk.FlagEphemeral = 1:短暂，session断开则改节点也被删除
func RegistServer(conn *zk.Conn, host string) (err error) {
	_, err = conn.Create("/go_servers/"+host, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))

	return
}

func handleClient(conn net.Conn, port string) {
	defer conn.Close()

	daytime := time.Now().String()

	conn.Write([]byte(port + ": " + daytime))	// 往接口写数据，发送给客户端
}


func main() {
	go startServer("127.0.0.1:8897")
	go startServer("127.0.0.1:8898")
	go startServer("127.0.0.1:8899")

	a := make(chan bool, 1)

	<-a	// 接收操作；
}




