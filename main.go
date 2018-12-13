package main

import (
	"flag"
	"fmt"
	"git.oschina.net/jscode/go-package-plantuml/codeanalysis"
	"os"
	"path"
)

var (
	codedir    = flag.String("codedir", "", "分析目标目录")
	outputfile = flag.String("outputfile", "", "分析结果保存到该文件")
	ignoredir  arrayFlags
)

func main() {
	flag.Var(&ignoredir, "", "不需要进行代码分析的目录")
	flag.Parse()
	parseCodeToTxt() // 第一步解析
	generateUmlPng()
}

func parseCodeToTxt() {
	if *codedir == "" {
		*codedir, _ = os.Getwd() //取当前目录
	}
	if *outputfile == "" {
		dir, _ := os.Getwd()
		*outputfile = path.Join(dir, "uml.txt") //取当前目录
	}
	gopath, _ := os.LookupEnv("GOPATH") // gopath

	config := codeanalysis.Config{
		CodeDir:    *codedir,
		GopathDir:  gopath,
		VendorDir:  path.Join(*codedir, "vendor"),
		IgnoreDirs: ignoredir,
	}

	result := codeanalysis.AnalysisCode(config)

	result.OutputToFile(*outputfile)
}

func generateUmlPng() {
	_, err := os.Stat(*outputfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	gopath, _ := os.LookupEnv("GOPATH") // gopath
	plantuml := path.Join(gopath, "src/github.com/learnergo/uml", "plantuml.jar")

	if execCommand("java", "-jar", plantuml, *outputfile) {
		dir, _ := os.Getwd()
		if execCommand("mv", "/tmp/uml.png", dir) {
			fmt.Println("成功！")
			os.Exit(0)
		}
	}
	os.Exit(1)
}
