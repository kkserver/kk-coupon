package coupon

import (
	"database/sql"
	"github.com/kkserver/kk-lib/kk"
	"github.com/kkserver/kk-lib/kk/app"
	"github.com/kkserver/kk-lib/kk/app/remote"
)

const CouponTypeCash = 0   //现金券
const CouponTypeRebate = 1 //折扣券

const CouponTimeTypeAbstract = 0 //绝对时间
const CouponTimeTypeRelative = 1 //相对领取时间

const CouponStatusNone = 0 //未上线
const CouponStatusIn = 200 //已上线

type Coupon struct {
	Id   int64 `json:"id"`
	Type int   `json:"type"`

	Count     int `json:"count"`     //已派发数量
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

	Status int `json:"status"`

	Ctime int64 `json:"ctime"`
}

const CouponReceiveStatusNone = 0      //未使用
const CouponReceiveStatusUse = 200     //已使用
const CouponReceiveStatusExpired = 300 //已过期

type CouponReceive struct {
	Id       int64 `json:"id"`
	CouponId int64 `json:"couponId"`
	Uid      int64 `json:"uid"`

	UseMaxCount int   `json:"useMaxCount"` //最大使用数量
	UseMinCount int   `json:"useMinCount"` //最小使用数量
	UseMaxValue int64 `json:"useMaxValue"` //最大使用金额
	UseMinValue int64 `json:"useMinValue"` //最小使用金额

	Value  int64 `json:"value"`  //抵用金额
	Rebate int64 `json:"rebate"` //折扣 (0~100)

	StartTime int64 `json:"startTime"` //开始有效时间
	EndTime   int64 `json:"endTime"`   //结束有效时间

	UseTime    int64  `json:"useTime"`    //使用时间
	UseType    string `json:"useType"`    //使用类型
	UseTradeNo string `json:"useTradeNo"` //使用流水号

	Status int `json:"status"`

	Ctime int64 `json:"ctime"`
}

type ICouponApp interface {
	app.IApp
	GetDB() (*sql.DB, error)
	GetPrefix() string
	GetCouponTable() *kk.DBTable
	GetCouponReceiveTable() *kk.DBTable
}

type CouponApp struct {
	app.App
	DB *app.DBConfig

	Remote *remote.Service

	Coupon *CouponService

	CouponTable        kk.DBTable
	CouponReceiveTable kk.DBTable
}

func (C *CouponApp) GetDB() (*sql.DB, error) {
	return C.DB.Get(C)
}

func (C *CouponApp) GetPrefix() string {
	return C.DB.Prefix
}

func (C *CouponApp) GetCouponTable() *kk.DBTable {
	return &C.CouponTable
}

func (C *CouponApp) GetCouponReceiveTable() *kk.DBTable {
	return &C.CouponReceiveTable
}
