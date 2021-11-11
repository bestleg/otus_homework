package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	fromTest = "testdata/input.txt"
	toTest   = "testdata/out1.txt"
)

func TestCopy(t *testing.T) {
	t.Run("offset > file size", func(t *testing.T) {
		err := Copy(fromTest, toTest, 1024*1024*500, limit)
		require.Error(t, err, ErrOffsetExceedsFileSize)
		_ = os.Remove(toTest)
	})
	t.Run("good job", func(t *testing.T) {
		err := Copy(fromTest, toTest, 0, 0)
		require.NoError(t, err)
		fiFrom, _ := os.Lstat(fromTest)
		fiTo, _ := os.Lstat(toTest)
		require.Equal(t, fiTo.Size(), fiFrom.Size())
		_ = os.Remove(toTest)
	})
	t.Run("limit > file Size", func(t *testing.T) {
		fi, err := os.Lstat(fromTest)
		err = Copy(fromTest, toTest, 0, fi.Size()+1000)
		require.NoError(t, err)
		fiFrom, _ := os.Lstat(fromTest)
		fiTo, _ := os.Lstat(toTest)
		require.Equal(t, fiTo.Size(), fiFrom.Size())
		_ = os.Remove(toTest)
	})
	t.Run("unsupported file", func(t *testing.T) {
		err := Copy("/dev/urandom", toTest, 0, 0)
		require.Error(t, err, ErrUnsupportedFile)
		_ = os.Remove(toTest)
	})
	t.Run("offset6000 limit1000", func(t *testing.T) {
		err := Copy(fromTest, toTest, 6000, 1000)
		require.NoError(t, err)
		fiFrom, _ := os.Lstat("testdata/out_offset6000_limit1000.txt")
		fiTo, _ := os.Lstat(toTest)
		require.Equal(t, fiTo.Size(), fiFrom.Size())
		_ = os.Remove(toTest)
	})
}
