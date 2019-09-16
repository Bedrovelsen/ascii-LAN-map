# ascii-LAN-map

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

Runs nmap on the LAN and then generates a basic ASCII network diagram, a 'hand drawn' network diagram and an bootstrap themed web report of the scan

---

![LAN_DRAWING](https://github.com/Bedrovelsen/ascii-LAN-map/blob/master/LAN_DRAWING.png)
