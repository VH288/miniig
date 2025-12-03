package main

import (
	"fmt"
	"log"

	"github.com/VH288/miniig/internal/configs"

	"github.com/VH288/miniig/internal/handlers/memberships"
	membershipRepo "github.com/VH288/miniig/internal/repository/memberships"
	membershipSvc "github.com/VH288/miniig/internal/service/memberships"
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
	membershipService := membershipSvc.NewService(membershipRepo)

	r := gin.Default()

	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	err = r.Run(cfg.Service.Port)
	if err != nil {
		fmt.Println(err)
	}
}
