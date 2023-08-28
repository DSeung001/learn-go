package cmd

import (
	"bufio"
	"fmt"
	"github.com/multiple-ssh/utils"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"os/exec"
	"strings"
)

var remoteFile string
var command string

// 여러 대의 PC에 SSH로 연결하고 명령을 실행하는 서브커맨드 정의
var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "여러 대의 PC에 SSH로 연결하고 명령을 실행",
	Run: func(cmd *cobra.Command, args []string) {
		if remoteFile == "" {
			fmt.Println("파일 경로를 --file 플래그로 제공해주세요")
			return
		}

		hosts, err := readRemoteFile(remoteFile)
		fmt.Println(hosts)

		if err != nil {
			fmt.Println("원격 파일을 읽는 데 실패했습니다:", err)
			return
		}
		//
		//for _, host := range hosts {
		//	runCommandOverSSH(host, command)
		//}
	},
}

// readRemoteFile : Remote 연결 정보를 담은 파일 읽기
func readRemoteFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	defer file.Close()
	utils.HandleErr(err)
	var hosts []string

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		line = strings.TrimSpace(line)
		if line != "" {
			hosts = append(hosts, line)
		}
	}

	return hosts, nil
}

// runCommandOverSSH : 넘어온 호스트 정보에 해당 명령어 실행
func runCommandOverSSH(host, command string) {
	// alias 인가?
	// Y -> 바로 접속
	// N -> user, host 분석

	// SSH 클라이언트 설정 생성
	config := &ssh.ClientConfig{
		User: "your_username", // SSH 사용자 이름으로 변경해주세요
		Auth: []ssh.AuthMethod{
			ssh.Password("your_password"), // SSH 비밀번호 또는 다른 인증 방법 사용하도록 변경해주세요
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 테스트 목적으로만 사용해주세요
	}

	// 원격 호스트에 연결
	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		fmt.Printf("%s에 연결할 수 없습니다: %s\n", host, err)
		return
	}
	defer client.Close()

	// 세션 생성
	session, err := client.NewSession()
	if err != nil {
		fmt.Printf("%s에 대한 세션 생성 실패: %s\n", host, err)
		return
	}
	defer session.Close()

	// 명령 실행
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()

	err = cmd.Run()
	if err != nil {
		fmt.Printf("%s에서 명령 실행 실패: %s\n", host, err)
	}
}
