<h1>Cisco Network Basic Notes

# Part 1.  Network Devices and Protocols

## Network Components

- End User Components:
  - Smartphones
  - Tablets
  - Printers

- Network Devices:
  - Servers
  - Switches
  - Routers
  - Satellite
  
## Network Connections

- Wired
  - Coaxial Cables [Copper cable used by TV companies.]
  - RJ45 5e Cables [4 Wires group together twisted in pairs.]
  - Fibre Optic    [Glass or Plastic cable.]

- Wireless
  - GPS
  - Wifi
  - Bluetooth

## Network Protocols

- What is a Protocol:
  - Protocol is rules of communication.
- What is a Protocol in Computing:
  - Protocols are used in networking to allows computers to communicate with each other.

### Types of Network Protocols

- Ethernet - Governs NIC (Network Interface Cards) to communicate between each other.
- IP - Internet Protocol
- TCP [Transmission Control Protocol]: Governs Information between networks
- HTTP [HyperText Transfer Protocol] : Governs HTML and WWW Protocols.

# Part 2.  Network Models

## The OSI Reference models

Types of Models:

- Protocol Model
  - Include a set of protocols used to communicate with the data network.
    - e.g. TCP/IP Model - It describes the functions at each layer of the protocol list.
- Reference Model
  - Describe that functions need to be completed at a specific layer but not exactly how they should be completed.
  - It does not provide a sufficient detail level but only how protocols should work with each other.
    - e.g. OSI Model Layer. [Open Systems Interconnection]
      - Used for data network design, operation specifications, and troubleshooting.

## OSI Model LayerIndicates

| TCP/IP  Layer | OSI Model Layer | Layer Description | Protocols Used |
| :---          |    :----:       |          ---: | -----------|
|Application    |7 - Application   |Applications can interact with network services [Process to Process]| HTTP, FTP, DNS, SNMP, Telnet|
|Application    |6-  Presentation  |Data is useable format [Data Encryption]    | SSL, TSL   |
|Application    |5 - Session       |Maintain connections and controls ports and sessions     | NetBIOS, PPTP|
|Transport      |4-  Transport     |Transmit data using transmission protocols.      | TCP, UDP   |
|Internet       |3-  Network       |Controls which physical path the data will take     | TP, ARP, ICMP, IPSec|
|Network Access |2-  Data Link     |Controls the network data format [Exchange data frames between devices]    | PPP, ATM, Ethernet|
|Network Access |1-  Physical      |Transmit raw bit stream over the physical medium      | USB, Bluetooth|

## Encapsulation

Definition:

- A message sent over the computer network, in which a message is inside another message package.

Example:

- A letter being inside a envelope to be sent on a delivery.

## De-Encapsulation

Definition:

- Reversed of Encapsulation. The package is open and the message is received.

Example:

- The letter is open by the recipient and removed from the envelope.

## Ethernet Frame

Definition:

- It defines many aspect of network communication.

### Ethernet Frame Table

Definition:

- This frame holds information to be delivered to the NIC (Network Interface card)

#### 1 byte = 8 bits

|Step     |Ethernet Frame            | Bytes    |Frame Description|
| --------------           | --------| ---------|-----------------|
|Step - 1 |Preamble                  |7 bytes        |Sync different NIC together|
|Step - 2 |Start Frame Delimiter     |1 byte   |Indicates following frame information.
|Step - 3 |Destination MAC Address   |6 bytes         |Indicates the network device end-point MAC address|
|Step - 4 |Source MAC Address        |6 bytes         |Indicates the network device start MAC address    |
|Step - 5 |Length Type               |2 bytes         |1. Length of the Data [Payload] 2. Type of Data [IPv4, IPv6]|
|Step - 6 |Data [Payload]            1|46 - 1500 bytes  |Encapsulated Data [IPv4, IPv6 with other protocols]|
|Step - 7 |Frame Check Sequence      |4 bytes         |Check for errors during transmission.

## The Access Layer

Definition:

- This layer provides connected devices (Ethernet connected devices) access to printers and files inside a network.

Obsolete Media :

- The Ethernet hub has multiple ports used to connect the hosts.
However, only one message can be simultaneously sent. Two or more messages can clog a network and cause collisions.
- Network switches are used instead.

### Network Switches

Switches are layer 2 devices [Data Link] that send messages between interconnected hosts.

A host sends a message to another connected host.

- The switch accepts the request and decodes the Ethernet frame to read the destination MAC address of the message.
- The switch table (MAC table) holds all attached MAC addresses between the host and the end-host
  - The message can also be encapsulated
- When a message is sent.
  - It checks whether the destination MAC address is on the table.
  - Build a connection (circuit) between source and destination.
  - It is sent and received simultaneously.

- A switch builds its MAC address table using
  - Examining the source MAC address of each ethernet frame that is sent between hosts
  - When a new host responds to a flooded message.
  - The switch learns its MAC address and connected ports.
  - The table updates each time a new MAC address is connected.
