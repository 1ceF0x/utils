package utils

import (
	"bufio"
	utils "github.com/1ceF0x/utils/rand"
	"io"
	"log"
	"os"
	"sync"
)

type Syncfile struct {
	mutex     *sync.Mutex
	iohandler *os.File
}

func NewSyncFile(filename string) (*Syncfile, error) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &Syncfile{mutex: &sync.Mutex{}, iohandler: f}, nil
}

func (sf *Syncfile) Write(content string) {
	sf.mutex.Lock()

	wbuf := bufio.NewWriterSize(sf.iohandler, len(content))
	wbuf.WriteString(content)
	wbuf.Flush()

	utils.RandSleep(1000)

	sf.mutex.Unlock()
}

// 判断文件是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
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
		line := scanner.Text() //获取每一行
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
		log.Println("读取文件失败: ", err)
		return nil, err
	}
	return data, nil
}

// 追加写入文件
func BufferWriteAppend(fileName string, param string) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_SYNC, 0660)
	if err != nil {
		return err
	}
	defer file.Close()
	// NewWriter 默认缓冲区大小是 4096
	// 需要使用自定义缓冲区的writer 使用 NewWriterSize()方法
	buf := bufio.NewWriter(file)
	// 字节写入
	//buf.Write([]byte(param))
	// 字符串写入
	buf.WriteString(param + "\n")
	// 将缓冲中的数据写入
	return buf.Flush()
}
