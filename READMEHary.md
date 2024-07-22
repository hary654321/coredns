./coredns -dns.port 54
.:54
CoreDNS-1.11.3
linux/amd64, go1.21.11, a6338e9-dirty
[INFO] 127.0.0.1:58270 - 16968 "A IN www.baidu.com. udp 42 false 4096" NOERROR qr,aa,rd 97 0.000888833s



dig @127.0.0.1 -p 54 www.baidu.com

; <<>> DiG 9.10.3-P4-Ubuntu <<>> @127.0.0.1 -p 54 www.baidu.com
; (1 server found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 16968
;; flags: qr aa rd; QUERY: 1, ANSWER: 0, AUTHORITY: 0, ADDITIONAL: 3
;; WARNING: recursion requested but not available

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 4096
;; QUESTION SECTION:
;www.baidu.com.			IN	A

;; ADDITIONAL SECTION:
www.baidu.com.		0	IN	A	127.0.0.1
_udp.www.baidu.com.	0	IN	SRV	0 0 58270 .

;; Query time: 1 msec
;; SERVER: 127.0.0.1#54(127.0.0.1)
;; WHEN: Fri Jul 05 01:54:15 PDT 2024
;; MSG SIZE  rcvd: 108
