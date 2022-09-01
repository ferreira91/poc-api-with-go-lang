package service

import "market-api/internal/core/domain"

type MarketDeleteService struct {
	Persistence domain.IMarketWriterPersistence
}

func NewMarketDeleteService(persistence domain.IMarketWriterPersistence) *MarketDeleteService {
	return &MarketDeleteService{Persistence: persistence}
}

func (s *MarketDeleteService) DeleteByRegistry(registry string) error {
	_, err := s.Persistence.DeleteByRegistry(registry)
	if err != nil {
		return err
	}
	return nil
}
