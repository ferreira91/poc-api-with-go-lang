package postgres

import (
	"database/sql"
	"market-api/internal/core/domain"
	"strconv"
)

type MarketEntity struct {
	ID                int64  `valid:"required"`
	Longitude         string `valid:"required"`
	Latitude          string `valid:"required"`
	CensusSector      string `valid:"required"`
	WeightingArea     string `valid:"required"`
	Township          string `valid:"required"`
	TownshipCode      string `valid:"required"`
	SubprefectureCode string `valid:"required"`
	Subprefecture     string `valid:"required"`
	Region5           string `valid:"required"`
	Region8           string `valid:"required"`
	Name              string `valid:"required"`
	Registry          string `valid:"required"`
	Street            string `valid:"required"`
	Number            sql.NullString
	District          sql.NullString
	Reference         sql.NullString
}

func (m MarketEntity) ToMarketDomain() *domain.Market {
	market := domain.NewMarket()
	market.ID = strconv.FormatInt(m.ID, 10)
	market.Longitude = m.Longitude
	market.Latitude = m.Latitude
	market.CensusSector = m.CensusSector
	market.WeightingArea = m.WeightingArea
	market.TownshipCode = m.TownshipCode
	market.Township = m.Township
	market.SubprefectureCode = m.SubprefectureCode
	market.Subprefecture = m.Subprefecture
	market.Region5 = m.Region5
	market.Region8 = m.Region8
	market.Name = m.Name
	market.Registry = m.Registry
	market.Street = m.Street
	market.Number = m.Number.String
	market.District = m.District.String
	market.Reference = m.Reference.String

	return market
}
