package cmd

import (
	"context"
	"log"
	"net"

	"github.com/jexodusmercado/blog-system-go/internal/api"
	"github.com/jexodusmercado/blog-system-go/internal/conf"
	"github.com/jexodusmercado/blog-system-go/internal/router"
	"github.com/jexodusmercado/blog-system-go/internal/storage"
	"github.com/spf13/cobra"
)

var serveCmd = cobra.Command{
	Use:  "serve",
	Long: "Start API server",
	Run: func(cmd *cobra.Command, args []string) {
		serve(cmd.Context())
	},
}

func serve(ctx context.Context) {
	config := conf.LoadConfig()

	db := storage.Dial(config)

	gRouter := router.NewRouter()
	router := gRouter.InitRouter()

	api := api.NewAPI(db, config, router)

	addr := net.JoinHostPort(config.Host, config.Port)
	log.Printf("Starting server on %s", addr)

	api.ListenAndServe(ctx, addr)

}