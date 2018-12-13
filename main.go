package main

import (
	"flag"
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
