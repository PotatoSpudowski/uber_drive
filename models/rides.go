package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "../utils"
)

type Ride struct {
	
	gorm.Model
	CustomerName string `json:"rideid"`
	StartingLat string `json:"startinglat"`
	StartingLng string `json:"startinglng"`
	DestinationLat string `json:"destinationlat"`
	DestinationLng string `json:"destinationlng"`
	Phone string `json:"phone"`
	DriverID uint `json:"driverid"`
	// CustomerID uint `json:"customerid"`
}

func (ride *Ride) Validate() (map[string]interface{}, bool) {

	if ride.DriverID <= 0 {
		return u.Message(false, "Driver is not recognised"), false
	}

	if ride.StartingLat == "" || ride.StartingLng == "" {
		return u.Message(false, "Starting location should be on the payload"), false
	}

	if ride.DestinationLat == "" || ride.DestinationLng == "" {
		return u.Message(false, "Destination location should be on the payload"), false
	}

	return u.Message(true, "Success"),true
}

func (ride *Ride) Create() (map[string]interface{}) {

	if resp, ok := ride.Validate(); !ok {
		return resp
	}

	GetDB().Create(ride)

	resp := u.Message(true, "success")
	resp["ride"] = ride
	return resp
}

func GetRide(id uint) (*Ride) {

	ride := &Ride{}
	err := GetDB().Table("rides").Where("id = ?", id).First(&ride).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return ride
}

func GetRides(driverId uint) ([]*Ride) {

	rides := make([]*Ride, 0)
	err := GetDB().Table("rides").Where("driverid = ?", driverId).Find(&rides).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return rides
}