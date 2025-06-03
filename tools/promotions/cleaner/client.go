package cleaner

import (
    "encoding/json"
    "time"

    "github.com/execut/ozon-reports-downloader/common"
)

type Client struct {
    companyID int64
    cookie    string
}

func NewClient(companyID int64, cookie string) *Client {
    return &Client{
        companyID: companyID,
        cookie:    cookie,
    }
}

func (c *Client) ActionList() (*ActionListResponse, error) {
    resp, err := common.DoGetRequest(struct{}{}, "https://seller.ozon.ru/api/site/sa-facade/v1/seller-actions?withHotsales=true&offset=0&limit=100", c.cookie, c.companyID)
    if err != nil {
        return nil, err
    }

    respObject := &ActionListResponse{}
    err = json.Unmarshal(resp, respObject)
    if err != nil {
        return nil, err
    }

    return respObject, nil
}

func (c *Client) ActiveToAutoAddProductList(actionID string, dateToAutoAdd time.Time) (*ActiveToAutoAddProductListResponse, error) {
    resp, err := common.DoGetRequest(struct{}{}, "https://seller.ozon.ru/api/site/sa-auto-add/v1/"+actionID+"/products-with-offset?offset=0&limit=100&autoAddDate="+dateToAutoAdd.Format(time.RFC3339), c.cookie, c.companyID)
    if err != nil {
        return nil, err
    }

    respObject := &ActiveToAutoAddProductListResponse{}
    err = json.Unmarshal(resp, respObject)
    if err != nil {
        return nil, err
    }

    return respObject, nil
}

func (c *Client) DeactivateToAutoAdd(actionID string, dateToAutoAdd time.Time, productIDList []string) (*DeactivateToAutoAddResponse, error) {
    request := DeactivateToAutoAddRequest{
        ProductIds:  productIDList,
        AutoAddDate: dateToAutoAdd,
    }
    resp, err := common.DoRequest(request, "https://seller.ozon.ru/api/site/sa-auto-add/v1/"+actionID+"/delete-products", c.cookie, c.companyID)
    if err != nil {
        return nil, err
    }

    respObject := &DeactivateToAutoAddResponse{}
    err = json.Unmarshal(resp, respObject)
    if err != nil {
        return nil, err
    }

    return respObject, nil
}
