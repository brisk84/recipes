package domain

import "io"

type FileInfoUpload struct {
	Id     string
	Step   string
	Reader io.Reader
	Size   int64
}

type FileInfoDownload struct {
	Id   string
	Step string
}
