# Monitor

Monitor is a client side monitoring suite that reports to a central server.
It is designed to run on predominantly linux hosts but there are also basic
checks included for windows machines.

The monitor wakes up periodically and runs it checks. If it finds a problem
it will raise an alert on the central server. The alert is formed in XML to
ensure that the complete message is transferred via the Internet connection.

# Documentation

This project is written in Go and includes suitable comments to enable `godoc`
to work well.
