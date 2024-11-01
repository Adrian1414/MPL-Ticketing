package seeder

import (
    "log"

    "example.com/mpl/database"
    "example.com/mpl/models"
)

func SeedData() {
    if database.DB == nil {
        log.Fatal("Database connection is nil.")
    }

    var count int64
    database.DB.Model(&models.Team{}).Count(&count)
    if count == 0 { 
        teams := []models.Team{
            {Name: "FNOC"},
            {Name: "Geek"},
            {Name: "TLID"},
        }
        for _, team := range teams {
            database.DB.Create(&team)
        }
        log.Println("Seeder berhasil dijalankan.")
    }
}