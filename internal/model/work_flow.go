package model

import (
	"time"

	"gorm.io/gorm"
)

type WorkFlow struct {
	ID        int64          `json:"id" gorm:"primary_key;column:id;comment:id"`               // ID
	Name      string         `json:"name" gorm:"column:name;notnull;comment:工作流名称"`            // 工作流名称
	Status    int            `json:"status" gorm:"column:status;notnull;comment:工作流状态"`        // 工作流状态
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;notnull;comment:创建时间"` // 创建时间
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;notnull;comment:更新时间"` // 更新时间
	DeleteAt  gorm.DeletedAt `json:"-" gorm:"column:delete_at;comment:删除时间"`                   // 删除时间
	WorkNodes []*WorkNode     `json:"work_nodes" gorm:"-"`                                      // 工作节点

}

func (WorkFlow) TableName() string {
	return "work_flow"
}
