package http

type Handler interface {
	Request() ([]byte, error)
}

type Contains struct {
	Method      string
	Content     any
	Url         string
	ContentType string
	Token       string
}
