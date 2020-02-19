# WifiDog Captive Portal Auth Server

This is yet another Authentication Server for the WifiDog Captive Portal Solution.

The server is written in Go.

## Installation
To start with there is included an example binary for testing the solution. Make sure you have your
wifidog gateway installed and running first.

The example binary is in the example folder. The binaries are compiled for Linux x86_64.
A few binaries included in the example folder. 

1. "basic" binary runs a very strip down version of the server. It will authenticated the user
    automatically without requiring any credentials or login page.

2. 