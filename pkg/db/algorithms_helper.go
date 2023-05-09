package db

import (
	"math/rand"
	"time"
)

func (d *DBController) EmptyOrders() ([]Order, error) {
	var orders []Order
	var emptyOrders []Order
	var err error

	if orders, err = d.Orders(); err != nil {
		return nil, err
	}

	for _, order := range orders {
		if !order.Taken {
			emptyOrders = append(emptyOrders, order)
		}
	}
	return emptyOrders, nil
}

func CheckTime(timeArrive, timeDue time.Time) bool {
	timeDiff := timeDue.Sub(timeArrive)
	if timeDiff.Minutes() > 10 { // 10 minutes to spare
		return true
	}
	return false
}

func (d *DBController) NeighboorsSimple(orders []Order, drivers []Driver) (map[*Driver]([]Order), []Order, error) {
	neighborhood := make(map[*Driver]([]Order))
	freeOrder := make([]Order, 0)

	numOrder := 0
	lenOrder := len(orders)
	for _, driver := range drivers {
		if numOrder > lenOrder {
			neighborhood[&driver] = nil
		} else {
			ord := []Order{orders[numOrder]}
			neighborhood[&driver] = ord
			numOrder += 1
		}
	}
	if numOrder < lenOrder {
		for i := numOrder; i < lenOrder; i++ {
			freeOrder = append(freeOrder, orders[i])
		}
	}

	return neighborhood, freeOrder, nil
}

func (d *DBController) AddOrdersByRemainingTime(driver *Driver, orders []Order, freeOrders []Order) ([]Order, []Order) {
	var timeLastOrder time.Time
	ordersAccepted := make([]Order, len(orders))
	ordersFree := make([]Order, len(freeOrders))

	timeLastOrder = time.Now()
	for _, order := range orders {
		order.TimeFinish = timeLastOrder.Add(
			d.AStar(*driver, driver.Position, order.PositionTake) +
				2*order.TimePack +
				d.AStar(*driver, order.PositionTake, order.PositionSend))

		if order.TimeEnd.Sub(order.TimeEnd).Hours() < 0 {
			ordersAccepted = append(ordersAccepted, freeOrders[0])
			freeOrders = freeOrders[1:]
			ordersFree = append(ordersFree, order)
		} else {
			ordersAccepted = append(ordersAccepted, order)
		}
	}
	ordersFree = append(ordersFree, freeOrders...)
	return ordersAccepted, ordersFree
}

func (d *DBController) Randomize(neighborhood map[*Driver]([]Order)) {
	rand.Seed(time.Now().UnixNano())
	keys := make([]Driver, 0, len(neighborhood))
	for k := range neighborhood {
		keys = append(keys, *k)
	}
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	val := neighborhood[&keys[0]]
	neighborhood[&keys[0]] = neighborhood[&keys[1]]
	neighborhood[&keys[1]] = val
}

func (d *DBController) ChangeNeighborhood(neighborhood map[*Driver]([]Order), freeOrders []Order) (map[*Driver]([]Order), error) {
	var orders []Order
	for k, v := range neighborhood {
		d.Randomize(neighborhood)
		orders, freeOrders = d.AddOrdersByRemainingTime(k, v, freeOrders)
		neighborhood[k] = orders
	}
	return neighborhood, nil
}
