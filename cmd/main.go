package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/itshadis/api-forum/internal/configs"
	"github.com/itshadis/api-forum/internal/handlers/memberships"
	"github.com/itshadis/api-forum/internal/handlers/posts"
	membershipRepo "github.com/itshadis/api-forum/internal/repositories/memberships"
	postRepo "github.com/itshadis/api-forum/internal/repositories/posts"
	membershipSvc "github.com/itshadis/api-forum/internal/services/memberships"
	postSvc "github.com/itshadis/api-forum/internal/services/posts"
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

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(cfg, membershipRepo)
	postService := postSvc.NewService(cfg, postRepo)

	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	postHandler := posts.NewHandler(r, postService)
	postHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
