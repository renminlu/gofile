package writeline

import (
	// "fmt"
	"os"
	"sync"
)

type WriteLine struct {
	osFile *os.File
}

var (
	single map[string]*WriteLine //单列
	lock   sync.Mutex            //声明一个全局的互斥锁【lock】
)

func init() {
	single = make(map[string]*WriteLine, 2)
}
