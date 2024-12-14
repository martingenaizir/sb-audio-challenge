package fsclients

type Client interface {
	Store() (filePath string, err error)
	Retrieve(filename string) (filePath string, err error)
	Remove(filePath string) error
}

type client struct {
}

func New() Client {
	return &client{}
}
