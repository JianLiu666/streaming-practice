package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"m7s.live/engine/v4"
	_ "m7s.live/plugin/rtmp/v4"
)

var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "",
	Long:  ``,
	RunE:  RunDemoCmd,
}

func init() {
	rootCmd.AddCommand(demoCmd)
}

func RunDemoCmd(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	engine.Run(ctx, "")
	return nil
}
