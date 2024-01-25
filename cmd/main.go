package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"taskema/adapter/fileadapter"
	"taskema/adapter/orgadapter"
	"taskema/adapter/useradapter"
	"taskema/adapter/workspaceadapter"
	"taskema/config"
	"taskema/datasource/mysql"
	"taskema/delivery/httpserver"
	"taskema/delivery/httpserver/authhandler"
	"taskema/delivery/httpserver/boardhandler"
	"taskema/delivery/httpserver/filehandler"
	"taskema/delivery/httpserver/orghandler"
	"taskema/delivery/httpserver/taskhandler"
	"taskema/delivery/httpserver/userhandler"
	"taskema/delivery/httpserver/workspacehandler"
	"taskema/repository/boardrepo"
	"taskema/repository/filerepo"
	"taskema/repository/orgrepo"
	"taskema/repository/taskrepo"
	"taskema/repository/userrepo"
	"taskema/repository/workspacerepo"
	"taskema/service/authservice"
	"taskema/service/boardservice"
	"taskema/service/fileservice"
	"taskema/service/hashingservice"
	"taskema/service/orgservice"
	"taskema/service/taskservice"
	"taskema/service/userservice"
	"taskema/service/workspaceservice"
	orgvalidation "taskema/validation/organization"
	"taskema/validation/uservalidation"
	"taskema/validation/workspacevalidation"
	"time"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg, cErr := config.Load()
	if cErr != nil {
		log.Fatal("failed to load config ", cErr)
	}

	mysql, mErr := mysql.New(cfg.MySql)
	if mErr != nil {
		log.Fatal("failed to initialize mysql ", mErr)
	}

	pErr := mysql.Conn().PingContext(ctx)
	if pErr != nil {
		log.Fatal("failed to connect to mysql: ", pErr)
	}

	defer func() {
		if err := mysql.Conn().Close(); err != nil {
			fmt.Println("failed to close mysql", err)
		}
	}()

	hashingSvc := hashingservice.New()
	authSvc := authservice.New(cfg.Auth)
	userRepo := userrepo.New(mysql)
	fileRepo := filerepo.New(mysql)
	fileAdapter := fileadapter.New(fileRepo)
	useradapter := useradapter.New(userRepo, hashingSvc)
	userValidator := uservalidation.New(useradapter, fileAdapter)
	userSvc := userservice.New(userRepo, hashingSvc, authSvc)
	userhandler := userhandler.New(userSvc, userValidator, authSvc, cfg.Auth)

	fileSvc := fileservice.New(hashingSvc, fileRepo)
	filehandler := filehandler.New(authSvc, cfg.Auth, fileSvc)

	authhandler := authhandler.New(authSvc, cfg.Auth)

	orgRepo := orgrepo.New(mysql)
	orgSvc := orgservice.New(orgRepo)
	orgAdapter := orgadapter.New(orgRepo)
	orgValidation := orgvalidation.New(fileAdapter, orgAdapter)
	orgHandler := orghandler.New(authSvc, cfg.Auth, orgSvc, orgValidation)

	workspaceRepo := workspacerepo.New(mysql)
	workspaceSvc := workspaceservice.New(workspaceRepo)
	workspaceAdapter := workspaceadapter.New(workspaceRepo)
	workspaceValidation := workspacevalidation.New(fileAdapter, orgAdapter, workspaceAdapter)
	workspaceHandler := workspacehandler.New(authSvc, cfg.Auth, workspaceSvc, workspaceValidation)

	boardRepo := boardrepo.New(mysql)
	boardSvc := boardservice.New(boardRepo)
	boardHandler := boardhandler.New(authSvc, cfg.Auth, boardSvc)

	taskRepo := taskrepo.New(mysql)
	taskSvc := taskservice.New(taskRepo)
	taskHandler := taskhandler.New(authSvc, cfg.Auth, taskSvc)

	server := httpserver.New(cfg.Server, userhandler, filehandler, authhandler, orgHandler, workspaceHandler, boardHandler, taskHandler)

	ch := make(chan error, 1)
	go func() {
		if err := server.Serve(); err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case chErr := <-ch:

		log.Fatal(chErr)
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		server.Router().Shutdown(timeout)
	}
}
