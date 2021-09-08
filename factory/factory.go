package factory

import (
	"gitlab.com/detl/detl-common/parsers"
	"gitlab.com/detl/detl-common/queues"
	"gitlab.com/detl/transform/handlers"
	"gitlab.com/detl/transform/maps"
)

func NewQueue(name string, conf map[string]string, handler *handlers.Handler) queues.Queue {
	switch name {
	case "rabbitQueue":
		queue := &queues.RabbitQueue{
			URL:           conf["url"],
			ReadQueue:     conf["readQueue"],
			WriteExchange: conf["writeExchange"],
			WriteKey:      conf["writeKey"],
			Handler:       handler,
			Conn:          nil,
		}
		queue.Init(conf)

		return queue

	default:
		return nil
	}
}

func NewParser(name string, conf map[string]string) parsers.Parser {
	switch name {
	case "jsonParser":
		parser := parsers.JSONParser{}

		return &parser

	default:
		return nil
	}
}

func NewHandler(name string,
	mapping map[string]interface{}, parser *parsers.Parser) handlers.Handler {
	switch name {
	case "mapHandler":
		handler := handlers.MapHandler{Mapping: mapping, Parser: parser, MapTree: maps.MapTree{}}

		return &handler

	default:
		return nil
	}
}
