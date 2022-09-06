package web

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"market-api/internal/core/domain"
	"net/http"
)

type (
	MarketRequestDTO struct {
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

	MarketResponseDTO struct {
		ID                string `json:"id"`
		Longitude         string `json:"longitude"`
		Latitude          string `json:"latitude"`
		CensusSector      string `json:"censusSector"`
		WeightingArea     string `json:"weightingArea"`
		TownshipCode      string `json:"townshipCode"`
		Township          string `json:"township"`
		SubPrefectureCode string `json:"subPrefectureCode"`
		SubPrefecture     string `json:"subPrefecture"`
		Region5           string `json:"region5"`
		Region8           string `json:"region8"`
		Name              string `json:"name"`
		Registry          string `json:"registry"`
		Street            string `json:"street"`
		Number            string `json:"number"`
		District          string `json:"district"`
		Reference         string `json:"reference"`
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

func (m MarketRequestDTO) ToMarketDomain() *domain.Market {
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

func ToMarketDTO(market domain.IMarket) MarketResponseDTO {
	m := MarketResponseDTO{
		ID:                market.GetID(),
		Longitude:         market.GetLongitude(),
		Latitude:          market.GetLatitude(),
		CensusSector:      market.GetCensusSector(),
		WeightingArea:     market.GetWeightingArea(),
		TownshipCode:      market.GetTownshipCode(),
		Township:          market.GetTownship(),
		SubPrefectureCode: market.GetSubPrefectureCode(),
		SubPrefecture:     market.GetSubPrefecture(),
		Region5:           market.GetRegion5(),
		Region8:           market.GetRegion8(),
		Name:              market.GetName(),
		Registry:          market.GetRegistry(),
		Street:            market.GetStreet(),
		Number:            market.GetNumber(),
		District:          market.GetDistrict(),
		Reference:         market.GetReference(),
	}
	return m
}

func ToMarketsDTO(markets []domain.IMarket) []MarketResponseDTO {
	ms := []MarketResponseDTO{}
	for _, market := range markets {
		m := MarketResponseDTO{
			ID:                market.GetID(),
			Longitude:         market.GetLongitude(),
			Latitude:          market.GetLatitude(),
			CensusSector:      market.GetCensusSector(),
			WeightingArea:     market.GetWeightingArea(),
			TownshipCode:      market.GetTownshipCode(),
			Township:          market.GetTownship(),
			SubPrefectureCode: market.GetSubPrefectureCode(),
			SubPrefecture:     market.GetSubPrefecture(),
			Region5:           market.GetRegion5(),
			Region8:           market.GetRegion8(),
			Name:              market.GetName(),
			Registry:          market.GetRegistry(),
			Street:            market.GetStreet(),
			Number:            market.GetNumber(),
			District:          market.GetDistrict(),
			Reference:         market.GetReference(),
		}
		ms = append(ms, m)
	}

	return ms
}
