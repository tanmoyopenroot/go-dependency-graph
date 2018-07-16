package main

import (
	"strings"
	"fmt"
)

func showSubDeps(pkgName string, padding int, lastPkg bool) {
	if _, pkgExist := pkgDeps[pkgName]; pkgExist {
		var prefix string
		if lastPkg {
			prefix = outputPrefixLast
		} else {
			prefix = outputPrefix
		}

		fmt.Println(strings.Repeat(outputPadding, padding), prefix, pkgName)
		
		pkgs := pkgDeps[pkgName]
		for index, subPack := range pkgs {
			showSubDeps(subPack, padding + 1, index == len(pkgs)-1)
		}
	}
}

func showDeps(pkgName string) {
	if _, pkgExist := pkgDeps[pkgName]; pkgExist {
		fmt.Println(pkgName)
		padding := 0
		pkgs := pkgDeps[pkgName]
		for index, subPack := range pkgs {
			showSubDeps(subPack, padding, index == len(pkgs)-1)
		}
	} else {
		fmt.Println("Uable to find %s in dependency list", pkgName)
	}
}