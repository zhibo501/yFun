package main

import (
	"log"
	//config
	"os"
	"path/filepath"
	//go filter
	"regexp"
	//ast
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	cfg, _ := getConfig()
	cfg.fileWalk(goDep)
}

//================================ go filter ================================
func goDep(path string, info os.FileInfo, err error) error {
	n := info.Name()

	if info.IsDir() {
		if ".git" == n {
			return filepath.SkipDir
		}
		return nil
	} else {
		ok, err := regexp.MatchString(".*\\.go$", n)
		if err != nil {
			return err
		} else if !ok {
			return nil
		}
	}
	//log.Printf("INFO file path : %s, name : %s\n", path, n)
	log.Printf("%s , %s\n", n, path)
	return goAst(path)
}

//================================ go ast ================================
func goAst(path string) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
	if err != nil {
		log.Printf("ERROR %v \n", err)
		return err
	}

	var v visitor
	ast.Walk(v, f)

	return nil
}

type visitor struct{}

func (v visitor) Visit(n ast.Node) ast.Visitor {
	if nil == n {
		return nil
	}

	//log.Printf("%T\n", n)
	switch d := n.(type) {
	case *ast.ImportSpec:
		log.Printf("\t%s\n", d.Path.Value)
	}

	return v
}

//================================ config ================================
type config struct {
	fs []*file
}
type file struct {
	path string
	info os.FileInfo
}

func (c *config) addFile(path string) {
	fi, err := os.Stat(path)
	if err != nil {
		log.Printf("WARNING ignor file[%s]. %v \n", path, err)
		return
	}

	c.fs = append(c.fs, &file{path, fi})
	log.Printf("INFO add file[%s].\n", path)
}
func (c *config) fileWalk(fn func(string, os.FileInfo, error) error) {
	for _, f := range c.fs {
		//fn(f.path, f.info, nil)
		filepath.Walk(f.path, fn)
	}
}

func getConfig() (*config, error) {
	cfg := &config{}

	if len(os.Args) < 2 {
		log.Printf("No para ...\n")
		os.Exit(1)
	}

	for _, arg := range os.Args[1:] {
		cfg.addFile(arg)
	}

	return cfg, nil
}
