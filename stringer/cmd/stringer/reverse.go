package stringer

import (
	"fmt"
	"github.com/spf13/cobra"
	"stringer/pkg/stringer"
)

// reverseCmd : 명령어 정의
var reverseCmd = &cobra.Command{
	Use:     "reverse",
	Aliases: []string{"rev"},
	Short:   "Reverse a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := stringer.Revers(args[0])
		fmt.Println(res)
	},
}

// init 함수는 패키지 로드시 실행
func init() {
	rootCmd.AddCommand(reverseCmd)
}
