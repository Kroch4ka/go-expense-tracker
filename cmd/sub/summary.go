package sub

import "github.com/spf13/cobra"

func init() {
	SummaryCmd.Flags().Int("month", 0, "month")
}

var SummaryCmd = &cobra.Command{
	Use: "summary",
	Run: func(cmd *cobra.Command, args []string) {
		// not implemented!
	},
}
