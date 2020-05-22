package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/pubsub"
)

func publish(w io.Writer, projectID, topicID, msg string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})

	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("Get: %v", err)
	}
	fmt.Fprintf(w, "Published a message; msg ID: %v\n", id)
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter message: ")
	text, _ := reader.ReadString('\n')

	err := publish(os.Stdout, os.Getenv("PROJECT"), "my-topic", text)

	if err != nil {
		fmt.Println(err)
	}
}
