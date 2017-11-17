package irc

import (
	"fmt"
	"strings"
	"sync"
)

//
// simple types
//

// ChannelNameMap holds a mapping of channel names to *Channel structs
// that is safe for concurrent readers and writers.
type ChannelNameMap struct {
	sync.RWMutex
	channels map[Name]*Channel
}

// NewChannelNameMap returns a new initialized *ChannelNameMap
func NewChannelNameMap() *ChannelNameMap {
	return &ChannelNameMap{
		channels: make(map[Name]*Channel),
	}
}

// Count returns the number of *Channel9s)
func (c *ChannelNameMap) Count() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.channels)
}

// Range ranges of the *Channels(s) calling f
func (c *ChannelNameMap) Range(f func(kay Name, value *Channel) bool) {
	c.Lock()
	defer c.Unlock()
	for k, v := range c.channels {
		if !f(k, v) {
			break
		}
	}
}

// Get returns a *Channel given a name if it exists or a zero-value *Channel
func (c *ChannelNameMap) Get(name Name) *Channel {
	c.RLock()
	defer c.RUnlock()
	return c.channels[name.ToLower()]
}

// Add adds a new *Channel if not already exists or an error otherwise
func (c *ChannelNameMap) Add(channel *Channel) error {
	c.Lock()
	defer c.Unlock()
	if c.channels[channel.name.ToLower()] != nil {
		return fmt.Errorf("%s: already set", channel.name)
	}
	c.channels[channel.name.ToLower()] = channel
	return nil
}

// Remove removes a *Channel if it exists or an error otherwise
func (c *ChannelNameMap) Remove(channel *Channel) error {
	c.Lock()
	defer c.Unlock()
	if channel != c.channels[channel.name.ToLower()] {
		return fmt.Errorf("%s: mismatch", channel.name)
	}
	delete(c.channels, channel.name.ToLower())
	return nil
}

type ChannelModeSet map[ChannelMode]bool

func (set ChannelModeSet) String() string {
	if len(set) == 0 {
		return ""
	}
	strs := make([]string, len(set))
	index := 0
	for mode := range set {
		strs[index] = mode.String()
		index += 1
	}
	return strings.Join(strs, "")
}

type ClientSet map[*Client]bool

func (clients ClientSet) Add(client *Client) {
	clients[client] = true
}

func (clients ClientSet) Remove(client *Client) {
	delete(clients, client)
}

func (clients ClientSet) Has(client *Client) bool {
	return clients[client]
}

type MemberSet map[*Client]ChannelModeSet

func (members MemberSet) Add(member *Client) {
	members[member] = make(ChannelModeSet)
}

func (members MemberSet) Remove(member *Client) {
	delete(members, member)
}

func (members MemberSet) Has(member *Client) bool {
	_, ok := members[member]
	return ok
}

func (members MemberSet) HasMode(member *Client, mode ChannelMode) bool {
	modes, ok := members[member]
	if !ok {
		return false
	}
	return modes[mode]
}

func (members MemberSet) AnyHasMode(mode ChannelMode) bool {
	for _, modes := range members {
		if modes[mode] {
			return true
		}
	}
	return false
}

type ChannelSet map[*Channel]bool

func (channels ChannelSet) Add(channel *Channel) {
	channels[channel] = true
}

func (channels ChannelSet) Remove(channel *Channel) {
	delete(channels, channel)
}

func (channels ChannelSet) First() *Channel {
	for channel := range channels {
		return channel
	}
	return nil
}

//
// interfaces
//

type Identifiable interface {
	Id() Name
	Nick() Name
}
