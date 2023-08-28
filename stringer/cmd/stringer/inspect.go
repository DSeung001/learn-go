package stringer

import (
	"fmt"
	"github.com/spf13/cobra"
	"stringer/pkg/stringer"
)

// 로컬 플래그 변수
var onlyDigits bool

var inspectCmd = &cobra.Command{
	Use:     "inspect",
	Aliases: []string{"insp"},
	Short:   "Inspect a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		i := args[0]
		res, kind := stringer.Inspect(i, onlyDigits)

		// 복수면 s를 추가
		plurals := "s"
		if res == 1 {
			plurals = ""
		}
		fmt.Printf("'%s' has a %d %s%s.\n", i, res, kind, plurals)
	},
}

// init 함수는 패키지 로드시 실행
func init() {
	// 로컬 플래그 추가
	inspectCmd.Flags().BoolVarP(&onlyDigits, "digits", "d", false, "Count only digits")
	rootCmd.AddCommand(inspectCmd)
}
