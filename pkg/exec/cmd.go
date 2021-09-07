package exec

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
)

func ExecCommand(commandName string, params []string, isPrint bool) (string, error) {
	cmd := exec.Command(commandName, params...)
	//显示运行的命令
	//fmt.Println(cmd.Args)
	stdout, err := cmd.StdoutPipe()
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	var lastOutput string
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		if isPrint {
			fmt.Print(line)
		}
		lastOutput = line
	}
	slurp, _ := ioutil.ReadAll(stderr)
	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("%s", slurp)
	}
	if isPrint && len(slurp) > 0 {
		fmt.Println(string(slurp))
	}
	return strings.TrimSpace(lastOutput), nil
}

func DockerLogin(register, user, pass string) error {
	cmd := fmt.Sprintf("docker login --username=%s --password=%s %s", user, pass, register)
	_, err := ExecCommand("bash", []string{"-c", cmd}, true)
	return err
}

func DockerLogout(register string) error {
	cmd := fmt.Sprintf("docker logout %s", register)
	_, err := ExecCommand("bash", []string{"-c", cmd}, true)
	return err
}

func DockerPull(image string) error {
	cmd := fmt.Sprintf("docker pull %s", image)
	_, err := ExecCommand("bash", []string{"-c", cmd}, true)
	return err
}

func DockerPush(image string) error {
	cmd := fmt.Sprintf("docker push %s", image)
	_, err := ExecCommand("bash", []string{"-c", cmd}, true)
	return err
}

func DockerRetag(oldImage, newImage string) error {
	cmd := fmt.Sprintf("docker tag %s %s", oldImage, newImage)
	_, err := ExecCommand("bash", []string{"-c", cmd}, true)
	return err
}

func DockerRM(oldImage string) error {
	cmd := fmt.Sprintf("docker rmi %s", oldImage)
	_, err := ExecCommand("bash", []string{"-c", cmd}, true)
	return err
}

func CtrPull(image, namespace string) error {
	if namespace != "" {
		namespace = "-n " + namespace
	}
	cmd := fmt.Sprintf("ctr image %s pull %s", namespace, image)
	_, err := ExecCommand("bash", []string{"-c", cmd}, true)
	return err
}

func CtrRetag(image, newImage, namespace string) error {
	if namespace != "" {
		namespace = "-n " + namespace
	}
	cmd := fmt.Sprintf("ctr image %s tag %s %s", namespace, image, newImage)
	_, err := ExecCommand("bash", []string{"-c", cmd}, true)
	return err
}
