package stringer

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// 빌드시 아래 명령어를 하면 version ㄴ을 바꿀 수 있음
// go build -o ./dist/stringer -ldflags="-X 'stringer/cmd/stringer.version=0.0.2'"
var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "stringer",
	Version: version,
	Short:   "stringer - a simple CLI to transform and inspect strings",
	Long:    "strings is a super fancy CLI\n\nOne can use stringer to modify or inspect strings straight from the terminal",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'\n", err)
		os.Exit(1)
	}
}
