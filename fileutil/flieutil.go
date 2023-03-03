package fileutil

import (
	"bufio"
	"github.com/1ceF0x/utils/randutil"
	"io"
	"os"
	"sync"
)

type Syncfile struct {
	mutex     *sync.Mutex
	ioHandler *os.File
}

func NewSyncFile(filename string) (*Syncfile, error) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &Syncfile{mutex: &sync.Mutex{}, ioHandler: f}, nil
}

func (sf *Syncfile) Write(content string) {
	sf.mutex.Lock()
	wBuf := bufio.NewWriterSize(sf.ioHandler, len(content))
	wBuf.WriteString(content)
	wBuf.Flush()
	randutil.RandSleep(1000)
	sf.mutex.Unlock()
}

// 判断文件是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// 判断是否是目录
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 读取文件每行内容
func ReadFileLines(fileName string) ([]string, error) {
	var result []string
	file, err := os.Open(fileName)
	if err != nil {
		return result, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result, err
}

// 读取文件返回scanner
func ReadFileScanner(fileName string) (*os.File, *bufio.Scanner, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return file, scanner, nil
}

// 读取文件内容
func ReadFileBytes(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 追加写入文件
func WriteFileAppend(fileName string, content string) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_SYNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	buf := bufio.NewWriter(file)
	buf.WriteString(content)
	return buf.Flush()
}

// 覆盖写入文件
func WriteFile(fileName string, content string) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_SYNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	buf := bufio.NewWriter(file)
	buf.WriteString(content)
	return buf.Flush()
}
