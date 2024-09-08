package ports

import (
	"context"

	"github.com/billowdev/go-fiber-e-commerce/internal/adapters/database/models"
	"github.com/billowdev/go-fiber-e-commerce/pkg/helpers/pagination"
	"github.com/billowdev/go-fiber-e-commerce/pkg/utils"
)

type IDocumentVersionRepository interface {
	GetDocumentVersion(ctx context.Context, id uint) (*models.DocumentVersion, error)
	GetDocumentVersions(ctx context.Context) (*pagination.Pagination[[]models.DocumentVersion], error)
	CreateDocumentVersion(ctx context.Context, payload *models.DocumentVersion) error
	UpdateDocumentVersion(ctx context.Context, payload *models.DocumentVersion) error
	DeleteDocumentVersion(ctx context.Context, id uint) error
}

type IDocumentVersionService interface {
	GetDocumentVersion(ctx context.Context, id uint) utils.APIResponse
	GetDocumentVersions(ctx context.Context) pagination.Pagination[[]models.DocumentVersion]
	CreateDocumentVersion(ctx context.Context, payload *models.DocumentVersion) utils.APIResponse
	UpdateDocumentVersion(ctx context.Context, payload *models.DocumentVersion) utils.APIResponse
	DeleteDocumentVersion(ctx context.Context, id uint) utils.APIResponse
}
