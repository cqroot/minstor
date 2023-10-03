package cmd

import (
	"github.com/cqroot/minstor/internal/obs"
	"github.com/spf13/cobra"
)

func newObsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "obs",
		Short: "obs",
		Long:  "obs",
		Run: func(cmd *cobra.Command, args []string) {
			obs, err := obs.InitObs()
			cobra.CheckErr(err)

			cobra.CheckErr(obs.Run())
		},
	}
	return cmd
}
