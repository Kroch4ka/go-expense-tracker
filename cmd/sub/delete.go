package sub

import (
	"fmt"

	"github.com/Kroch4ka/go-expense-tracker/storage"
	"github.com/spf13/cobra"
)

var (
	defaultStorage = storage.CSVStorage{}
	expenses       = defaultStorage.Load()
)

func init() {
	DeleteCmd.Flags().Int("id", 0, "id")
}

var DeleteCmd = &cobra.Command{
	Use: "delete",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			fmt.Println("error during parsing id")
			return
		}
		err = expenses.Delete(id)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			defaultStorage.Unload(expenses)
			fmt.Println("Success!")
		}
	},
}
