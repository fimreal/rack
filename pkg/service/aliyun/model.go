package aliyun

type SGRule struct {
	IP     string `json:"ip" validate:"required,ip"`
	SGID   string `json:"sgid" validate:"required"`
	Remark string `json:"remark" validate:"required"`
	Policy string `json:"policy"` // 默认 accept, 可选 drop
}
