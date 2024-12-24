package data

import (
	"pigeon2/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type workFlowRepo struct {
	data *Data
	log  *log.Helper
}

func NewWorkFlowRepo(data *Data, logger log.Logger) biz.WorkFlowRepo {
	return &workFlowRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
