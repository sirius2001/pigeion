package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
)

// Greeter is a Greeter model.
type WorkFlow struct {
}

// TemproalRepo is a Greater repo.
type TemproalRepo interface {
}

type WorkFlowRepo interface {
}

// GreeterUsecase is a Greeter usecase.
type WorkFlowUsecase struct {
	temporal TemproalRepo
	repo     WorkFlowRepo
	log      *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewWorkFlowUsecase(repo WorkFlowRepo, temporal TemproalRepo, logger log.Logger) *WorkFlowUsecase {
	return &WorkFlowUsecase{repo: repo, temporal: temporal, log: log.NewHelper(logger)}
}
