package main

import (
	"github.com/swafran/detl-common/parsers"
	"github.com/swafran/detl-common/queues"
	"github.com/swafran/detl-transform/handlers"
	"github.com/swafran/detl-transform/maps"
)

func NewQueue(name string, conf map[string]string, handler queues.Handler) queues.Queue {
	switch name {
	case "rabbitQueue":
		queue := queues.RabbitQueue{
			URL:           conf["url"],
			ReadQueue:     conf["readQueue"],
			WriteExchange: conf["writeExchange"],
			Handler:       handler,
			Conn:          nil,
		}
		queue.Init(conf)

		return &queue

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
	mapping maps.Mapping, parser parsers.Parser) queues.Handler {
	switch name {
	case "mapHandler":
		handler := handlers.MapHandler{Mapping: mapping, Parser: parser, MapTree: &maps.MapTree{}}

		return &handler

	default:
		return nil
	}
}
