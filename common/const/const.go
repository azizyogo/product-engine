package constanta

const (

	// HTTP Methods
	HTTPGet    string = "GET"
	HTTPPost   string = "POST"
	HTTPPut    string = "PUT"
	HTTPPatch  string = "PATCH"
	HTTPDelete string = "DELETE"

	// MongoDB
	DBName            string = "product-engine"
	CollectionProduct string = "products"
	CollectionUser    string = "users"

	// Context
	CLAIMS_CONTEXT_KEY string = "claims"

	// Error
	ErrMsgUnauthorize   string = "unauthorize"
	ErrMsgInternalError string = "internal server errror"
	ErrMsgBadRequest    string = "bad request"
	ErrMsgNotFound      string = "not found"
)
