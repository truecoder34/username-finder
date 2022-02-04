# username-finder
First app written on Golang: just to touch Goroutines and Mocks
Currently only one endpoint is impemented. WIP. List of endpoints described below
## API methods:
	POST: /username
		input : JSON body ["url1", "url2", ..."urln"]
		outout : JSON  ["valid-url1", "valid-url2", ..."valid-urln"]

	POST: /qr
		input : JSON body ["url1", "url2", ..."urln"]
		return : qr code in text format for valid urls

### Rabbit MQ
1 - in helper added sender of rabbit mq  messages. If endpoint used, it will be send to consumer
2 - consumer is in /rabbit-mq-receiver folder



## How to run :
Docker RabbitMQ:
```
	#t0> docker run --detach --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```

Consumer :
```
	#t1> cd rabbit-mq-receiver
	#t1> go run .\receive-mq.go
```

Server : 
```
	#t2> cd server
	#t2> go run .\main.go
```




# Practice with rabbit-mq
folder : ```go-tabbit-mq/```
copy pasted from : https://github.com/Pungyeon/go-rabbitmq-example 
## added lib with settings
...
## sender
```
	#t1> go run consumer.go log.WARN log.ERROR
	#t2> go run consumer.go log.*
	#t3> go run consumer.go *
```

## consumer
```
	#t3> go run sender.go RAP
	#t3> go run sender.go TEST
	#t3> go run sender.go log.WARN
```
