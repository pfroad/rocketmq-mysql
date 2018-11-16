package producer

import (
	"github.com/pfroad/rocketmq-mysql/config"
	"github.com/pfroad/rocketmq-mysql/util"
	"github.com/sevennt/rocketmq"
)

// import rocketmq "github.com/apache/rocketmq-externals/rocketmq-go"

// func NewMQProducer() *rocketmq.DefaultMQProducer {
// 	// create a mqClientManager instance
// 	var mqClientConfig = &rocketmq.MqClientConfig{}
// 	var mqClientManager = rocketmq.NewMqClientManager(mqClientConfig)

// 	mqProducerConfig := &rocketmq.MqProducerConfig{}
// 	mqProducer := &rocketmq.DefaultMQProducer{
// 		G
// 	}
// }
const (
	BINLOG_PRODUCER_GROUP = "BINLOG_PRODUCER_GROUP"
)

func NewMQProducer() (rocketmq.Producer, error) {
	ip, _ := util.GetIPAddr()
	conf := &rocketmq.Config{
		Namesrv:      config.GetConfig().MQNameSvcAddr,
		ClientIp:     ip,
		InstanceName: "binlog_producer",
	}

	producer, err := rocketmq.NewDefaultProducer(BINLOG_PRODUCER_GROUP, conf)
	if err != nil {
		return nil, err
	}
	producer.Start()
	return producer, nil
}
