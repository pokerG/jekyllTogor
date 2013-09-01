package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// octopress 转 gor 用的小工具，用来在头部插入 date 标签
// 用法： mkdir tmp
//        blogtrans *.md
// 运行后，会在tmp目录下生成插入了 date标签的md 文件。
// 其实，这个用 sed 什么的秒杀就做了，但是为了练手 go 语言，所以。。
// 第一个go程序，感觉还不顺手，还需要慢慢体会。

func main() {
	// 没有检查，直接读参数
	for _, fname := range os.Args[1:] {
		content, err := ioutil.ReadFile(fname)
		if err != nil {
			fmt.Println(fname, err)
			return
		}

		outname := "tmp/" + fname
		fout, errout := os.Create(outname)
		if errout != nil {
			fmt.Println(outname, err)
			return
		}
		defer fout.Close()

		lines := strings.Split(string(content), "\n")
		fout.Write([]byte(lines[0] + "\n"))
		fout.Write(insert_head(fname))

		for _, line := range lines[1:] {
			output, result := change_img(line)
			if result < 0 {
				output, result = change_img_legacy(line)
			}
			fout.Write(output)
		}
	}
}

func insert_head(fname string) (head []byte) {

	tmp := "date: "
	re := regexp.MustCompile(`\d+-\d+-\d+`)
	data := re.FindString(fname)
	tmp += data + "\n"
	fmt.Printf("%q\n", tmp)
	head = []byte(tmp)

	return
}

func change_img(line string) ([]byte, int) {

	re := regexp.MustCompile(`{% *img */rc/(\S+) (\d*) *(\d*) *\"(.*?)\" *\"(.*?)\" *%}`)
	result := re.FindStringSubmatch(line)
	if len(result) < 6 {
		return []byte(line + "\n"), -1
	}

	var height, width string
	h, herr := strconv.Atoi(result[2])
	w, werr := strconv.Atoi(result[3])
	if herr == nil && werr == nil {
		height = fmt.Sprintf("height=\"%d\"", h)
		width = fmt.Sprintf("width=\"%d\"", w)
	} else if herr == nil {
		width = fmt.Sprintf("width=\"%d\"", h) // 第一个参数作为width使用
	} else {
		width = fmt.Sprintf("width=\"%d\"", 600) // 第一个参数作为width使用
	}

	tmp := fmt.Sprintf("<img class=\"imgbox\" src=\"http://iwood.qiniudn.com/media/%s\" alt=\"%s\" %s %s >%s</img>",
		result[1], result[4], height, width, result[5])
	println(tmp)

	return []byte(tmp + "\n"), 1
}

func change_img_legacy(line string) ([]byte, int) {

	re := regexp.MustCompile(`{% *img */rc/(\S+) +(\d*) *%}`)
	result := re.FindStringSubmatch(line)
	if len(result) < 3 {
		return []byte(line + "\n"), -1
	}
	i, _ := strconv.Atoi(result[2])
	width := fmt.Sprintf("width=\"%d\"", i)
	if result[2] == "" {
		width = ""
	}

	tmp := fmt.Sprintf("<img class=\"imgbox\" src=\"http://iwood.qiniudn.com/media/%s\" %s ></img>", result[1], width)
	println(tmp)

	return []byte(tmp + "\n"), 1
}
