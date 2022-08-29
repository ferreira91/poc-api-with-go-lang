package service

import "market-api/internal/core/domain"

type MarketUpdateService struct {
	Persistence domain.IMarketWriterPersistence
}

func (s *MarketUpdateService) Update(id string, market domain.IMarket) (domain.IMarket, error) {
	_, err := market.IsValid()
	if err != nil {
		return &domain.Market{}, err
	}
	result, err := s.Persistence.Update(id, market)
	if err != nil {
		return &domain.Market{}, err
	}
	return result, nil
}
