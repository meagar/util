package util

import (
	"io"
	"os"
	"testing"
)

func TestWriteFile(t *testing.T) {
	testdata := "this is a test"
	tmpfile, err := os.CreateTemp("", "writefiletest")
	if err != nil {
		panic(err)
	}

	tmpfile.Close()

	defer cleanup(tmpfile.Name())

	err = WriteFile(tmpfile.Name(), testdata)

	if err != nil {
		t.Errorf("WriteFile: Got %s want nil", err.Error())
	}

	file, err := os.Open(tmpfile.Name())
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	if string(data) != testdata {
		t.Errorf("WriteFile: Reading back data: Got '%s', want '%s'", data, testdata)
	}
}

func cleanup(path string) {
	err := os.Remove(path)
	if err != nil {
		panic(err)
	}
}
