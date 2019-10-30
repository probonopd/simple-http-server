# simple-http-server

A simple HTTP server that advertises itself to the network with Zeroconf (Bonjour).

It serves the root directory (`/`) to the network without any security protections. Do not use it if this is not what you want.

## Download

For Windows, Linux, macOS, OpenWrt, precompiled versions are available from

https://github.com/probonopd/simple-http-server/releases/tag/continuous

## Usage

On all systems except Windows the server needs to be invoked with `sudo` since it uses port 80.

chmod +x ./server-*
sudo ./server
