The printed number is a random number between -1,000,000 and 1,000,000, not 0 as expected.
This is because the two threads share a resource, i. This means that one thread might be blocking
the other thread from adding/subtracting from i while it does its work. Therefore we probably a number otherthan 0.





