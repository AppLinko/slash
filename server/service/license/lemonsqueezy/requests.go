package lemonsqueezy

const (
	// The base API URL for the Lemon Squeezy API.
	baseAPIURL = "https://api.lemonsqueezy.com"
	// The store ID for the yourselfhosted store.
	// Link: https://yourselfhosted.lemonsqueezy.com
	storeID = 15634
	// The product ID for the subscription pro product.
	// Link: https://yourselfhosted.lemonsqueezy.com/checkout/buy/d03a2696-8a8b-49c9-9e19-d425e3884fd7
	subscriptionProProductID = 98995
)

type LicenseKey struct {
	ID        int32   `json:"id"`
	Status    string  `json:"status"`
	Key       string  `json:"key"`
	CreatedAt string  `json:"created_at"`
	ExpiresAt *string `json:"updated_at"`
}

type LicenseKeyMeta struct {
	StoreID       int32  `json:"store_id"`
	OrderID       int32  `json:"order_id"`
	OrderItemID   int32  `json:"order_item_id"`
	ProductID     int32  `json:"product_id"`
	ProductName   string `json:"product_name"`
	VariantID     int32  `json:"variant_id"`
	VariantName   string `json:"variant_name"`
	CustomerID    int32  `json:"customer_id"`
	CustomerName  string `json:"customer_name"`
	CustomerEmail string `json:"customer_email"`
}

type ValidateLicenseKeyResponse struct {
	Valid      bool            `json:"valid"`
	Error      *string         `json:"error"`
	LicenseKey *LicenseKey     `json:"license_key"`
	Meta       *LicenseKeyMeta `json:"meta"`
}

type ActiveLicenseKeyResponse struct {
	Activated  bool            `json:"activated"`
	Error      *string         `json:"error"`
	LicenseKey *LicenseKey     `json:"license_key"`
	Meta       *LicenseKeyMeta `json:"meta"`
}

func ValidateLicenseKey(licenseKey string, instanceName string) (*ValidateLicenseKeyResponse, error) {
    return &ValidateLicenseKeyResponse{
        Valid: true,
        LicenseKey: &LicenseKey{
            Status:    "active",
            ExpiresAt: nil,
        },
    }, nil
}

func ActiveLicenseKey(licenseKey string, instanceName string) (*ActiveLicenseKeyResponse, error) {
    return &ActiveLicenseKeyResponse{
        Activated: true,
    }, nil
}