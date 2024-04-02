d_build:
	docker build -t todo-online:1.0.0 .

d_run:
	docker run -p 8000:8080 -e MONGODB_CONNECTION_STRING="mongodb+srv://duong:duong-todo-online@todo-online.fpmhk1z.mongodb.net/?retryWrites=true&w=majority" -e MONGODB_DATABASE_NAME=todo-online todo-online:1.0.0