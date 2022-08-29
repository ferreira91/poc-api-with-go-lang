package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mock_domain "market-api/test"
	"testing"
)

func TestMarketDeleteService_DeleteByRegistry(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	persistenceMock := mock_domain.NewMockIMarketWriterPersistence(ctrl)
	persistenceMock.EXPECT().DeleteByRegistry(gomock.Any()).Return(int64(1), nil).AnyTimes()

	service := MarketDeleteService{
		Persistence: persistenceMock,
	}

	err := service.DeleteByRegistry("registry")
	require.Nil(t, err)
}
