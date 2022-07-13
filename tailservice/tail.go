package tailservice

import (
	"github.com/hpcloud/tail"
	"log"
	"logagent/conf"
)

// InitTail 使用tail读取日志，返回tail对象
func InitTail(config conf.Collect) *tail.Tail{
	cfg := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}

	tailObj, err := tail.TailFile(config.LogfilePath, cfg)
	if err != nil {
		log.Fatalf("read %s failed err: %s", config.LogfilePath, err.Error())
	}
	return tailObj
}
