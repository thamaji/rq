package rq

type readCloser struct {
	read  func([]byte) (int, error)
	close func() error
}

func (rc readCloser) Read(p []byte) (int, error) {
	return rc.read(p)
}
func (rc readCloser) Close() error {
	return rc.close()
}
