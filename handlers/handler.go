package handlers

//Handler processes data from input all the way through transformation
type Handler interface {
	Handle(string)
}
