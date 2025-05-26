package cleaner

import "time"

type DeactivateToAutoAddRequest struct {
    ProductIds    []string  `json:"product_ids"`
    DateToAutoAdd time.Time `json:"date_to_auto_add"`
}
