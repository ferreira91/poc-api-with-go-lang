package web

import (
	"errors"
	"github.com/go-playground/validator"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"market-api/internal/core/domain"
	mock_domain "market-api/test"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestCreateMarket(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serviceMock := mock_domain.NewMockIMarketService(ctrl)
	serviceMock.EXPECT().Create(gomock.Any()).Return("1", nil).AnyTimes()

	var marketJson = `{
		"longitude": "-11111111",
		"latitude": "-22222222",
		"censusSector": "333333333333333",
		"weightingArea": "4444444444444",
		"townshipCode": "555555555",
		"Township": "township",
		"subPrefectureCode": "66",
		"subPrefecture": "subPrefecture",
		"region5": "region",
		"region8": "region8",
		"name": "name",
		"registry": "666669",
		"street": "street",
		"number": "777777777777777",
		"district": "district",
		"reference": "reference"
	}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/market", strings.NewReader(marketJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	e.Validator = &CustomValidator{validator: validator.New()}
	s := &Server{serviceMock}

	if assert.NoError(t, CreateMarket(s, c)) {
		headers := rec.Header()
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, echo.MIMEApplicationJSONCharsetUTF8, headers.Get("Content-Type"))
		assert.Equal(t, "example.com/market/1", headers.Get("Location"))
	}

	var marketJsonBadRequest = `{
		"longitude": "-11111111",
		"latitude": "-22222222",
		"censusSector": "333333333333333",
		"weightingArea": "4444444444444",
		"townshipCode": "555555555",
		"Township": "township",
		"subPrefectureCode": "66",
		"subPrefecture": "subPrefecture",
		"region5": "region",
		"region8": "region8",
		"name": "name",
		"registry": "",
		"street": "street",
		"number": "777777777777777",
		"district": "district",
		"reference": "reference"
	}`

	e = echo.New()
	req = httptest.NewRequest(http.MethodPost, "/market", strings.NewReader(marketJsonBadRequest))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	e.Validator = &CustomValidator{validator: validator.New()}
	assert.EqualError(t, CreateMarket(s, c), "code=400, message={invalid_request Invalid request}")

	serviceMockError := mock_domain.NewMockIMarketService(ctrl)
	serviceMockError.EXPECT().Create(gomock.Any()).Return("", errors.New("Error ")).AnyTimes()

	e = echo.New()
	req = httptest.NewRequest(http.MethodPost, "/market", strings.NewReader(marketJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	e.Validator = &CustomValidator{validator: validator.New()}
	s = &Server{serviceMockError}

	assert.EqualError(t, CreateMarket(s, c), "code=500, message={server_error Oops! Something went wrong...}")
}

func TestGetMarketByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	market := domain.NewMarket()
	market.ID = "1"
	market.Longitude = "-11111111"
	market.Latitude = "-22222222"
	market.CensusSector = "333333333333333"
	market.WeightingArea = "4444444444444"
	market.TownshipCode = "555555555"
	market.Township = "township"
	market.SubprefectureCode = "66"
	market.Subprefecture = "subPrefecture"
	market.Region5 = "region5"
	market.Region8 = "region8"
	market.Name = "name"
	market.Registry = "666666"
	market.Street = "street"
	market.Number = "777777777777777"
	market.District = "district"
	market.Reference = "reference"

	serviceMock := mock_domain.NewMockIMarketService(ctrl)
	serviceMock.EXPECT().GetByID(gomock.Any()).Return(market, nil).AnyTimes()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/market", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	s := &Server{serviceMock}

	var marketJson = `{
    	"id": "1",
		"longitude": "-11111111",
		"latitude": "-22222222",
		"censusSector": "333333333333333",
		"weightingArea": "4444444444444",
		"townshipCode": "555555555",
		"township": "township",
		"subPrefectureCode": "66",
		"subPrefecture": "subPrefecture",
		"region5": "region5",
		"region8": "region8",
		"name": "name",
		"registry": "666666",
		"street": "street",
		"number": "777777777777777",
		"district": "district",
		"reference": "reference"
	}`

	if assert.NoError(t, GetMarketByID(s, c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, echo.MIMEApplicationJSONCharsetUTF8, rec.Header().Get("Content-Type"))
		assert.Equal(t,
			removeSpacesAndBreakLineAndHorizontalTab(rec.Body.String()),
			removeSpacesAndBreakLineAndHorizontalTab(marketJson),
		)
	}

	serviceMockError := mock_domain.NewMockIMarketService(ctrl)
	serviceMockError.EXPECT().GetByID(gomock.Any()).Return(nil, errors.New("Error ")).AnyTimes()

	e = echo.New()
	req = httptest.NewRequest(http.MethodGet, "/market", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	s = &Server{serviceMockError}

	assert.EqualError(t, GetMarketByID(s, c), "code=404, message={not_found Market not found}")
}

func TestGetMarkets(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	market := domain.NewMarket()
	market.ID = "1"
	market.Longitude = "-11111111"
	market.Latitude = "-22222222"
	market.CensusSector = "333333333333333"
	market.WeightingArea = "4444444444444"
	market.TownshipCode = "555555555"
	market.Township = "township"
	market.SubprefectureCode = "66"
	market.Subprefecture = "subPrefecture"
	market.Region5 = "region5"
	market.Region8 = "region8"
	market.Name = "name"
	market.Registry = "666666"
	market.Street = "street"
	market.Number = "777777777777777"
	market.District = "district"
	market.Reference = "reference"

	serviceMock := mock_domain.NewMockIMarketService(ctrl)
	var markets []domain.IMarket
	markets = append(markets, market)
	serviceMock.EXPECT().GetAll().Return(markets, nil).AnyTimes()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/market", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	s := &Server{serviceMock}

	var marketsJson = `
	[{
    	"id": "1",
		"longitude": "-11111111",
		"latitude": "-22222222",
		"censusSector": "333333333333333",
		"weightingArea": "4444444444444",
		"townshipCode": "555555555",
		"township": "township",
		"subPrefectureCode": "66",
		"subPrefecture": "subPrefecture",
		"region5": "region5",
		"region8": "region8",
		"name": "name",
		"registry": "666666",
		"street": "street",
		"number": "777777777777777",
		"district": "district",
		"reference": "reference"
	}]`

	if assert.NoError(t, GetMarkets(s, c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, echo.MIMEApplicationJSONCharsetUTF8, rec.Header().Get("Content-Type"))
		assert.Equal(t,
			removeSpacesAndBreakLineAndHorizontalTab(rec.Body.String()),
			removeSpacesAndBreakLineAndHorizontalTab(marketsJson),
		)
	}

	serviceMockError := mock_domain.NewMockIMarketService(ctrl)
	serviceMockError.EXPECT().GetAll().Return(nil, errors.New("Error ")).AnyTimes()

	e = echo.New()
	req = httptest.NewRequest(http.MethodGet, "/market", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	s = &Server{serviceMockError}

	assert.EqualError(t, GetMarkets(s, c), "code=500, message={server_error Oops! Something went wrong...}")

	a := gomock.Any()
	serviceMock.EXPECT().Get(a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a).Return(markets, nil).AnyTimes()

	q := make(url.Values)
	q.Add("township", "test")
	q.Add("region5", "test")
	q.Add("name", "test")
	q.Add("district", "test")
	req = httptest.NewRequest(http.MethodGet, "/market?"+q.Encode(), nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	s = &Server{serviceMock}

	if assert.NoError(t, GetMarkets(s, c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, echo.MIMEApplicationJSONCharsetUTF8, rec.Header().Get("Content-Type"))
		assert.Equal(t,
			removeSpacesAndBreakLineAndHorizontalTab(rec.Body.String()),
			removeSpacesAndBreakLineAndHorizontalTab(marketsJson),
		)
	}

	serviceMockError = mock_domain.NewMockIMarketService(ctrl)
	serviceMockError.EXPECT().Get(a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a).Return(nil, errors.New("Error ")).AnyTimes()

	req = httptest.NewRequest(http.MethodGet, "/market?"+q.Encode(), nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	s = &Server{serviceMockError}

	assert.EqualError(t, GetMarkets(s, c), "code=500, message={server_error Oops! Something went wrong...}")
}

func TestUpdateMarket(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	market := domain.NewMarket()
	market.ID = "1"
	market.Longitude = "-11111111"
	market.Latitude = "-22222222"
	market.CensusSector = "333333333333333"
	market.WeightingArea = "4444444444444"
	market.TownshipCode = "555555555"
	market.Township = "township"
	market.SubprefectureCode = "66"
	market.Subprefecture = "subPrefecture"
	market.Region5 = "region"
	market.Region8 = "region8"
	market.Name = "name"
	market.Registry = "666666"
	market.Street = "street"
	market.Number = "777777777777777"
	market.District = "district"
	market.Reference = "reference"

	serviceMock := mock_domain.NewMockIMarketService(ctrl)
	a := gomock.Any()
	serviceMock.EXPECT().Update(a, a).Return(market, nil).AnyTimes()

	var marketJson = `{
		"id": "1",
		"longitude": "-11111111",
		"latitude": "-22222222",
		"censusSector": "333333333333333",
		"weightingArea": "4444444444444",
		"townshipCode": "555555555",
		"township": "township",
		"subPrefectureCode": "66",
		"subPrefecture": "subPrefecture",
		"region5": "region",
		"region8": "region8",
		"name": "name",
		"registry": "666666",
		"street": "street",
		"number": "777777777777777",
		"district": "district",
		"reference": "reference"
	}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/market", strings.NewReader(marketJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	e.Validator = &CustomValidator{validator: validator.New()}
	s := &Server{serviceMock}

	if assert.NoError(t, UpdateMarket(s, c)) {
		headers := rec.Header()
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, echo.MIMEApplicationJSONCharsetUTF8, headers.Get("Content-Type"))
		assert.Equal(t,
			removeSpacesAndBreakLineAndHorizontalTab(rec.Body.String()),
			removeSpacesAndBreakLineAndHorizontalTab(marketJson),
		)
	}

	serviceMockError := mock_domain.NewMockIMarketService(ctrl)
	serviceMockError.EXPECT().Update(a, a).Return(nil, errors.New("Error ")).AnyTimes()

	e = echo.New()
	req = httptest.NewRequest(http.MethodPost, "/market", strings.NewReader(marketJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	e.Validator = &CustomValidator{validator: validator.New()}
	s = &Server{serviceMockError}

	assert.EqualError(t, UpdateMarket(s, c), "code=404, message={not_found Market not found}")
}

func TestDeleteMarket(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serviceMock := mock_domain.NewMockIMarketService(ctrl)
	serviceMock.EXPECT().DeleteByRegistry(gomock.Any()).Return(nil).AnyTimes()

	e := echo.New()
	q := make(url.Values)
	q.Set("registry", "test")
	req := httptest.NewRequest(http.MethodGet, "/market?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	s := &Server{serviceMock}

	if assert.NoError(t, DeleteMarket(s, c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
		assert.Equal(t, echo.MIMEApplicationJSONCharsetUTF8, rec.Header().Get("Content-Type"))
	}

	serviceMockError := mock_domain.NewMockIMarketService(ctrl)
	serviceMockError.EXPECT().DeleteByRegistry(gomock.Any()).Return(errors.New("Error ")).AnyTimes()

	e = echo.New()
	q = make(url.Values)
	q.Set("registry", "test")
	req = httptest.NewRequest(http.MethodGet, "/market?"+q.Encode(), nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	s = &Server{serviceMockError}

	assert.EqualError(t, DeleteMarket(s, c), "code=404, message={not_found Market not found}")
}

func removeSpacesAndBreakLineAndHorizontalTab(data string) string {
	data = strings.ReplaceAll(data, " ", "")
	data = strings.ReplaceAll(data, "\n", "")
	data = strings.ReplaceAll(data, "\t", "")
	return data
}
