# Description #

We have a door. The door has a magnetic latch lock. There is a cable. And that's it...

It's one of those system where you swipe a RFID card and the door goes *bzzzz* and you can push it open.

Only all the brains were taken out by the last tenant and I didn't really feel like buying a whole new system.

And I wanted to have some fun.

So got myself (from home garage or office stuff):
- 220V AC to 12V DC converter, to power the door magnet
- RPi3
- 4 port Sainsmart relay
- some wires and a breadboard

Wire up the RPi to drive the relay from a GPIO and 5V VCC.

The program has a loop to open the lock and hold it open for 3 seconds, then close it back. And a web server with a ded simple html page to trigger the operation. Security is out of scope and provided by the network.

Problems I had:
- Tried to drive the relay board with 3.3V and the switches did not flip, but led indicator did...
- Wrong GPIO wiring, door always open!
