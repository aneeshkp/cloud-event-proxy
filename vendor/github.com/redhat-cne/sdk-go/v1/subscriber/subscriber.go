package subscriber

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/pkg/errors"
	"github.com/redhat-cne/sdk-go/pkg/pubsub"
	"github.com/redhat-cne/sdk-go/pkg/store"

	log "github.com/sirupsen/logrus"

	subscriberStore "github.com/redhat-cne/sdk-go/pkg/store/subscriber"
	"github.com/redhat-cne/sdk-go/pkg/subscriber"
)

// API ... api methods  for publisher subscriber
type API struct {
	subscriberStore  *subscriberStore.Store //  each client will have one store
	storeFilePath    string                 // subscribers
	transportEnabled bool                   //  http  is enabled
}

var instance *API
var once sync.Once
var mu sync.Mutex

// NewSubscriber create new subscribers connections
func NewSubscriber(clientID string) subscriber.Subscriber {
	return subscriber.Subscriber{
		ClientID:    clientID,
		SubStore:    &store.PubSubStore{},
		EndPointURI: nil,
		Status:      1,
	}
}

// New creates empty publisher or subscriber
func New() subscriber.Subscriber {
	return subscriber.Subscriber{}
}

// GetAPIInstance get event instance
func GetAPIInstance(storeFilePath string) *API {
	once.Do(func() {
		instance = &API{
			transportEnabled: true,
			subscriberStore: &subscriberStore.Store{
				RWMutex: sync.RWMutex{},
				Store:   map[string]*subscriber.Subscriber{},
			},
			storeFilePath: storeFilePath,
		}
		hasDir(storeFilePath)
		instance.ReloadStore()
	})
	return instance
}

// ReloadStore reload store if there is any change or refresh is required
func (p *API) ReloadStore() {
	// load for file
	if files, err := loadFileNamesFromDir(p.storeFilePath); err == nil {
		for _, f := range files {
			if b, err := loadFromFile(fmt.Sprintf("%s/%s", p.storeFilePath, f)); err == nil {
				if len(b) > 0 {
					var sub subscriber.Subscriber
					if err := json.Unmarshal(b, &sub); err == nil {
						p.subscriberStore.Set(sub.ClientID, sub)
					}
				}
			}
		}
	}
}

// HasTransportEnabled flag to indicate if amqp is enabled
func (p *API) HasTransportEnabled() bool {
	return p.transportEnabled
}

// DisableTransport disables usage of amqp
func (p *API) DisableTransport() {
	p.transportEnabled = false
}

// EnableTransport enable usage of amqp
func (p *API) EnableTransport() {
	p.transportEnabled = true
}

// ClientCount .. client cound
func (p *API) ClientCount() int {
	return len(p.subscriberStore.Store)
}

// GetSubFromSubscriptionsStore get data from publisher store
func (p *API) GetSubFromSubscriptionsStore(clientID, address string) (pubsub.PubSub, error) {
	if subscriber, ok := p.HasClient(clientID); ok {
		for _, sub := range subscriber.SubStore.Store {
			if sub.GetResource() == address {
				return pubsub.PubSub{
					ID:          sub.ID,
					EndPointURI: sub.EndPointURI,
					URILocation: sub.URILocation,
					Resource:    sub.Resource,
				}, nil
			}
		}
	}

	return pubsub.PubSub{}, fmt.Errorf("publisher not found for address %s", address)
}

// HasSubscription check if the subscriptionOne is already exists in the store/cache
func (p *API) HasSubscription(clientID, address string) (pubsub.PubSub, bool) {
	if sub, err := p.GetSubFromSubscriptionsStore(clientID, address); err == nil {
		return sub, true
	}
	return pubsub.PubSub{}, false
}

// HasClient check if  client is already exists in the store/cache
func (p *API) HasClient(clientID string) (*subscriber.Subscriber, bool) {
	if subscriber, ok := p.subscriberStore.Store[clientID]; ok {
		return subscriber, true
	}
	return nil, false
}

