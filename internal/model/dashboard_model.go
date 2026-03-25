package model

type DashboardSummary struct {
	TotalSupplier    int64 `json:"total_supplier"`
	SupplierPending  int64 `json:"supplier_pending"`
	SupplierApproved int64 `json:"supplier_approved"`
	SupplierRejected int64 `json:"supplier_rejected"`
	TotalSPPG        int64 `json:"total_sppg"`
	TotalOrder       int64 `json:"total_order"`
	OrderPending     int64 `json:"order_pending"`
	OrderProcessing  int64 `json:"order_processing"`
	OrderShipped     int64 `json:"order_shipped"`
	OrderCompleted   int64 `json:"order_completed"`
	OrderCancelled   int64 `json:"order_cancelled"`
}

type DashboardResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func DashboardOK(message string, data interface{}) DashboardResponse {
	return DashboardResponse{Success: true, Message: message, Data: data}
}

func DashboardFail(message string) DashboardResponse {
	return DashboardResponse{Success: false, Message: message}
}