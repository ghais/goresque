# goresque
[![Build Status](https://travis-ci.org/ghais/goresque.png?branch=master)](https://travis-ci.org/ghais/goresque)
[![Coverage Status](https://coveralls.io/repos/ghais/goresque/badge.png)](https://coveralls.io/r/ghais/goresque)

A Simple [Resque](https://github.com/resque/resque) queue client for [Go](http://golang.org).

## Installation
```
go get github.com/ghais/goresque
```

## Usage

Let's assume that you have such Resque Job (taken from Resque examples):

```ruby
module Demo
  class Job
    def self.Perform(params)
      puts "Doing!"
    end
  end
end
```

So, we can enqueue this job from Go.

```go
package main

import (
  "github.com/ghais/goresque" // Import this package
)

func main() {
  client, err := goresque.Dial("127.0.0.1:6379") // get a client
  if err != nil {
      //handle error
  }
  // Enqueue the job into the go queue.
  resque.Enqueue("Demo::Job", "default")

  // Enqueue into the "default" queue with passing one parameter to the Demo::Job.
  resque.Enqueue("Demo::Job", "default", 1)

  // Enqueue into the "default" queue with passing multiple
  // parameters to the Demo::Job.perform so it will fail
  resque.Enqueue("Demo::Job", "default", 1, 2, "woot")
}
```
## Documentation 

- [Online Documentation (godoc.org)](http://godoc.org/github.com/ghais/goresque)

## Contributing

Fork and send a pull request.

## Note

This package was based on [go-resque](https://github.com/kavu/go-resque)
