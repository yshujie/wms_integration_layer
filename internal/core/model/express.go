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

// ExpressPackage 快递包裹
type ExpressPackage struct {
	Length         float64
	Width          float64
	Height         float64
	Weight         float64
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
