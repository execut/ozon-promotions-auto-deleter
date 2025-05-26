package main

import (
    "log"
    "ozon-promotions-auto-deleter/config"
    promotionsCleaner "ozon-promotions-auto-deleter/tools/promotions/cleaner"
)

func main() {
    c := config.ReadConfig()
    client := promotionsCleaner.NewClient(c.CompanyID, c.Cookie)
    cleaner := promotionsCleaner.New(client)
    err := cleaner.Run()
    if err != nil {
        log.Fatal(err)
    }
}
