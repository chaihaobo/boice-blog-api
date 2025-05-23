package cmder

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/chaihaobo/boice-blog-api/cmd/core"
	"github.com/chaihaobo/boice-blog-api/transport/http"
	"github.com/chaihaobo/boice-blog-api/utils"
)

// rootCmd represents the base command when called without any subcommands
var httpCmd = &cobra.Command{
	Use:   "serveHTTP",
	Short: "will start the http process",
}

func NewHTTP() core.Cmder {
	return core.CmderFunc(func(ctx *core.Context) *cobra.Command {
		httpCmd.Run = func(cmd *cobra.Command, args []string) {
			listenHTTP(ctx, ctx.Transport.HTTP())
		}
		return httpCmd
	})
}

func listenHTTP(ctx *core.Context, http http.Transport) {
	go func() {
		if err := http.Serve(); err != nil {
			ctx.Resource.Logger().Error(context.Background(), "listen http failed", err)
		}
	}()
	utils.GracefulShutdown(http.Shutdown, ctx.Resource.Close)
}
