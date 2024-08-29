package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)
// 专门往kafka写日志

var (
	client sarama.SyncProducer   //声明一个全局的连接的kafka的生产者client
)

func Init(addrs []string)(err error){
	// 构建 生产者
    // 生成 生产者配置文件
    config := sarama.NewConfig()
    // 设置生产者 消息 回复等级 0 1 all
    config.Producer.RequiredAcks = sarama.WaitForAll
    // 设置生产者 成功 发送消息 将在什么 通道返回
    config.Producer.Return.Successes = true
    // 设置生产者 发送的分区
    config.Producer.Partitioner = sarama.NewRandomPartitioner
    // 连接 kafka
    client, err = sarama.NewSyncProducer(addrs, config)
    if err != nil {
        fmt.Println("连接err:%v",err)
        return
    }
    return
}

func SendMsgToKafka(topic,data string){
	msg := &sarama.ProducerMessage{}
    msg.Topic = topic
    msg.Value = sarama.StringEncoder(data)
    message, offset, err := client.SendMessage(msg)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(message, " ", offset)
}