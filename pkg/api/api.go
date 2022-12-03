package api

import (
	"blog/pkg/db"
	indexTransport "blog/pkg/services/index/transport"
	postTransport "blog/pkg/services/posts/transport"
	userTransport "blog/pkg/services/users/transport"
	"blog/pkg/utils/sessions"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"log"
	"os"
)

type Config struct {
	DBHost       string
	DBPort       int
	DBUser       string
	DBPassword   string
	DBName       string
	RunMigration bool

	AppHost string
	AppPort int
}

// Start initializes the API server, adding the reuired middlewares and dependent services
func Start(cfg *Config) {
	//ctx := context.Background()
	conn, err := db.GetConnection(
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName)
	//if err != nil {
	//	log.Error(ctx, "unable to establish a database connection: %s", err.Error())
	//}

	if cfg.RunMigration && conn != nil {
		err := db.Migrate(conn)
		if err != nil {
			//log.Error(ctx, "unable to complete auto migration", err.Error())
			log.Fatal("unable to complete auto migration")
		}
	}

	//TEMPLATE_ENGINE_DIR="./src/views/templates"

	//templateEngineDir := os.Getenv("TEMPLATE_ENGINE_DIR")
	//./../../src/views/templates
	os.Getenv("TEMPLATE_ENGINE_DIR")
	engine := html.New("./pkg/services", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	Session := sessions.Instance
	Session.SetSession()
	//app.Use(middlewares.PersistContext())
	//app.Use(middlewares.RequestLogger())
	//app.Use(middlewares.ForceJSON())
	//app.Use(middlewares.Recover())

	//app.NoRoute(middlewares.NoRoute())
	//app.NoMethod(middlewares.NoMethod())

	indexTransport.Activate(app, conn)
	postTransport.Activate(app, conn)
	userTransport.Activate(app, conn)

	err = app.Listen(fmt.Sprintf("%s:%d", cfg.AppHost, cfg.AppPort))
	if err != nil {
		panic(err)
	}

}
