package coupon

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type CouponCreateTaskResult struct {
	app.Result
	Coupon *Coupon `json:"coupon,omitempty"`
}

type CouponCreateTask struct {
	app.Task

	Type int `json:"type"`

	MaxCount  int `json:"maxCount"`  //最大派发数量
	UMaxCount int `json:"umaxCount"` //用户领取最大数量

	UseMaxCount int   `json:"useMaxCount"` //最大使用数量
	UseMinCount int   `json:"useMinCount"` //最小使用数量
	UseMaxValue int64 `json:"useMaxValue"` //最大使用金额
	UseMinValue int64 `json:"useMinValue"` //最小使用金额

	MinValue int64 `json:"minValue"` //最小金额
	MaxValue int64 `json:"maxValue"` //最大金额

	MinRebate int64 `json:"minRebate"` //最小折扣 (0~100)
	MaxRebate int64 `json:"maxRebate"` //最大折扣 (0~100)

	StartTimeType int   `json:"startTimeType"` //开始有效时间类型
	StartTime     int64 `json:"startTime"`     //开始有效时间
	EndTimeType   int   `json:"endTimeType"`   //开始有效时间类型
	EndTime       int64 `json:"endTime"`       //结束有效时间

	Title   string `json:"title"`
	Summary string `json:"summary"`
	Remark  string `json:"remark"`

	Result CouponCreateTaskResult
}

func (task *CouponCreateTask) GetResult() interface{} {
	return &task.Result
}

func (task *CouponCreateTask) GetInhertType() string {
	return "coupon"
}

func (task *CouponCreateTask) GetClientName() string {
	return "Coupon.Create"
}
