package main

import (
	"gitlab.com/detl/detl-common/parsers"
	"gitlab.com/detl/detl-common/queues"
	"gitlab.com/detl/transform/handlers"
)

type factory struct{}

func (*factory) NewQueue(name string, conf map[string]string) queues.Queue {
	switch name {
	case "rabbitQueue":
		queue := &queues.RabbitQueue{
			URL:           conf["url"],
			ReadQueue:     conf["readQueue"],
			WriteExchange: conf["writeExchange"],
			WriteKey:      conf["writeKey"],
			Conn:          nil,
		}
		queue.Init(conf)

		return queue

	default:
		return nil
	}
}

func (*factory) NewParser(name string, conf map[string]string) parsers.Parser {
	switch name {
	case "jsonParser":
		parser := parsers.JSONParser{}

		return parser

	default:
		return nil
	}
}

func (*factory) NewHandler(name string,
	mapping map[string]interface{}, parser *parsers.Parser) handlers.Handler {
	switch name {
	case "mapHandler":
		handler := &handlers.MapHandler{Mapping: mapping, Parser: parser}

		return handler

	default:
		return nil
	}
}
