package client

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/yuhu-tech/uche-data-sync/internal/collection/config"
	"github.com/yuhu-tech/uche-data-sync/internal/common/log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"gopkg.in/mgo.v2"
)

var (
	conf *config.Config

	quitChan  <-chan bool
	doneChan  chan<- bool
	fatalChan chan<- bool

	pool *poolStruct

	url              string
	db               string
	mongodbPoolLimit int
	session          *mgo.Session

	messageHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
		pool.works <- msg
	}
)

type poolStruct struct {
	wg       sync.WaitGroup
	works    chan MQTT.Message
	poolSize int
}

func StartMQTTClient(config *config.Config, qChan <-chan bool, dChan chan<- bool, fChan chan<- bool) {
	conf = config
	quitChan = qChan
	doneChan = dChan
	fatalChan = fChan

	url = conf.URL()
	db = conf.DB()
	mongodbPoolLimit = conf.MongoDBPoolLimit()

	dialInfo := &mgo.DialInfo{
		Addrs:     []string{url},
		PoolLimit: mongodbPoolLimit,
	}
	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Error(err)
	}
	session = s

	disconnectWiatingTime := conf.DisconnectWaitingTime()

	opts := MQTT.NewClientOptions()
	opts.AddBroker(conf.Server()).SetClientID(conf.ClientID()).SetUsername(conf.Username()).SetPassword(conf.Password()).SetAutoReconnect(true)

	opts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		log.Infof("topic: %s", msg.Topic())
	})

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Error(token.Error())
	}

	filters, err := conf.TopicandQoS()
	if err != nil {
		log.Error(err)
		fatalChan <- true
	}

	client.SubscribeMultiple(filters, messageHandler)

	pool = newWorkPool()
	pool.start()
	for {
		select {
		case <-quitChan:
			client.Disconnect(disconnectWiatingTime)
			pool.stop()
		}
	}
}

// newWorkPool is the function to return a new work pool
func newWorkPool() *poolStruct {
	workSize := conf.WorkSize()
	poolSize := conf.PoolSize()
	return &poolStruct{
		works:    make(chan MQTT.Message, workSize),
		poolSize: poolSize,
	}
}

func (w *poolStruct) start() {
	w.wg.Add(w.poolSize)
	for i := 0; i < w.poolSize; i++ {
		go func() {
			for message := range w.works {
				work(message)
			}
			w.wg.Done()
		}()
	}
}

func (w *poolStruct) stop() {
	close(w.works)
	w.wg.Wait()
	doneChan <- true
}

func work(message MQTT.Message) {
	switch message.Topic() {
	case "gps":
		workWithGPS(message)
	default:
		workWithOther(message)
	}
}

func workWithGPS(m MQTT.Message) {
	s := session.Copy()
	defer s.Close()

	topic := m.Topic()
	payload := m.Payload()

	v := &GPS{}
	err := json.Unmarshal(payload, v)
	if err != nil {
		log.Errorf("Unmarshal the payload failed, err: %s", err)
	}

	for i := range v.Data {
		v.Data[i].Boxid = v.Boxid
		if err != nil {
			log.Errorf("Marmal the data failed, err: %s", err)
		}

		cname := topic + "|" + "20" + v.Data[i].Tick[:6]

		c := s.DB(db).C(cname)
		err = c.Insert(v.Data[i])
		if err != nil {
			log.Errorf("Insert payload into mongoDB failed, topic: %s, err: %s", topic, err)
		}
	}
}

func workWithOther(m MQTT.Message) {
	s := session.Copy()
	defer s.Close()

	topic := m.Topic()
	payload := m.Payload()

	tv := &struct {
		Tick string `json:"tick"`
	}{}
	if err := json.Unmarshal(payload, tv); err != nil {
		log.Errorf("Unmarshal the payload failed, err: %s", err)
	}
	tick := tv.Tick

	var v interface{}
	switch topic {
	case "boxver":
		v = &Boxver{}
	case "protocol":
		v = &Protocol{}
	case "dbinfo":
		v = &Dbinfo{}
	case "boxstate":
		v = &Boxstate{}
	case "event/engstate":
		v = &Engstate{}
	case "event/drive":
		v = &Drive{}
	case "event/run":
		v = &Run{}
	case "event/traffic":
		v = &Traffic{}
	case "event/box":
		v = &Box{}
	case "event/tyre":
		v = &Tyre{}
	case "event/fuel":
		v = &Fuel{}
	case "event/dtc":
		v = &Dtc{}
	case "event/area":
		v = &Area{}
	case "statistic":
		v = &Statistic{}
	case "boxbody":
		v = &Boxbody{}
	case "cardata":
		v = &Cardata{}
	}

	err := json.Unmarshal(payload, v)
	if err != nil {
		log.Errorf("Unmarshal the payload failed, err: %s", err)
	}

	cname := topic + "|" + "20" + tick[:6]

	c := s.DB(db).C(cname)
	err = c.Insert(v)
	if err != nil {
		log.Errorf("Insert payload into mongoDB failed, topic: %s, err: %s", topic, err)
	}
}
