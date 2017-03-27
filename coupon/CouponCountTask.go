package coupon

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type CouponCountTaskResult struct {
	app.Result
	Count int `json:"count"`
}

type CouponCountTask struct {
	app.Task
	Id      int64  `json:"id"`
	Status  string `json:"status"`
	Keyword string `json:"q"`
	Result  CouponCountTaskResult
}

func (task *CouponCountTask) GetResult() interface{} {
	return &task.Result
}

func (task *CouponCountTask) GetInhertType() string {
	return "coupon"
}

func (task *CouponCountTask) GetClientName() string {
	return "Coupon.Count"
}
