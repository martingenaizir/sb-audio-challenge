package recordings

func (d *domain) RemoveFile(storedFile StoredFile) {
	_ = d.fsClient.Remove(storedFile.Path())
}
