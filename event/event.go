// Package event implements an event bus.
// for a great introduction to the event bus pattern in go, see:
// https://levelup.gitconnected.com/lets-write-a-simple-event-bus-in-go-79b9480d8997
package event

import (
	"context"
	"fmt"
	"sync"
	"time"

	golog "github.com/ipfs/go-log"
)

var (
	log = golog.Logger("event")

	// ErrBusClosed indicates the event bus is no longer coordinating events
	// because it's parent context has closed
	ErrBusClosed = fmt.Errorf("event bus is closed")
	// NowFunc is the function that generates timestamps (tests may override)
	NowFunc = time.Now
)

// Topic is the set of all kinds of events emitted by the bus. Use the "Topic"
// type to distinguish between different events. Event emitters should
// declare Topics as constants and document the expected payload type.
type Topic string

// Event represents an event that subscribers will receive from the bus
type Event struct {
	Topic     Topic
	Timestamp int64
	SessionID string
	Payload   interface{}
}

// Handler is a function that will be called by the event bus whenever a
// matching event is published. Handler calls are blocking, called in order
// of subscription. Any error returned by a handler is passed back to the
// event publisher.
// The handler context originates from the publisher, and in practice will often
// be scoped to a "request context" like an HTTP request or CLI command
// invocation.
// Generally, even handlers should aim to return quickly, and only delegate to
// goroutines when the publishing event is firing on a long-running process
type Handler func(ctx context.Context, e Event) error

// Publisher is an interface that can only publish an event
type Publisher interface {
	Publish(ctx context.Context, t Topic, payload interface{}) error
	PublishID(ctx context.Context, t Topic, sessionID string, payload interface{}) error
}

// Bus is a central coordination point for event publication and subscription
// zero or more subscribers register eventTopics to be notified of, a publisher
// writes a topic event to the bus, which broadcasts to all subscribers of that
// topic
type Bus interface {
	// Publish an event to the bus
	Publish(ctx context.Context, t Topic, data interface{}) error
	// PublishID publishes an event with an arbitrary session id
	PublishID(ctx context.Context, t Topic, sessionID string, data interface{}) error
	// Subscribe to one or more eventTopics with a handler function that will be called
	// whenever the event topic is published
	SubscribeTopics(handler Handler, eventTopics ...Topic)
	// SubscribeID subscribes to only events that have a matching session id
	SubscribeID(handler Handler, sessionID string)
	// SubscribeAll subscribes to all events
	SubscribeAll(handler Handler)
	// NumSubscriptions returns the number of subscribers to the bus's events
	NumSubscribers() int
}

// NilBus replaces a nil value. it implements the bus interface, but does
// nothing
var NilBus = nilBus{}

type nilBus struct{}

// assert at compile time that nilBus implements the Bus interface
var _ Bus = (*nilBus)(nil)

// Publish does nothing with the event
func (nilBus) Publish(_ context.Context, _ Topic, _ interface{}) error {
	return nil
}

// PublishID does nothing with the event
func (nilBus) PublishID(_ context.Context, _ Topic, _ string, _ interface{}) error {
	return nil
}

// SubscribeTopics does nothing
func (nilBus) SubscribeTopics(handler Handler, eventTopics ...Topic) {}

func (nilBus) SubscribeID(handler Handler, id string) {}

func (nilBus) SubscribeAll(handler Handler) {}

func (nilBus) NumSubscribers() int {
	return 0
}

type bus struct {
	lk      sync.RWMutex
	closed  bool
	subs    map[Topic][]Handler
	allSubs []Handler
	idSubs  map[string][]Handler
}

// assert at compile time that bus implements the Bus interface
var _ Bus = (*bus)(nil)

// NewBus creates a new event bus. Event busses should be instantiated as a
// singleton. If the passed in context is cancelled, the bus will stop emitting
// events and close all subscribed channels
//
// TODO (b5) - finish context-closing cleanup
func NewBus(ctx context.Context) Bus {
	b := &bus{
		subs:    map[Topic][]Handler{},
		idSubs:  map[string][]Handler{},
		allSubs: []Handler{},
	}

	go func(b *bus) {
		<-ctx.Done()
		log.Debugf("close bus")
		b.lk.Lock()
		b.closed = true
		b.lk.Unlock()
	}(b)

	return b
}

// Publish sends an event to the bus
func (b *bus) Publish(ctx context.Context, topic Topic, payload interface{}) error {
	return b.publish(ctx, topic, "", payload)
}

// Publish sends an event with a given sessionID to the bus
func (b *bus) PublishID(ctx context.Context, topic Topic, sessionID string, payload interface{}) error {
	return b.publish(ctx, topic, sessionID, payload)
}

func (b *bus) publish(ctx context.Context, topic Topic, sessionID string, payload interface{}) error {
	b.lk.RLock()
	defer b.lk.RUnlock()
	log.Debugw("publish", "topic", topic, "payload", payload)

	if b.closed {
		return ErrBusClosed
	}

	e := Event{
		Topic:     topic,
		Timestamp: NowFunc().UnixNano(),
		SessionID: sessionID,
		Payload:   payload,
	}

	for _, handler := range b.subs[topic] {
		if err := handler(ctx, e); err != nil {
			return err
		}
	}

	if sessionID != "" {
		for _, handler := range b.idSubs[sessionID] {
			if err := handler(ctx, e); err != nil {
				return err
			}
		}
	}

	for _, handler := range b.allSubs {
		if err := handler(ctx, e); err != nil {
			return err
		}
	}

	return nil
}

// Subscribe requests events from the given topic, returning a channel of those events
func (b *bus) SubscribeTopics(handler Handler, eventTopics ...Topic) {
	b.lk.Lock()
	defer b.lk.Unlock()
	log.Debugf("Subscribe to topics: %v", eventTopics)

	for _, topic := range eventTopics {
		b.subs[topic] = append(b.subs[topic], handler)
	}
}

// SubscribeID requests events that match the given sessionID
func (b *bus) SubscribeID(handler Handler, sessionID string) {
	b.lk.Lock()
	defer b.lk.Unlock()
	log.Debugf("Subscribe to ID: %v", sessionID)
	b.idSubs[sessionID] = append(b.idSubs[sessionID], handler)
}

// SubscribeAll requets all events from the bus
func (b *bus) SubscribeAll(handler Handler) {
	b.lk.Lock()
	defer b.lk.Unlock()
	log.Debugf("Subscribe All")
	b.allSubs = append(b.allSubs, handler)
}

// NumSubscribers returns the number of subscribers to the bus's events
func (b *bus) NumSubscribers() int {
	b.lk.Lock()
	defer b.lk.Unlock()
	total := 0
	for _, handlers := range b.subs {
		total += len(handlers)
	}
	for _, handlers := range b.idSubs {
		total += len(handlers)
	}
	total += len(b.allSubs)
	return total
}
