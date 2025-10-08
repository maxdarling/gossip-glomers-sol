package main

import (
	"encoding/json"
	"log"
	"maps"
	"slices"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()

	// STATE
	var (
		mu        sync.Mutex
		values    = make(map[int]struct{})
		neighbors = make([]string, 0)
	)

	// BROADCAST
	n.Handle("broadcast", func(msg maelstrom.Message) error {
		type BroadcastMsg struct {
			Type string `json:"type"`
			Val  int    `json:"message"`
		}
		var body BroadcastMsg
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// store and broadcast new values
		mu.Lock()
		if _, ok := values[body.Val]; !ok {
			values[body.Val] = struct{}{}
			for _, peerID := range neighbors {
				n.Send(peerID, body)
			}
		}
		mu.Unlock()

		respBody := map[string]any{
			"type": "broadcast_ok",
		}
		return n.Reply(msg, respBody)
	})

	// BROADCAST_OK (ignore responses from neighbors)
	n.Handle("broadcast_ok", func(msg maelstrom.Message) error {
		return nil
	})

	// READ
	n.Handle("read", func(msg maelstrom.Message) error {
		mu.Lock()
		keys := slices.Collect(maps.Keys(values))
		mu.Unlock()

		respBody := map[string]any{
			"type":     "read_ok",
			"messages": keys,
		}
		return n.Reply(msg, respBody)
	})

	// TOPOLOGY
	n.Handle("topology", func(msg maelstrom.Message) error {
		type TopologyMsg struct {
			Type     string              `json:"type"`
			Topology map[string][]string `json:"topology"`
		}

		var body TopologyMsg
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		topology := body.Topology
		mu.Lock()
		neighbors = topology[n.ID()]
		mu.Unlock()

		respBody := map[string]any{
			"type": "topology_ok",
		}
		return n.Reply(msg, respBody)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
