package cmder

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/chaihaobo/boice-blog-api/cmd/core"
	"github.com/chaihaobo/boice-blog-api/transport/grpc"
	"github.com/chaihaobo/boice-blog-api/utils"
)

// grpcCmd will start the grpc server
var grpcCmd = &cobra.Command{
	Use:   "serveGrpc",
	Short: "will start the grpc process",
}

func NewGrpc() core.Cmder {
	return core.CmderFunc(func(ctx *core.Context) *cobra.Command {
		httpCmd.Run = func(cmd *cobra.Command, args []string) {
			listenGrpc(ctx, ctx.Transport.Grpc())
		}
		return httpCmd
	})
}

func listenGrpc(ctx *core.Context, grpc grpc.Transport) {
	go func() {
		if err := grpc.Serve(); err != nil {
			ctx.Resource.Logger().Error(context.Background(), "listen grpc failed", err)
		}
	}()
	utils.GracefulShutdown(func() error {
		grpc.GracefulStop()
		return nil
	}, ctx.Resource.Close)
}
