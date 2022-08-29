package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"market-api/internal/core/domain"
	mock_domain "market-api/test"
	"testing"
)

func TestMarketUpdateService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	marketMock := mock_domain.NewMockIMarket(ctrl)
	persistenceMock := mock_domain.NewMockIMarketWriterPersistence(ctrl)
	persistenceMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(marketMock, nil).AnyTimes()

	service := MarketUpdateService{
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

	result, err := service.Update("123", market)
	require.Nil(t, err)
	require.Equal(t, marketMock, result)
}
