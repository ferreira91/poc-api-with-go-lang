package repository

import (
	"database/sql"
	"market-api/internal/core/domain"
)

func (m *MarketDb) FindByID(id int) (domain.IMarket, error) {
	var market domain.Market
	stmt, err := m.db.Prepare(
		"select " +
			"id, " +
			"longitude, " +
			"latitude, " +
			"census_sector, " +
			"weighting_area, " +
			"township_code, " +
			"township, " +
			"subprefecture_code, " +
			"subprefecture, " +
			"region_5, " +
			"region_8, " +
			"name, " +
			"registry, " +
			"street, " +
			"number, " +
			"district, " +
			"reference " +
			"from market " +
			"where id=$1",
	)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(
		&market.ID,
		&market.Longitude,
		&market.Latitude,
		&market.CensusSector,
		&market.WeightingArea,
		&market.Township,
		&market.TownshipCode,
		&market.SubPrefectureCode,
		&market.SubPrefecture,
		&market.Region5,
		&market.Region8,
		&market.Name,
		&market.Registry,
		&market.Street,
		&market.Number,
		&market.District,
		&market.Reference,
	)
	if err != nil {
		return nil, err
	}
	return &market, nil
}

func (m *MarketDb) FindAll() ([]domain.IMarket, error) {
	stmt, err := m.db.Prepare(
		"select " +
			"id, " +
			"longitude, " +
			"latitude, " +
			"census_sector, " +
			"weighting_area, " +
			"township_code, " +
			"township, " +
			"subprefecture_code, " +
			"subprefecture, " +
			"region_5, " +
			"region_8, " +
			"name, " +
			"registry, " +
			"street, " +
			"number, " +
			"district, " +
			"reference " +
			"from market ",
	)
	if err != nil {
		return nil, err
	}

	var rows *sql.Rows
	rows, err = stmt.Query()
	if err != nil {
		return nil, err
	}

	var markets []domain.IMarket
	for rows.Next() {
		var market domain.Market
		if err := rows.Scan(
			&market.ID,
			&market.Longitude,
			&market.Latitude,
			&market.CensusSector,
			&market.WeightingArea,
			&market.Township,
			&market.TownshipCode,
			&market.SubPrefectureCode,
			&market.SubPrefecture,
			&market.Region5,
			&market.Region8,
			&market.Name,
			&market.Registry,
			&market.Street,
			&market.Number,
			&market.District,
			&market.Reference,
		); err != nil {
			return nil, err
		}
		markets = append(markets, &market)
	}

	return markets, nil
}

func (m *MarketDb) Find(query map[string]string) ([]domain.IMarket, error) {
	stmt, err := m.db.Prepare(
		"select " +
			"id, " +
			"longitude, " +
			"latitude, " +
			"census_sector, " +
			"weighting_area, " +
			"township_code, " +
			"township, " +
			"subprefecture_code, " +
			"subprefecture, " +
			"region_5, " +
			"region_8, " +
			"name, " +
			"registry, " +
			"street, " +
			"number, " +
			"district, " +
			"reference " +
			"from market " + generateQueryFilter(query),
	)
	if err != nil {
		return nil, err
	}

	var rows *sql.Rows
	rows, err = stmt.Query()
	if err != nil {
		return nil, err
	}

	var markets []domain.IMarket
	for rows.Next() {
		var market domain.Market
		if err := rows.Scan(
			&market.ID,
			&market.Longitude,
			&market.Latitude,
			&market.CensusSector,
			&market.WeightingArea,
			&market.Township,
			&market.TownshipCode,
			&market.SubPrefectureCode,
			&market.SubPrefecture,
			&market.Region5,
			&market.Region8,
			&market.Name,
			&market.Registry,
			&market.Street,
			&market.Number,
			&market.District,
			&market.Reference,
		); err != nil {
			return nil, err
		}
		markets = append(markets, &market)
	}

	return markets, nil
}

func generateQueryFilter(query map[string]string) string {
	where := "where 1=1"

	if longitude, ok := query["Longitude"]; ok {
		where += " AND longitude='" + longitude + "'"
	}

	if latitude, ok := query["Latitude"]; ok {
		where += " AND latitude='" + latitude + "'"
	}

	if weightingArea, ok := query["WeightingArea"]; ok {
		where += " AND weighting_area='" + weightingArea + "'"
	}

	if townshipCode, ok := query["TownshipCode"]; ok {
		where += " AND township_code='" + townshipCode + "'"
	}

	if township, ok := query["Township"]; ok {
		where += " AND township='" + township + "'"
	}

	if subprefectureCode, ok := query["SubprefectureCode"]; ok {
		where += " AND subprefecture_code='" + subprefectureCode + "'"
	}

	if subprefecture, ok := query["Subprefecture"]; ok {
		where += " AND subprefecture='" + subprefecture + "'"
	}

	if region5, ok := query["Region5"]; ok {
		where += " AND region_5='" + region5 + "'"
	}

	if region8, ok := query["Region8"]; ok {
		where += " AND region_8='" + region8 + "'"
	}

	if name, ok := query["Name"]; ok {
		where += " AND name='" + name + "'"
	}

	if registry, ok := query["Registry"]; ok {
		where += " AND registry='" + registry + "'"
	}

	if street, ok := query["Street"]; ok {
		where += " AND street='" + street + "'"
	}

	if number, ok := query["Number"]; ok {
		where += " AND number='" + number + "'"
	}

	if district, ok := query["District"]; ok {
		where += " AND district='" + district + "'"
	}

	if reference, ok := query["Reference"]; ok {
		where += " AND reference='" + reference + "'"
	}

	return where
}
