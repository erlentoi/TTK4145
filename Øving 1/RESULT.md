The printed number is a random number between -1,000,000 and 1,000,000, not 0 as expected.
This is because the two threads share a resource, i. This means that one thread might be reading i while the other thread is incrementing/decrementing it. This means that the second thread will write the value it calculated, ignoring what the other thread did. Therefore we probably get a number other than 0.





