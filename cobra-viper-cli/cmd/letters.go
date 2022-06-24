package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	countFlagLetters int
	langFlagLetters  string
)

var lettersCmd = &cobra.Command{
	Use:   "letters",
	Short: "Returns random letters",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("letters mode")
		fmt.Println("--count:", countFlagLetters)
		fmt.Println("--lang:", langFlagLetters)
		fmt.Println("--verbose:", verbose)
	},
}

func init() {
	rootCmd.AddCommand(lettersCmd)
	var Count int
	lettersCmd.Flags().IntVarP(
		&Count, "count", "c", 0,
		"A count of random letters",
	)
	lettersCmd.MarkFlagRequired("count")

	var Lang string
	lettersCmd.Flags().StringVarP(
		&Lang, "lang", "l", "en",
		"A language. Optional",
	)
}
