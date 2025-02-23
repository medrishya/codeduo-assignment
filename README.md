Codeduo assignment task application

The application is split into two parts, api (in GO) for api server and ui for the frontend appllication ( in React ).

To setup and run the api server,

1. Run `cd api` and run `go mod tidy` or `go get` to install the modules.
2. Run `fresh` to start the api server at 8080 port.

To setup and run the frontend application

1. Run `cd ui` and run `npm install` to install the node packages.
2. Run `npm start` to start the application server at 3000 port.

Enjoy the task application.
Thank you!

Run `cd api` and run `go build -o task_manager` to build the project.

## Expected Execution Modes

| Mode              | Command                          | Description                                       |
| ----------------- | -------------------------------- | ------------------------------------------------- |
| CLI Add Task      | `./task_manager add "Task Name"` | Adds a new task                                   |
| CLI Process Tasks | `./task_manager process`         | Starts the consumer to complete a task from stack |
| CLI List Tasks    | `./task_manager list`            | Lists all tasks                                   |
| Run API           | `./task_manager api`             | Starts the REST API                               |
