package emstockbasicinfo

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-efinance-api/internal/logic/emstockbasicinfo"
	"github.com/suyuan32/simple-admin-efinance-api/internal/svc"
	"github.com/suyuan32/simple-admin-efinance-api/internal/types"
)

// swagger:route post /em_stock_basic_info/delete emstockbasicinfo DeleteEmStockBasicInfo
//
// Delete em stock basic info information | 删除EmStockBasicInfo信息
//
// Delete em stock basic info information | 删除EmStockBasicInfo信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsInt32Req
//
// Responses:
//  200: BaseMsgResp

func DeleteEmStockBasicInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsInt32Req
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := emstockbasicinfo.NewDeleteEmStockBasicInfoLogic(r.Context(), svcCtx)
		resp, err := l.DeleteEmStockBasicInfo(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
