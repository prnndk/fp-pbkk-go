package main

import (
	"log"
	"os"

	"github.com/prnndk/final-project-golang-pbkk/command"
	"github.com/prnndk/final-project-golang-pbkk/config"
	"github.com/prnndk/final-project-golang-pbkk/controller"
	"github.com/prnndk/final-project-golang-pbkk/middleware"
	"github.com/prnndk/final-project-golang-pbkk/repository"
	"github.com/prnndk/final-project-golang-pbkk/routes"
	"github.com/prnndk/final-project-golang-pbkk/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetUpDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	if len(os.Args) > 1 {
		flag := command.Commands(db)
		if !flag {
			return
		}
	}

	var (
		jwtService service.JWTService = service.NewJWTService()

		// Implementation Dependency Injection
		// Repository
		userRepository       repository.UserRepository        = repository.NewUserRepository(db)
		userTicketRepository repository.EventTicketRepository = repository.NewEventTicketRepository(db)
		pembayaranRepository repository.PembayaranRepository  = repository.NewPembayaranRepository(db)
		eventRepository      repository.EventRepository       = repository.NewEventRepository(db)

		// Service
		userService       service.UserService       = service.NewUserService(userRepository, jwtService)
		userTicketService service.UserTicketService = service.NewUserTicketService(userTicketRepository, eventRepository)
		pembayaranService service.PembayaranService = service.NewPembayaranService(pembayaranRepository, userTicketRepository)
		eventService      service.EventService      = service.NewEventService(eventRepository)
		storageService    service.StorageService    = service.NewStorageService()

		// Controller
		userController       controller.UserController       = controller.NewUserController(userService)
		userTicketController controller.UserTicketController = controller.NewUserTicketController(userTicketService)
		pembayaranController controller.PembayaranController = controller.NewPembayaranController(pembayaranService)
		eventController      controller.EventController      = controller.NewEventController(eventService)
		storageController    controller.StorageController    = controller.NewStorageController(storageService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	// routes
	routes.User(server, userController, jwtService)
	routes.UserTicket(server, userTicketController, jwtService)
	routes.Pembayaran(server, pembayaranController, jwtService)
	routes.Event(server, eventController, jwtService)
	routes.Storage(server, storageController, jwtService)

	server.Static("/assets", "./assets")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
