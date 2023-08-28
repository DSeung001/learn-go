package cmd

import (
	"github.com/multiple-ssh/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "msh",
	Short: "한번에 여러개의 ssh 접속 후 같은 명령어 실행할 수 있는 cli 도구",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	utils.HandleErr(rootCmd.Execute())
}

func init() {

	// ssh 커맨드에 플래그 추가
	sshCmd.Flags().StringVarP(&remoteFile, "file", "f", "remote.txt", "원격 파일 경로 (예: remote.txt)")
	sshCmd.Flags().StringVarP(&command, "command", "c", "", "실행할 명령")

	rootCmd.AddCommand(sshCmd)
}
