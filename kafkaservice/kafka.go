package kafkaservice

import (
	"github.com/Shopify/sarama"
	"log"
)

type producer struct {
	client sarama.SyncProducer
}

// InitProducer 初始化kafka生产者对象
func InitProducer(addrs []string) *producer {
	//1.生产者配置
	cfg := sarama.NewConfig()
	cfg.Producer.RequiredAcks = sarama.WaitForAll          //客户端,主分片,副本都分别ack
	cfg.Producer.Partitioner = sarama.NewRandomPartitioner //随机分片（每个分片消息是按顺序存储的,随机的话我理解消息不能保证顺序存储）
	cfg.Producer.Return.Successes = true                   //成功交付的信息

	//2.链接kafka
	client, err := sarama.NewSyncProducer(addrs, cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	p := new(producer)
	p.client = client
	return p
}

// Close 关闭生产者客户端
func (p *producer) Close() {
	p.Close()
}

// SendMes 向kafka发送消息
func (p *producer) SendMes() chan<- *sarama.ProducerMessage {
	msgChat := make(chan *sarama.ProducerMessage)
	go func() {
		for {
			msg := <-msgChat
			pid, offset, err := p.client.SendMessage(msg)
			if err != nil {
				log.Printf("message send failed err: %s", err.Error())
			}
			log.Printf("message send success pid: %v offset: %v\n", pid, offset)
		}
	}()
	return msgChat
}
