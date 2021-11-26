// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Ticker string

const (
	BTCUSDTicker    Ticker = "BTC_USD"
	SOURCE_CAPACITY        = 100
	CALC_INTERVAL          = 60
	PRICE_UPDATE_INTERVAL  = 5 

)

type TickerPrice struct {
	Ticker Ticker
	Time   time.Time
	Price  string // decimal value. example: "0", "10", "12.2", "13.2345122"
}

type PriceStreamSubscriber interface {
	SubscribePriceStream(Ticker, chan TickerPrice)
}

type Subscriber struct {
	id int
}

func (s Subscriber) SubscribePriceStream(Ticker, result chan TickerPrice) {
	i := 0
	for {
		// genereting random value at this point of time [95-105 USD] -> testing purpouses
		rand := strconv.Itoa(rand.Intn(105-95) + 95)
		result <- TickerPrice{Ticker: BTCUSDTicker, Time: time.Now(), Price: rand + ".00"}
		time.Sleep(4 * time.Second)
		i++
	}
}

func main() {
	// Init buffers
	var source [SOURCE_CAPACITY]Subscriber
	var results [SOURCE_CAPACITY]chan TickerPrice

	for i := 0; i < SOURCE_CAPACITY; i++ {
		source[i] = Subscriber{id: i}
		results[i] = make(chan TickerPrice, 1)
		go sub.SubscribePriceStream(results[i], results[i])
	}

	for {

		time.Sleep(CALC_INTERVAL * time.Second)
		fmt.Println(calculatePrice(results))
	}
}

// Calculating AVG price for our exchange 
func calculatePrice(res [SOURCE_CAPACITY]chan TickerPrice) string {
	sum := 0.0
	for i := 0; i < SOURCE_CAPACITY; i++ {
		select {
		case x, ok := <-res[i]:
			if ok {
				floatNum, err := strconv.ParseFloat(x.Price, 64)
				if err != nil {
					fmt.Println(err)
				}
				sum += floatNum

			} else {
				//fmt.Println("Channel closed!")
			}
		default:
			//fmt.Println("No value ready, moving on.")
		}

	}
	return fmt.Sprintf(strconv.Itoa(int(time.Now().Unix()))+" %f", sum/SOURCE_CAPACITY)

}
