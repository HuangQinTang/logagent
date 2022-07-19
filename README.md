# logagent
日志收集

- demo大致流程：从etcd中获取要读取的日志文件路径，以及要发送到kafka哪个topic下
- etct配置的key设置为`collect_log_%s_conf`,%s替换为ip地址
- etcd和kafka请在/conf/config.ini中配置
- 启动`go run main.go`
- 模仿 [李文州老师](https://www.bilibili.com/video/BV1Df4y1C7o5?spm_id_from=333.337.search-card.all.click) 这个案例写的，原demo最后是把kafka中的数据消费到es中，我这里没有实现这一步