// CreateSubscription create a subscriptionOne and store it in a file and cache
func (p *API) CreateSubscription(clientID string, sub subscriber.Subscriber) (subscriptionClient *subscriber.Subscriber, err error) {

	var ok bool
	if subscriptionClient, ok = p.HasClient(clientID); !ok {
		subscriptionClient = subscriber.New(clientID)
	}
	subscriptionClient.SetEndPointURI(sub.GetEndPointURI())
	subscriptionClient.SetStatus(subscriber.Active)
	pubStore := subscriptionClient.GetSubStore()
	var hasResource bool
	for key, value := range sub.SubStore.Store {
		hasResource = false
		for _, p := range pubStore.Store {
			if p.Resource == value.Resource {
				hasResource = true
				continue
			}
		}
		if !hasResource {
			subscriptionClient.SubStore.Set(key, *value)
		}

	}
	p.subscriberStore.Set(clientID, *subscriptionClient)
	// persist the subscriptionOne -
	//TODO:might want to use PVC to live beyond pod crash
	err = writeToFile(*subscriptionClient, fmt.Sprintf("%s/%s", p.storeFilePath, fmt.Sprintf("%s.json", clientID)))
	if err != nil {
		log.Errorf("error writing to a store %v\n", err)
		return nil, err
	}
	log.Infof("subscriptionOne persisted into a file %s", fmt.Sprintf("%s/%s  - content %s", p.storeFilePath, fmt.Sprintf("%s.json", clientID), subscriptionClient.String()))
	// store the publisher

	return subscriptionClient, nil
}

// GetSubscriptionClient  get a clientID by id
func (p *API) GetSubscriptionClient(clientID string) (subscriber.Subscriber, error) {
	if sub, ok := p.subscriberStore.Store[clientID]; ok {
		return *sub, nil
	}
	return subscriber.Subscriber{}, fmt.Errorf("subscriber data was not found for id %s", clientID)
}

// GetSubscriptionsFromFile  get subscriptions data from the file store
func (p *API) GetSubscriptionsFromFile(clientID string) ([]byte, error) {
	b, err := loadFromFile(fmt.Sprintf("%s/%s", p.storeFilePath, fmt.Sprintf("%s.json", clientID)))
	return b, err
}

// GetSubscriptions  get all subscriptionOne inforamtions
func (p *API) GetSubscriptions(clientID string) (sub map[string]*pubsub.PubSub) {
	if sub, ok := p.subscriberStore.Store[clientID]; ok {
		return sub.SubStore.Store
	}
	return
}

// GetSubscription  get  subscriptionOne inforamtions
func (p *API) GetSubscription(clientID, subID string) (sub pubsub.PubSub) {

	if subscriber, ok := p.subscriberStore.Store[clientID]; ok {
		if subscription, ok2 := subscriber.SubStore.Store[subID]; ok2 {
			return *subscription
		}
	}
	return
}

// GetSubscriberURLByResource  get  subscriptionOne inforamtions
func (p *API) GetSubscriberURLByResource(resource string) (urls []string) {
	for _, subscriber := range p.subscriberStore.Store {
		for _, sub := range subscriber.SubStore.Store {
			if sub.GetResource() == resource {
				urls = append(urls, subscriber.GetEndPointURI())
			}
		}
	}
	return urls
}

// DeleteSubscription delete a subscriptionOne by id
func (p *API) DeleteSubscription(clientID, subscriptionID string) error {
	log.Info("deleting subscriptionOne")
	if subStore, ok := p.subscriberStore.Store[clientID]; ok { // client found
		if sub, ok2 := subStore.SubStore.Store[subscriptionID]; ok2 {
			err := deleteFromFile(*sub, fmt.Sprintf("%s/%s", p.storeFilePath, fmt.Sprintf("%s.json", clientID)))
			p.subscriberStore.Store[clientID].SubStore.Delete(subscriptionID)
			return err
		}
	}
	return nil
}

