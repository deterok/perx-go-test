package views

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

	"perx-go-test/resources"
)

type CodesView struct {
	codeSize int
	codesRes *resources.CodesResource
}

func NewCodesView(codeSize int, codesRes *resources.CodesResource) *CodesView {
	return &CodesView{
		codeSize: codeSize,
		codesRes: codesRes,
	}
}

func (v *CodesView) Post(c echo.Context) error {
	result, err := v.codesRes.Create(v.codeSize)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	resp := struct {
		Code      string    `json:"code"`
		CreatedAt time.Time `json:"created_at"`
	}{
		Code:      result.Code,
		CreatedAt: result.CreatedAt,
	}

	return c.JSON(http.StatusCreated, resp)
}
