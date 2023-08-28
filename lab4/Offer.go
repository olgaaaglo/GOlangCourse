package main

import (
	"fmt"
	"encoding/json"
	"errors"
)

type Offer struct {
	Id int
	Position string
	AgeRange string
	Education string 
	EmpId int
}

func (offer Offer) ToString() string {
	return fmt.Sprintf("Offer: id=%d, position=%s, age=%d, ed=%s, empid=%d", offer.Id, offer.Position, offer.AgeRange, offer.Education, offer.EmpId);
}

type OffersStruct struct {
	Offers []*Offer
}

func (offers* OffersStruct) createElements(data []byte, count int) error {
	err := json.Unmarshal(data[:count], offers)
    if err != nil {
		return err
    }
	if len(offers.Offers) == 0 {
		return errors.New("Error creating offers")
	}
	fmt.Println("Create Offers: ")
	for _, o := range offers.Offers {
		fmt.Println(*o)
	}
	return nil
}

func (offers OffersStruct) ToString() string {
	ret := "Offers: "
	for _, p := range offers.Offers {
		ret += p.ToString() + "\n        "
	}
	return ret
}