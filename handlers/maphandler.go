package handlers

import (
	"gitlab.com/detl/detl-common/parsers"
	"gitlab.com/detl/detl-common/queues"
)

//MapHandler is the default handler, applies map to data; implements Handler
type MapHandler struct {
	Mapping map[string]interface{}
	Parser  *parsers.Parser
}

//Handle processes data, from message in to message out
func (*MapHandler) Handle(*queues.Queue) {}
