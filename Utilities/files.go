package Utilities

import "os"

func OpenFile(fileName string) []byte {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	bytesRead, err := file.Read(buffer)
	if err != nil {
		panic(err)
	}

	if int64(bytesRead) != fileSize {
		panic("file read error")
	}

	return buffer
}
