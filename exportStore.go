package main

import (
	"context"
	"time"

	"github.com/patrickalin/nest-client-go/pkg/ring"

	nest "github.com/patrickalin/nest-api-go"
	"github.com/sirupsen/logrus"
)

type store struct {
	in     chan nest.Nest
	stores map[string]*ring.Ring
}

type measure struct {
	Timestamp time.Time
	value     float64
}

func (m measure) TimeStamp() time.Time {
	return m.Timestamp
}

/**
* Measure  represents a measure that has a GetValue
 */

func (m measure) Value() float64 {
	return m.value
}

//InitConsole listen on the chanel
func createStore(messages chan nest.Nest) (store, error) {
	stores := make(map[string]*ring.Ring)
	stores["temperatureCelsius"] = &ring.Ring{}
	stores["windGustkmh"] = &ring.Ring{}
	return store{in: messages, stores: stores}, nil

}

func (c *store) listen(context context.Context) {

	go func() {

		log.WithFields(logrus.Fields{
			"fct": "exportStore.listen",
		}).Info("init")

		for {
			msg := <-c.in
			log.WithFields(logrus.Fields{
				"fct": "exportStore.listen",
			}).Debug("Receive message")
			c.stores["temperatureCelsius"].Enqueue(measure{time.Now(), msg.GetTemperatureCelsius()})
			c.stores["windGustkmh"].Enqueue(measure{time.Now(), msg.GetWindGustkmh()})
		}
	}()

}

func (c *store) GetValues(name string) []ring.TimeMeasure {
	return c.stores[name].Values()
}

func (c *store) String(name string) string {
	s, _ := c.stores[name].DumpLine()
	return s
}
