package main

import (
	"testing"
)

func TestCmd(t *testing.T) {
	javaHome := "/Library/Java/JavaVirtualMachines/jdk1.8.0_162.jdk/Contents/Home/jre"

	cmd := parseCmd()
	cmd.XjreOption = javaHome
	cmd.class = "java.lang.Object"
	if cmd.versionFlag {
		t.Log("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
		t.Log("get ClassFile")
	}
}
