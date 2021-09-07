package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"ipull/pkg/code"
	"ipull/pkg/exec"
	"log"
	"os"
	"strings"
)

/*
通过github action下载海外镜像，然后推送到阿里云的容器镜像仓库
 */

var imageFile string

var (
	repo = "registry.cn-shanghai.aliyuncs.com/shuaiyy/2233"
	registry = "registry.cn-shanghai.aliyuncs.com"
	user = os.Getenv("secrets.DOCKER_USERNAME")
	key = os.Getenv("secrets.DOCKER_PASSWORD")
)


func main() {
	if len(os.Args) < 2 {
		return
	}
	imageFile = os.Args[1]
	fmt.Println(imageFile)
	c, err := ioutil.ReadFile(imageFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	// step 1 login
	if err = exec.DockerLogin(registry, user, key); err != nil{
		log.Println(err)
		//return
	}
	reader := bufio.NewReader(bytes.NewReader(c))
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil  {
			if line == "" {
				log.Print("read line error", err2)
				break
			}
		}
		image := strings.TrimSpace(line)
		if !strings.Contains(image, "gcr.io") &&
			!strings.Contains(image, "quay.io") &&
			!strings.Contains(image, "docker.io") {
			log.Print("image url must contains: gcr.io |quay.io|docker.io", image)
			continue
		}
		tag := code.UrlEncode(image)
		newImage := fmt.Sprintf("%s:%s", repo, tag)
		// step 2 拉镜像
		if err = exec.DockerPull(image); err != nil{
			log.Println(image, err)
			continue
		}
		// step 3 retag
		if err = exec.DockerRetag(image, newImage); err != nil{
			log.Println(image, newImage, err)
			continue
		}
		// step 4 push
		if err = exec.DockerPush(newImage); err != nil{
			log.Println(newImage, err)
			continue
		}
	}
}
