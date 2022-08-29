package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"market-api/internal/core/domain"
	mock_domain "market-api/test"
	"testing"
)

func TestMarketFindService_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	marketMock := mock_domain.NewMockIMarket(ctrl)
	persistenceMock := mock_domain.NewMockIMarketReaderPersistence(ctrl)
	persistenceMock.EXPECT().FindByID(gomock.Any()).Return(marketMock, nil).AnyTimes()

	service := MarketFindService{
		Persistence: persistenceMock,
	}

	result, err := service.GetById("123")
	require.Nil(t, err)
	require.Equal(t, marketMock, result)
}

func TestMarketFindService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	marketMock := mock_domain.NewMockIMarket(ctrl)
	markets := append(make([]domain.IMarket, 0), marketMock)

	persistenceMock := mock_domain.NewMockIMarketReaderPersistence(ctrl)
	persistenceMock.EXPECT().FindAll().Return(markets, nil).AnyTimes()

	service := MarketFindService{
		Persistence: persistenceMock,
	}

	result, err := service.GetAll()
	require.Nil(t, err)
	require.Equal(t, markets, result)
}

func TestMarketFindService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	marketMock := mock_domain.NewMockIMarket(ctrl)
	markets := append(make([]domain.IMarket, 0), marketMock)

	persistenceMock := mock_domain.NewMockIMarketReaderPersistence(ctrl)
	persistenceMock.EXPECT().Find(gomock.Any()).Return(markets, nil).AnyTimes()

	service := MarketFindService{
		Persistence: persistenceMock,
	}

	result, err := service.Get("township", "region5", "name", "district")
	require.Nil(t, err)
	require.Equal(t, markets, result)
}
