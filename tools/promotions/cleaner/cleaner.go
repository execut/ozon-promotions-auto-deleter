package cleaner

import "fmt"

func New(client *Client) *Cleaner {
    return &Cleaner{
        client: client,
    }
}

type Cleaner struct {
    client *Client
}

func (c *Cleaner) Run() error {
    actionListResponse, err := c.client.ActionList()
    if err != nil {
        return err
    }

    for _, action := range actionListResponse.Actions {
        if action.NextAutoAddProductCount == "0" {
            continue
        }

        if action.DateToNextAutoAdd == nil {
            fmt.Printf("Date to next is empty for action %v\n", action.Id)
            continue
        }
        productListResponse, err := c.client.ActiveToAutoAddProductList(action.Id, *action.DateToNextAutoAdd)
        if err != nil {
            return err
        }

        var productIDList []string
        productIDMap := make(map[string]struct{}, len(productListResponse.Products))
        for _, product := range productListResponse.Products {
            productIDList = append(productIDList, product.Id)
            productIDMap[product.Id] = struct{}{}
        }

        deactivateResponse, err := c.client.DeactivateToAutoAdd(action.Id, *action.DateToNextAutoAdd, productIDList)
        if err != nil {
            return err
        }

        for _, productID := range deactivateResponse.ProductIds {
            if _, ok := productIDMap[productID]; !ok {
                fmt.Printf("Product %v not founded after deactivation\n", productID)
            }
        }

        fmt.Printf("Успешно удалено %v автодобавленых товаров из акции №%v\n", len(actionListResponse.Actions), action.Id)
    }

    fmt.Printf("Успешно очишено %v акций\n", len(actionListResponse.Actions))

    return nil
}
