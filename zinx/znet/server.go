package znet

import (
	"fmt"
	"net"
	"time"
	"zinx/utils"
	"zinx/ziface"
)

// IServer 接口实现 定义一个Server服务类
type Server struct {
	//服务器的名称
	Name string
	//tcp4 or other
	IPVersion string
	//服务绑定的IP地址
	IP string
	//服务绑定的端口
	Port   int
	Router ziface.IRouter
}

func NewServer() ziface.IServer {
	//先初始化全局配置文件
	utils.GlobalObject.Reload()
	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:    nil,
	}
	return s
}

// ********** 实现 ziface.IServer 里的全部接口方法 **********
// 开启服务网络
func (s *Server) Start() {
	fmt.Printf("IP为: %s, 端口为 %d,的服务器列表服务器正在启动\n", s.IP, s.Port)
	go func() {
		// 1 获取一个tcp的addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}
		// 2 监听服务器的地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}
		//已经监听成功
		fmt.Println("开启 Zinx 服务  ", s.Name, " 成功,监听中...")
		var cid uint32 = 0
		// 3 阻塞的等待客户端链接 处理客户端链接的业务
		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("accept err", err)
				continue
			}
			//不断的循环从客户端获取数据
			// go func() {
			// 	for {
			// 		buf := make([]byte, 512)
			// 		cnt, err := conn.Read(buf)
			// 		if err != nil {
			// 			fmt.Println("recv buf err ", err)
			// 			continue
			// 		}
			// 		//回显
			// 		if _, err := conn.Write(buf[:cnt]); err != nil {
			// 			fmt.Println("write back buf err ", err)
			// 			continue
			// 		}
			// 	}
			// }()

			dealConn := NewConntion(conn, cid, s.Router)
			cid++
			// 启动当前链接的处理业务
			dealConn.Start()

		}
	}()
}

// 停止服务
func (s *Server) Stop() {
	fmt.Println("停止 Zinx 服务 , 名称 ", s.Name)
	//TODO  Server.Stop() 将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
}

// 启动业务
func (s *Server) Serve() {
	s.Start()

	//TODO Server.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加

	//阻塞,否则主Go退出， listenner的go将会退出
	for {
		time.Sleep(10 * time.Second)
	}
}

// 路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router

	fmt.Println("Add Router 成功! ")
}

// ******定义当前客户端链接的handle api ******
// func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
// 	if _, err := conn.Write(data[:cnt]); err != nil {
// 		fmt.Println("write back buf err ", err)
// 		return errors.New("CallBackToClient error")
// 	}
// 	return nil
// }
