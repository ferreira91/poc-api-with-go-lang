package repository

import (
	"github.com/orlangure/gnomock"
	"github.com/stretchr/testify/require"
	"market-api/test/postgres"
	"testing"
)

func TestMarketDb_Find(t *testing.T) {
	container, db, _ := postgres.SetUp("market")
	t.Cleanup(func() { _ = gnomock.Stop(container) })

	marketDb := NewMarketDb(db)

	query := make(map[string]string)
	query["Longitude"] = "-1111111"
	query["Latitude"] = "-9999999"
	query["CensusSector"] = "census_sector"
	query["WeightingArea"] = "weighting_are"
	query["TownshipCode"] = "123"
	query["Township"] = "township"
	query["SubprefectureCode"] = "00"
	query["Subprefecture"] = "subprefecture"
	query["Region5"] = "region"
	query["Region8"] = "region_"
	query["Name"] = "name"
	query["Registry"] = "123456"
	query["Street"] = "street"
	query["Number"] = "123"
	query["District"] = "district"
	query["Reference"] = "reference"

	result, err := marketDb.Find(query)
	require.Nil(t, err)
	require.Equal(t, result[0].GetID(), int64(1))
	require.Equal(t, len(result), 1)
}

func TestMarketDb_FindAll(t *testing.T) {
	container, db, _ := postgres.SetUp("market")
	t.Cleanup(func() { _ = gnomock.Stop(container) })

	marketDb := NewMarketDb(db)

	result, err := marketDb.FindAll()
	require.Nil(t, err)
	require.Equal(t, result[0].GetID(), int64(1))
	require.Equal(t, result[1].GetID(), int64(2))
	require.Equal(t, len(result), 2)
}

func TestMarketDb_FindByID(t *testing.T) {
	container, db, _ := postgres.SetUp("market")
	t.Cleanup(func() { _ = gnomock.Stop(container) })

	marketDb := NewMarketDb(db)

	result, err := marketDb.FindByID(1)
	require.Nil(t, err)
	require.Equal(t, result.GetRegistry(), "123456")

	result, err = marketDb.FindByID(999)
	require.Error(t, err, "sql: no rows in result set")
}
