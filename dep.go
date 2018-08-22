package main

import (
	"fmt"
	"regexp"

	_ "github.com/ServiceComb/go-chassis/core/config"
	_ "github.com/ServiceComb/go-chassis/core/registry"
	_ "github.com/coreos/go-systemd/daemon"
)

var str = "jakee and zz done.<p>zzz</p> <p style=\"text-align: center;\" align=\"justify\">==TUIGUANG===</p>\n<p align=\"justify\"><img style=\"width: auto;\" src=\"https://imag.xxxx.com/b931035c-5bf2-4553-9593-07480f74a5c6.png\" /></p> this should not dispeared <p>ALL MY ...</p> <p > See you</p> behan"

func main() {
	fmt.Printf("hello world!\n")

	ok, err := regexp.MatchString("<p[^<]*TUIGUANG.*</p>", str)
	fmt.Printf("ok : %v \nerr : %v\n", ok, err)

	reg := regexp.MustCompile(`<p[^<]*TUIGUANG.*</p>`)
	fmt.Println(reg.FindString(str))
}
