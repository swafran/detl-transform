package main

import "gitlab.com/detl/detl-common/queues"

// NewService returns a new service instance
func NewService(name string, conf map[string]string) interface{} {
	switch name {
	case "rabbitQueue":
		queue := queues.RabbitQueue{}
		queue.Init(conf)
		return queue

	default:

		// TODO: return error, empty interface
		return queues.RabbitQueue{}
	}

}
