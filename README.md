# Monitor

Monitor is a client side monitoring suite that reports to a central server.
It is designed to run on predominantly linux hosts but there are also basic
checks included for windows machines.

The monitor wakes up periodically and runs it's checks. If it finds a problem
it then raises an alert on the central server (monsrv). The alert is formed in 
XML to ensure that the complete message is transferred via the Internet 
connection.

# Documentation

This project is written in Go and includes suitable comments to enable `godoc`
to provide accurate documentation.
