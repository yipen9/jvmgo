package main

import "flag"
import "fmt"
import "os"

type Cmd struct {
	helpFlag bool
	versionFlag bool
	cpOptions string
	class string
	args	[]string
}


//定义* Cmd的指针，&表示是对应变量的内存地址
func parseCmd() * Cmd {
	cmd := &Cmd{}	//定义变量cmd是结构体Cmd的内存地址
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag,"help",false, "Print help message")  //指针变量的存储地址
	flag.BoolVar(&cmd.helpFlag,"?" , false, "Print help message")
	flag.BoolVar(&cmd.versionFlag,"version",false, "print version and exit")
	flag.StringVar(&cmd.cpOptions,"classpath","", "classpath")
	flag.StringVar(&cmd.cpOptions,"cp","", "classpath")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

//定义printUsage，打印输入第0个参数
func printUsage()  {
	fmt.Printf("Usage: %s [-options] class [args ....]\n",os.Args[0])
}