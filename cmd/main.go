package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/itshadis/api-forum/internal/configs"
	"github.com/itshadis/api-forum/internal/handlers/memberships"
	membershipRepo "github.com/itshadis/api-forum/internal/repositories/memberships"
	membershipSvc "github.com/itshadis/api-forum/internal/services/memberships"
	"github.com/itshadis/api-forum/pkg/internalsql"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}
	cfg = configs.Get()
	log.Println("config", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisiasi database", err)
	}

	membershipRepo := membershipRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(cfg, membershipRepo)

	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
