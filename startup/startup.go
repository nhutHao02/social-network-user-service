package startup

import (
	"github.com/nhutHao02/social-network-user-service/config"
	http "github.com/nhutHao02/social-network-user-service/internal/http"
)

func StartServer() {
	// load congig
	cfg := config.LoadConfig()

	// database setup

	// setup route

	// setup server
	runServer(cfg)

}

func runServer(*config.Config) {

	// run http server
	http.RunHTTPServer()

	// run grpc server
}
