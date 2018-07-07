package main

import (
	"fmt"
	"os"

	"pget"
)

func main() {

	//新建pget对象
	cli := pget.New()
	if err := cli.Run(); err != nil {
		if cli.Trace {
			fmt.Fprintf(os.Stderr, "Error:\n%+v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "Error:\n  %v\n", cli.ErrTop(err))
		}
		os.Exit(1)
	}

	//程序正确退出，想看fmt的时候，不用exit
	os.Exit(0)
}
