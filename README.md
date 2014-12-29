# gopro-client
A native desktop client for automatically downloading images and videos from a WiFi-enabled *GoPro* unit. Obviously, this utility requires a computer with a WiFi-receiver as you need to connect to a network hosted by the unit itself. At the moment, we're using a separate Goroutine for downloading each file. This will most likely get replaced in the future as it doesn't scale very well to a lot of files due to the amount of I/O calls.

### Usage
...

### Todo's
 - GUI (Qt?)
 - Check for existing files before overwriting
 - Users should be able to set their own output directory
 - Do some proper error handling
 - Make the program function autonomously in the background like a daemon
 - Write "usage"
