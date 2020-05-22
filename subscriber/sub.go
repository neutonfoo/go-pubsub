package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"cloud.google.com/go/pubsub"
)

func pullMsgs(w io.Writer, projectID, subID string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}

	fmt.Fprintf(w, "Subscriber %q will consume 5 messages.\n", subID)

	// Consume 5 messages.
	var mu sync.Mutex
	received := 0
	sub := client.Subscription(subID)
	cctx, cancel := context.WithCancel(ctx)
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Fprintf(w, "Got message: %q\n", string(msg.Data))
		msg.Ack()
		mu.Lock()
		defer mu.Unlock()
		received++
		if received == 5 {
			cancel()
		}
	})
	if err != nil {
		return fmt.Errorf("Receive: %v", err)
	}
	return nil
}

func main() {
	pullMsgs(os.Stdout, os.Getenv("PROJECT"), "my-sub")
}
