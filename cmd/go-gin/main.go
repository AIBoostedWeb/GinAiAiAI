package main

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/internal/config"
	"go-gin/internal/engine"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go-gin/internal/pkg/util"
)

func main() {
	var mainHandler *gin.Engine
	var logger *zap.Logger
	var err error
	mainHandler, logger, err = initServer()
	if err != nil {
		log.Fatal(err)
		return
	}

	engine.RegisterRoutes(mainHandler, nil)
	srv := &http.Server{
		Addr:    ":" + config.CONFIG.Server.Port,
		Handler: mainHandler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("服务启动失败:" + err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("正在关闭服务...")
}

func initServer() (*gin.Engine, *zap.Logger, error) {
	var err error
	var wd string
	var errs []error

	var logger *zap.Logger
	var mainHandler *gin.Engine

	wd, err = util.FindProjectRoot("go.mod")
	if err != nil {
		errs = append(errs, err)

	}

	err = config.InitConf(wd, "config", "yaml")
	if err != nil {
		errs = append(errs, err)
		return nil, nil, err

	}

	logger, err = util.InitLogger(config.CONFIG.Global.RunMode)

	if err != nil {
		errs = append(errs, err)

	}

	mainHandler = engine.InitMainHandler(config.CONFIG.Global.RunMode, logger)

	if len(errs) > 0 {
		return nil, nil, fmt.Errorf("初始化服务失败:\n%w", errors.Join(errs...))
	}
	return mainHandler, logger, nil
}
