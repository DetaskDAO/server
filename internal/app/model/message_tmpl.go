package model

type MessageTmpl struct {
	ID       uint   `gorm:"primarykey"`                                                                   // 主键ID
	Issuer   string `gorm:"size:500;comment:甲方模版" json:"issuer" form:"issuer"`                            // 甲方模版
	IssuerZh string `gorm:"size:500;comment:中文甲方模版" json:"issuer_zh" form:"issuer_zh"`                    // 中文甲方模版
	Worker   string `gorm:"size:500;comment:乙方模版" json:"worker" form:"worker"`                            // 乙方模版
	WorkerZh string `gorm:"size:500;comment:中文乙方模版" json:"worker_zh" form:"worker_zh"`                    // 中文乙方模版
	Status   string `gorm:"column:status;unique;size:30;comment:事件状态" json:"status" form:"status"`        // 事件状态
	Type     uint8  `gorm:"column:type;default:0;comment:消息类型 0: 正常消息 1: 提示消息" json:"type" form:"type"`   // 消息类型 0: 正常消息 1: 提示消息
	Disable  bool   `gorm:"disable;default:false;comment:是否禁用 0: 启用 1: 禁用" json:"disable" form:"disable"` // 是否禁用 0: 启用 1: 禁用
}
