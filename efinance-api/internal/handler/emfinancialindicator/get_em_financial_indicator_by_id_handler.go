package emfinancialindicator

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-efinance-api/internal/logic/emfinancialindicator"
	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
)

// swagger:route post /em_financial_indicator emfinancialindicator GetEmFinancialIndicatorById
//
// Get em financial indicator by ID | 通过ID获取EmFinancialIndicator信息
//
// Get em financial indicator by ID | 通过ID获取EmFinancialIndicator信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDInt32Req
//
// Responses:
//  200: EmFinancialIndicatorInfoResp

func GetEmFinancialIndicatorByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDInt32Req
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := emfinancialindicator.NewGetEmFinancialIndicatorByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetEmFinancialIndicatorById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
