package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

var moduleNameParam = flag.String("module", "", "moduleName参数")
var modelNameParam = flag.String("model", "", "modelName参数")

func main() {

	flag.Parse()

	moduleName := *moduleNameParam
	modelName := *modelNameParam
	if len(moduleName) == 0 || len(modelName) == 0 {
		panic("参数不能为空")
	}

	prefix := "ctr"
	ctrTmpl := []string{"_i", "_proxy", ""}
	for _, tmpl := range ctrTmpl {
		targetFile := fmt.Sprintf("../%s/controller/%s.go", moduleName, strings.ToLower(modelName)+tmpl)
		parse(modelName, moduleName, prefix+tmpl, targetFile)
	}

	prefix = "repo"
	repoTmpl := []string{"_i", "_proxy", ""}
	for _, tmpl := range repoTmpl {
		targetFile := fmt.Sprintf("../%s/repository/%s.go", moduleName, strings.ToLower(modelName)+tmpl)
		parse(modelName, moduleName, prefix+tmpl, targetFile)
	}

	prefix = "serv"
	servTmpl := []string{"_i", "_proxy", ""}
	for _, tmpl := range servTmpl {
		targetFile := fmt.Sprintf("../%s/service/%s.go", moduleName, strings.ToLower(modelName)+tmpl)
		parse(modelName, moduleName, prefix+tmpl, targetFile)
	}

}

func parse(beanName string, moduleName string, tmplName string, targetFile string) {
	data := map[string]interface{}{
		"name":       beanName,
		"moduleName": moduleName,
	}

	tmpl, err := template.ParseFiles("./" + tmplName + ".tmpl")
	if err != nil {
		fmt.Println("template parsefile", tmplName, err)
		return
	}

	_, err = os.Stat(targetFile)
	if err == nil {
		fmt.Println("File exist", targetFile)
		return
	}

	target, err := os.Create(targetFile)
	if err != nil {
		fmt.Println("os.Create", targetFile, err)
		return
	}
	defer target.Close()

	tmpl.Execute(target, data)
}
