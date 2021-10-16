# mq-go

A local mq module with retry and statistics.

functionality

Todo List

- [ ] single message queue
- [ ] multiple message queue manager
    - [ ] insert mq balancer
    - [ ] remove mq balancer
- [ ] queue adapter
    - [ ] insert queue callback
    - [ ] remove queue callback
- [ ] alerts
- [ ] statistics
    - [ ] insert count
    - [ ] insert time cost
    - [ ] remove count
    - [ ] remove time cost
    - [ ] data size

#### message queue

Local message queue is a queue to store data in the order of they insert and dequeue data as needed.


#### queue adapter

The adapter can alter the queue with a remote storage other than local memory. It can save data to whatever the user want by implementing the adapter methods.
For example, adapter can be implemented by local file or kafka to sustain data.

#### max size of queue

By default the size of the queue is unlimited (or is only limited by mem size). When set the max size, the alert function will be triggered every `ALERT_INTERVAL` and every time a message is inserted.

```
code sample
```

#### statistics

The queue can collect multiple statistics and generate a detailed report.
You can set the report interval.
