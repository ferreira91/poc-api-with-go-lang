package web

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"market-api/internal/core/domain"
	"net/http"
)

type (
	MarketDTO struct {
		Longitude         string `json:"longitude" validate:"required,max=10"`
		Latitude          string `json:"latitude" validate:"required,max=10"`
		CensusSector      string `json:"censusSector" validate:"required,max=15"`
		WeightingArea     string `json:"weightingArea" validate:"required,max=13"`
		TownshipCode      string `json:"townshipCode" validate:"required,max=9"`
		Township          string `json:"township" validate:"required,max=18"`
		SubPrefectureCode string `json:"subPrefectureCode" validate:"required,max=2"`
		SubPrefecture     string `json:"subPrefecture" validate:"required,max=25"`
		Region5           string `json:"region5" validate:"required,max=6"`
		Region8           string `json:"region8" validate:"required,max=7"`
		Name              string `json:"name" validate:"required,max=30"`
		Registry          string `json:"registry" validate:"required,max=6"`
		Street            string `json:"street" validate:"required,max=34"`
		Number            string `json:"number,max=15"`
		District          string `json:"district" validate:"required,max=20"`
		Reference         string `json:"reference,max=30"`
	}

	MarketGetParam struct {
		Township string `query:"township"`
		Region5  string `query:"region5"`
		Name     string `query:"name"`
		District string `query:"district"`
	}

	MarketDeleteParam struct {
		Registry string `query:"registry"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (c CustomValidator) Validate(i interface{}) error {
	if err := c.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (m MarketDTO) ToMarketDomain() *domain.Market {
	market := domain.NewMarket()
	market.Longitude = m.Longitude
	market.Latitude = m.Latitude
	market.CensusSector = m.CensusSector
	market.WeightingArea = m.WeightingArea
	market.TownshipCode = m.TownshipCode
	market.Township = m.Township
	market.SubPrefectureCode = m.SubPrefectureCode
	market.SubPrefecture = m.SubPrefecture
	market.Region5 = m.Region5
	market.Region8 = m.Region8
	market.Name = m.Name
	market.Registry = m.Registry
	market.Street = m.Street
	market.Number = m.Number
	market.District = m.District
	market.Reference = m.Reference

	return market
}