// DeleteAllSubscriptions  delete all subscriptionOne information
func (p *API) DeleteAllSubscriptions(clientID string) error {
	log.Info("deleting all subscriptionOne")
	if subStore, ok := p.subscriberStore.Store[clientID]; ok { // client found
		if err := deleteAllFromFile(fmt.Sprintf("%s/%s", p.storeFilePath, fmt.Sprintf("%s.json", clientID))); err != nil {
			return err
		}
		subStore.SubStore = &store.PubSubStore{
			RWMutex: sync.RWMutex{},
			Store:   map[string]*pubsub.PubSub{},
		}
	}
	return nil
}

// DeleteClient  delete all subscriptionOne information
func (p *API) DeleteClient(clientID string) error {
	log.Info("deleting  client")
	if subStore, ok := p.subscriberStore.Store[clientID]; ok { // client found
		if err := deleteAllFromFile(fmt.Sprintf("%s/%s", p.storeFilePath, fmt.Sprintf("%s.json", clientID))); err != nil {
			return err
		}
		subStore.SubStore = &store.PubSubStore{
			RWMutex: sync.RWMutex{},
			Store:   map[string]*pubsub.PubSub{},
		}
		delete(p.subscriberStore.Store, clientID)
	}
	return nil
}

// UpdateStatus .. update status
func (p *API) UpdateStatus(clientID string, status subscriber.Status) error {
	if subStore, ok := p.subscriberStore.Store[clientID]; ok {
		subStore.SetStatus(status)
		// do not write to file , if restarts it will consider all client are active
	} else {
		return errors.New("failed to update subscriber status")
	}
	return nil
}

// deleteAllFromFile deletes  publisher and subscriptionOne information from the file system
func deleteAllFromFile(filePath string) error {
	if err := os.Remove(filePath); err != nil {
		return err
	}
	return nil
}

// DeleteFromFile is used to delete subscriptionOne from the file system
func deleteFromFile(sub pubsub.PubSub, filePath string) error {
	var persistedSubClient subscriber.Subscriber
	//open file
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	//read file and unmarshall json file to slice of users
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	if len(b) > 0 {
		err = json.Unmarshal(b, &persistedSubClient)
		if err != nil {
			return err
		}
	}
	if _, ok := persistedSubClient.SubStore.Store[sub.ID]; ok {
		delete(persistedSubClient.SubStore.Store, sub.ID)
	}

	newBytes, err := json.MarshalIndent(&persistedSubClient, "", " ")
	if err != nil {
		log.Errorf("error deleting sub %v", err)
		return err
	}
	if err := ioutil.WriteFile(filePath, newBytes, 0666); err != nil {
		return err
	}
	return nil
}

// loadFromFile is used to read subscriptionOne/publisher from the file system
func loadFromFile(filePath string) (b []byte, err error) {
	//open file
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	//read file and unmarshall json file to slice of users
	b, err = ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func loadFileNamesFromDir(filePath string) (subFiles []string, err error) {
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		return subFiles, err
	}
	for _, file := range files {
		if !file.IsDir() {
			subFiles = append(subFiles, file.Name())
		}
	}
	return
}

// writeToFile writes subscriptionOne data to a file
func writeToFile(subscriberClient subscriber.Subscriber, filePath string) error {
	//open file
	mu.Lock()
	defer mu.Unlock()
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	//read file and unmarshall json file to slice of users
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	var persistedSubClient subscriber.Subscriber
	if len(b) > 0 {
		err = json.Unmarshal(b, &persistedSubClient)
		if err != nil {
			return err
		}
	} else {
		persistedSubClient = *subscriber.New(subscriberClient.ClientID)
	} // no  file found
	persistedSubClient.SetEndPointURI(subscriberClient.GetEndPointURI())
	persistedSubClient.SetStatus(subscriber.Active)
	for subID, sub := range subscriberClient.SubStore.Store {
		persistedSubClient.SubStore.Store[subID] = sub
	}

	newBytes, err := json.MarshalIndent(&persistedSubClient, "", " ")
	if err != nil {
		return err
	}
	log.Infof("persisting following contents %s to a file %s\n", string(newBytes), filePath)
	if err := ioutil.WriteFile(filePath, newBytes, 0666); err != nil {
		return err
	}
	return nil
}

func hasDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, 0700)
	}
}
