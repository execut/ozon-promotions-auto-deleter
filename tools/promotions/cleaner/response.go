package cleaner

import "time"

type DeactivateToAutoAddResponse struct {
    ProductIds []string `json:"productIds"`
}

type ActionListResponse struct {
    Actions []struct {
        Id             string    `json:"id"`
        Title          string    `json:"title"`
        SubTitle       string    `json:"subTitle"`
        DateStart      time.Time `json:"dateStart"`
        DateEnd        time.Time `json:"dateEnd"`
        SkuCount       string    `json:"skuCount"`
        Discount       int       `json:"discount"`
        IsParticipated bool      `json:"isParticipated"`
        Highlights     []struct {
            Id          string        `json:"id"`
            Category    string        `json:"category"`
            MinDiscount int           `json:"minDiscount"`
            Parameters  []interface{} `json:"parameters"`
        } `json:"highlights"`
        ParticipatedProductCount string        `json:"participatedProductCount"`
        Description              string        `json:"description"`
        ActionType               string        `json:"actionType"`
        SkuBannedCount           string        `json:"skuBannedCount"`
        Segments                 []interface{} `json:"segments"`
        SegmentGroups            []interface{} `json:"segmentGroups"`
        HasAutoAddedProducts     bool          `json:"hasAutoAddedProducts"`
        DiscountType             string        `json:"discountType"`
        DiscountValue            int           `json:"discountValue"`
        OrderAmount              int           `json:"orderAmount"`
        NthQuantity              string        `json:"nthQuantity"`
        IsBatchParticipation     bool          `json:"isBatchParticipation"`
        DecorationUrl            string        `json:"decorationUrl"`
        ActionPriority           int           `json:"actionPriority"`
        AdditionalDescription    string        `json:"additionalDescription"`
        FreezeDate               interface{}   `json:"freezeDate"`
        Badges                   []struct {
            Title string `json:"title"`
            Color string `json:"color"`
        } `json:"badges"`
        CanEditProductRange         bool        `json:"canEditProductRange"`
        DefectRateTaskId            string      `json:"defectRateTaskId"`
        ShouldCheckDefectRateStatus bool        `json:"shouldCheckDefectRateStatus"`
        FinalDefectRate             string      `json:"finalDefectRate"`
        DefectRateFinalizedAt       interface{} `json:"defectRateFinalizedAt"`
        IsVoucherEnabled            bool        `json:"isVoucherEnabled"`
        StockChangedAt              interface{} `json:"stockChangedAt"`
        HasDMP                      bool        `json:"hasDMP"`
        HasWarehouses               bool        `json:"hasWarehouses"`
        AdditionalBenefits          []struct {
            OrderAmount   int    `json:"orderAmount"`
            DiscountValue int    `json:"discountValue"`
            DiscountType  string `json:"discountType"`
            NthQuantity   string `json:"nthQuantity"`
            Period        string `json:"period"`
            Commission    int    `json:"commission"`
        } `json:"additionalBenefits"`
        HasAddresses bool `json:"hasAddresses"`
        Labels       []struct {
            Title     string `json:"title"`
            Color     string `json:"color"`
            TextColor string `json:"textColor"`
        } `json:"labels"`
        IsStockRequired              bool        `json:"isStockRequired"`
        IsHotsale                    bool        `json:"isHotsale"`
        HotsaleId                    string      `json:"hotsaleId"`
        IsPriorityAction             bool        `json:"isPriorityAction"`
        FairDiscountLimitDate        interface{} `json:"fairDiscountLimitDate"`
        FairDiscountFailed           bool        `json:"fairDiscountFailed"`
        IsStockRecommended           bool        `json:"isStockRecommended"`
        IsUnlimitedAddition          bool        `json:"isUnlimitedAddition"`
        HasFbsScheduleWarning        bool        `json:"hasFbsScheduleWarning"`
        FbsScheduleDateFrom          interface{} `json:"fbsScheduleDateFrom"`
        FbsScheduleDateTo            interface{} `json:"fbsScheduleDateTo"`
        HasUploadedAutoAddedProducts bool        `json:"hasUploadedAutoAddedProducts"`
        PriceListMechanics           bool        `json:"priceListMechanics"`
        Promotion                    bool        `json:"promotion"`
        AutoAddedProductCount        string      `json:"autoAddedProductCount"`
        DecorationTemplate           struct {
            Id               string `json:"id"`
            MainColor        string `json:"mainColor"`
            AdditionalColor  string `json:"additionalColor"`
            WidgetStyle      string `json:"widgetStyle"`
            ElementColor     string `json:"elementColor"`
            ElementTextColor string `json:"elementTextColor"`
        } `json:"decorationTemplate"`
        CandidateProductCount     string     `json:"candidateProductCount"`
        Period                    string     `json:"period"`
        Commission                int        `json:"commission"`
        DateToNextAutoAdd         *time.Time `json:"dateToNextAutoAdd"`
        NextAutoAddProductCount   string     `json:"nextAutoAddProductCount"`
        ShowMpsp                  bool       `json:"showMpsp"`
        MaxActionPriceCheckFailed bool       `json:"maxActionPriceCheckFailed"`
    } `json:"actions"`
    Total string `json:"total"`
}

type ActiveToAutoAddProductListResponse struct {
    Products []struct {
        Id                        string   `json:"id"`
        Price                     int      `json:"price"`
        BasePrice                 int      `json:"basePrice"`
        Name                      string   `json:"name"`
        OfferID                   string   `json:"offerID"`
        ItemType                  string   `json:"itemType"`
        OzonStock                 string   `json:"ozonStock,omitempty"`
        SellerStock               string   `json:"sellerStock"`
        Skus                      []string `json:"skus"`
        MaxDiscountPrice          int      `json:"maxDiscountPrice"`
        MinSellerPrice            int      `json:"minSellerPrice"`
        Thumbnail                 string   `json:"thumbnail"`
        TotalStock                string   `json:"totalStock"`
        Currency                  string   `json:"currency"`
        OzonSku                   string   `json:"ozonSku"`
        DescriptionCategoryId     string   `json:"descriptionCategoryId"`
        ActionPriceToAutoAdd      int      `json:"actionPriceToAutoAdd"`
        MarketplaceSellerPrice    int      `json:"marketplaceSellerPrice"`
        AllowStockEdit            bool     `json:"allowStockEdit"`
        PriceReferenceForBoosting struct {
            BoostingX2Price  int `json:"boostingX2Price"`
            BoostingX3Price  int `json:"boostingX3Price"`
            BoostingX4Price  int `json:"boostingX4Price"`
            MaxBoostingPrice int `json:"maxBoostingPrice"`
        } `json:"priceReferenceForBoosting"`
        BoostingInSearch struct {
        } `json:"boostingInSearch"`
        ActionPrice int `json:"actionPrice,omitempty"`
    } `json:"products"`
    Total string `json:"total"`
}
