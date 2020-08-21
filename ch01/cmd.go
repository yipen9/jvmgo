package main

import "flag"
import "fmt"
import "os"

type Cmd struct {
	helpFlag bool
	versionFlag bool
	cpOption string
	class string
	args	[]string
}


func parseCmd() *Cmd{
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag,"help",false, "print help message")
	flag.BoolVar(&cmd.helpFlag,"?",false, "print help message")
	flag.BoolVar(&cmd.versionFlag,"version",false, "print version and exit")
	flag.StringVar(&cmd.cpOption,"classpath","", "classpath")
	flag.StringVar(&cmd.cpOption, "cp","", "classpath")
	flag.Parse()
	args := flag.Args()	//如果解析成功，调用flag.Args()函数可以捕获其他没有被解析的参数
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}


func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n",os.Args[0])
}