package postgres

import (
	"database/sql"
	"fmt"
	"market-api/internal/core/domain"
	"market-api/utils"
	"strconv"
)

type MarketDb struct {
	db *sql.DB
}

func NewMarketDb(db *sql.DB) *MarketDb {
	return &MarketDb{db: db}
}

func (m *MarketDb) Save(market domain.IMarket) (string, error) {
	utils.LoggerInfo("persistence - save market start")
	stmt, err := m.db.Prepare(
		`INSERT INTO market (
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
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)`,
	)
	if err != nil {
		utils.LoggerError("persistence - create market error generate query", err)
		return "", err
	}

	id, err := m.getId()
	if err != nil {
		utils.LoggerError("persistence - create market error get id", err)
		return "", err
	}

	_, err = stmt.Exec(
		id,
		market.GetLongitude(),
		market.GetLatitude(),
		market.GetCensusSector(),
		market.GetWeightingArea(),
		market.GetTownshipCode(),
		market.GetTownship(),
		market.GetSubprefectureCode(),
		market.GetSubprefecture(),
		market.GetRegion5(),
		market.GetRegion8(),
		market.GetName(),
		market.GetRegistry(),
		market.GetStreet(),
		market.GetNumber(),
		market.GetDistrict(),
		market.GetReference(),
	)

	if err != nil {
		utils.LoggerError("persistence - create market error execute query", err)
		return "", err
	}

	if err = stmt.Close(); err != nil {
		return "", err
	}

	return strconv.FormatInt(id, 10), nil
}

func (m *MarketDb) FindByID(id string) (domain.IMarket, error) {
	utils.LoggerInfo("persistence - find market by id start")
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
		utils.LoggerError("persistence - find market by id error generate query", err)
		return nil, err
	}

	var entity MarketEntity
	err = stmt.QueryRow(id).Scan(
		&entity.ID,
		&entity.Longitude,
		&entity.Latitude,
		&entity.CensusSector,
		&entity.WeightingArea,
		&entity.TownshipCode,
		&entity.Township,
		&entity.SubprefectureCode,
		&entity.Subprefecture,
		&entity.Region5,
		&entity.Region8,
		&entity.Name,
		&entity.Registry,
		&entity.Street,
		&entity.Number,
		&entity.District,
		&entity.Reference,
	)
	if err != nil {
		utils.LoggerError("persistence - find market by id error execute query", err)
		return nil, err
	}
	return entity.ToMarketDomain(), nil
}

func (m *MarketDb) FindAll() ([]domain.IMarket, error) {
	utils.LoggerInfo("persistence - find all markets start")
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
		utils.LoggerError("persistence - find all market error generate query", err)
		return nil, err
	}

	var rows *sql.Rows
	rows, err = stmt.Query()
	if err != nil {
		return nil, err
	}

	var markets []domain.IMarket
	for rows.Next() {
		var entity MarketEntity
		if err = rows.Scan(
			&entity.ID,
			&entity.Longitude,
			&entity.Latitude,
			&entity.CensusSector,
			&entity.WeightingArea,
			&entity.TownshipCode,
			&entity.Township,
			&entity.SubprefectureCode,
			&entity.Subprefecture,
			&entity.Region5,
			&entity.Region8,
			&entity.Name,
			&entity.Registry,
			&entity.Street,
			&entity.Number,
			&entity.District,
			&entity.Reference,
		); err != nil {
			utils.LoggerError("persistence - find all markets error execute query", err)
			return nil, err
		}
		markets = append(markets, entity.ToMarketDomain())
	}

	return markets, nil
}

