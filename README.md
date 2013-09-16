monitor
=======

Client daemon for server monitoring system for *NIX servers.

Checks
------

* fsmon - raise a warning or critical alert depending on how full the filesystem is
* procmon - raise an alert if a given process is not running
* httpmon - raise an alert if a URL is not available
* logmon - raise an alert if "error" is found in one of the configured log files
