package graph

import (
	"context"
	"time"
)

// LastPing is the resolver for the lastPing field.
func (r *queryResolver) LastPing(ctx context.Context) (string, error) {
	if r.lastPing == "" {
		return "No pings received yet", nil
	}
	return r.lastPing, nil
}

// Ping is the resolver for the ping field.
func (r *mutationResolver) Ping(ctx context.Context) (string, error) {
	timestamp := time.Now().Format(time.RFC3339)
	r.lastPing = timestamp

	// Notify all subscribers
	r.subscriberMu.Lock()
	for _, ch := range r.subscribers {
		select {
		case ch <- timestamp:
		default:
		}
	}
	r.subscriberMu.Unlock()

	return timestamp, nil
}

// PingReceived is the resolver for the pingReceived field.
func (r *subscriptionResolver) PingReceived(ctx context.Context) (<-chan string, error) {
	id := time.Now().String()
	ch := make(chan string, 1)

	r.subscriberMu.Lock()
	r.subscribers[id] = ch
	r.subscriberMu.Unlock()

	go func() {
		<-ctx.Done()
		r.subscriberMu.Lock()
		delete(r.subscribers, id)
		r.subscriberMu.Unlock()
		close(ch)
	}()

	return ch, nil
}
