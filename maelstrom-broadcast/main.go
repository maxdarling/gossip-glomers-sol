package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)



func main() {
	n := maelstrom.NewNode()

	messages := make([]int, 0)

	// BROADCAST
	n.Handle("broadcast", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}




		val := int(body["message"].(float64))
		messages = append(messages, val)

		respBody := map[string]any{
			"type" : "broadcast_ok",
		}
		return n.Reply(msg, respBody)
	})

	// READ
	n.Handle("read", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		respBody := map[string]any{
			"type" : "read_ok",
			"messages" : messages,
		}
		return n.Reply(msg, respBody)
	})

	// TOPOLOGY
	n.Handle("topology", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// todo: do topology stuff
		// ...

		respBody := map[string]any{
			"type" : "topology_ok",
		}
		return n.Reply(msg, respBody)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
