package handler

type resultInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	CodeID  string `json:"codeId"`
}

type data struct {
	CodeID              string      `json:"codeId"`
	Url                 string      `json:"url"`
	Deeplink            string      `json:"deeplink"`
	ExpiryDate          int         `json:"expiryDate"`
	MerchantPaymentId   string      `json:"merchantPaymentId"`
	Amount              amount      `json:"amount"`
	OrderDescription    string      `json:"orderDescription"`
	OrderItems          []orderItem `json:"orderItems"`
	CodeType            string      `json:"codeType"`
	StoreInfo           string      `json:"storeInfo"`
	StoreID             string      `json:"storeId"`
	TerminalID          string      `json:"terminalId"`
	RequestedAt         int         `json:"requestedAt"`
	RedirectUrl         string      `json:"redirectUrl"`
	RedirectType        string      `json:"redirectType"`
	IsAuthorization     bool        `json:"isAuthorization"`
	AuthorizationExpiry int         `json:"authorizationExpiry"`
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
