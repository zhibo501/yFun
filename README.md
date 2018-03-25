# yFun

*findd.go* is used to find dependent package in go source files.

```bash
xxx$ ./findd hello demo *.go
WARNING ignor file[hello]. stat hello: no such file or directory 
INFO add file[demo].
INFO add file[findd.go].
INFO add file[rcu.go].
INFO add file[rcu_test.go].
main.go , demo/main.go
	"fmt"
	"github.com/zhibo501/yFun/ben"
	"sync"
findd.go , findd.go
	"log"
	"os"
	"path/filepath"
	"regexp"
	"go/ast"
	"go/parser"
	"go/token"
rcu.go , rcu.go
	"sync"
	"sync/atomic"
	"unsafe"
rcu_test.go , rcu_test.go
	"testing"
	"unsafe"
xxx$ ls
Makefile	ben		findd		rcu.go
README.md	demo		findd.go	rcu_test.go
xxx$ 
```

@2018.3.25


=============================

only support :
```bash
go test -race -bench . -benchtime 5s
```
@2017.9.10


==============================

do something funny ^_^

welcom~~~~~~

2015.11.7
