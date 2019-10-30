# simple-http-server

A simple HTTP server that advertises itself to the network with [DNS Service Discovery (DNS-SD)](http://www.dns-sd.org/), also known as Zeroconf (Bonjour, previously Rendezvous). It serves the root directory (`/`) to the network without any security protections. Do not use it if this is not what you want. Since the server is written in Go, it can easily be adjusted to your own needs by editing the source code.

## Download

Precompiled binaries for Windows, Linux, Raspberry Pi, macOS, OpenWrt, precompiled versions are available from https://github.com/probonopd/simple-http-server/releases/tag/continuous

## Usage

On all systems except Windows the server needs to be invoked with `sudo` since it uses port 80.

```
chmod +x ./server-*
sudo ./server-*
```

Browsers like Safari on the Mac can show it like this:

![](https://user-images.githubusercontent.com/2480569/67892246-395d1a80-fb4c-11e9-9aea-aca5a196495f.png)
