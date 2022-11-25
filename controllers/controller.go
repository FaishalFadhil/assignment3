package controllers

import (
	"assignment3/database"
	"assignment3/models"
	"fmt"
	"gorm.io/gorm/clause"
	"math/rand"
	"time"
)

func CreateData(t time.Time) {
	var db = database.GetDB()
	rand.Seed(time.Now().UTC().UnixNano())
	var wind = rand.Intn(20)
	var water = rand.Intn(10)
	var windStatus, waterStatus string
	if wind > 15 {
		windStatus = "Bahaya"
	} else if wind > 6 {
		windStatus = "Siaga"
	} else {
		windStatus = "Aman"
	}

	if water > 8 {
		waterStatus = "Bahaya"
	} else if water > 5 {
		waterStatus = "Siaga"
	} else {
		waterStatus = "Aman"
	}

	input := models.Wheater{Id: 1, Water: int32(water), Wind: int32(wind), WaterStatus: waterStatus, WindStatus: windStatus}
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},                                                      // key colume
		DoUpdates: clause.AssignmentColumns([]string{"water", "wind", "water_status", "wind_status"}), // column needed to be updated
	}).Create(&input)

	fmt.Printf("\nLOG: Data Updated on %v\n", t)
	fmt.Printf("Water: %v, Status: %v \n", water, waterStatus)
	fmt.Printf("Wind: %v, Status: %v \n", wind, windStatus)
}

func UpdateData() {
	CreateData(time.Now())
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(150 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			CreateData(t)
		}
	}

}
