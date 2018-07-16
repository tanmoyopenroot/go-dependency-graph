package main

import (
	"os"
	"fmt"
	"flag"
	"go/build"

	"github.com/tanmoyopenroot/go-dependency-graph/cmd/const/main"
	"github.com/tanmoyopenroot/go-dependency-graph/cmd/dot-graph/main"
	"github.com/tanmoyopenroot/go-dependency-graph/cmd/generate-dependency-graph/main"
	"github.com/tanmoyopenroot/go-dependency-graph/cmd/show-dependency-tree/main"
)

var (
	showStdLib = flag.Bool("show-std", false, "Show dependencies os Standard Library")
	depLevel = flag.Int("level", 10, "Dept of Dependency Graph")

	pkgList = map[string]bool{}
	graphList = map[string]bool{}
	pkgDeps = make(map[string][]string)
	buildContext = build.Default
)

func getImports(pkg *build.Package) []string{
	allImports := pkg.Imports
	// fmt.Println("All Imports", allImports)
	return allImports
}

func processEachPackage(dir string, pkgName string) error {
	// fmt.Println("Directory: ", dir)
	// fmt.Println("Current Package Processing: ", pkgName)

	pkg, err := buildContext.Import(pkgName, dir, 0)
	if err != nil {
		return fmt.Errorf("Failed to import: %s", pkgName)
	}

	pkgList[pkgName] = true

	if pkg.Goroot && !*showStdLib {
		return nil
	}

	pkgImports := getImports(pkg)
	pkgDeps[pkgName] = pkgImports

	for _, subPack := range pkgImports  {
		if _, pkgExist := pkgList[subPack]; !pkgExist {
			if err := processEachPackage(pkg.Dir, subPack); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Invalid Arguments!")
	} else {
		fmt.Println("Arguments: ", args)
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current working directory")
	} 



	for _, pkgName := range args {
		err := processEachPackage(cwd, pkgName)
		if err != nil {
			fmt.Println("Error while processing the: ", pkgName, err)
		}
	}

	// fmt.Println("Package Details:")
	// fmt.Println(pkgDeps)


	for _, pkgName := range args {
		showDeps(pkgName)
		processGraph(pkgName)
	}
}