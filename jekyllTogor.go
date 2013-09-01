package main

import (
	"fmt"
	//"io"
	"io/ioutil"
	"os"
	"path/filepath"
	//"bytes"
)

func main() {
	if len(os.Args) > 1 {
		err := Tree(os.Args[1], 1)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Finsh")

}
func Tree(dirname string, curHier int) error {
	dirAbs, err := filepath.Abs(dirname)
	if err != nil {
		return err
	}
	fileInfos, err := ioutil.ReadDir(dirAbs)
	if err != nil {
		return err
	}
	//fileNum := len(fileInfos)
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			Tree(filepath.Join(dirAbs, fileInfo.Name()), curHier+1)
		} else {
			err := Dealwith(filepath.Join(dirAbs, fileInfo.Name()))
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func Dealwith(fpath string) error {
	file, err := os.OpenFile(fpath, os.O_RDWR, 0666)
	if err != nil {
		return err

	}
	defer file.Close()
	s := []byte("Hello World!")
	file.Write(s)
	return nil
}
