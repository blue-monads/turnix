
start_dev_frontend:
	cd frontend && npm run dev
start_dev_backend:
	FRONTEND_DEV_SERVER="http://localhost:5173" go run -v main.go
run_tests:
	cd tmp && rm -rf data.db && go run ../tests/*.go

play:
	FRONTEND_DEV_SERVER="http://localhost:5173" go run cli/play/play.go