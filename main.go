package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	HomeEnv          = "HOME"
	KubeconfigSuffix = ".kubeconfig"
)

var (
	KubeconfigPath = os.Getenv(HomeEnv) + "/.kube"
)

func GetKubeDirectory() {
	stat, err := os.Stat(KubeconfigPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s is a directory? %t \n", stat.Name(), stat.IsDir())

	files, err := os.ReadDir(KubeconfigPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), KubeconfigSuffix) && file.Type().IsRegular() {
			fmt.Println(file.Name())
		}
	}

}

func main() {
	GetKubeDirectory()
}
