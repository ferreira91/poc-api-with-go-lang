package web

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"market-api/internal/core/domain"
	"net/http"
)

func CreateMarket(s *Server, ctx echo.Context) (err error) {
	dto := new(MarketRequestDTO)

	if err = ctx.Bind(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, apiError(err.Error()))
	}
	if err = ctx.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, apiError(err.Error()))
	}

	res, err := s.Service.Create(dto.ToMarketDomain())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, apiError(err.Error()))
	}

	req := ctx.Request()
	location := fmt.Sprintf("%s/%s/%s", req.Host, req.RequestURI, res)
	h := ctx.Response().Header()
	h.Add(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	h.Add("Location", location)
	ctx.NoContent(http.StatusCreated)
	return
}

func GetMarketByID(s *Server, ctx echo.Context) (err error) {
	id := ctx.Param("id")

	res, err := s.Service.GetByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, apiError(err.Error()))
	}
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return ctx.JSON(http.StatusOK, ToMarketDTO(res))
}

func GetMarkets(s *Server, ctx echo.Context) (err error) {
	var params MarketGetParam
	if err = (&echo.DefaultBinder{}).BindQueryParams(ctx, &params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, apiError(err.Error()))
	}

	var res []domain.IMarket
	if params == (MarketGetParam{}) {
		res, err = s.Service.GetAll()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, apiError(err.Error()))
		}
	} else {
		res, err = s.Service.Get(params.Township, params.Region5, params.Name, params.District)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, apiError(err.Error()))
		}
	}

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return ctx.JSON(http.StatusOK, ToMarketsDTO(res))
}

func UpdateMarket(s *Server, ctx echo.Context) (err error) {
	id := ctx.Param("id")
	dto := new(MarketRequestDTO)

	if err = ctx.Bind(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, apiError(err.Error()))
	}
	if err = ctx.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, apiError(err.Error()))
	}

	market := dto.ToMarketDomain()

	res, err := s.Service.Update(id, market)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, apiError(err.Error()))
	}

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return ctx.JSON(http.StatusOK, ToMarketDTO(res))
}

func DeleteMarket(s *Server, ctx echo.Context) (err error) {
	var param MarketDeleteParam
	if err = (&echo.DefaultBinder{}).BindQueryParams(ctx, &param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, apiError(err.Error()))
	}

	if err := s.Service.DeleteByRegistry(param.Registry); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, apiError(err.Error()))
	}

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	ctx.NoContent(http.StatusNoContent)
	return
}
