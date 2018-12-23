# eavesdrop

 ## Synopsis
A sniffing app to pull redundant copies of submission forms or 
other Ethernet traffic. Part of Unhackable Server Project. 

## Code Information

Govesdrop, sniffs the network, extracts text and other useful information.
    Using tshark, it initiates a packet capture that pipes the stdout into the Go process.
    The capture pipes the ENTIRE packet and only saves it if the user wants to. 
    
##User Interface
   The user interface, at this point only allows the user (let's call him Dave) to chose with device he want't to sniff on.
   I mean to add some nore functionality:
 
  - [ ] Dave can choose which device to sniff on
  
  - [ ] Dave can choose what type of payload he wants to consider.
  
  -[ ] Dave can go in and see how much data leakage is happening. How many packets have plain text in them and a list of  sources.
       
  - [ ] Dave can see how many total packets are passing through
  
  - [ ] Dave can choose where the payload is being stored on his system
  
  - [ ] Dave can kill the sniff
  
  - [ ] Dave can schedule a sniff 
  
  - [ ] Dave can count/ see a graph of malformed packets to see if Hal is misbehaving
  

## Packaging Eavesdrop:
Using Go has made packaging a lot easier. All I will have to do is compile it. `go build main.go`
## Running the application

The system must have tshark on it but that's it and you must have sudo privilages.
