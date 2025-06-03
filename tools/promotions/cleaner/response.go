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
        OfferId                   string   `json:"offerId"`
        Skus                      []string `json:"skus"`
        OzonSku                   string   `json:"ozonSku"`
        Name                      string   `json:"name"`
        ItemType                  string   `json:"itemType"`
        DescriptionCategoryId     string   `json:"descriptionCategoryId"`
        Price                     int      `json:"price"`
        ActionPrice               int      `json:"actionPrice"`
        BasePrice                 int      `json:"basePrice"`
        MaxDiscountPrice          int      `json:"maxDiscountPrice"`
        MinSellerPrice            int      `json:"minSellerPrice"`
        MarketplaceSellerPrice    int      `json:"marketplaceSellerPrice"`
        ActionPriceToAutoAdd      int      `json:"actionPriceToAutoAdd"`
        OzonStock                 string   `json:"ozonStock"`
        SellerStock               string   `json:"sellerStock"`
        TotalStock                string   `json:"totalStock"`
        MinActionQuantity         string   `json:"minActionQuantity"`
        Quantity                  string   `json:"quantity"`
        QuantityToAutoAdd         string   `json:"quantityToAutoAdd"`
        Currency                  string   `json:"currency"`
        Thumbnail                 string   `json:"thumbnail"`
        PriceReferenceForBoosting struct {
            BoostingX2Price  int `json:"boostingX2Price"`
            BoostingX3Price  int `json:"boostingX3Price"`
            BoostingX4Price  int `json:"boostingX4Price"`
            MaxBoostingPrice int `json:"maxBoostingPrice"`
        } `json:"priceReferenceForBoosting"`
        BoostingInSearch struct {
            BoostingPercent    int    `json:"boostingPercent"`
            MaxBoostingPercent int    `json:"maxBoostingPercent"`
            BoostingScaleColor string `json:"boostingScaleColor"`
        } `json:"boostingInSearch"`
        IsManuallyAdded bool `json:"isManuallyAdded"`
    } `json:"products"`
    Total string `json:"total"`
}
