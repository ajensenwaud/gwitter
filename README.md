Gwitter is a Twitter client for the UNIX command line. Gwitter is written in Go. 

Current dependencies:
* Go oauth library written by mrjones (https://github.com/mrjones/oauth) 
* Gcfg Go ini config file library (https://code.google.com/p/gcfg/)
* Nothing else.

This project is currently a toy, but will be expanded over time into a fully functioning client. Contributions are much welcome!

To-do list: 
* Implement proper deserialization of JSON objects to enum for key objects
* Once done, implement "get latest tweets" and write to command line
* If xterm/rxvt-color TERMTYPE is supported, use proper colour for showing latest tweets
