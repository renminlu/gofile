package writeLines

import (
	"github.com/renminlu/goconst"
)

func (w *WriteLine) Write(str string) {
	str += goconst.EOL
	lock.Lock() //加锁
	w.osFile.WriteString(str)
	lock.Unlock() //解锁
	// fmt.Println("writeLines调试：", str)
}
