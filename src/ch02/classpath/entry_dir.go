package classpath

import "io/ioutil"
import "path/filepath"


//Go结构体不需要显示的实现接口，只要方法匹配即可
type DirEntry struct {
	absDir string
}


func newDirEntry(path string) *DirEntry {
	absDir , err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}


func (self *DirEntry) readClass(className string) ([]byte , Entry , error){
	fileName := filepath.Join(self.absDir, className)
	data , err := ioutil.ReadFile(fileName)
	return data , self , err
}

func (self *DirEntry) String() string {
	return self.absDir
}




