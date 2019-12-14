package classpath

import (
	"os"
	"path/filepath"
)

type ClassPath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClassPath Entry
}

func Parse(jarOption, cpOption string) *ClassPath {
	cp := &ClassPath{}
	cp.parseBootAndExtClasspath(jarOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *ClassPath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClassPath.readClass(className)
}

func (self *ClassPath) String() string {
	return self.bootClasspath.String()
}

func (self *ClassPath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//如果用户没有提供-classpath/-cp选项，则使用当前目录作为用户类路径
func (self *ClassPath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClassPath = newEntry(cpOption)
}
