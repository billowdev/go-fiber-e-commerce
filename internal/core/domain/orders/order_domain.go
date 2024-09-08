package domain

import (
	"time"

	"github.com/billowdev/go-fiber-e-commerce/internal/adapters/database/models"
)

type OrderDomain struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt          time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	OrderType          string    `gorm:"orderType" json:"order_type"`
	PortOfLoading      string    `json:"port_of_loading"`
	PortOfDestination  string    `json:"port_of_destination"`
	DescriptionOfGoods string    `json:"description_of_goods"`
}

func ToOrderDomain(data *models.Order) OrderDomain {
	if data == nil {
		return OrderDomain{}
	}

	return OrderDomain{
		ID:                 data.ID,
		OrderType:          data.OrderType,
		CreatedAt:          data.CreatedAt,
		UpdatedAt:          data.UpdatedAt,
		PortOfLoading:      data.PortOfLoading,
		PortOfDestination:  data.PortOfDestination,
		DescriptionOfGoods: data.DescriptionOfGoods,
	}
}

func ToOrderModel(data OrderDomain) *models.Order {
	return &models.Order{
		ID:                 data.ID,
		OrderType:          data.OrderType,
		CreatedAt:          data.CreatedAt,
		UpdatedAt:          data.UpdatedAt,
		PortOfLoading:      data.PortOfLoading,
		PortOfDestination:  data.PortOfDestination,
		DescriptionOfGoods: data.DescriptionOfGoods,
	}
}
