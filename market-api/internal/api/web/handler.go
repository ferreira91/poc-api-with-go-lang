package web

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"market-api/internal/core/domain"
	"market-api/utils"
	"net/http"
)

func CreateMarket(s *Server, ctx echo.Context) (err error) {
	utils.LoggerInfo("handler - create market start")
	dto := new(MarketRequestDTO)

	if err = ctx.Bind(dto); err != nil {
		utils.LoggerError("handler - create market error bind", err)
		return echo.NewHTTPError(http.StatusBadRequest, apiError(ErrBadRequest, BadRequest))
	}
	if err = ctx.Validate(dto); err != nil {
		utils.LoggerError("handler - create market error validate", err)
		return echo.NewHTTPError(http.StatusBadRequest, apiError(ErrBadRequest, BadRequest))
	}

	res, err := s.Service.Create(dto.ToMarketDomain())
	if err != nil {
		utils.LoggerError("handler - create market error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, apiError(ErrInternalServerError, InternalServerError))
	}

	req := ctx.Request()
	location := fmt.Sprintf("%s%s/%s", req.Host, req.RequestURI, res)
	h := ctx.Response().Header()
	h.Add(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	h.Add("Location", location)
	_ = ctx.NoContent(http.StatusCreated)
	return
}

func GetMarketByID(s *Server, ctx echo.Context) (err error) {
	utils.LoggerInfo("handler - get market by ID start")
	id := ctx.Param("id")

	res, err := s.Service.GetByID(id)
	if err != nil {
		utils.LoggerError("handler - get market by ID error", err)
		return echo.NewHTTPError(http.StatusNotFound, apiError(ErrNotFound, MarketNotFound))
	}
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return ctx.JSON(http.StatusOK, ToMarketDTO(res))
}

func GetMarkets(s *Server, ctx echo.Context) (err error) {
	utils.LoggerInfo("Get markets start")
	var params MarketGetParam
	if err = (&echo.DefaultBinder{}).BindQueryParams(ctx, &params); err != nil {
		utils.LoggerError("handler - get markets bind error", err)
		return echo.NewHTTPError(http.StatusBadRequest, apiError(ErrBadRequest, err.Error()))
	}

	var res []domain.IMarket
	if params == (MarketGetParam{}) {
		res, err = s.Service.GetAll()
		if err != nil {
			utils.LoggerError("handler - get markets error", err)
			return echo.NewHTTPError(http.StatusInternalServerError, apiError(ErrInternalServerError, InternalServerError))
		}
	} else {
		res, err = s.Service.Get(
			"",
			"",
			"",
			"",
			params.Township,
			"",
			"",
			"",
			params.Region5,
			"",
			params.Name,
			"",
			"",
			"",
			params.District,
			"",
		)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, apiError(ErrInternalServerError, InternalServerError))
		}
	}

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return ctx.JSON(http.StatusOK, ToMarketsDTO(res))
}

func UpdateMarket(s *Server, ctx echo.Context) (err error) {
	utils.LoggerInfo("handler - update market start")
	id := ctx.Param("id")
	dto := new(MarketRequestDTO)

	if err = ctx.Bind(dto); err != nil {
		utils.LoggerError("handler - update market error bind", err)
		return echo.NewHTTPError(http.StatusBadRequest, apiError(ErrBadRequest, BadRequest))
	}
	if err = ctx.Validate(dto); err != nil {
		utils.LoggerError("handler - update market error validate", err)
		return echo.NewHTTPError(http.StatusBadRequest, apiError(ErrBadRequest, BadRequest))
	}

	market := dto.ToMarketDomain()

	res, err := s.Service.Update(id, market)
	if err != nil {
		utils.LoggerError("handler - update market error", err)
		return echo.NewHTTPError(http.StatusNotFound, apiError(ErrNotFound, MarketNotFound))
	}

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return ctx.JSON(http.StatusOK, ToMarketDTO(res))
}

func DeleteMarket(s *Server, ctx echo.Context) (err error) {
	utils.LoggerInfo("handler - delete market start")
	var param MarketDeleteParam
	if err = (&echo.DefaultBinder{}).BindQueryParams(ctx, &param); err != nil {
		utils.LoggerError("handler - delete market error bind", err)
		return echo.NewHTTPError(http.StatusBadRequest, apiError(ErrBadRequest, BadRequest))
	}

	if err = s.Service.DeleteByRegistry(param.Registry);
		err != nil {
		utils.LoggerError("handler - delete market error", err)
		return echo.NewHTTPError(http.StatusNotFound, apiError(ErrNotFound, MarketNotFound))
	}

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	_ = ctx.NoContent(http.StatusNoContent)
	return
}
