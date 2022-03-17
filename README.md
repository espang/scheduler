# scheduler

Scheduler is a system that is used to schedule tasks. The system has 3 parts:

* http server that receives tasks, validates them and publishs them to SQS
* main system that
  * receives tasks from SQS and writes them to persistent storage
  * reads tasks to be executed from persistent storage
  * emits tasks to SQS
* http client that receives tasks from SQS and makes POST requests



A task has:

* time when the task should be executed
* label
* url
* payload
* additional metadata that can be used to query it (?)
* partition

A task can be:

* scheduled via http endpoint
* schedulde via sqs message
* deleted via http endpoint
* checked via http endpoint


The scheduler will emit the event to an sqs queue on execution time.


