package service

import "market-api/internal/core/domain"

type MarketFindService struct {
	Persistence domain.IMarketReaderPersistence
}

func (s *MarketFindService) GetById(id string) (domain.IMarket, error) {
	result, err := s.Persistence.FindByID(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *MarketFindService) GetAll() ([]domain.IMarket, error) {
	result, err := s.Persistence.FindAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *MarketFindService) Get(township string, region5 string, name string, district string) ([]domain.IMarket, error) {
	query := make(map[string]string)
	if township != "" {
		query["township"] = township
	}
	if region5 != "" {
		query["region_5"] = region5
	}
	if name != "" {
		query["name"] = name
	}
	if district != "" {
		query["district"] = district
	}
	result, err := s.Persistence.Find(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}
