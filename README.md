# go-dependency-graph

Purpose of this tool is to provide a visual dependency overview of your program.

## Install

```sh
sudo apt-get install graphviz
go get github.com/tanmoyopenroot/go-dependency-graph
```
### Command-Line

Simply execute main.go with one or more package names to visualize.

```sh
$ go run main.go -show-std=true strings
Arguments:  [strings]
strings
 ├ errors
 ├ io
   ├ errors
   └ sync
     ├ internal/race
       └ unsafe
     ├ runtime
       ├ runtime/internal/atomic
         └ unsafe
       ├ runtime/internal/sys
       └ unsafe
     ├ sync/atomic
       └ unsafe
     └ unsafe
 ├ unicode
 └ unicode/utf8
```

This generates a Graphviz's dot format file, to view execute the following: 
```sh
$ dot -Tpng strings.dot -o strings.png
$ eog strings.png
```

![picture](screenshot/strings.png)