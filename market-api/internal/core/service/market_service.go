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

func (s *MarketService) Get(
	longitude string,
	latitude string,
	censusSector string,
	weightingArea string,
	township string,
	townshipCode string,
	subprefectureCode string,
	subprefecture string,
	region5 string,
	region8 string,
	name string,
	registry string,
	street string,
	number string,
	district string,
	reference string,
) ([]domain.IMarket, error) {
	result, err := s.Persistence.Find(
		longitude,
		latitude,
		censusSector,
		weightingArea,
		township,
		townshipCode,
		subprefectureCode,
		subprefecture,
		region5,
		region8,
		name,
		registry,
		street,
		number,
		district,
		reference,
	)
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
