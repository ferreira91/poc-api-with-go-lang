package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"market-api/internal/core/domain"
	mockdomain "market-api/test"
	"testing"
)

func TestMarketService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var id = "1"
	persistenceMock := mockdomain.NewMockIMarketPersistence(ctrl)
	persistenceMock.EXPECT().Save(gomock.Any()).Return(id, nil).AnyTimes()
	service := MarketService{Persistence: persistenceMock}

	market := domain.NewMarket()
	market.Longitude = "-11111111"
	market.Latitude = "-22222222"
	market.CensusSector = "333333333333333"
	market.WeightingArea = "4444444444444"
	market.TownshipCode = "555555555"
	market.Township = "township"
	market.SubprefectureCode = "66"
	market.Subprefecture = "subPrefecture"
	market.Region5 = "region"
	market.Region8 = "region8"
	market.Name = "name"
	market.Registry = "666666"
	market.Street = "street"
	market.Number = "777777777777777"
	market.District = "district"
	market.Reference = "reference"

	result, err := service.Create(market)
	require.Nil(t, err)
	require.Equal(t, id, result)
}

func TestMarketService_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	marketMock := mockdomain.NewMockIMarket(ctrl)
	persistenceMock := mockdomain.NewMockIMarketPersistence(ctrl)
	persistenceMock.EXPECT().FindByID(gomock.Any()).Return(marketMock, nil).AnyTimes()
	service := MarketService{Persistence: persistenceMock}

	result, err := service.GetByID("123")
	require.Nil(t, err)
	require.Equal(t, marketMock, result)
}

func TestMarketService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	marketMock := mockdomain.NewMockIMarket(ctrl)
	markets := append(make([]domain.IMarket, 0), marketMock)

	persistenceMock := mockdomain.NewMockIMarketPersistence(ctrl)
	persistenceMock.EXPECT().FindAll().Return(markets, nil).AnyTimes()
	service := MarketService{Persistence: persistenceMock}

	result, err := service.GetAll()
	require.Nil(t, err)
	require.Equal(t, markets, result)
}

func TestMarketService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	marketMock := mockdomain.NewMockIMarket(ctrl)
	markets := append(make([]domain.IMarket, 0), marketMock)

	persistenceMock := mockdomain.NewMockIMarketPersistence(ctrl)
	a := gomock.Any()
	persistenceMock.EXPECT().Find(
		a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a,
	).Return(markets, nil).AnyTimes()
	service := MarketService{Persistence: persistenceMock}

	result, err := service.Get("", "", "", "", "", "", "", "", "", "", "", "registry", "", "", "", "")
	require.Nil(t, err)
	require.Equal(t, markets, result)
}

func TestMarketService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	marketMock := mockdomain.NewMockIMarket(ctrl)
	persistenceMock := mockdomain.NewMockIMarketPersistence(ctrl)
	persistenceMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(marketMock, nil).AnyTimes()
	service := MarketService{Persistence: persistenceMock}

	market := domain.NewMarket()
	market.Longitude = "-11111111"
	market.Latitude = "-22222222"
	market.CensusSector = "333333333333333"
	market.WeightingArea = "4444444444444"
	market.TownshipCode = "555555555"
	market.Township = "township"
	market.SubprefectureCode = "66"
	market.Subprefecture = "subPrefecture"
	market.Region5 = "region"
	market.Region8 = "region8"
	market.Name = "name"
	market.Registry = "666666"
	market.Street = "street"
	market.Number = "777777777777777"
	market.District = "district"
	market.Reference = "reference"

	result, err := service.Update("123", market)
	require.Nil(t, err)
	require.Equal(t, marketMock, result)
}

func TestMarketService_DeleteByRegistry(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	persistenceMock := mockdomain.NewMockIMarketPersistence(ctrl)
	persistenceMock.EXPECT().DeleteByRegistry(gomock.Any()).Return(nil).AnyTimes()
	service := MarketService{Persistence: persistenceMock}

	err := service.DeleteByRegistry("registry")
	require.Nil(t, err)
}
