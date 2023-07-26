package model

type StockInType string
type StockOutType string

const (
	PurchaseInStock    StockInType = "CQRK" // 采购入库
	ProfitCheckInStock StockInType = "PYRK" // 盘盈入库
	OrderBackInStock   StockInType = "THRK" // 退货入库
)

const (
	PurchaseReturnOutStock StockOutType = "CQCK" // 采购退货出库
	ProfitCheckOutStock    StockOutType = "PYCK" // 盘亏出库
	OrderOutStock          StockOutType = "XSCK" // 销售出库
)

// StockInRecord 入库记录
type StockInRecord struct {
	ID              int
	Type            StockInType
	FinishTime      string
	ExpressSender   ExpressSender   // 寄件人
	ExpressAddress  ExpressAddress  // 寄件地址
	ExpressDelivery ExpressDelivery // 快递单
	Goods           []StockGood
}

// StockOutRecord 出库记录
type StockOutRecord struct {
	ID              int
	Type            StockOutType
	FinishTime      string
	ExpressReceiver ExpressReceiver // 收件人
	ExpressAddress  ExpressAddress  // 收件地址
	ExpressDelivery ExpressDelivery // 快递单
	Goods           []StockGood
}
