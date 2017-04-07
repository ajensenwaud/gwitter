Gwitter is a Twitter client for the UNIX command line. Gwitter is written in Go. 

Current dependencies:
* Go oauth library written by mrjones (https://github.com/mrjones/oauth) 
* Gcfg Go ini config file library (https://gopkg.in/gcfg.v1)
* Nothing else.

This project is currently a toy, but will be expanded over time into a fully functioning client. Contributions are much welcome!

To-do list: 
* Non-colour flag for better $PAGER support (piping through less looks weird)
* 'tail -f' style support where gwitter will run in the background and output new tweets as they appear in the main feed
* Simple search functionality
