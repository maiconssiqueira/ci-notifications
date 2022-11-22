package http

type HttpHandlers interface {
	HandlerPost() (raw []byte, pretty string, err error)
}

type Methods struct {
	Post
}

type Post struct {
	Content     any
	Url         string
	ContentType string
	Token       string
}
