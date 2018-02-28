package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/engineyard/scaley/workflows"
)

var scaleCmd = &cobra.Command{
	Use: "scale <group>",

	Short: "Scale a group up or down",

	Long: `Scale a group up or down

Given a group description file, writen in YAML, determine if the group in
question should be scaled up or down. If yes, scale the group in the
proper direction.`,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Usage: scaley scale groupname")
		}

		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		return workflows.Perform(
			&workflows.ScalingAGroup{
				GroupName: args[0],
			},
		)
	},

	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	RootCmd.AddCommand(scaleCmd)
}
