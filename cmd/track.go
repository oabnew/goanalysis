package cmd

import (
	"errors"

	"github.com/oabnew/goanalysis/pkg/track"
	"github.com/spf13/cobra"
)

/**
@file:
@author: levi.Tang
@time: 2024/10/28 22:51
@description:
**/

// analysisCmd represents the analysis command
var trackCmd = &cobra.Command{
	Use:   "track <programDir>",
	Short: "Rewrite the contents of the catalog to implement function tracing.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please input program dir")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		track.RewriteDir(args[0])
	},
}

func init() {
	rootCmd.AddCommand(trackCmd)
}
