package main

import (
	_ "github.com/nhutHao02/social-network-common-service/model"
	"github.com/nhutHao02/social-network-user-service/startup"
)

// @title			Social Network Service
// @description	This is user service of the social network implament using Go
// @version		1.0
// @BasePath		/api/v1
func main() {
	startup.Start()
}
