package sse

import (
	"sync"
	"time"
)

// Channel represents a server sent events channel.
type Channel struct {
	mu          sync.RWMutex
	lastEventID string
	name        string
	clients     map[*Client]bool
	logger      LogPrinter
}

func newChannel(name string, logger LogPrinter) *Channel {
	return &Channel{
		mu:          sync.RWMutex{},
		lastEventID: "",
		name:        name,
		clients:     make(map[*Client]bool),
		logger:      logger,
	}
}

// SendMessage broadcast a message to all clients in a channel.
func (c *Channel) SendMessage(message *Message) {
	c.lastEventID = message.id

	c.mu.RLock()

	for cl, open := range c.clients {
		if open {
			select {
			case cl.send <- message:
			case <-time.After(2000 * time.Millisecond):
				// todo: logger.Warn
				c.logger.Warn("SSE Timed out sending message to client")
			}
		}
	}

	c.mu.RUnlock()
}

// Close closes the channel and disconnect all clients.
func (c *Channel) Close() {
	// Kick all clients of this channel.
	for client := range c.clients {
		c.removeClient(client)
	}
}

// ClientCount returns the number of clients connected to this channel.
func (c *Channel) ClientCount() int {
	c.mu.RLock()
	count := len(c.clients)
	c.mu.RUnlock()

	return count
}

// LastEventID returns the ID of the last message sent.
func (c *Channel) LastEventID() string {
	return c.lastEventID
}

func (c *Channel) addClient(client *Client) {
	c.mu.Lock()
	c.clients[client] = true
	c.mu.Unlock()
}

func (c *Channel) removeClient(client *Client) {
	c.mu.Lock()
	c.clients[client] = false
	delete(c.clients, client)
	c.mu.Unlock()
	close(client.send)
}
