# Go Pubsub

Basic implementation of Google Cloud's Pub/Sub in Go. Subscriber consumes 5 messages before exiting.

## Set Up

Enable the Pub/Sub API for the project.

Create Server Account under Project with roles: "Pub/Sub Publisher" and "Pub/Sub Subscriber". Generate key and save into `keys` folder.

```sh
# Create topic
gcloud pubsub topics create my-topic

# Create subscriber and subscribe to topic
gcloud pubsub subscriptions create my-sub --topic my-topic
```

## Executing

Open two Terminal windows and export `GOOGLE_APPLICATION_CREDENTIALS` and `PROJECT`.

```sh
cd go-pubsub
export GOOGLE_APPLICATION_CREDENTIALS=$(pwd)/keys/keys.json
export PROJECT=`gcloud config get-value project`
```

First execute subscriber

```sh
go run subscriber/sub.go

```

Then execute publisher

```sh
go run publisher/pub.go

```