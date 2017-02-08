package coupon

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type CouponQueryCounter struct {
	PageIndex int `json:"p"`
	PageSize  int `json:"size"`
	PageCount int `json:"count"`
	RowCount  int `json:"rowCount"`
}

type CouponQueryTaskResult struct {
	app.Result
	Counter *CouponQueryCounter `json:"counter,omitempty"`
	Coupons []Coupon            `json:"coupons,omitempty"`
}

type CouponQueryTask struct {
	app.Task
	Id        int64  `json:"id"`
	Status    string `json:"status"`
	Keyword   string `json:"q"`
	OrderBy   string `json:"orderBy"` // desc, asc
	PageIndex int    `json:"p"`
	PageSize  int    `json:"size"`
	Counter   bool   `json:"counter"`
	Result    CouponQueryTaskResult
}

func (task *CouponQueryTask) GetResult() interface{} {
	return &task.Result
}

func (task *CouponQueryTask) GetInhertType() string {
	return "coupon"
}

func (task *CouponQueryTask) GetClientName() string {
	return "Coupon.Query"
}
