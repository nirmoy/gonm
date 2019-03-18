# gonm
Network manager written in Golang

## get started
```
go get github.com/nirmoy/gonm; $GOPATH/bin/gonm list
```

## gonm list 
```
gonm list 
+-----------------+------------------------+--------------------------------------------+
|      NAME       |         FLAGS          |                    ADDR                    |
+-----------------+------------------------+--------------------------------------------+
| lo              | up|loopback            | 127.0.0.1/8                                |
|                 |                        | ::1/128                                    |
| wlp1s0          | up|broadcast|multicast | 192.168.178.24/24                          |
|                 |                        | fe80::d00e:8508:5d24:3827/64               |
| br-0c99c62e5376 | up|broadcast|multicast | 172.19.0.1/16                              |
| docker0         | up|broadcast|multicast | 172.17.0.1/16                              |
| br-9fe258918a99 | up|broadcast|multicast | 172.18.0.1/16                              |
| virbr0          | up|broadcast|multicast | 192.168.122.1/24                           |
+-----------------+------------------------+--------------------------------------------+
```
## gonm top
![picture](https://i.ibb.co/NVBx0WD/Screenshot-from-2019-03-17-15-01-33.png)
