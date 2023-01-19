package converter

import (
	"github.com/spf13/cobra"
	"path/filepath"
	"smartdom/services/executer/internal/usecase/converter"
)

var ConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "A note can be anything you'd like to study and review",
	Long:  `A note can be anything you'd like to study and review`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filepath.Walk(args[0], converter.Run)
	},
}
