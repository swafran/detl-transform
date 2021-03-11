package handlers

import "gitlab.com/detl/detl-common/queues"

//Handler processes data from input all the way through transformation
type Handler interface {
	Handle(*queues.Queue)
}
