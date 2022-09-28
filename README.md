```
 ▄▄▄· .▄▄ ·  ▄▄· ▪  ▪      ▄▄▌   ▄▄▄·  ▐ ▄     • ▌ ▄ ·.  ▄▄▄·  ▄▄▄·
▐█ ▀█ ▐█ ▀. ▐█ ▌▪██ ██     ██•  ▐█ ▀█ •█▌▐█    ·██ ▐███▪▐█ ▀█ ▐█ ▄█
▄█▀▀█ ▄▀▀▀█▄██ ▄▄▐█·▐█·    ██▪  ▄█▀▀█ ▐█▐▐▌    ▐█ ▌▐▌▐█·▄█▀▀█  ██▀·
▐█ ▪▐▌▐█▄▪▐█▐███▌▐█▌▐█▌    ▐█▌▐▌▐█ ▪▐▌██▐█▌    ██ ██▌▐█▌▐█ ▪▐▌▐█▪·•
 ▀  ▀  ▀▀▀▀ ·▀▀▀ ▀▀▀▀▀▀    .▀▀▀  ▀  ▀ ▀▀ █▪    ▀▀  █▪▀▀▀ ▀  ▀ .▀   
```

```

		+-----------------------------+
		|                             |
		|     <192.68.1.64>           |
		|                             |
		|                       *     |
		+-----------------------+-----+
		                        |
		                        |
		                        |
		+-----------------------|-----+
		|                       *     |
		|     <192.68.1.66>           |
		| [+] 80:http                 |
		| [+] 111:rpcbind             |
		|                             |
		|                       *     |
		+-----------------------+-----+
		                        |
		                        |
		                        |
		+-----------------------|-----+
		|                       *     |
		|     <192.68.1.109>          |
		| [+] 53:domain dnsmasq       |
		| [+] 111:rpcbind             |
		|                             |
		|                       *     |
		+-----------------------+-----+
		                        |
		                        |
		                        |
		+-----------------------|-----+
		|                       *     |
		|     <192.68.1.254>          |
		| [+] 80:http                 |
		| [+] 443:http                |
		| [+] 8200:upnp               |
		|                             |
		|                             |
		+-----------------------------+
```

---

1. Runs nmap using the awesome [Ullaakut/nmap](https://github.com/Ullaakut/nmap) Idiomatic nmap library for go on the LAN.
2. Generates a basic ASCII network diagram
3. Generates a 'hand drawn' version of the ASCII network diagram using [esimov/diagram](https://github.com/esimov/diagram)
4. Generates a bootstrap themed web report from the LAN scan using [honze-net/nmap-bootstrap-xsl](https://github.com/honze-net/nmap-bootstrap-xsl) and xsltproc

---

### How to use:
#### Pre-requisits
1. If go is not installed, [install go](https://golang.org/doc/install)
2. If diagram is not installed, [install diagram](https://github.com/esimov/diagram)

#### Usage from source

1. clone repo

```
git clone https://github.com/Bedrovelsen/ascii-LAN-map
cd ascii-LAN-map
```

2a. Build binary and execute

```
go build .
./asciiLANmap
```
OR

2b. run with go run directly

```
go run main.go
```

---
![LAN_DRAWING](https://github.com/Bedrovelsen/ascii-LAN-map/blob/master/LAN_DRAWING.png)

---
Notes:
1. Requires [esimov/diagram](https://github.com/esimov/diagram).
2. Built in LAN CIDR network range discovery (Tested under debian and previously under Mac OS X, don't hesitate to make issue if fails to fetch CIDR on your setup)
