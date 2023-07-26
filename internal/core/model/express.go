package model

// ExpressSender 快递寄件人
type ExpressSender struct {
	Name   string
	Mobile string
}

// ExpressReceiver 快递收件人
type ExpressReceiver struct {
	Name   string
	Mobile string
}

// ExpressDelivery 快递单
type ExpressDelivery struct {
	ExpressNo      string
	ExpressCompany ExpressCompany
}

// ExpressAddress 快递地址
type ExpressAddress struct {
	Province string
	City     string
	County   string
	Address  string
}

// ExpressCompany 快递公司
type ExpressCompany struct {
	Name string
	Code string
}
