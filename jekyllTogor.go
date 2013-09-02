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
	"errors"
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

//When this md isn't Blog will return errNoMeta
var errNoMeta = errors.New("Don't find meta data")

// the num of meta data
var numMeta int

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

//get gor.date from filename
func (g *Gor)getDate(filepath string){
	paths := strings.Split(filepath, "/")
	fname := paths[len(paths) - 1]
	re := regexp.MustCompile(`\d+-\d+-\d+`)
	data := re.FindString(fname)
	//fmt.Printf("%q\n", data)
	g.date= data

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
					if err == errNoMeta{
						fmt.Println(err.Error())
					}else{
						return err
					}
					
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
	

	//Creat tmp file
	outname := file.Name() + ".tmp"
	fmt.Println(outname)
	fout, err := os.Create(outname)
	if err != nil {
		fmt.Println(outname, err)
		return err
	}
	//Close the src file ,operate the tmp file
	file.Close()
	
	defer fout.Close()

	// let jekyll's meta data to gor's meta data
	content, _ := ioutil.ReadFile(file.Name())
	lines := strings.Split(string(content), "\n")
	numMeta,err = parseJekyll(lines)
	if err != nil{
		return err
	}
	jekyllToGor(file.Name())	


	//	file.Write()
	err = writeFile(fout,lines)
	if err != nil{
		return err
	}
	return nil
}

func parseJekyll(lines []string)(int,error){
	ResetJekyll(myJekyll)
	for i,line := range lines{
		//fmt.Println(line)
		if strings.Contains(line, metaEOF){
			return i,nil
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
	return 0,errNoMeta
}

func jekyllToGor(filename string){
	ResetGor(myGor)
	myGor.layout = myJekyll.layout
	myGor.title = myJekyll.title
	myGor.categories[0] = myJekyll.category
	myGor.tags = myJekyll.tags
	myGor.tagline = myJekyll.tagline
	myGor.getDate(filename)
	//fmt.Println(myGor)
}

//There use refelect should be more pithy
func writeFile(file *os.File,lines []string) error{
	file.Write([]byte("---\n"))
	file.Write([]byte("date: " + myGor.date + "\n"))
	file.Write([]byte("layout: " + myGor.layout + "\n"))
	file.Write([]byte("title: " + myGor.title + "\n"))
	file.Write([]byte("tagline: " + myGor.tagline + "\n"))
	file.Write([]byte("categories:\n"))
	file.Write([]byte("- " + myGor.categories[0] + "\n"))
	file.Write([]byte("tags:\n"))
	for _,tag := range myGor.tags{
		file.Write([]byte("- " + tag + "\n"))
	}
	file.Write([]byte("---\n"))
	for _,line := range lines[numMeta + 1:]{
		file.Write([]byte(line + "\n"))
	}
	return nil
}