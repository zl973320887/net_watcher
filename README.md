## Usage ↯↯↯
```
help cmd:

172-1-1-120:watcher liangzhang$ ./watcher -help
Usage of ./watcher:
  -count int
        Tcp test conuts (default -1)
  -host string
        Dest host ipaddress
  -port int
        Dest host port
  -protocol string
        Set network protocol (default "tcp")
  -timeout int
        Tcp test timeout (default -1)

```

## Example ↯↯↯

- always ping

```
172-1-1-120:watcher liangzhang$ ./watcher -host 101.36.113.254 -port 9092
2021-08-19 11:18:54 Connected to  101.36.113.254:9092 protocol=tcp timeout=53.506ms
2021-08-19 11:18:55 Connected to  101.36.113.254:9092 protocol=tcp timeout=50.345ms
2021-08-19 11:18:56 Connected to  101.36.113.254:9092 protocol=tcp timeout=47.366ms
2021-08-19 11:18:57 Connected to  101.36.113.254:9092 protocol=tcp timeout=47.629ms
2021-08-19 11:18:58 Connected to  101.36.113.254:9092 protocol=tcp timeout=53.256ms
2021-08-19 11:18:59 Connected to  101.36.113.254:9092 protocol=tcp timeout=46.845ms
2021-08-19 11:19:00 Connected to  101.36.113.254:9092 protocol=tcp timeout=48.671ms
2021-08-19 11:19:01 Connected to  101.36.113.254:9092 protocol=tcp timeout=51.739ms
2021-08-19 11:19:02 Connected to  101.36.113.254:9092 protocol=tcp timeout=46.863ms
2021-08-19 11:19:03 Connected to  101.36.113.254:9092 protocol=tcp timeout=43.103ms
2021-08-19 11:19:04 Connected to  101.36.113.254:9092 protocol=tcp timeout=45.633ms
2021-08-19 11:19:06 Connected to  101.36.113.254:9092 protocol=tcp timeout=45.025ms
2021-08-19 11:19:07 Connected to  101.36.113.254:9092 protocol=tcp timeout=45.463ms
2021-08-19 11:19:08 Connected to  101.36.113.254:9092 protocol=tcp timeout=47.411ms
2021-08-19 11:19:09 Connected to  101.36.113.254:9092 protocol=tcp timeout=45.782ms
2021-08-19 11:19:10 Connected to  101.36.113.254:9092 protocol=tcp timeout=59.995ms
2021-08-19 11:19:11 Connected to  101.36.113.254:9092 protocol=tcp timeout=44.642ms
2021-08-19 11:19:12 Connected to  101.36.113.254:9092 protocol=tcp timeout=46.045ms
.
.
.
```
- ping once
```
172-1-1-120:watcher liangzhang$ ./watcher -host 101.36.113.254 -port 9092 -count=1
2021-08-19 11:18:45 Connected to  101.36.113.254:9092 protocol=tcp timeout=52.707ms
```
