package service

import "market-api/internal/core/domain"

type MarketCreateService struct {
	Persistence domain.IMarketWriterPersistence
}

func NewMarketCreateService(persistence domain.IMarketWriterPersistence) *MarketCreateService {
	return &MarketCreateService{Persistence: persistence}
}

func (s *MarketCreateService) Create(market domain.IMarket) (int64, error) {
	_, err := market.IsValid()
	if err != nil {
		return 0, err
	}
	result, err := s.Persistence.Save(market)
	if err != nil {
		return 0, err
	}
	return result, nil
}
