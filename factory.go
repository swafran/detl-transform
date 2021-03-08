package main

import (
	"gitlab.com/detl/detl-common/queues"
)

type factory struct{}

func (*factory) NewQueue(name string, conf map[string]string) queues.Queue {
	switch name {
	default:
		queue := &queues.RabbitQueue{}
		queue.Init(conf)

		return queue
	}
}
