package main

import "time"


type Order struct{
	Size float64
	Limit *Limit
	Timestamp int64
	Bid bool
}

//Limit is a group of orders at a certain price level
type Limit struct{
	Price float64
	Orders []*Order
	TotalVolume float64
}

type Orderbook struct{
	Asks []*Limit
	Bids []*Limit
}

func NewLimit(price float64) *Limit{
	return &Limit{
		Price: price,
		Orders: []*Order{},
	}
}


func NewOrder(bid bool,size float64) *Order{
	return &Order{
		Size: size,
		Bid:bid,
		Timestamp: time.Now().UnixNano(),
	}
}
