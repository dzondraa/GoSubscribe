# GoSubscribe
Consuming streaming data from multiple sources 

## Idea:
- Creating buffers for streaming (creating channel and go routine for each exchange)
```
var source [SOURCE_CAPACITY]Subscriber
var results [SOURCE_CAPACITY]chan TickerPrice
```
- Exchanges can update the price any time (example 4s)

```
rand := strconv.Itoa(rand.Intn(105-95) + 95)
result <- TickerPrice{Ticker: BTCUSDTicker, Time: time.Now(), Price: rand + ".00"}
time.Sleep(4 * time.Second)
```
- Calculating the index price from most relevant data from all the soruces
```
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
```