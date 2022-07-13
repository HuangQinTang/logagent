package main

import (
	"github.com/Shopify/sarama"
	"log"
	"logagent/conf"
	"logagent/kafkaservice"
	"logagent/tailservice"
)

// 日志收集客户端(收集指定目录下的日志文件，发送到kafka中)
// 类似开源的项目还有filebeat
func main() {
	//1.读配置文件
	cfg := conf.GetCfg()

	//2.连接kafka
	client := kafkaservice.InitProducer(cfg.Kafka.ProducerAddr)
	defer client.Close()
	msgChat := client.SendMes() //返回一个管道，往该管道发送的数据将存入kafka

	//3.获取tail对象(该对象可以读取日志文件)
	tailObj := tailservice.InitTail(cfg.Collect)

	//4.读取日志并发送到kafka
	for {
		line, ok := <-tailObj.Lines
		if !ok { //读取错误
			log.Printf("tail file close reopen, filename:%s\n", tailObj.Filename)
			continue
		}
		msg := &sarama.ProducerMessage{
			Topic: cfg.Kafka.LogTopic,
			Value: sarama.StringEncoder(line.Text),
		}
		go func(msg *sarama.ProducerMessage) {
			msgChat <- msg
		}(msg)
	}
}
