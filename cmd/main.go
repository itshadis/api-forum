package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itshadis/api-forum/internal/handlers/memberships"
)

func main() {
	r := gin.Default()

	membershipHandler := memberships.NewHandler(r)
	membershipHandler.RegisterRoute()

	r.Run()
}
