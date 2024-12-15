package recordings

type StoredFile struct {
	name string
	path string
}

func (sf StoredFile) Name() string {
	return sf.name
}

func (sf StoredFile) Path() string {
	return sf.path
}
