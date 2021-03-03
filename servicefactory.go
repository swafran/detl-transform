package main

import (
	detl "gitlab.com/detl/detl-common"
)


// NewService returns a new service instance
func NewService(name string, conf map[string]string) interface{} {
	switch name {
	case "rabbitQueue":
		return detl.queue.RabbitQueue{conf["url"], conf["exchange"], conf["key"]}
	default:
		return detl.queue.RabbitQueue{conf["url"], conf["exchange"], conf["key"]}
	}

}
