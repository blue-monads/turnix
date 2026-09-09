default:
	just --list
frontend_start_dev:
	cd frontend && bun run dev
backend_start_dev:
	# export POTATO_DEV_SPACES="cimple-table:5174"
	export FRONTEND_DEV_SERVER="http://localhost:7779" && go run -v cmd/dev/*.go
litestream_start_dev:
	litestream replicate -config contrib/litestream/litestream_dev.yml
labs_start_dev:
	go run -v backend/labs/labs.go
run_tests:
	cd tmp && rm -rf data.db && go run ../tests/*.go

play:
	FRONTEND_DEV_SERVER="http://localhost:7779" go run cli/play/play.go

build:
	cd frontend && bun run build
	go build -v -o ./tmp/potatoverse ./main.go

build_frontend:
	cd frontend && bun run build

build_backend:
	go build -v -o tmp/potatoverse ./main.go

copy_dev_binary:
	cd frontend && mv output/build old_build && bun run build && rm -rf output/build && mv old_build output/build
	go build -v -o tmp/potatoverse ./main.go
	cp tmp/potatoverse ~/go/bin

run_cli:
	cd tmp && FRONTEND_DEV_SERVER="http://localhost:7779"  go run ../main.go 

run_full_server:
  cd tmp && go run ../main.go server init-and-start
  
run_http_server:
	cd contrib && python3 -m http.server 8080

start_test_server:
	FRONTEND_DEV_SERVER="http://localhost:7779" TURNIX_DEV_MODE="true" TURNIX_DEV_ABC_SERVER="http://localhost:8080" go run -v cmd/dev/main.go

run_cloudy:
    export CLOUDY_MASTER_SECRET="potatoverse" && export TURSO_AUTH_TOKEN="eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhIjoicnciLCJpYXQiOjE3ODA0NzYyMDksImlkIjoiMDE5ZThjYTYtYTIwMS03OTY1LWEzNzItNTBmZmNjN2YxNWM1IiwicmlkIjoiODkyNGVlNjEtNTgzYS00MGUwLTk3OTAtZmVlZDZiN2MzYThhIn0.q0reuJ0v5pMqpaqFqmPtRLKpwgunwhAyzUQ8KrBn21CkI9oReWVOX0l74x0H_juNECwJB8aH-gLfdNKlymdSAg" && export TURSO_REMOTE_URL="https://cloudy-bornjre.aws-eu-west-1.turso.io" && cd tmp && go run ../cmd/cloudy/main.go


citest:
	./ci.sh