package model

import "gorm.io/gorm"

type OrderField struct {
	WorkOrderID int64  `json:"work_order_id" gorm:"column:work_order_id;notnull;comment:工单id"` // 工单ID
	Label       string `json:"label" gorm:"column:label;notnull;comment:字段名称"`                 // 字段名称
	Key         string `json:"key" gorm:"column:key;notnull;comment:字段key"`                    // 字段key
	Value       string `json:"value" gorm:"column:value;notnull;comment:字段值"`                  // 字段值
}

type WorkOrder struct {
	ID          int64          `json:"id" gorm:"primary_key;column:id;comment:id"`                     // ID
	Creator     string         `json:"creator" gorm:"column:creator;notnull;comment:创建人"`              // 创建人
	Name        string         `json:"name" gorm:"column:name;notnull;comment:工单名称"`                   // 工单名称
	ApproveUser string         `json:"approve_user" gorm:"column:approve_user;notnull;comment:审批人"`    // 审批人
	WorkNodeID  int64          `json:"work_node_id" gorm:"column:work_node_id;notnull;comment:工作节点id"` // 工作节点ID
	WorkFlowID  int64          `json:"work_flow_id" gorm:"column:work_flow_id;notnull;comment:工作流id"`  // 工作流ID
	Status      int            `json:"status" gorm:"column:status;notnull;comment:工单状态"`               // 工单状态
	CreatedAt   string         `json:"created_at" gorm:"column:created_at;notnull;comment:创建时间"`       // 创建时间
	UpdatedAt   string         `json:"updated_at" gorm:"column:updated_at;notnull;comment:更新时间"`       // 更新时间
	DeleteAt    gorm.DeletedAt `json:"-" gorm:"column:delete_at;comment:删除时间"`                         // 删除时间
	OrderField  []*OrderField  `json:"order_field,omitempty" gorm:"-"`                                 // 工单字段
}

func (WorkOrder) TableName() string {
	return "work_order"
}
