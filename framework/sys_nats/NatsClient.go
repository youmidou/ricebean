package sys_nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"phoenix-tudou/framework/log4"
	"time"
)

type NatsMsgPack struct {
	ServerId int32  //服务器Id
	Type     string //服务器类型

	NatsKey string
	MsgId   int32
	MsgData []byte
}

func NewNatsMsgPack(natsKey string, msgId int32, msgData []byte) *NatsMsgPack {
	t := &NatsMsgPack{}
	t.NatsKey = natsKey
	t.MsgId = msgId
	t.MsgData = msgData
	return t
}

type NatsClient struct {
	nc            *nats.Conn
	subscriptions map[string]*nats.Subscription // 保存所有订阅
	list          map[string]func(msgId string, bytes []byte)
}

func NewNatsClient() *NatsClient {
	t := &NatsClient{}
	t.list = make(map[string]func(msgId string, bytes []byte))
	t._init()
	return t
}
func (t *NatsClient) _init() {

}

func (t *NatsClient) Start(url string) {
	// Do something with the connection
	// 连接到 NATS 服务器
	nc, err := nats.Connect(nats.DefaultURL,
		nats.Name("API Name Option Example"),
		//nats.Timeout(10*time.Second),
		nats.PingInterval(20*time.Second), nats.MaxPingsOutstanding(5),
		//nats.NoEcho(), //自己不需要收到
	)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	t.nc = nc
	//注册订阅
	// 订阅主题 "updates"
	_, err = nc.Subscribe("updates", func(m *nats.Msg) {
		fmt.Printf("Received message: %s\n", string(m.Data))
		key := ""
		fun := t.list[key]
		d := m.Data
		if fun != nil {
			fun(key, d)
		} else {

		}

		m.Respond([]byte(fmt.Sprintf("data->%s", string(m.Data))))
	})
	if err != nil {
		log.Fatal(err)
	}
	// 使用 select 保持连接活跃
	select {}
}

func (t *NatsClient) Publish(key string, data []byte) {
	/*
		defer func() {
			if r := recover(); r != nil {
				err := debug.Stack()
				log4.Error("OnClientToGateCloseHandler:%s", err)
			}
		}()
	*/
	// 发送请求并等待响应"updates"
	//msg, err := t.nc.Request(key, []byte(fmt.Sprintf("Hello, NATS! %d", num)), time.Second)
	msg, err := t.nc.Request(key, data, time.Second)
	if err != nil {
		log4.Error("Request:%s", err)
		return
	}
	str := string(msg.Data)
	// 打印接收到的响应消息
	fmt.Printf("Received response: %s", str)

	/*
		//t.nc.ChanSubscribe()
		// 发布消息到主题 "updates"
		err := t.nc.Publish("updates", []byte(fmt.Sprintf("Hello, NATS! %d", num)))
		if err != nil {
			log4.Error("Publish:%s", err)
			//log.Fatal(err)
		}
	*/
	/*
		// 确保消息被处理
		//t.nc.Flush()
		if err := t.nc.LastError(); err != nil {
			log.Fatal(err)
		}
	*/
	fmt.Println("Message published and received successfully.")
}

func (t *NatsClient) OnSubscribeReceiveMessages(key int32, fun func(msgId int32, bytes []byte)) {

}

/*
func (t *NatsClient) Start(url string) {
	// 连接到 NATS 服务器
	nc, err := nats.Connect(url,
		nats.Name("API Name Option Example"),
		nats.PingInterval(20*time.Second),
		nats.MaxPingsOutstanding(5),
	)
	if err != nil {
		log.Fatal(err)
	}
	t.nc = nc
	t.subscriptions = make(map[string]*nats.Subscription)

	fmt.Println("NATS client started...")
}*/

/*
// 动态创建订阅
func (t *NatsClient) AddSubscription(subject string) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	sub, err := t.nc.Subscribe(subject, func(m *nats.Msg) {
		// 接收到消息后动态返回
		response := fmt.Sprintf("Received data: %s", string(m.Data))
		fmt.Println(response)

		// 动态响应消息
		err := m.Respond([]byte(response))
		if err != nil {
			log.Println("Error sending response:", err)
		}
	})

	if err != nil {
		return err
	}

	// 将订阅保存到 map 中，以便后续删除
	t.subscriptions[subject] = sub
	fmt.Printf("Subscribed to subject: %s\n", subject)
	return nil
}

// 动态删除订阅
func (t *NatsClient) RemoveSubscription(subject string) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	sub, ok := t.subscriptions[subject]
	if !ok {
		return fmt.Errorf("no subscription found for subject %s", subject)
	}

	// 取消订阅
	err := sub.Unsubscribe()
	if err != nil {
		return err
	}

	// 从 map 中删除订阅
	delete(t.subscriptions, subject)
	fmt.Printf("Unsubscribed from subject: %s\n", subject)
	return nil
}
*/
