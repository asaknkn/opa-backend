# opa-backend
A simple application to call Online Payment Api(OPA) provided by PayPay.

# Provided Apis
- /v2/codes
- /v2/codes/payments/:merchantPaymentId
- /v1/qr/sessions

# Configurations
Set environment variables to access to PayPay.<br>
You can get theses values in [PayPay for Developer](https://developer.paypay.ne.jp/)
```
export OPA_APIKEY=ApiKey
export OPA_APISECRET=ApiSecret
export OPA_BASEURL=PayPay's URL
export OPA_ASSUMEMERCHANT=The Merchant Identifier
```

# How to Build a Container
Recommend [VSCode Remote Container](https://code.visualstudio.com/docs/remote/containers#_quick-start-open-an-existing-folder-in-a-container)


# How to Run the API Server
```
go run cmd/server.go 
```
Local server should be running on http://localhost:8080.