package service

import "market-api/internal/core/domain"

type MarketService struct {
	Persistence domain.IMarketPersistence
}

func (s *MarketService) Create(market domain.IMarket) (string, error) {
	_, err := market.IsValid()
	if err != nil {
		return "", err
	}
	result, err := s.Persistence.Save(market)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (s *MarketService) GetByID(id string) (domain.IMarket, error) {
	result, err := s.Persistence.FindByID(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *MarketService) GetAll() ([]domain.IMarket, error) {
	result, err := s.Persistence.FindAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *MarketService) Get(township string, region5 string, name string, district string) ([]domain.IMarket, error) {
	query := make(map[string]string)
	if township != "" {
		query["Township"] = township
	}
	if region5 != "" {
		query["Region5"] = region5
	}
	if name != "" {
		query["Name"] = name
	}
	if district != "" {
		query["District"] = district
	}
	result, err := s.Persistence.Find(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *MarketService) Update(id string, market domain.IMarket) (domain.IMarket, error) {
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

func (s *MarketService) DeleteByRegistry(registry string) error {
	err := s.Persistence.DeleteByRegistry(registry)
	if err != nil {
		return err
	}
	return nil
}
