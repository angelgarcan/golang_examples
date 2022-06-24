package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	countFlagNumbers int
	rangeFlagNumbers []string
)
var numbersCmd = &cobra.Command{
	Use:   "numbers",
	Short: "Returns random numbers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("numbers mode")
		fmt.Println("--count:", countFlagNumbers)
		fmt.Println("--range:", rangeFlagNumbers)
		fmt.Println("--verbose:", verbose)
	},
}

func init() {
	rootCmd.AddCommand(numbersCmd)

	var countFlag int
	numbersCmd.Flags().IntVarP(
		&countFlag, "count", "c", 0,
		"A count of random numbers",
	)

	numbersCmd.MarkFlagRequired("count")

	var rangeFlag []string
	numbersCmd.Flags().StringSliceVarP(
		&rangeFlag, "range", "r", []string{"1:100"},
		"Range of numbers. Optional",
	)
}
