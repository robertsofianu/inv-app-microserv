run:
	cd invetory-app/api_server/CRUD && go run . &
	cd invetory-app/api_server/home_svc && go run . &
	cd invetory-app/api_server/createNirSvc && go run . &
	cd invetory-app/api_server/productsSvc && go run . &
	wait