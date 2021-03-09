package main

import (
	"gitlab.com/detl/detl-common/parsers"
	"gitlab.com/detl/detl-common/queues"
)

type factory struct{}

func (*factory) NewQueue(name string, conf map[string]string) queues.Queue {
	switch name {
	default:
		queue := &queues.RabbitQueue{
			URL:           conf["url"],
			ReadQueue:     conf["readQueue"],
			WriteExchange: conf["writeExchange"],
			WriteKey:      conf["writeKey"],
			Conn:          nil,
		}
		queue.Init(conf)

		return queue
	}
}

func (*factory) NewParser(name string, conf map[string]string) parsers.Parser {
	switch name {
	default:
		parser := &parsers.JSONParser{}

		return parser
	}
}
