package errors

import "errors"

var (
	ErrIOWriteImage = errors.New("can't write image into storage")
	ErrImageDecode  = errors.New("an error occurred during the decoding process")
)
