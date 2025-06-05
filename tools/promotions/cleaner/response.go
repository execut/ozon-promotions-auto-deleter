package cleaner

import "time"

type DeactivateToAutoAddResponse struct {
    ProductIds []string `json:"productIds"`
}

type ActionListResponse struct {
    Actions []struct {
        Id                      string     `json:"id"`
        DateToNextAutoAdd       *time.Time `json:"dateToNextAutoAdd"`
        NextAutoAddProductCount string     `json:"nextAutoAddProductCount"`
    } `json:"actions"`
    Total string `json:"total"`
}

type ActiveToAutoAddProductListResponse struct {
    Products []struct {
        Id string `json:"id"`
    } `json:"products"`
    Total string `json:"total"`
}
