# QueuedJobs
Simple job processor distributed in queues.

## Install
To install the package just run:

    $ go get github.com/bfb/queued_jobs

## Configuring

Import the package:

    import "github.com/bfb/queued_jobs"

The first step is to configure the Redis connection:

    qj.Setup("redis://root:ce511ac0b7e0@54.12.129.10:6379/0", "test", 10)

The first parameter is the uri of Redis server, the second parameter is the namespace for all keys and the last is the connections pool size.

To configure a worker just use the Workers function:

    qj.Workers("critical", MyProcessFunc, 25)

The first parameter is the queue name, the second is the worker function and the last is the number of concurrent jobs.

Just start all workers:

    qj.Start()

## Example

    package main

    import "github.com/bfb/queued_jobs"

    func main() {
      // setup redis
      qj.Setup("redis://localhost:6379/0", "test", 20)

      // configure workers
      qj.Workers("critical", MyProcessFunc, 10)
      qj.Workers("normal", MySecondProcessFunc, 10)

      // run
      qj.Start()
    }

    func MyProcessFunc(args ...interface{}) {
      // hard work goes here
      // you can access the func params like args[0], args[1]
    }

    func MySecondProcessFunc(args ...interface{}) {
      // hard work goes here
    }

## Posting jobs

There are not a client to post messages yet, but you can post encoded messages in json into a list in Redis.
The list should be named like namespace:queue_name. The message structure is:

    {
      "jid" : "91659f84-c8c7-11e48731-1681e6b88ec1",
      "queue" : "critical",
      "args" : [
        "parameter1",
        "paramter2"
      ],
      posted_at: 1401584726
    }


## License

This package is licensed under the MIT License.
