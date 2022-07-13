package events

import (
	"log"
	"reflect"
	"time"
)

const (
	idleEventHubSleepTime time.Duration = 5 * time.Millisecond
	eventQueueCapacity                  = 300000
)

type TimeTickListener interface {
	HandleTimeTick(*TimeTick)
}

type UserInputListener interface {
	HandleUserInput(*UserInput)
}

type EventHub struct {
	// Event queues
	eventQueue chan any
	// Listeners list
	timeTickListener  []TimeTickListener
	userInputListener []UserInputListener
}

func NewEventHub() *EventHub {
	return &EventHub{
		eventQueue:        make(chan any, eventQueueCapacity),
		timeTickListener:  []TimeTickListener{},
		userInputListener: []UserInputListener{},
	}
}

// Register listener
func (h *EventHub) RegisterTimeTickListener(listener TimeTickListener) {
	h.timeTickListener = append(h.timeTickListener, listener)
}

func (h *EventHub) RegisterUserInputListener(listener UserInputListener) {
	h.userInputListener = append(h.userInputListener, listener)
}

func (h *EventHub) RunEventLoop() {
	for {
		event := <-h.eventQueue
		h.handle(event)
		time.Sleep(idleEventHubSleepTime)
	}
}

func (h *EventHub) handle(event any) {
	switch reflect.TypeOf(event) {
	case reflect.TypeOf(&TimeTick{}):
		e, ok := event.(*TimeTick)
		if !ok {
			log.Println("Event timetick error")
			break
		}
		for _, listener := range h.timeTickListener {
			listener.HandleTimeTick(e)
		}

		break
	case reflect.TypeOf(&UserInput{}):
		e, ok := event.(*UserInput)
		if !ok {
			log.Println("Event error")
			break
		}
		for _, listener := range h.userInputListener {
			listener.HandleUserInput(e)
		}
		break
	default:
		log.Println("Event type not supported")
		break

	}
}

func (h *EventHub) FireEvent(event any) {
	h.eventQueue <- event
}
