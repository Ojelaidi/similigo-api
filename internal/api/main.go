package api

import (
	"context"
	"fmt"
	"github.com/Ojelaidi/similigo-api/config"
	"github.com/Ojelaidi/similigo-api/internal/api/router"
	similigo_api "github.com/Ojelaidi/similigo-api/internal/api/similigo-api"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	config config.ServerConfiguration
	router *gin.Engine
}

func New() *App {
	app := &App{}
	app.setup()
	return app
}

func (app *App) setup() {
	// Load config
	cfg := config.LoadConfig()

	r := gin.Default()

	similiService := similigo_api.NewService()
	router.SetupRoutes(r, similiService)

	app.config = cfg
	app.router = r

	//url := ginSwagger.URL(fmt.Sprintf("http://%s:%s/swagger/doc.json", cfg.Host, cfg.Port))
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

}

func (app *App) Run() {

	// https://gin-gonic.com/docs/examples/graceful-restart-or-stop/
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.config.Port),
		Handler: app.router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Run ListenAndServe", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Print("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Run Shutdown", zap.Error(err))
	}

	log.Print("Server exiting")
}
