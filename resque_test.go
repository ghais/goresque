package goresque

func ExampleDial() {
	c, err := Dial("localhost:6379")
	if err != nil {
		// handle error
	}
	defer c.Close()
}

func ExampleClient() {
	// Given the following Ruby code and the default queue
	// module Module
	//    class Job
	//        def self.Do(params)
	//            puts "Doing!"
	//        end
	//    end
	// end
	//
	c, err := Dial("localhost:6379")
	if err != nil {
		//handle error
	}

	// Enqueue with no params
	if err := c.Enqueue("Module::Job", "default"); err != nil {
		//handle error
	}

	// Enqueue with 1 parameter of type int
	if err := c.Enqueue("Module::Job", "default", 1); err != nil {
		//handle error
	}

	// Enqueue with 2 parameters an int and a string
	if err := c.Enqueue("Module::Job", "default", 1, "str"); err != nil {
		//handle error
	}
}
