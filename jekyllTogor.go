package main

import (
	"fmt"
	//"io"
	"io/ioutil"
	"os"
	"path/filepath"
	//"bytes"
	"regexp"
)
type Jekyll struct{
	layout string
	title string
	category string
	tagline string
	tags []string
}

type Gor struct{
	layout string
	date string
	title string
	//permalink string
	categories []string
	tags []string
}
func Newjekyll() *Jekyll{
	return &Jekyll{layout:"post"}
}
func NewGor() *Gor{
	return &Gor{layout:"post"}
}
var myJekyll *Jekyll
var myGor *Gor
func main() {
	if len(os.Args) > 1 {
		myJekyll = Newjekyll()
		myGor = NewGor()
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
			b := []byte(fileInfo.Name())
			fmt.Println(fileInfo.Name())
			matched,_ := regexp.Match("[.](md)$",b )
			fmt.Println(matched)
			if matched{
				err := Dealwith(filepath.Join(dirAbs, fileInfo.Name()))
				if err != nil {
					return err
				}
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
