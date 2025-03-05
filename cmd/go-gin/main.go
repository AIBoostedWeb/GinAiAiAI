package main

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-gin/internal/config"
	"go-gin/internal/engine"
	"go-gin/internal/repository"
	util2 "go-gin/pkg/util"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {
	var wd string
	var mainHandler *gin.Engine
	var logger *zap.Logger
	var err_list []string
	var err error
	var CONFIG config.Config

	var db *gorm.DB

	wd, err = util2.FindProjectRoot("go.mod")
	wd = filepath.Join(wd, "configs")
	if err != nil {
		panic(err)
	}

	err = config.InitConf(wd, "config", "yaml", &CONFIG)
	if err != nil {
		print(err.Error())
		panic(err)
	}

	logger, err = util2.InitLogger(CONFIG.Global.RunMode)

	if err != nil {
		err_list = append(err_list, err.Error())
		for _, errInfo := range err_list {
			print(errInfo)
		}
		print(err.Error())
		panic(err)

	}

	db, err = repository.InitDB(&CONFIG.Database, logger)
	if err != nil {
		err_list = append(err_list, err.Error())
	}

	mainHandler = engine.InitMainHandler(CONFIG.Global.RunMode, logger)
	engine.RegisterRoutes(&CONFIG.Service, &CONFIG.Server, mainHandler, db)

	srv := &http.Server{
		Addr:    CONFIG.Server.Host + ":" + CONFIG.Server.Port,
		Handler: mainHandler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("服务启动失败:" + err.Error())
		}
	}()

	for _, errInfo := range err_list {
		logger.Error(errInfo)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("正在关闭服务...")
}
