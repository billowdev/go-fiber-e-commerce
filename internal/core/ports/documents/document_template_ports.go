package ports

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/billowdev/exclusive-go-hexa/pkg/utils"
)

type IDocumentTemplateRepository interface {
	GetDocumentTemplate(ctx context.Context, id uint) (*models.DocumentTemplate, error)
	GetDocumentTemplates(ctx context.Context) (*pagination.Pagination[[]models.DocumentTemplate], error)
	CreateDocumentTemplate(ctx context.Context, payload *models.DocumentTemplate) error
	UpdateDocumentTemplate(ctx context.Context, payload *models.DocumentTemplate) error
	DeleteDocumentTemplate(ctx context.Context, id uint) error
}

type IDocumentTemplateService interface {
	GetDocumentTemplate(ctx context.Context, id uint) utils.APIResponse
	GetDocumentTemplates(ctx context.Context) pagination.Pagination[[]models.DocumentTemplate]
	CreateDocumentTemplate(ctx context.Context, payload *models.DocumentTemplate) utils.APIResponse
	UpdateDocumentTemplate(ctx context.Context, payload *models.DocumentTemplate) utils.APIResponse
	DeleteDocumentTemplate(ctx context.Context, id uint) utils.APIResponse
}