func (m *MarketDb) Find(
	longitude string,
	latitude string,
	censusSector string,
	weightingArea string,
	township string,
	townshipCode string,
	subprefectureCode string,
	subprefecture string,
	region5 string,
	region8 string,
	name string,
	registry string,
	street string,
	number string,
	district string,
	reference string,
) ([]domain.IMarket, error) {
	utils.LoggerInfo("persistence - find markets start")
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
					FROM market %s`, generateQueryFilter(
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
			),
		),
	)
	if err != nil {
		utils.LoggerError("persistence - find market error generate query", err)
		return nil, err
	}

	var rows *sql.Rows
	rows, err = stmt.Query()
	if err != nil {
		return nil, err
	}

	markets := []domain.IMarket{}
	for rows.Next() {
		var entity MarketEntity
		if err = rows.Scan(
			&entity.ID,
			&entity.Longitude,
			&entity.Latitude,
			&entity.CensusSector,
			&entity.WeightingArea,
			&entity.TownshipCode,
			&entity.Township,
			&entity.SubprefectureCode,
			&entity.Subprefecture,
			&entity.Region5,
			&entity.Region8,
			&entity.Name,
			&entity.Registry,
			&entity.Street,
			&entity.Number,
			&entity.District,
			&entity.Reference,
		); err != nil {
			utils.LoggerError("persistence - find markets error execute query", err)
			return nil, err
		}
		markets = append(markets, entity.ToMarketDomain())
	}

	return markets, nil
}

func (m *MarketDb) Update(id string, market domain.IMarket) (domain.IMarket, error) {
	utils.LoggerInfo("persistence - update market start")
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
		utils.LoggerError("persistence - update market error generate query", err)
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
		market.GetSubprefectureCode(),
		market.GetSubprefecture(),
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
		&marketUpdated.SubprefectureCode,
		&marketUpdated.Subprefecture,
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
		utils.LoggerError("persistence - update market error execute query", err)
		return nil, err
	}

	return &marketUpdated, nil
}

func (m *MarketDb) DeleteByRegistry(registry string) error {
	utils.LoggerInfo("persistence - delete market by registry start")
	var id *string
	stmt, err := m.db.Prepare(`SELECT id FROM market WHERE registry=$1`)
	if err != nil {
		utils.LoggerError("persistence - delete market by registry error generate query", err)
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
		utils.LoggerError("persistence - delete market by registry error execute query", err)
		return err
	}
	return nil
}

func generateQueryFilter(
	longitude string,
	latitude string,
	censusSector string,
	weightingArea string,
	township string,
	townshipCode string,
	subprefectureCode string,
	subprefecture string,
	region5 string,
	region8 string,
	name string,
	registry string,
	street string,
	number string,
	district string,
	reference string,
) string {
	where := "where 1=1"

	if ok := longitude != ""; ok {
		where += " AND longitude='" + longitude + "'"
	}

	if ok := latitude != ""; ok {
		where += " AND latitude='" + latitude + "'"
	}

	if ok := censusSector != ""; ok {
		where += " AND census_sector='" + censusSector + "'"
	}

	if ok := weightingArea != ""; ok {
		where += " AND weighting_area='" + weightingArea + "'"
	}

	if ok := townshipCode != ""; ok {
		where += " AND township_code='" + townshipCode + "'"
	}

	if ok := township != ""; ok {
		where += " AND township='" + township + "'"
	}

	if ok := subprefectureCode != ""; ok {
		where += " AND subprefecture_code='" + subprefectureCode + "'"
	}

	if ok := subprefecture != ""; ok {
		where += " AND subprefecture='" + subprefecture + "'"
	}

	if ok := region5 != ""; ok {
		where += " AND region_5='" + region5 + "'"
	}

	if ok := region8 != ""; ok {
		where += " AND region_8='" + region8 + "'"
	}

	if ok := name != ""; ok {
		where += " AND name='" + name + "'"
	}

	if ok := registry != ""; ok {
		where += " AND registry='" + registry + "'"
	}

	if ok := street != ""; ok {
		where += " AND street='" + street + "'"
	}

	if ok := number != ""; ok {
		where += " AND number='" + number + "'"
	}

	if ok := district != ""; ok {
		where += " AND district='" + district + "'"
	}

	if ok := reference != ""; ok {
		where += " AND reference='" + reference + "'"
	}

	return where
}

func (m *MarketDb) getId() (int64, error) {
	stmt, err := m.db.Prepare(`SELECT SETVAL('market_id_seq', (SELECT MAX(id) FROM market)+1)`)
	if err != nil {
		return 0, err
	}
	var id int64
	err = stmt.QueryRow().Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err
}
