package model

type WorkNode struct {
	ID         int64        `json:"id" gorm:"primary_key;column:id;comment:id"`                    // ID
	ParentID   int          `json:"parent_id" gorm:"column:parent_id;comment:父节点id"`               // 父节点ID
	WorkFlowID int64        `json:"work_flow_id" gorm:"column:work_flow_id;notnull;comment:工作流id"` // 工作流ID
	Type       int          `json:"type" gorm:"column:type;notnull;comment:工作节点类型"`                // 节点类型
	Name       string       `json:"name" gorm:"column:name;notnull;comment:节点名称"`                  // 节点名称
	Condition  []*Condition `json:"condition,,omitempty" gorm:"-"`                                 // 条件
}

func (WorkNode) TableName() string {
	return "work_node"
}

type Condition struct {
	ID              int64  `json:"id" gorm:"primary_key;column:id;comment:id"`                                        // ID
	WorkNodeID      int    `json:"work_node_id" gorm:"column:work_node_id;notnull;comment:工作节点id"`                    // 工作节点ID
	Key             string `json:"key" gorm:"column:key;notnull;comment:条件key"`                                       // 条件key
	Operator        string `json:"operator" gorm:"column:operator;notnull;comment:条件操作符"`                             // 条件操作符
	Value           string `json:"value" gorm:"column:value;notnull;comment:条件值"`                                     // 条件值
	FitNextNodeID   int64  `json:"fit_next_node_id" gorm:"column:fit_next_node_id;notnull;comment:符合条件跳转节点id"`        // 符合条件跳转节点ID
	UnFitNextNodeID int64  `json:"un_fit_next_node_id" gorm:"column:un_fit_next_node_id;notnull;comment:不符合条件跳转节点id"` // 不符合条件跳转节点ID
}

func (Condition) TableName() string {
	return "condition"
}
