package main

import (
	"strings"
	"fmt"
)

func showSubDeps(pkgName string, padding int, lastPkg bool, level int) {
	if _, pkgExist := pkgDeps[pkgName]; pkgExist {
		if (level != -1 && level < 1) {
			return
		} 

		if (level != -1) {
			level = level - 1
		}

		var prefix string
		if lastPkg {
			prefix = outputPrefixLast
		} else {
			prefix = outputPrefix
		}

		fmt.Println(strings.Repeat(outputPadding, padding), prefix, pkgName)
		
		pkgs := pkgDeps[pkgName]
		for index, subPack := range pkgs {
			showSubDeps(subPack, padding + 1, index == len(pkgs)-1, level)
		}
	}
}

// ShowGoDeps ... Show the Dependency in CMD
func ShowGoDeps(pkgName string, level int) {
	if _, pkgExist := pkgDeps[pkgName]; pkgExist {
		fmt.Println(pkgName)
		padding := 0
		pkgs := pkgDeps[pkgName]
		for index, subPack := range pkgs {
			showSubDeps(subPack, padding, index == len(pkgs)-1, level)
		}
	}
}