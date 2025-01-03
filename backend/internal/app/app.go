package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/snakoner/dex/internal/config"
	"github.com/snakoner/dex/internal/eth"
)

type App struct {
	logger *logrus.Logger
	config *config.Config
	router *mux.Router
	server *http.Server
	ethSrv *eth.EthereumServer
}

func New(config *config.Config, logger *logrus.Logger) (*App, error) {
	ethSrv, err := eth.New(config)
	if err != nil {
		return nil, err
	}

	app := &App{
		logger: logger,
		config: config,
		router: mux.NewRouter(),
		server: &http.Server{},
		ethSrv: ethSrv,
	}

	app.setRoutes()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	app.server.Handler = handlers.CORS(headers, methods, origins)(app.router)
	app.server.Addr = fmt.Sprintf("%s:%s", app.config.Service.Host, app.config.Service.Port)

	return app, nil
}

func (a *App) setRoutes() {
	a.router.HandleFunc("/round/{number}", a.ethSrv.RoundHandler).Methods("GET")
	a.router.HandleFunc("/winner/{number}", a.ethSrv.WinnerHandler).Methods("GET")
	a.router.HandleFunc("/all-time-reward", a.ethSrv.AllTimeRewardHander).Methods("GET")
}

func (a *App) Run(ctx context.Context) error {
	go func() {
		a.logger.Info("start http server")
		if err := a.server.ListenAndServe(); err != nil {
			a.logger.Error(err)
		}
	}()

	// go func() {
	// 	a.logger.Info("start SubscribeBidEvent()")
	// 	a.ethSrv.SubscribeBidEvent(ctx)
	// }()

	// go func() {
	// 	a.logger.Info("start SelectWinner()")
	// 	a.ethSrv.SelectWinner(ctx)
	// }()

	<-ctx.Done()
	ctxShutdown, cancel := context.WithTimeout(context.Background(), time.Second*3)
	a.ethSrv.Stop("")
	defer cancel()

	if err := a.server.Shutdown(ctxShutdown); err != nil {
		a.logger.Error(err)
		return err
	}

	a.logger.Info("http server successfully stopped")

	return nil
}
