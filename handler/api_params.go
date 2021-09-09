package handler

type resultInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	CodeID  string `json:"codeId"`
}

type amount struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

type orderItem struct {
	Name      string `json:"name"`
	Category  string `json:"category"`
	Quantity  int    `json:"quantity"`
	ProductID string `json:"productId"`
	UnitPrice amount `json:"unit_price"`
}

type orderRequestItem struct {
	Name      string `json:"name"`
	Category  string `json:"category"`
	Quantity  int    `json:"quantity"`
	ProductID string `json:"productId"`
	UnitPrice amount `json:"unitPrice"`
}

type refunds struct {
	Data []refund `json:"data"`
}

type refund struct {
	Status           string `json:"status"`
	AcceptedAt       int    `json:"acceptedAt"`
	MerchantRefundId string `json:"merchantRefundId"`
	PaymentId        string `json:"paymentId"`
	Amount           amount `json:"amount"`
	RequestedAt      int    `json:"requestedAt"`
	Reason           string `json:"reason"`
}

type captures struct {
	Data []capture `json:"data"`
}

type capture struct {
	AcceptedAt        int    `json:"acceptedAt"`
	MerchantCaptureId string `json:"merchantCaptureId"`
	Amount            amount `json:"amount"`
	OrderDescription  string `json:"orderDescription"`
	RequestedAt       int    `json:"requestedAt"`
	ExpiresAt         int    `json:"expiresAt"`
	Status            string `json:"status"`
}

type revert struct {
	AcceptedAt       int    `json:"acceptedAt"`
	MerchantRevertId string `json:"merchantRevertId"`
	RequestedAt      int    `json:"requestedAt"`
	Reason           string `json:"reason"`
}
