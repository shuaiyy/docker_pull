package main

import (
	"crypto/sha256"
	"fmt"
	"ipull/pkg/exec"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ./pull {IMAGE}")
		return
	}
	image := strings.TrimSpace(os.Args[1])
	aliImage := fmt.Sprintf("registry.cn-shanghai.aliyuncs.com/shuaiyy/2233:%s", Sha256(image))
	if err := exec.DockerPull(aliImage); err != nil	{
		fmt.Println(err)
		return
	}
	newImage := strings.Split(image, "@")[0]
	if err := exec.DockerRetag(aliImage, newImage); err != nil	{
		fmt.Println(err)
		return
	}
	fmt.Println("ReTag:")
	fmt.Println(newImage)
	exec.DockerRM(aliImage)
}

func Sha256(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}