package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	live "github.com/iyear/biligo-live"
	"github.com/iyear/biligo-live/message"
	"github.com/iyear/biligo-live/utils"
)

// 同 README.md 的快速开始

func main() {
	const room int64 = 573893

	// 获取一个Live实例
	// debug: debug模式，输出一些额外的信息
	// heartbeat: 心跳包发送间隔。不发送心跳包，70 秒之后会断开连接，通常每 30 秒发送 1 次
	// cache: Rev channel 的缓存
	// recover: panic recover后的操作函数
	l := live.NewLive(true, 30*time.Second, 0, func(err error) {
		log.Println("panic:", err)
		// do something...
	})

	// 连接ws服务器
	// dialer: ws dialer
	// host: bilibili live ws host
	if err := l.Conn(websocket.DefaultDialer, utils.WsDefaultHost); err != nil {
		log.Fatal(err)
		return
	}

	ctx, stop := context.WithCancel(context.Background())

	ifError := make(chan error)
	go func() {
		// 进入房间
		// room: room id(真实ID，短号需自行转换)
		// key: 用户标识，可留空
		// uid: 用户UID，可随机生成
		if err := l.Enter(ctx, 12345678, int(room), "", ""); err != nil {
			log.Println("Error Encountered: ", err)
			log.Println("Room Disconnected")
			ifError <- err
			return
		}
	}()

	go rev(ctx, l)

	// 15s的演示
	after := time.After(15 * time.Second)
	select {
	case <-after:
		fmt.Println("I want to stop")
		// 关闭ws连接与相关协程
		stop()
		break
	case err := <-ifError:
		fmt.Println("I don't want to stop, but I encountered an error: ", err)
		break
	}
	// 五秒時間讓他關閉其他 goroutine
	<-time.After(5 * time.Second)
}

func rev(ctx context.Context, l *live.Live) {
	for {
		select {
		case tp := <-l.Rev:
			if tp.Error != nil {
				// do something...
				log.Println(tp.Error)
				continue
			}
			handle(tp.Msg)
		case <-ctx.Done():
			log.Println("rev func stopped")
			return
		}
	}
}

func handle(msg message.Msg) {
	// 使用 msg.(type) 进行事件跳转和处理，常见事件基本都完成了解析(Parse)功能，不常见的功能有一些实在太难抓取
	// 更多注释和说明等待添加
	switch msg := msg.(type) {
	// 心跳回应直播间人气值
	case message.HeartbeatReply:
		log.Printf("hot: %d\n", msg.GetHot())
	// 弹幕消息
	case message.Danmaku:
		fmt.Printf("弹幕: %s (%d:%s) 【%s】| %d\n", msg.Content, msg.MID, msg.Uname, msg.MedalName, msg.Time)
	// 礼物消息
	case message.SendGift:
		fmt.Printf("%s: %s %d个%s\n", msg.Action, msg.Uname, msg.Num, msg.GiftName)
	// 直播间粉丝数变化消息
	case message.FansUpdate:
		fmt.Printf("room: %d,fans: %d,fansClub: %d\n", msg.RoomID, msg.Fans, msg.FansClub)
	// case:......

	// Raw 表示live未实现的CMD命令，请自行处理raw数据。也可以提issue更新这个CMD
	case message.Raw:
		fmt.Println("unknown msg type|raw:", string(msg))
	}
}
