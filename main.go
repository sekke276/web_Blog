package main

import (
	"fmt"
	"log"
	"web_Blogs/api/configs"
	"web_Blogs/api/handler"
	"web_Blogs/api/mongo_repository"
	"web_Blogs/api/routes"
	_ "web_Blogs/docs"
	"web_Blogs/pkg/usecase/user"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg, err := configs.EnvMongoURI()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	db, err := mongo_repository.ConnectMongo(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to mongo db")
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))
	app.Get("/swagger/doc.json", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/swagger.json")
	})

	api := app.Group("/v1")
	userRepo := mongo_repository.NewUserMongoRepository(db)
	userUC := user.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUC)
	routes.UserRouter(api, *userHandler)
	port := fmt.Sprintf(":%v", cfg.Port)
	app.Listen(port)
	log.Printf("Server started on port %v", cfg.Port)
}
