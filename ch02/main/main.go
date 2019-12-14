package main

import (
	"MyJVM/ch02/classpath"
	"fmt"
	"strings"
)

func main() {

	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) ([]byte, error) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath: %v class: %v args:%v \n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find load main class %s \n", cmd.class)
		return nil, err
	}
	fmt.Printf("class data:%v\n", classData)
	return classData, nil
}
