package repository

import "market-api/internal/core/domain"

func (m *MarketDb) Save(market domain.IMarket) (int, error) {
	stmt, err := m.db.Prepare(
		`INSERT INTO market (
                    longitude,
                    latitude,
                    census_sector,
                    weighting_area,
                    township_code,
                    township,
                    subprefecture_code,
                    subprefecture,
                    region_5,
                    region_8,
                    name,
                    registry,
                    street,
                    number,
                    district,
                    reference
              ) 
              VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16) 
              RETURNING id`,
	)
	if err != nil {
		return 0, err
	}
	var lastInsertId int
	err = stmt.QueryRow(
		market.GetLongitude(),
		market.GetLatitude(),
		market.GetCensusSector(),
		market.GetWeightingArea(),
		market.GetTownshipCode(),
		market.GetTownship(),
		market.GetSubPrefectureCode(),
		market.GetSubPrefecture(),
		market.GetRegion5(),
		market.GetRegion8(),
		market.GetName(),
		market.GetRegistry(),
		market.GetStreet(),
		market.GetNumber(),
		market.GetDistrict(),
		market.GetReference(),
	).Scan(&lastInsertId)

	if err != nil {
		return 0, err
	}
	err = stmt.Close()
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}
func (m *MarketDb) Update(id int64, market domain.IMarket) (domain.IMarket, error) {
	stmt, err := m.db.Prepare(
		`UPDATE market SET 
					longitude = $3,
                    latitude = $4,
                    census_sector = $5,
                    weighting_area = $6,
                    township_code = $7,
                    township = $8,
                    subprefecture_code = $9,
                    subprefecture = $10,
                    region_5 = $11,
                    region_8 = $12,
                    name = $13,
                    street = $14,
                    number = $15,
                    district = $16,
                    reference = $17
				WHERE id = $1 AND registry = $2
				RETURNING id,longitude,latitude,census_sector,weighting_area,township_code,township,subprefecture_code,subprefecture,region_5,region_8,name,registry,street,number,district,reference`,
	)
	if err != nil {
		return nil, err
	}

	var marketUpdated domain.Market

	err = stmt.QueryRow(
		id,
		market.GetRegistry(),
		market.GetLongitude(),
		market.GetLatitude(),
		market.GetCensusSector(),
		market.GetWeightingArea(),
		market.GetTownshipCode(),
		market.GetTownship(),
		market.GetSubPrefectureCode(),
		market.GetSubPrefecture(),
		market.GetRegion5(),
		market.GetRegion8(),
		market.GetName(),
		market.GetStreet(),
		market.GetNumber(),
		market.GetDistrict(),
		market.GetReference(),
	).Scan(
		&marketUpdated.ID,
		&marketUpdated.Longitude,
		&marketUpdated.Latitude,
		&marketUpdated.CensusSector,
		&marketUpdated.WeightingArea,
		&marketUpdated.TownshipCode,
		&marketUpdated.Township,
		&marketUpdated.SubPrefectureCode,
		&marketUpdated.SubPrefecture,
		&marketUpdated.Region5,
		&marketUpdated.Region8,
		&marketUpdated.Name,
		&marketUpdated.Registry,
		&marketUpdated.Street,
		&marketUpdated.Number,
		&marketUpdated.District,
		&marketUpdated.Reference,
	)
	if err != nil {
		return nil, err
	}

	return &marketUpdated, nil
}
func (m *MarketDb) DeleteByRegistry(registry string) (int64, error) {
	stmt, err := m.db.Prepare(`DELETE FROM market WHERE registry=$1`)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(registry)
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return count, nil
}
