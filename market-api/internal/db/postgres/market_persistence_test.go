package postgres

import (
	"github.com/orlangure/gnomock"
	"github.com/stretchr/testify/require"
	"market-api/internal/core/domain"
	"market-api/test/postgres"
	"testing"
)

func TestMarketDb_Save(t *testing.T) {
	container, db, _ := postgres.SetUp()
	t.Cleanup(func() { _ = gnomock.Stop(container) })

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

	marketDb := NewMarketDb(db)

	result, err := marketDb.Save(market)
	require.Nil(t, err)
	require.Equal(t, result, "3")

	result, err = marketDb.Save(market)
	require.Error(t, err, "duplicate key value violates unique constraint \"market_registry_key\"")
}

func TestMarketDb_FindByID(t *testing.T) {
	container, db, _ := postgres.SetUp()
	t.Cleanup(func() { _ = gnomock.Stop(container) })

	marketDb := NewMarketDb(db)

	result, err := marketDb.FindByID("1")
	require.Nil(t, err)
	require.Equal(t, result.GetRegistry(), "123456")

	result, err = marketDb.FindByID("999")
	require.Error(t, err, "sql: no rows in result set")
}

func TestMarketDb_FindAll(t *testing.T) {
	container, db, _ := postgres.SetUp()
	t.Cleanup(func() { _ = gnomock.Stop(container) })

	marketDb := NewMarketDb(db)

	result, err := marketDb.FindAll()
	require.Nil(t, err)
	require.Equal(t, result[0].GetID(), "1")
	require.Equal(t, result[1].GetID(), "2")
	require.Equal(t, len(result), 2)
}

func TestMarketDb_Find(t *testing.T) {
	container, db, _ := postgres.SetUp()
	t.Cleanup(func() { _ = gnomock.Stop(container) })

	marketDb := NewMarketDb(db)

	longitude := "-1111111"
	latitude := "-9999999"
	censusSector := "census_sector"
	weightingArea := "weighting_are"
	townshipCode := "123"
	township := "township"
	subprefectureCode := "00"
	subprefecture := "subprefecture"
	region5 := "region"
	region8 := "region_"
	name := "name"
	registry := "123456"
	street := "street"
	number := "123"
	district := "district"
	reference := "reference"

	result, err := marketDb.Find(
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
	require.Nil(t, err)
	require.Equal(t, result[0].GetID(), "1")
	require.Equal(t, len(result), 1)

	registry = "unknown"
	result, err = marketDb.Find(
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
	require.Nil(t, err)
	require.Empty(t, result)
}

func TestMarketDb_Update(t *testing.T) {
	container, db, _ := postgres.SetUp()
	t.Cleanup(func() { _ = gnomock.Stop(container) })

	market := domain.NewMarket()
	market.Longitude = "-11111111"
	market.Latitude = "-22222222"
	market.CensusSector = "test"
	market.WeightingArea = "test"
	market.TownshipCode = "555555555"
	market.Township = "test"
	market.SubprefectureCode = "55"
	market.Subprefecture = "test"
	market.Region5 = "test"
	market.Region8 = "test"
	market.Name = "test"
	market.Registry = "123456"
	market.Street = "test"
	market.Number = "777777777777777"
	market.District = "test"
	market.Reference = "test"

	marketDb := NewMarketDb(db)

	result, err := marketDb.Update("1", market)
	market.ID = "1"
	require.Nil(t, err)
	require.Equal(t, result, market)

	market.Registry = "555555"
	result, err = marketDb.Update("1", market)
	require.Error(t, err, "sql: no rows in result set")
}

func TestMarketDb_DeleteByRegistry(t *testing.T) {
	container, db, _ := postgres.SetUp()
	t.Cleanup(func() { _ = gnomock.Stop(container) })

	marketDb := NewMarketDb(db)

	err := marketDb.DeleteByRegistry("654321")
	require.Nil(t, err)
	require.Equal(t, err, nil)

	err = marketDb.DeleteByRegistry("654321")
	require.Error(t, err, "sql: no rows in result set")
}
