package core

import "github.com/yshujie/wms_integration_layer/internal/core/model"

type WMSClient interface {
	// PushGood 推送商品信息
	PushGood(good model.StockGood) error
	// QueryGood 查询商品信息
	QueryGood(skuCode string) (model.StockGood, error)
	// QueryInventory 查询库存信息
	QueryInventory(skuCode string) (model.StockGood, error)

	// PushInboundApplication 推送入库单
	PushInboundApplication(record model.StockInRecord) error
	// CancelInboundApplication 取消入库单
	CancelInboundApplication(record model.StockInRecord) error

	// PushOutboundApplication 推送出库单
	PushOutboundApplication(record model.StockOutRecord) error
	// CancelOutboundApplication 取消出库单
	CancelOutboundApplication(record model.StockOutRecord) error

	// PushSalesOrder 推送销售订单
	PushSalesOrder(record model.StockOutRecord) error
	// CancelSalesOrder 取消销售订单
	CancelSalesOrder(record model.StockOutRecord) error

	// PushReturnOrder 推送退货单
	PushReturnOrder(record model.StockInRecord) error
	// CancelReturnOrder 取消退货单
	CancelReturnOrder(record model.StockInRecord) error
}
