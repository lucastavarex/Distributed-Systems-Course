# TP1

## Table of Contents

<!--ts-->
   * [About](#about)
   * [Requirements](#requirements)
   * [Setting up Program](#setup)
   * [Technologies](#technologies)
<!--te-->

## About

The objective of this project is to become familiar with the main Interprocess Communication (IPC) mechanisms based on message passing. For each part, was required to develop a program in any programming language (I chose golang ), provided it directly supports IPC mechanisms. It was recommended to use C or C++, considering the close integration of their libraries with the operating system, offering the developer greater control over these operations.

In addition to implementation, I tested my program by running case studies. In this task I were also required to prepare a report, with a maximum length of 5 pages, detailing the design and implementation decisions for the specified functionalities, as well as the evaluation of the case studies. The report should include the URL to the source code of your implementation. This project must be completed in pairs.

## Requirements
1 - ```git clone https://github.com/lucastavarex/Distributed-Systems-Course.git``` <br>
2 - Install golang in your PC

### Setup

```bash
# Clone the repository
$ git clone https://github.com/lucastavarex/Distributed-Systems-Course.git

# Go to the project folder
$ cd Distributed-Systems-Course/TP1/

# Run the main program
$ go run main.go

# Once the code starts, then you will be requested to pick which program to run. 

# The 3 implementations are:
# IPC using Pipes (Run the pipe code)
# IPC using Signals (Run the main program 2x, one for signal_rec and other for signal_sen)
# IPC using Sockets (Run the main program 2x, one for socket_server and other for socket_client)
```

## Technologies

* Golang
* Go modules
