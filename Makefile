clean:
	rm -f build/run-go-easy-frame
	rm -f build/cron
	rm -f build/consumer
	
build_web:
	go build -ldflags '-extldflags "-static"'  -o build/web-go-easy-frame go-easy-frame/entry

build_cron:
	go build -ldflags '-extldflags "-static"'  -o build/cron go-easy-frame/crontab

build_consumer:
	go build -ldflags '-extldflags "-static"'  -o build/consumer go-easy-frame/consumer

