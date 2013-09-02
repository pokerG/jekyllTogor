package main

import (
	"fmt"
	//"io"
	"io/ioutil"
	"os"
	"path/filepath"
	//"bytes"
	"regexp"
	"strings"
)

//Jekyll's Meta data
type Jekyll struct {
	layout   string
	title    string
	category string
	tagline  string
	tags     []string
}

//Gor's Meta data
type Gor struct {
	layout string
	date   string
	title  string
	//permalink string
	categories []string
	tags       []string
}

//Because my blogs have this sentence which don't need
//so I delete it
//In face the metaEoF should be the second "---"
var metaEOF string = "{% include JB/setup %}"

func Newjekyll() *Jekyll {
	return &Jekyll{layout: "post"}
}
func NewGor() *Gor {
	return &Gor{layout: "post"}
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

//list files under the dir
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
			//fmt.Println(fileInfo.Name())
			matched, _ := regexp.Match("[.](md)$", b)
			//fmt.Println(matched)
			if matched {
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

	//Creat tmp file
	outname := file.Name() + ".tmp"
	fmt.Println(outname)
	fout, err := os.Create(outname)
	if err != nil {
		fmt.Println(outname, err)
		return err
	}
	defer fout.Close()

	content, _ := ioutil.ReadFile(file.Name())
	lines := strings.Split(string(content), "\n")
	CreatJekyll(lines)

	//	file.Write(s)
	return nil
}

func CreatJekyll(lines []string) {
	//myJekyll =
	for _,line := range lines{
		//fmt.Println(line)
		if strings.Contains(line, metaEOF){
			return
		}
		meta := strings.Split(line, ":")
		switch{
		case strings.Contains(meta[0], "layout"):
			myJekyll.layout = meta[1]
		case strings.Contains(meta[0], "title"):
			myJekyll.title = meta[1]
		case strings.Contains(meta[0], "tagline"):
			myJekyll.tagline = meta[1]	
		case strings.Contains(meta[0], "category"):
			myJekyll.category = meta[1]
		case strings.Contains(meta[0], "tags"):
			meta[1] = strings.TrimSpace(meta[1])
			meta[1] = strings.TrimLeft(meta[1],"[")
			meta[1] = strings.TrimRight(meta[1],"]")
			fmt.Println(meta[1])
			data := strings.Split(meta[1], ",")
			myJekyll.tags = data
		}
	}
	fmt.Println(myJekyll)
}