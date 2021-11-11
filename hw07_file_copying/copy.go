package main

import (
	"errors"
	"io"
	"os"
	"time"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fi, err := os.Lstat(fromPath)
	if err != nil {
		return err
	}
	if !fi.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	if offset > fi.Size() {
		return ErrOffsetExceedsFileSize
	}

	if limit == 0 || offset+limit > fi.Size() {
		limit = fi.Size() - offset
	}

	bar := pb.Full.Start64(limit)
	bar.SetRefreshRate(10 * time.Millisecond)
	fileToCopyFrom, err := os.OpenFile(fromPath, os.O_RDONLY, fi.Mode().Perm())
	barReader := bar.NewProxyReader(fileToCopyFrom)
	defer fileToCopyFrom.Close()
	if err != nil {
		return err
	}

	if offset != 0 {
		_, err = fileToCopyFrom.Seek(offset, io.SeekStart)
		if err != nil {
			return err
		}
	}

	fileToCopyTo, err := os.Create(toPath)
	defer fileToCopyTo.Close()
	if err != nil {
		return err
	}

	if _, err = io.CopyN(fileToCopyTo, barReader, limit); err != nil {
		return err
	}
	bar.Finish()

	return nil
}
