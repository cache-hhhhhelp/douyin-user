package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"microservice/user/api/internal/logic"
	"microservice/user/api/internal/svc"
	"microservice/user/api/internal/types"
)

func loginUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterOrLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginUserLogic(r.Context(), svcCtx)
		resp, err := l.LoginUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
