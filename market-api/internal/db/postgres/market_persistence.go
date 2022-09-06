package postgres

import (
	"database/sql"
	"fmt"
	"market-api/internal/core/domain"
	"strconv"
)

type MarketDb struct {
	db *sql.DB
}

func NewMarketDb(db *sql.DB) *MarketDb {
	return &MarketDb{db: db}
}

func (m *MarketDb) Save(market domain.IMarket) (string, error) {
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
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING id`,
	)
	if err != nil {
		return "", err
	}
	var lastInsertId int64
	if err = stmt.QueryRow(
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
	).Scan(&lastInsertId); err != nil {
		return "", err
	}

	if err = stmt.Close(); err != nil {
		return "", err
	}

	id := strconv.FormatInt(lastInsertId, 10)
	return id, nil
}

func (m *MarketDb) FindByID(id string) (domain.IMarket, error) {
	var market domain.Market
	stmt, err := m.db.Prepare(
		`SELECT 
			id,
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
		FROM market
		WHERE id=$1`,
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
		`SELECT 
			id,
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
		FROM market`,
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
		fmt.Sprintf(
			`SELECT 
						id,
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
					FROM market %s`, generateQueryFilter(query),
		),
	)
	if err != nil {
		return nil, err
	}

	var rows *sql.Rows
	rows, err = stmt.Query()
	if err != nil {
		return nil, err
	}

	markets := []domain.IMarket{}
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

func (m *MarketDb) Update(id string, market domain.IMarket) (domain.IMarket, error) {
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
		WHERE id = $1 AND registry = $2 RETURNING id,longitude,latitude,census_sector,weighting_area,township_code,township,subprefecture_code,subprefecture,region_5,region_8,name,registry,street,number,district,reference`,
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

func (m *MarketDb) DeleteByRegistry(registry string) error {
	var id *string
	stmt, err := m.db.Prepare(`SELECT id FROM market WHERE registry=$1`)
	if err != nil {
		return err
	}

	if err = stmt.QueryRow(registry).Scan(&id); err != nil {
		return err
	}

	stmt, err = m.db.Prepare(`DELETE FROM market WHERE id=$1`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
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
