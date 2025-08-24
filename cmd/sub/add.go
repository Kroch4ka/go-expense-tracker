package sub

import (
	"fmt"

	"github.com/Kroch4ka/go-expense-tracker/domain"
	"github.com/spf13/cobra"
)

func init() {
	AddCmd.Flags().Int("amount", 0, "expense amount")
	AddCmd.Flags().String("description", "", "expense description")
}

var AddCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		amount, err := cmd.Flags().GetInt("amount")
		if err != nil {
			fmt.Println("error during parsing amount")
			return
		}
		description, err := cmd.Flags().GetString("description")
		if err != nil {
			fmt.Println("error during parsing description")
		}
		expenses.Add(domain.Expense{
			Amount:      amount,
			Description: description,
		})
		defaultStorage.Unload(expenses)
	},
}
