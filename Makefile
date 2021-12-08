all: server client

server:
	go build -o echo-server echo-server.go

client:
	go build -o echo-client echo-client.go

clean:
	rm ./echo-server
	rm ./echo-client
