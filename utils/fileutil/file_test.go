package fileutil

import "testing"

func TestWriteToFileWithRename(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := WriteToFileWithRename("./download/1", "test.jpg", []byte("hello"))
		if err != nil {
			t.Fatal(err)
		}
	}

}
