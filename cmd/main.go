package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

var IDLPath = flag.String("idl", "", "IDL path")

func main() {
	flag.Parse()

	idl := path.Base(*IDLPath)

	AbsPath, err := filepath.Abs(*IDLPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	dirPath := filepath.Dir(AbsPath)

	gopath := GetGoPATH()
	log.Println("use GOPATH:", gopath)

	// 生成protoc，grpc代码
	cmd := exec.Command("protoc", "--go_out="+gopath, "--go-grpc_out="+gopath, fmt.Sprintf("--proto_path=%s", dirPath), idl)
	log.Println(cmd.String())

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err, ":", string(output))
		return
	}

	// 创建git仓库,add,commit,push
	// var idlWithoutType string
	// if strings.HasSuffix(idl, ".proto") {
	// 	idlWithoutType = idl[:strings.LastIndex(idl, ".proto")]
	// }

}

func GetGoPATH() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = os.Getenv("HOME") + "/go/src"
	}
	return gopath
}

func ExtracProtocGoPackage(proto string) (string, error) {
	f, err := os.Open(proto)
	if err != nil {
		return "", err
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				return "", errors.New("go package not found in proto")
			}
			return "", err
		}
		var packageName string
		_, err = fmt.Sscanf(line, "option go_package = \"%s\";\n", &packageName)
		if err == nil {
			return packageName, nil
		}
	}
}
