package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"market-api/internal/core/domain"
	mock_domain "market-api/test"

	"testing"
)

func TestCreateMarketService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var id int64 = 1
	persistenceMock := mock_domain.NewMockIMarketWriterPersistence(ctrl)
	persistenceMock.EXPECT().Save(gomock.Any()).Return(id, nil).AnyTimes()

	service := MarketCreateService{
		Persistence: persistenceMock,
	}

	market := domain.NewMarket()
	market.Longitude = "-11111111"
	market.Latitude = "-22222222"
	market.CensusSector = "333333333333333"
	market.WeightingArea = "4444444444444"
	market.TownshipCode = "555555555"
	market.Township = "township"
	market.SubPrefectureCode = "66"
	market.SubPrefecture = "subPrefecture"
	market.Region5 = "region5"
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
