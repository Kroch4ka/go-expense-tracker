package sub

import (
	"fmt"

	"github.com/Kroch4ka/go-expense-tracker/domain"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ID  Date       Description  Amount")
		expenses.ForEach(func(e domain.Expense) {
			printExpense(e)
		})
	},
}

func printExpense(e domain.Expense) {
	fmt.Printf("%d %s %s %s%d\n", e.Id, e.Date.Format("2006-01-02"), e.Description, e.Currency.Format(), e.Amount)
}
