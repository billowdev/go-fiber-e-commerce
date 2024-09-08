package ports

import (
	"context"

	"github.com/billowdev/go-fiber-e-commerce/internal/adapters/database/models"
	domain "github.com/billowdev/go-fiber-e-commerce/internal/core/domain/documents"
	"github.com/billowdev/go-fiber-e-commerce/pkg/helpers/pagination"
	"github.com/billowdev/go-fiber-e-commerce/pkg/utils"
)

type IDocumentRepository interface {
	GetDocument(ctx context.Context, id uint) (*models.Document, error)
	GetDocuments(ctx context.Context) (*pagination.Pagination[[]models.Document], error)
	CreateDocument(ctx context.Context, payload *models.Document) error
	UpdateDocument(ctx context.Context, payload *models.Document) error
	DeleteDocument(ctx context.Context, id uint) error
}

type IDocumentService interface {
	GetDocument(ctx context.Context, id uint) utils.APIResponse
	GetDocuments(ctx context.Context) pagination.PaginationResponse
	CreateDocument(ctx context.Context, payload domain.DocumentDomain) utils.APIResponse
	UpdateDocument(ctx context.Context, payload domain.DocumentDomain) utils.APIResponse
	DeleteDocument(ctx context.Context, id uint) utils.APIResponse
}
