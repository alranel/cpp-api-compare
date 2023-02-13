package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                   "cpp-api-compare",
	Short:                 "Compare multiple C++ libraries implementing the same API.",
	Long:                  "This tool compares two or more C++ libraries supposed to implement the same API, in order to check their interoperability.",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	//rootCmd.PersistentFlags().String("datadir", "", "The directory where test results are stored.")
}

// Execute starts the cobra command parsing chain.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
