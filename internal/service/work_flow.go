package service

import "pigeon2/internal/biz"

type WorkFlowService struct {
	uc *biz.WorkFlowUsecase
}

func NewWorkService(uc *biz.WorkFlowUsecase) *WorkFlowService {
	return &WorkFlowService{uc: uc}
}

