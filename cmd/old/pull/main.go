package main

/**
通过阿里云容器仓库拉gcr镜像；
需要先通过github action 将目标镜像同步至阿里云
 */

import (
	"flag"
	"fmt"
	"ipull/pkg/code"
	"log"
	"strings"

	"ipull/pkg/exec"
)

var (
	cri   string
	image string
	namespace string
)

const aliRepo = "registry.cn-shanghai.aliyuncs.com/shuaiyy/2233"

func init() {
	flag.StringVar(&cri, "runtime", "docker", "容器运行时:docker|containerd|ctr|podman")
	flag.StringVar(&image, "image", "", "镜像")
	flag.StringVar(&namespace, "namespace", "", "镜像空间，for containerd")
}

func main() {
	flag.Parse()
	image = strings.TrimSpace(image)
	if image == "" {
		fmt.Println("please give a image")
	}
	tag := code.UrlEncode(image)
	fmt.Println(image, tag)
	aliImage := fmt.Sprintf("%s:%s", aliRepo, tag)
	switch cri {
	case "docker":
		if err := exec.DockerPull(aliImage); err != nil{
			log.Println(aliImage, err)
			return
		}
		if err := exec.DockerRetag(aliImage, image); err != nil{
			log.Println(aliImage, image, err)
			return
		}
		if err := exec.DockerRM(aliImage); err != nil{
			log.Println(aliImage, err)
			return
		}
	case "ctr", "containerd":
		if err := exec.CtrPull(aliImage, namespace); err != nil{
			log.Println(aliImage, namespace, err)
			return
		}
		if err := exec.CtrRetag(aliImage, image, namespace); err != nil{
			log.Println(aliImage, image, namespace, err)
			return
		}
	}
}

