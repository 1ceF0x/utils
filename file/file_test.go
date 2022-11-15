package file

import (
	"fmt"
	"strings"
	"testing"
)

func TestFile(t *testing.T) {
	fileHas := "rand_test.go"
	fileNot := "notexists.go"
	fmt.Println(Exists(fileHas))
	fmt.Println(Exists(fileNot))
	fmt.Println(IsDir("../file"))

	strs, _ := ReadFileLines(fileHas)
	for _, v := range strs {
		fmt.Println(v)
	}
}

func TestReadLines(t *testing.T) {
	strs, _ := ReadFileLines("rand_test.go")
	for _, v := range strs {
		fmt.Println(v)
	}
}

func TestReadBytes(t *testing.T) {
	bs, _ := ReadFileBytes("rand_test.go")
	fmt.Println(string(bs))
}

func TestReadScanner(t *testing.T) {
	file, scanner, _ := ReadFileScanner("rand_test.go")
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			fmt.Println(line)
		}
	}
	file.Close()
}

func TestBufferWriteAppend(t *testing.T) {
	strs := "thisistest"
	BufferWriteAppend("test.txt", strs)
}
