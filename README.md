# Ymir
Ymir is designed as a distributed script tool running on kubernetes.

## Design 
master - manage test task
agent  - get/run test work from master, agent's life cycle is same with a single task
dashboard - ymir frontend, on which user can manage and start task task, watch test metric and node load.

## What is Ymir
Ymir is designed as distributed script tool which can be used as user load testing tool like [locust](https://locust.io/). 
- Ymir is running on kubernetes, it can test service target in or out kubernetes.
- Ymir is first designed for user load test, but as a general script tool, it can be used to run any distributed script 
- Ymir use golang as test script, you can edit and rerun online.
- Ymir job defined with ymir script and node you selected to run, after a test job finish, it collect running status and runing metrics such as status code, avg/max/min latency.


## Ymir Script
Ymir Script is user's golang script which implented interface [task](https://github.com/arlert/ymir/blob/master/task/task.go), here are [examples](https://github.com/arlert/ymir/tree/master/taskset).

- A Ymir Script may contain one or more tasksets, taskset show be registered and ymir will run parallelly. Load will be distribute to tasksets by taskset's Weight
- A taskset may contain one or more tasks and ymir will run tasks in a taskset serially
- Ymir Script can import and use any package in golang std library

```
// TTaskSet ...
type TTaskSet interface {
	Name() string  // taskset name
	Weight() int   // default 1
	Tasks() []TTask // tasks
}

//TTask ....
type TTask interface {
	Name() string  // task name
	Run() int     // entrypoint
}

// Register your taskset with a NewFunc ...
func Register(newfunc NewFunc) {
}

// Set Run Time for your script ...
func SetRunTime(t time.Duration) {
}

// Set Routine Count for your script ...
func SetRoutineCount(c int) {
}

```

## Demo
![Alt text](/demo/demo-0.png)
![Alt text](/demo/demo-1.png)
![Alt text](/demo/demo-2.png)
![Alt text](/demo/demo-3.png)


## Deploy / Develop
deploy

```
cd deploy && kubectl create -f server_deploy.yaml
```

develop

```
# change ImageDestBase and Tag in makefile
make build  # build dashboard and binary
make publish_docker # build and push server docker image
make publish_docker_agent # build and push agent docker image
```

## RoadMap
- datamodel: done
- task management: done
- task flow control: 90%
- node/pod monitor: get node monitor metrics during a task
