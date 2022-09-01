package repository

import (
	"github.com/orlangure/gnomock"
	"github.com/stretchr/testify/require"
	"market-api/internal/core/domain"
	"market-api/test/postgres"
	"testing"
)

func TestMarketDb_DeleteByRegistry(t *testing.T) {
	container, db, _ := postgres.SetUp("market")
	t.Cleanup(func() { _ = gnomock.Stop(container) })

	marketDb := NewMarketDb(db)

	result, err := marketDb.DeleteByRegistry("654321")
	require.Nil(t, err)
	require.Equal(t, result, int64(1))

	result, err = marketDb.DeleteByRegistry("654321")
	require.Nil(t, err)
	require.Equal(t, result, int64(0))
}

func TestMarketDb_Save(t *testing.T) {
	container, db, _ := postgres.SetUp("market")
	t.Cleanup(func() { _ = gnomock.Stop(container) })

	market := domain.NewMarket()
	market.Longitude = "-11111111"
	market.Latitude = "-22222222"
	market.CensusSector = "333333333333333"
	market.WeightingArea = "4444444444444"
	market.TownshipCode = "555555555"
	market.Township = "township"
	market.SubPrefectureCode = "66"
	market.SubPrefecture = "subPrefecture"
	market.Region5 = "region"
	market.Region8 = "region8"
	market.Name = "name"
	market.Registry = "666666"
	market.Street = "street"
	market.Number = "777777777777777"
	market.District = "district"
	market.Reference = "reference"

	marketDb := NewMarketDb(db)

	result, err := marketDb.Save(market)
	require.Nil(t, err)
	require.Equal(t, result, 3)

	result, err = marketDb.Save(market)
	require.Error(t, err, "duplicate key value violates unique constraint \"market_registry_key\"")
}

func TestMarketDb_Update(t *testing.T) {
	container, db, _ := postgres.SetUp("market")
	t.Cleanup(func() { _ = gnomock.Stop(container) })

	market := domain.NewMarket()
	market.Longitude = "-11111111"
	market.Latitude = "-22222222"
	market.CensusSector = "test"
	market.WeightingArea = "test"
	market.TownshipCode = "555555555"
	market.Township = "test"
	market.SubPrefectureCode = "55"
	market.SubPrefecture = "test"
	market.Region5 = "test"
	market.Region8 = "test"
	market.Name = "test"
	market.Registry = "123456"
	market.Street = "test"
	market.Number = "777777777777777"
	market.District = "test"
	market.Reference = "test"

	marketDb := NewMarketDb(db)

	result, err := marketDb.Update(1, market)
	market.ID = 1
	require.Nil(t, err)
	require.Equal(t, result, market)

	market.Registry = "555555"
	result, err = marketDb.Update(1, market)
	require.Error(t, err, "sql: no rows in result set")
}