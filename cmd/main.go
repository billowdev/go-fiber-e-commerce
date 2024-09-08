package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	// graphqlAdapter "gqlhanon/internal/adapters/gql"
	// graphqlHandler "github.com/99designs/gqlgen/graphql/handler"
	// repositories "github.com/billowdev/go-fiber-e-commerce/internal/adapters/repositories/system_fields"
	// services "github.com/billowdev/go-fiber-e-commerce/internal/core/services/system_fields"

	"github.com/billowdev/go-fiber-e-commerce/internal/adapters/app"
	"github.com/billowdev/go-fiber-e-commerce/internal/adapters/database"
	"github.com/billowdev/go-fiber-e-commerce/internal/adapters/temporal/dto"
	"github.com/billowdev/go-fiber-e-commerce/internal/adapters/temporal/workflows"
	"github.com/billowdev/go-fiber-e-commerce/pkg/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
	temporalLog "go.temporal.io/sdk/log"
)

func main() {
	params := configs.NewFiberHttpServiceParams()
	fiberHTTP := configs.NewFiberHTTPService(params)
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal("Failed to start Database:", err)
	}

	logger := temporalLog.NewStructuredLogger(slog.Default())
	hostPort := client.DefaultHostPort
	if configs.TEMPORAL_CLIENT_URL != "" {
		hostPort = configs.TEMPORAL_CLIENT_URL
	}

	temporalClient, err := client.Dial(client.Options{
		HostPort: hostPort,
		Logger:   logger,
	})
	if err != nil {
		log.Fatal("Failed to connect Temporal client:", err)
	}

	app.AppContainer(fiberHTTP, db)
	portString := fmt.Sprintf(":%v", params.Port)

	fiberHTTP.Post("/buy-ticket", func(c *fiber.Ctx) error {
		request := new(dto.TicketPurchaseRequest)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
		}
		// Start the Temporal Workflow
		workflowOptions := client.StartWorkflowOptions{
			ID:        "ticket_purchase_" + request.UserID + "_" + request.ConcertID + "_" + uuid.New().String(),
			TaskQueue: "TICKET_TASK_QUEUE",
		}
		we, err := temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, workflows.TicketPurchaseWorkflow, *request)
		if err != nil {
			log.Println("Unable to start workflow", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to start workflow")
		}

		log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
		return c.SendString("Ticket purchase initiated successfully")
	})

	fiberHTTP.Post("/registers", func(c *fiber.Ctx) error {
		request := new(dto.RegistrationData)
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
		}
		// Start the Temporal Workflow
		workflowOptions := client.StartWorkflowOptions{
			ID:        "regis_" + request.IDCard + "_" + uuid.New().String(),
			TaskQueue: "REGISTER_TASK_QUEUE",
		}

		ctx := context.Background()

		if err := database.InjectTx(ctx, db); err != nil {
			return c.SendString("Register initiated failed")
		}

		we, err := temporalClient.ExecuteWorkflow(ctx, workflowOptions, workflows.RegistrationWorkflow, *request)
		if err != nil {
			log.Println("Unable to start workflow", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to start workflow")
		}

		log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
		return c.SendString("Register initiated successfully")
	})

	// systemFieldRepo := repositories.NewSystemFieldRepo(db)
	// trans := database.NewTransactorRepo(db)
	// systemFieldService := services.NewSystemFieldService(systemFieldRepo, trans)

	// resolver := &graphqlAdapter.Resolver{
	// 	SystemFieldService: systemFieldService,
	// 	TransactorRepo:     trans,
	// }

	// queryHandler := graphqlHandler.NewDefaultServer(
	// 	graphqlAdapter.NewExecutableSchema(
	// 		graphqlAdapter.Config{
	// 			Resolvers: resolver,
	// 		},
	// 	),
	// )
	// fiberHTTP.Post("/query", func(c *fiber.Ctx) error {
	// 	fasthttpadaptor.NewFastHTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
	// 		queryHandler.ServeHTTP(writer, request)
	// 	})(c.Context())

	// 	return nil
	// })

	err = fiberHTTP.Listen(portString)

	if err != nil {
		log.Fatal("Failed to start golang Fiber server:", err)
	}

}
