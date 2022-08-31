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

	var rowsUpdated int64 = 1
	persistenceMock := mock_domain.NewMockIMarketWriterPersistence(ctrl)
	persistenceMock.EXPECT().DeleteByRegistry(gomock.Any()).Return(rowsUpdated, nil).AnyTimes()

	service := MarketDeleteService{
		Persistence: persistenceMock,
	}

	err := service.DeleteByRegistry("registry")
	require.Nil(t, err)
}
