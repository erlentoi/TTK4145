The printed number is a random number between -1,000,000 and 1,000,000, not 0 as expected.
This is because the two threads share a resource, i, without locking or synchronization. This means that one thread might be reading i while the other thread is incrementing/decrementing it. This means that the second thread will write the value it calculated to the variable i, ignoring what the other thread wrote. Therefore we probably get a number other than 0.
https://en.wikipedia.org/wiki/Race_condition#Example






