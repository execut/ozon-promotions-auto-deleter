package cleaner

import "time"

type DeactivateToAutoAddRequest struct {
    ProductIds  []string  `json:"product_ids"`
    AutoAddDate time.Time `json:"auto_add_date"`
}
