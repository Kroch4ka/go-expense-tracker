package cmd

import (
	"github.com/Kroch4ka/go-expense-tracker/cmd/sub"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sub.AddCmd, sub.DeleteCmd, sub.ListCmd, sub.SummaryCmd)
}

var rootCmd = &cobra.Command{
	Use: "expense-tracker",
}

func Execute() {
	rootCmd.Execute()
}
