# overlap

*overlap check if multiple IP ranges overlap.*

## Supported IP range formats

- `172.18.0.1` / `fd00::1`
- `172.18.0.0/24` / `fd00::/64`
- `172.18.0.1-10` / `fd00::1-a`
- `172.18.0.1-172.18.1.10` / `fd00::1-fd00::1:a`

## Install

```shell
go install github.com/iiiceoo/overlap
```

## Help

```shell
$ overlap
Usage: overlap [-v] [-f file] [IP ranges...]
       overlap [-V]

Options:
  -V    Display the version of overlap.
  -f string
        The file path of the IP ranges list, which supports the following
        formats of IP ranges:
            (IPv4)                  (IPv6)
            172.18.0.1              fd00::1
            172.18.0.0/24           fd00::/64
            172.18.0.1-10           fd00::1-a
            172.18.0.1-172.18.1.10  fd00::1-fd00::1:a
  -v    Be verbose, display details of overlaping IP ranges
```

## Example

```shell
$ cat << EOF > list
172.18.40.1
172.18.40.0/24
172.18.0.1-20
EOF

$ overlap -f list
Overlaping :(

$ overlap -v -f list
Overlaping :(

Details:
172.18.40.1 and 172.18.40.0/24 overlap at [172.18.40.1]

$ overlap -v -f list 172.18.0.10-15
Overlaping :(

Details:
172.18.0.10 and 172.18.0.1-20 overlap at [172.18.0.10-172.18.0.15]
172.18.40.1 and 172.18.40.0/24 overlap at [172.18.40.1]
```

## License

Package iprange is MIT-Licensed.
