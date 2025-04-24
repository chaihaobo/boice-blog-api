package cmder

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/chaihaobo/boice-blog-api/cmd/core"
	"github.com/chaihaobo/boice-blog-api/utils"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "be-template",
	Short: "will start the all process",
}

// NewRoot initializes the root commander
func NewRoot() core.Cmder {
	return core.CmderFunc(func(ctx *core.Context) *cobra.Command {
		rootCmd.AddCommand(NewHTTP().Command(ctx))
		rootCmd.AddCommand(NewGrpc().Command(ctx))
		rootCmd.Run = func(cmd *cobra.Command, args []string) {
			listenRoot(ctx)
		}
		return rootCmd
	})
}

func listenRoot(ctx *core.Context) {
	go func() {
		if err := ctx.Transport.ServeAll(); err != nil {
			ctx.Resource.Logger().Error(context.Background(), "failed to listen root", err)
		}
	}()

	utils.GracefulShutdown(ctx.Transport.ShutdownAll)

}
