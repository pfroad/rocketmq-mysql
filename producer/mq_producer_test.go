package producer

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/pfroad/rocketmq-mysql/util"
	"github.com/sevennt/rocketmq"
)

func TestNewMQProducer(t *testing.T) {
	producer, err := NewMQProducer()
	defer producer.Shutdown()
	if err != nil {
		t.Error(err)
	}
	msg := rocketmq.NewMessage("binlog_topic", []byte("Hello RocketMQ!"))
	if sendResult, err := producer.Send(msg); err != nil {
		t.Error(errors.New("Sync sending fail!"))
	} else {
		fmt.Println("Sync sending success!, ", sendResult)
	}

	consumer(t)
}

func consumer(t *testing.T) {
	var timeSleep = 20 * time.Second
	ip, _ := util.GetIPAddr()
	conf := &rocketmq.Config{
		Namesrv:      "10.35.22.61:9876",
		ClientIp:     ip,
		InstanceName: "DEFAULT",
	}

	consumer, err := rocketmq.NewDefaultConsumer(BINLOG_PRODUCER_GROUP, conf)
	defer consumer.Shutdown()
	if err != nil {
		t.Error(err)
	}
	consumer.Subscribe("binlog_topic", "*")
	consumer.RegisterMessageListener(
		func(msgs []*rocketmq.MessageExt) error {
			for i, msg := range msgs {
				fmt.Println("msg", i, msg.Topic, msg.Flag, msg.Properties, string(msg.Body))
			}
			fmt.Println("Consume success!")
			return nil
		})
	consumer.Start()

	defer func() {
		if e := recover(); e != nil {
			// t.Error(e)
		}
	}()
	time.Sleep(timeSleep)
}
