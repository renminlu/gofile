//功能：【单文件压缩】和【目录递归压缩】
//
package zip

import (
	"archive/zip"
	"errors"

	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"qqqkkk/file/fileFunc"
)

type Wzip struct {
	handle          *zip.Writer
	rootPath        string
	keepFile        string   //保存的路径文件名，不用带后缀zip
	SourceFilefirst string   //第一个压缩源
	sourceFileList  []string //待压缩文件列表

	invalidFileList []string //无效的异常的待压缩列表
}

//SourceFilefirst：待压缩的代表文件
//keepFile：保存的路径文件名，不用带后缀zip
//如果keepFile已存在，会覆盖

//压缩单个目标
func Zip(SourceFile string, keepFile string) error {
	wz := Wzip{}
	wz.keepFile = keepFile
	wz.SourceFilefirst = SourceFile
	ok := wz.Begin()
	if ok != nil {
		return ok
	}
	ok = wz.MakeSourceFileList(SourceFile)
	if ok != nil {
		return ok
	}
	for _, v := range wz.sourceFileList {
		fmt.Println(v)
		ok := wz.MakeYasuo(v)
		if ok != nil {
			return ok
		}
	}
	defer wz.Close()

	return nil
}

//压缩多个目标
func Zips(SourceFile []string, keepFile string) error {
	return nil
}

//压缩源：预处理 beforehand
func (w *Wzip) Begin() error {

	//获取&制作：当前压缩项目的源文件的根目录
	w.rootPath = filepath.Dir(w.SourceFilefirst)

	//生成一个压缩对象
	buf, ok := os.OpenFile(w.keepFile+".zip", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if ok != nil {
		return ok
	}

	//func NewWriter(w io.Writer) *Writer
	w.handle = zip.NewWriter(buf) // 实例一个写对象

	return nil
}

//
//制作生成：待压缩资源文件列表
func (w *Wzip) MakeSourceFileList(file string) error {

	//判断：既不是目录也不是文件：则在此退出
	if !fileFunc.FileExist(file) {
		return errors.New("fileFunc.FileExist(" + file + ") is false")
	}
	//如果该目标是文件：直接加入列表
	if fileFunc.IsFile(file) {
		w.sourceFileList = append(w.sourceFileList, file)
		return nil
	}

	ok := filepath.Walk(
		file,
		func(path string, f os.FileInfo, err error) error {

			//处理无效文件列表（= nil表示：异常存在的文件）可能是非法的文件名导致的
			if f == nil {
				w.AddInvalid(path) //添加到无效列表

				return nil //retrun err会退出循环遍历，nil不会退出循环
			}

			w.sourceFileList = append(w.sourceFileList, path)

			return nil
		})
	return ok
}

//压缩处理
func (w *Wzip) MakeYasuo(pfilename string) error {

	//制作该文件：相对路径
	//func Rel(basepath, targpath string) (string, error)
	createname, ok := filepath.Rel(w.rootPath, pfilename)
	if ok != nil {
		w.AddInvalid(pfilename) //添加到无效列表
		return ok
	}

	//判断：如果是目录：则在此退出
	if fileFunc.IsDir(pfilename) {

		//判断：如果不是空目录：则在此退出
		dir, _ := ioutil.ReadDir(pfilename)
		if len(dir) > 0 {
			return nil
		}

		//创建空目录
		// finfo, _ := os.Stat(createname)
		// fih, _ := zip.FileInfoHeader(finfo)
		// fih.SetMode(os.ModeDir)
		_, err := w.handle.Create(createname)
		if err != nil {
			w.AddInvalid(pfilename) //添加到无效列表
			return err
		}
		return nil
	}

	//在压缩包中创建这个文件
	//本方法返回一个io.Writer接口（用于写入新添加文件的内容）
	//文件名必须是相对路径，不能以设备或斜杠开始，只接受'/'作为路径分隔
	f, err := w.handle.Create(createname)
	if err != nil {
		w.AddInvalid(pfilename) //添加到无效列表
		return err
	}

	//读取文件
	content, err2 := ioutil.ReadFile(pfilename) //不适合读大文件
	if err2 != nil {
		w.AddInvalid(pfilename) //添加到无效列表
		return err2
	}
	_, err = f.Write(content) //注意参数：[]byte
	if err != nil {
		w.AddInvalid(pfilename) //添加到无效列表
		return err
	}

	return nil
}

//关闭new创建的资源
func (w *Wzip) Close() error {
	return w.handle.Close()
}

//压缩失败时：删除已压缩的文件
func (w *Wzip) Del() error {
	return nil
}

//添加：无效的压缩目标
func (w *Wzip) AddInvalid(pf string) {
	w.invalidFileList = append(w.invalidFileList, pf)
}
