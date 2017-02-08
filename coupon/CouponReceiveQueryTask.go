package coupon

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type CouponReceiveQueryCounter struct {
	PageIndex int `json:"p"`
	PageSize  int `json:"size"`
	PageCount int `json:"count"`
	RowCount  int `json:"rowCount"`
}

type CouponReceiveQueryTaskResult struct {
	app.Result
	Counter  *CouponReceiveQueryCounter `json:"counter,omitempty"`
	Receives []CouponReceive            `json:"receives,omitempty"`
}

type CouponReceiveQueryTask struct {
	app.Task
	Uid       int64  `json:"uid"`
	CouponId  int64  `json:"couponId"`
	Status    string `json:"status"`
	OrderBy   string `json:"orderBy"` // desc, asc, value, endTime
	PageIndex int    `json:"p"`
	PageSize  int    `json:"size"`
	Counter   bool   `json:"counter"`
	Result    CouponReceiveQueryTaskResult
}

func (task *CouponReceiveQueryTask) GetResult() interface{} {
	return &task.Result
}

func (task *CouponReceiveQueryTask) GetInhertType() string {
	return "coupon"
}

func (task *CouponReceiveQueryTask) GetClientName() string {
	return "Coupon.ReceiveQuery"
}
