package ls

import (
	"os"

	"github.com/ncw/rclone/cmd"
	"github.com/ncw/rclone/fs"
	"github.com/spf13/cobra"
)

func init() {
	cmd.Root.AddCommand(lsCmd)
}

var lsCmd = &cobra.Command{
	Use:   "ls remote:path",
	Short: `List all the objects in the the path with size and path.`,
	Run: func(command *cobra.Command, args []string) {
		cmd.CheckArgs(1, 1, command, args)
		fsrc := cmd.NewFsSrc(args)
		cmd.Run(false, command, func() error {
			return fs.List(fsrc, os.Stdout)
		})
	},
}
