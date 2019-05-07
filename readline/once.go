package readline

import (
	"bufio"
	// "fmt"
	"io"
	"os"
	"strings"
)

//【带注释#】一次读取，存放到字符串切片中
func GetUsable(filename *string) (*[]string, error) {
	var slice []string //注意：此处不可分配空间
	f, err := os.Open(*filename)
	if err != nil {
		return &slice, err
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		line = strings.TrimSpace(line) //去除两头空格
		//fmt.Println("line", len(line))
		//过滤空
		if len(line) == 0 {
			continue
		}
		//过滤带注释的行
		if strings.HasPrefix(line, "#") {
			continue
		}
		slice = append(slice, line)
	}
	return &slice, nil
}

//【不带注释】一次读取，存放到字符串切片中
func GetAll(filename *string) (*[]string, error) {
	var slice []string //注意：此处不可分配空间
	f, err := os.Open(*filename)
	if err != nil {
		return &slice, err
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		line = strings.TrimSpace(line)
		// fmt.Println("line", len(line))
		if len(line) == 0 {
			continue
		}
		slice = append(slice, line)
	}
	return &slice, nil
}
