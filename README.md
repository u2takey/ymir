# Ymir
Ymir is designed as a distributed, user load testing tool running on kubernetes.

## Design 
master - manage test task
agent  - get/run test work from master, agent's life cycle is same with a single task
dashboard - ymir frontend, on which user can manage and start task task, watch test metric and node load.

## RoadMap
- datamodel: for test task, designed as a test platform, we prefer a abstract test model in instead of a test script
- task management: use custom resource in kubernetes
- task flow control: use job or deployment, master slave architecture, master-slave communication use queue or rpc.
- node/pod monitor: get node monitor metrics during a task
