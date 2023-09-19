package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Problem struct {
	Num       int32
	Question  string
	QUrl      string
	Answer    string
	ClassType string
	Level     string
}

func main() {
	content, err := checkReadMe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	content += "| 题号 | 题目 | 题解 | 类别 | 难度 | 标识 |  \n"
	content += "|:---|:---|:---:|:---:|:---:|:---:|\n"

	solves := scanProblem()

	for _, s := range solves {
		content += fmt.Sprintf("|%d|[%v](%v)|[GO](%v)|%v|%v| |\n", s.Num, s.Question, s.QUrl, s.Answer, s.ClassType, s.Level)
	}

	ioutil.WriteFile("README_tmp.md", []byte(content), os.ModePerm)
}

func checkReadMe() (string, error) {
	all, err := ioutil.ReadFile("README.md")
	if err != nil {
		return "", err
	}

	data := string(all)
	content := strings.Split(data, "\n")

	var ret string

	for _, i := range content {
		if strings.Contains(i, "题目") && strings.Contains(i, "题解") && strings.Contains(i, "类别") && strings.Contains(i, "难度") {
			return ret, nil
		} else {
			ret += i + "\n"
		}
	}

	return "", errors.New("not title match")
}

func scanProblem() []*Problem {
	solve := make([]*Problem, 0)
	ret, _ := ioutil.ReadDir(".")
	for _, v := range ret {
		if v.IsDir() {
			ret, ok := checkPath(".", v.Name())
			if ok {
				solve = append(solve, ret...)
			}
		}
	}

	sort.Slice(solve, func(i, j int) bool {
		return solve[i].Num < solve[j].Num
	})

	return solve
}

func checkPath(father, path string) (ret []*Problem, ok bool) {
	ret = make([]*Problem, 0)
	filedir, _ := ioutil.ReadDir(father + "/" + path)
	for _, v := range filedir {
		if v.IsDir() {
			dirRet, ok := checkPath(father+"/"+path, v.Name())
			if ok {
				ret = append(ret, dirRet...)
			}
		} else if v.Name() == "p"+path+".go" {
			fileRet, err := checkFile(father, path)
			if err == nil {
				ret = append(ret, fileRet)
			}
		}
	}
	return ret, true
}

func checkFile(father, path string) (*Problem, error) {
	filedata, err := ioutil.ReadFile(father + "/" + path + "/README.md")
	if err != nil {
		return nil, err
	}

	result := &Problem{}
	result.Answer = strings.Trim(father+"/"+path, ".")

	num, _ := strconv.ParseInt(path, 10, 64)
	result.Num = int32(num)

	data := string(filedata)
	content := strings.Split(data, "\n")

	line1 := content[0]
	reg := regexp.MustCompile(`(.*?)\[(.*?)\]\((.*?)\)`)
	ret := reg.FindStringSubmatch(line1)
	if len(ret) >= 3 {
		result.Question = ret[2]
	}
	if len(ret) >= 4 {
		result.QUrl = ret[3]
	}

	if len(content) > 2 {
		result.ClassType = content[len(content)-2]
		result.Level = content[len(content)-1]
	}

	return result, nil
}
