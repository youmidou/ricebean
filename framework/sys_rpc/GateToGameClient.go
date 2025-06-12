/*
* rpc
added by yh @ 2023/5/19 10:16
注意:辉哥提醒 rpc只能单向请求 只适合 单向逻辑运算服务器连接
*/
package sys_rpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"phoenix-tudou/framework/log4"
	"phoenix-tudou/z_Tools/ProtoToCS/Protocal/pb"
)

type GateToGameClient struct {
	conn     *grpc.ClientConn
	c        pb.GateToGameClient
	ctx      context.Context
	queueReq chan *pb.GateToGameReq
	queueRet chan *pb.GateToGameRet
}

func (t *GateToGameClient) Connect(ipaddress string) {
	// 建立GRPC连接
	conn, err := grpc.Dial(ipaddress, grpc.WithInsecure())
	if err != nil {
		log4.Error("GateToGame Connection failed...")
		return
	}
	defer conn.Close()
	t.conn = conn
	//ctx, cancel := context.WithTimeout(t.ctx, time.Second*5)
	//defer cancel()
	log4.Info("GateToGame Connect successfully.. ")
	//# 建立链接 socket
	t.c = pb.NewGateToGameClient(t.conn)
	//创建连接
	//t.ctx = context.Background()
	t.ctx = context.WithValue(context.Background(), "userid", "1234")
	t.queueReq = make(chan *pb.GateToGameReq)
	t.queueRet = make(chan *pb.GateToGameRet)

	for {
		select {
		case req := <-t.queueReq:
			//fmt.Println("处理对象 " + req.Name)
			t.__processReq(req)
		case ret := <-t.queueRet:
			t.processRet(ret)
		default:
			//fmt.Println("队列为空，等待中...")
			//time.Sleep(time.Second) // 模拟等待时间
		}
	}
}
func (t *GateToGameClient) RequestProcessor(req *pb.GateToGameReq) (*pb.GateToGameRet, error) {
	ret, err := t.c.ProcessGateToGame(t.ctx, req)
	if err != nil {
		fmt.Printf("无法发起gRPC请求: %v", err)
	}
	fmt.Printf("返回的消息: %s", ret.RetId)
	return ret, err
}

func (t *GateToGameClient) SendProcess(req *pb.GateToGameReq) *pb.GateToGameRet {

	ret, err := t.c.ProcessGateToGame(t.ctx, req)
	if err != nil {
		fmt.Printf("无法发起gRPC请求: %v", err)
	}
	fmt.Printf("返回的消息: %s", ret.RetId)
	return ret
}

func (t *GateToGameClient) __processReq(req *pb.GateToGameReq) {

	//todata := &pb.GateToGameReq{Name: "GateToGame->hello"}
	//todata.ReqId = "wwwwwwww"
	defer func() {
		if err2 := recover(); err2 != nil {
			fmt.Println(" test3出错了，错误信息为：", err2)
			//panic(err2) // 将异常继续向上抛出
		}
	}()
	ret, err := t.c.ProcessGateToGame(t.ctx, req)
	if err != nil {
		fmt.Printf("无法发起gRPC请求: %v", err)
	}
	fmt.Printf("返回的消息: %s", ret.RetId)
	t.queueRet <- ret
}
func (t *GateToGameClient) processRet(ret *pb.GateToGameRet) {
	fmt.Printf("processRet处理消息: %s", ret.RetId)
}
