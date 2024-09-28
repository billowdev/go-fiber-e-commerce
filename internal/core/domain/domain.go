package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("--------------JSONB Value-----------------")
			err := fmt.Errorf("panic occurred: %v", r)
			log.Println(err)
			log.Println("-------------------------------")
		}
	}()
	valueString, err := json.Marshal(j)
	return string(valueString), err
}
func (j *JSONB) Scan(value interface{}) error {
	defer func() {
		if r := recover(); r != nil {
			log.Println("--------------JSONB Scan-----------------")
			err := fmt.Errorf("panic occurred: %v", r)
			log.Println(err)
			log.Println("-------------------------------")
		}
	}()

	if data, ok := value.([]byte); ok {
		if err := json.Unmarshal(data, j); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unexpected type for JSONB: %T", value)
	}

	return nil
}

func ConvertToJSONB(oldData, newData interface{}) (JSONB, JSONB, error) {
	oldDataJSON, err := json.Marshal(oldData)
	if err != nil {
		return nil, nil, err
	}

	newDataJSON, err := json.Marshal(newData)
	if err != nil {
		return nil, nil, err
	}

	var oldDataMap map[string]interface{}
	if err := json.Unmarshal(oldDataJSON, &oldDataMap); err != nil {
		return nil, nil, err
	}

	var newDataMap map[string]interface{}
	if err := json.Unmarshal(newDataJSON, &newDataMap); err != nil {
		return nil, nil, err
	}

	return JSONB(oldDataMap), JSONB(newDataMap), nil
}
