package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "minstor",
		Short: "minstor",
		Long:  "minstor",
		Run:   runRootCmd,
	}

	cmd.AddCommand(newObsCmd())

	return cmd
}

func runRootCmd(cmd *cobra.Command, args []string) {
	fmt.Println("Hello, minstor!")
}

func Execute() {
	cobra.CheckErr(newRootCmd().Execute())
}
