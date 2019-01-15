# Python 3.3.3 and 2.7.6
# python fo.py

from threading import Thread


i = 0

def incrementingFunction():
    global i

    for a in range(1,1000000):
        i+=1


def decrementingFunction():
    global i
    for a in range(1,1000000):
        i-=1



def main():
    # TODO: Something is missing here (needed to print i)
    global i #ikke nødvendig? (fungerer også uten)

    #lager thread objecter
    incrementing = Thread(target = incrementingFunction, args = (),)
    decrementing = Thread(target = decrementingFunction, args = (),)
    
    #starter threads
    incrementing.start()
    decrementing.start()

    #venter til begge tråder er ferdig
    incrementing.join()
    decrementing.join()
    
    print("The magic number is %d" % (i))


main()
