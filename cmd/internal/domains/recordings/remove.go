package recordings

func (d *domain) RemoveFile(storedFile StoredFile) {
	_ = d.storage.Remove(storedFile.Path())
}
