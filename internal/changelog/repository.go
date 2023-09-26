package changelog

import "github.com/VadimGossip/oraScriptTool/internal/domain"

type Service interface {
	Get() []domain.ChangelogItem
}

type service struct {
}

var _ Service = (*service)(nil)

func NewService() *service {
	return &service{}
}

func (s *service) Get() []domain.ChangelogItem {
	return nil
}
