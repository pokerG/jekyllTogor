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
	tagline string
	categories []string
	tags       []string
}

//Because my blogs have this sentence which don't need
//so I delete it
//In face the metaEoF should be the second "---"
var metaEOF string = "{% include JB/setup %}"

func NewJekyll() *Jekyll {
	return &Jekyll{layout: "post"}
}
func NewGor() *Gor {
	return &Gor{layout: "post"}
}

func ResetJekyll(j *Jekyll){
		j.layout = "post"
		j.title = ""
		j.category = ""
		j.tagline = ""
		j.tags = []string{}
}

func ResetGor(g * Gor){
	g.layout = "post"
	g.title = ""
	g.date = ""
	g.categories = make([]string,1)
	g.tags = []string{}
	g.tagline = ""
}

func (g *Gor)getDate(filepath string){
	paths := strings.Split(filepath, "/")
	fname := paths[len(paths) - 1]
	fmt.Println(fname)

}
var myJekyll *Jekyll
var myGor *Gor

func main() {
	if len(os.Args) > 1 {
		myJekyll = NewJekyll()
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
	parseJekyll(lines)
	jekyllToGor(file.Name())	
	//	file.Write(s)
	return nil
}

func parseJekyll(lines []string) error{
	ResetJekyll(myJekyll)
	for _,line := range lines{
		//fmt.Println(line)
		if strings.Contains(line, metaEOF){
			return nil
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
			//fmt.Println(meta[1])
			data := strings.Split(meta[1], ",")
			myJekyll.tags = data
		}
	}
	return nil
}

func jekyllToGor(filename string){
	ResetGor(myGor)
	myGor.layout = myJekyll.layout
	myGor.title = myJekyll.title
	myGor.categories[0] = myJekyll.category
	myGor.tags = myJekyll.tags
	myGor.tagline = myJekyll.tagline
	myGor.getDate(filename)
}