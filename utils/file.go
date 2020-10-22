package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
)

//GetSize：获取文件大小
//GetExt：获取文件后缀
//CheckExist：检查文件是否存在
//CheckPermission：检查文件权限
//IsNotExistMkDir：如果不存在则新建文件夹
//MkDir：新建文件夹
//Open：打开文件

func GetSize(file multipart.File) (int, error) {
	content, err := ioutil.ReadAll(file)
	return len(content), err
}

func GetExt(fileName string) string {
	return path.Ext(fileName)
}

func CheckNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

//返回一个布尔值说明该错误是否表示因权限不足要求被拒绝。
//ErrPermission和一些系统调用错误会使它返回真。
func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

func IsNotExistMkDir(src string) error {
	if exist := CheckNotExist(src); exist == true {
		if err := Mkdir(src); err != nil {
			return err
		}
	}
	return nil
}

func Mkdir(filepath string) error {
	if _, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(filepath, os.ModePerm); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	file, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// MustOpen maximize trying to open the file
func MustOpen(fileName, filePath string) (*os.File, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := path.Join(pwd, filePath)
	file := path.Join(pwd, filePath, fileName)
	perm := CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := Open(file, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}

/*************************************************/
//创建文件 此处传入相对路径 public/xxx
func CreateFile(dir, filename string) (string, error) {
	file := path.Join(dir, filename) // **pwd/{{dir}}/{{name}}

	_, err := os.Stat(file)
	if !os.IsNotExist(err) {
		return file + "文件已经存在", err
	}

	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		if os.IsPermission(err) {
			return "权限不足以创建文件", err
		}
		return "其他错误", err
	}

	if _, err = os.Create(file); err != nil {
		return "创建文件失败", err
	}
	return file, nil
}

//写入文件 此处传入相对路径 public/xxx
func WriteFile(dir, filename, content string) error {
	file, err := MustOpen(filename, dir)
	if err != nil {
		return fmt.Errorf("打开文件错误：%s", err.Error())
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write([]byte(content))
	if err != nil {
		return fmt.Errorf("写入文件错误：%s", err.Error())
	}
	return writer.Flush()
}

//使用io.Copy
func CopyFile(src, des string) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	//获取源文件的权限
	fi, _ := srcFile.Stat()
	perm := fi.Mode()

	//desFile, err := os.Create(des)  //无法复制源文件的所有权限
	desFile, err := os.OpenFile(des, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm) //复制源文件的所有权限
	if err != nil {
		return 0, err
	}
	defer desFile.Close()

	return io.Copy(desFile, srcFile)
}

func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

func CopyDir(srcPath, desPath string) error {
	//检查目录是否正确
	if srcInfo, err := os.Stat(srcPath); err != nil {
		return err
	} else {
		if !srcInfo.IsDir() {
			return errors.New("源路径不是一个正确的目录！")
		}
	}

	if desInfo, err := os.Stat(desPath); err != nil {
		return err
	} else {
		if !desInfo.IsDir() {
			return errors.New("目标路径不是一个正确的目录！")
		}
	}

	if strings.TrimSpace(srcPath) == strings.TrimSpace(desPath) {
		return errors.New("源路径与目标路径不能相同！")
	}

	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		//复制目录是将源目录中的子目录复制到目标路径中，不包含源目录本身
		if path == srcPath {
			return nil
		}

		//生成新路径
		destNewPath := strings.Replace(path, srcPath, desPath, -1)

		if !f.IsDir() {
			CopyFile(path, destNewPath)
		} else {
			if !FileIsExisted(destNewPath) {
				return Mkdir(destNewPath)
			}
		}
		return nil
	})

	return err
}
