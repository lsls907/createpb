package createpb

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	files                  []string
	IgnoreDir              []string
	IgnoreReplaceOmitempty []string
)

func init() {
	IgnoreDir = append(IgnoreDir, "google", ".git", ".idea")
}

func Generate() error {
	exe := func(name string) error {
		cmd := exec.Command("cmd.exe", "/C", "protoc --go_out=plugins=grpc:. "+name)
		w := bytes.NewBuffer(nil)
		cmd.Stderr = w
		if err := cmd.Run(); err != nil {
			fmt.Println(fmt.Sprintf("generate %s pb.go file error，%s %s", name, err.Error(), string(w.Bytes())))
			return errors.New(string(w.Bytes()))
		}
		log.Println(fmt.Sprintf("success %s => %s.pb.go", name, strings.TrimRight(name, ".proto")))
		return nil
	}

	recursionReadFile(".")

	var err error
	for _, v := range files {
		if strings.HasSuffix(v, ".proto") {
			if err = exe(v); err != nil {
				break
			}
		}
	}

	if err == nil {
		for _, v := range files {
			if strings.HasSuffix(v, ".pb.go") && !isExistsArrary(v, IgnoreReplaceOmitempty) {
				if f, err := os.OpenFile(v, os.O_RDWR, os.ModePerm); err != nil {
					log.Println(err.Error())
				} else {
					defer f.Close()
					if b, err := ioutil.ReadAll(f); err != nil {
						println("err", err.Error())
					} else {
						n := strings.ReplaceAll(string(b), ",omitempty", "")
						os.Truncate(v, 0)
						f.WriteAt([]byte(n), 0)
					}
				}
			}
		}
	}
	return err
}

func recursionReadFile(dirname string) {
	file, _ := ioutil.ReadDir(dirname)
	for _, v := range file {
		if v.IsDir() {
			if !isExistsArrary(v.Name(), IgnoreDir) {
				recursionReadFile(dirname + "\\" + v.Name())
			}
		} else {
			files = append(files, dirname+"\\"+v.Name())
		}
	}
}

func isExistsArrary(s string, arr []string) bool {
	for _, v := range arr {
		if s == v {
			return true
		}
	}
	return false
}
