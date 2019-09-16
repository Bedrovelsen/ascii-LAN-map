```
 ▄▄▄· .▄▄ ·  ▄▄· ▪  ▪      ▄▄▌   ▄▄▄·  ▐ ▄     • ▌ ▄ ·.  ▄▄▄·  ▄▄▄·
▐█ ▀█ ▐█ ▀. ▐█ ▌▪██ ██     ██•  ▐█ ▀█ •█▌▐█    ·██ ▐███▪▐█ ▀█ ▐█ ▄█
▄█▀▀█ ▄▀▀▀█▄██ ▄▄▐█·▐█·    ██▪  ▄█▀▀█ ▐█▐▐▌    ▐█ ▌▐▌▐█·▄█▀▀█  ██▀·
▐█ ▪▐▌▐█▄▪▐█▐███▌▐█▌▐█▌    ▐█▌▐▌▐█ ▪▐▌██▐█▌    ██ ██▌▐█▌▐█ ▪▐▌▐█▪·•
 ▀  ▀  ▀▀▀▀ ·▀▀▀ ▀▀▀▀▀▀    .▀▀▀  ▀  ▀ ▀▀ █▪    ▀▀  █▪▀▀▀ ▀  ▀ .▀   
```

```
		+-----------------------------+
		|     <172.16.1.66>           |
		| [+] 80:http                 |
		| [+] 111:rpcbind             |
		|                             |
		+---------------+-------------+
				|
				|
				|
				V
		+-----------------------------+
		|     <172.16.1.111>          |
		| [+] 62078:tcpwrapped        |
		|                             |
		+---------------+-------------+
				|
				|
				|
				V
		+-----------------------------+
		|     <172.16.1.115>          |
		| [+] 53:domain dnsmasq       |
		| [+] 111:rpcbind             |
		|                             |
		+---------------+-------------+
				|
				|
				|
				V
		+-----------------------------+
		|     <172.16.1.121>          |
		| [+] 22:ssh OpenSSH          |
		| [+] 111:rpcbind             |
		| [+] 443:http Golang n       |
		| [+] 8083:us srv             |
		|                             |
		+---------------+-------------+
				|
				|
				|
				V
		+-----------------------------+
		|     <172.16.1.254>          |
		| [+] 80:http 2Wire Ho        |
		| [+] 443:http 2Wire Ho       |
		| [+] 8200:upnp               |
		|                             |
		+---------------+-------------+
```

---

1. Runs nmap using the awesome [Ullaakut/nmap](https://github.com/Ullaakut/nmap) Idiomatic nmap library for go on the LAN.
2. Generates a basic ASCII network diagram
3. Generates a 'hand drawn' version of the ASCII network diagram using [esimov/diagram](https://github.com/esimov/diagram)
4. Generates a bootstrap themed web report from the LAN scan using [honze-net/nmap-bootstrap-xsl](https://github.com/honze-net/nmap-bootstrap-xsl) and xsltproc

---

![LAN_DRAWING](https://github.com/Bedrovelsen/ascii-LAN-map/blob/master/LAN_DRAWING.png)
