// Package main
package main

import (
	"fmt"
	"log"

	"github.com/VH288/miniig/internal/configs"

	"github.com/VH288/miniig/internal/handlers/memberships"
	"github.com/VH288/miniig/internal/handlers/posts"
	membershipRepo "github.com/VH288/miniig/internal/repository/memberships"
	postRepo "github.com/VH288/miniig/internal/repository/posts"
	membershipSvc "github.com/VH288/miniig/internal/service/memberships"
	postSvc "github.com/VH288/miniig/internal/service/posts"
	"github.com/VH288/miniig/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	var cfg *configs.Config

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("Gagal inisiai config", err)
	}

	cfg = configs.Get()
	log.Printf("Configs: %+v", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisiai database", err)
	}

	membershipRepo := membershipRepo.NewRepository(db)
	membershipService := membershipSvc.NewService(cfg, membershipRepo)

	postRepo := postRepo.NewRepository(db)
	postSvc := postSvc.NewService(cfg, postRepo)

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	postHandler := posts.NewHandler(r, postSvc)
	postHandler.RegisterRoute()

	err = r.Run(cfg.Service.Port)
	if err != nil {
		fmt.Println(err)
	}
}
