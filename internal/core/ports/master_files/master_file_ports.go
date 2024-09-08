package ports

import (
	"context"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	"github.com/billowdev/exclusive-go-hexa/pkg/helpers/pagination"
	"github.com/billowdev/exclusive-go-hexa/pkg/utils"
)

type IMasterFileRepository interface {
	GetMasterFile(ctx context.Context, id uint) (*models.MasterFile, error)
	GetMasterFiles(ctx context.Context) (*pagination.Pagination[[]models.MasterFile], error)
	CreateMasterFile(ctx context.Context, payload *models.MasterFile) error
	UpdateMasterFile(ctx context.Context, payload *models.MasterFile) error
	DeleteMasterFile(ctx context.Context, id uint) error
}

type IMasterFileService interface {
	GetMasterFile(ctx context.Context, id uint) utils.APIResponse
	GetMasterFiles(ctx context.Context) pagination.Pagination[[]models.MasterFile]
	CreateMasterFile(ctx context.Context, payload *models.MasterFile) utils.APIResponse
	UpdateMasterFile(ctx context.Context, payload *models.MasterFile) utils.APIResponse
	DeleteMasterFile(ctx context.Context, id uint) utils.APIResponse
}
