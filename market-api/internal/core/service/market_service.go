package service

import "market-api/internal/core/domain"

type MarketService struct {
	Persistence domain.IMarketPersistence
}

func NewMarketService(persistence domain.IMarketPersistence) *MarketService {
	return &MarketService{Persistence: persistence}
}

func (s *MarketService) Create(market domain.IMarket) (int64, error) {
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

func (s *MarketService) GetByID(id int64) (domain.IMarket, error) {
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

func (s *MarketService) Update(id int64, market domain.IMarket) (domain.IMarket, error) {
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
	_, err := s.Persistence.DeleteByRegistry(registry)
	if err != nil {
		return err
	}
	return nil
}
