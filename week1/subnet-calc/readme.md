# Subnet Calculator

CLI tool that takes a CIDR block and outputs network details.


This Lab was done to learn about how IP addresses are handled in GO. 
I accepted and IP Address range and decoded all the details such as - 
Gateway IP
Total IP
Network range
First usable and Last usable address within the network range.

# how to run it

go run main.go 10.0.0.0/16

# Example Ouput 

Parsing: 10.0.0.0/16
IP: 10.0.0.0
Network: 10.0.0.0
Mask: ffff0000
Broadcast  10.0.255.255
Total Hosts  65536
firstUsable IP :  10.0.0.1
lastUsable IP :  10.0.255.254