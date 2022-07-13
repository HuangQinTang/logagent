package conf

type Config struct {
	Kafka   Kafka   `ini:"kafka"`
	Collect Collect `ini:"collect"`
}

// Kafka 配置
type Kafka struct {
	ProducerAddr []string `ini:"producerAddr"`
	LogTopic     string   `ini:"logTopic"`
}

// Collect 日志收集配置
type Collect struct {
	LogfilePath string `ini:"logfilePath"` //收集日志的路径
}
