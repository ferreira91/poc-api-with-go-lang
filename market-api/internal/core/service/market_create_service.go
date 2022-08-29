package service

import "market-api/internal/core/domain"


type MarketCreateService struct {
	Persistence domain.IMarketWriterPersistence
}

func (s *MarketCreateService) Create(market domain.IMarket) (domain.IMarket, error) {
	_, err := market.IsValid()
	if err != nil {
		return &domain.Market{}, err
	}
	result, err := s.Persistence.Save(market)
	if err != nil {
		return &domain.Market{}, err
	}
	return result, nil
}

