
seqshard1: go run cmd/shard/shard.go -isRoot -ipAddr localhost:4000


recshard1: go run cmd/shard/shard.go -rec -parentIpAddr localhost:4000 -ipAddr localhost:6000 -storagePath tmp/shard1
recshard2: go run cmd/shard/shard.go -rec -parentIpAddr localhost:4000 -ipAddr localhost:6001 -storagePath tmp/shard2

client1: go run cmd/client/client.go -parentIpAddr localhost:6000
client2: go run cmd/client/client.go -parentIpAddr localhost:6000
client3: go run cmd/client/client.go -parentIpAddr localhost:6000
client4: go run cmd/client/client.go -parentIpAddr localhost:6001
client5: go run cmd/client/client.go -parentIpAddr localhost:6001
#client3: go run cmd/readclient/readclient.go -parentIpAddr localhost:6000