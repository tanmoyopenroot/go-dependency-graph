package main

import (
	"os"
	"fmt"
	"bytes"
	"strings"
	"text/template"
)

func WriteDot(g DotGraph) {
	file := strings.Join([]string{g.Title, ".dot"}, "")
	f, err := os.Create(file)
	if err != nil {
		fmt.Println("Unable to create %s file", file)
		return
	}

	t := template.Must(template.New("dot.template").ParseFiles("template/dot.template"))
	err = t.Execute(f, g)
	if err != nil {
		fmt.Println("Error while excuting the graph template")
	} 

	f.Close()
}


func concatDeps(pkgName string) string {
	var buffer bytes.Buffer
	buffer.WriteString("\"")
	buffer.WriteString(pkgName)
	buffer.WriteString("\" -> { ")

	for _, dep := range pkgDeps[pkgName] {
		buffer.WriteString("\"")
		buffer.WriteString(dep)
		buffer.WriteString("\" ")
	}

	buffer.WriteString("} ")

	return buffer.String()
}

func processSubGraph(pkgName string) string {
	var buffer bytes.Buffer
	pkgs := pkgDeps[pkgName]

	if len(pkgs) < 1 {
		return ""
	}

	if _, pkgExist := graphList[pkgName]; pkgExist {
		return ""
	}

	graphList[pkgName] = true

	buffer.WriteString(concatDeps(pkgName))
	for _, subPack := range pkgs {
		buffer.WriteString(processSubGraph(subPack))
	}
	return buffer.String()
}

func processGraph(pkgName string) {
	if _, pkgExist := pkgDeps[pkgName]; pkgExist {
		pkgs := pkgDeps[pkgName]
		graphList = map[string]bool{}
		var buffer bytes.Buffer

		graphList[pkgName] = true

		buffer.WriteString(concatDeps(pkgName))
		for _, pkg := range pkgs {
			buffer.WriteString(processSubGraph(pkg))	
		}

		// fmt.Println(buffer.String())
		dot := DotGraph{
			Title: pkgName, 
			DepsPath: buffer.String(),
		}

		WriteDot(dot)
	} else {
		fmt.Println("Uable to find %s in dependency list", pkgName)
	}
}