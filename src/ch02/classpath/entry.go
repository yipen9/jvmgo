package classpath

import "os"
import "strings"

//定义常量
const  pathListSeparator = string(os.PathListSeparator)

//定义接口,按照惯例，使用最后一个返回值作为错误信息
type Entry interface {
	//方法返回多个参数
	readClass(classname string) ([]byte,Entry, error)
	//返回变量的字符串
	String() string

}

func newEntry(path string) Entry  {
	if strings.Contains(path, pathListSeparator) {
		//返回compositeEntry
	}

	if strings.HasSuffix(path, "*") {
		//返回wildcard
	}

	if strings.HasSuffix(path,".jar") || strings.HasSuffix(path,".JAR") ||
		strings.HasSuffix(path,".jar") || strings.HasSuffix(path,".JAR"){
		// newZip

	}
}

