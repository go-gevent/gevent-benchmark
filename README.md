## gevent benchmark

benchmark tool:

-  [tcpkali](https://github.com/machinezone/tcpkali) for Echo

| OS       | Package manager                         | Command                |
| -------- | --------------------------------------- | ---------------------- |
| Mac OS X | [Homebrew](http://brew.sh/)             | `brew install tcpkali` |
| Mac OS X | [MacPorts](https://www.macports.org/)   | `port install tcpkali` |
| FreeBSD  | [pkgng](https://wiki.freebsd.org/pkgng) | `pkg install tcpkali`  |
| Linux    | [nix](https://nixos.org/nix/)           | `nix-env -i tcpkali`   |

Run：

```bash
 ./bench.sh      
```

## Note

-  Running inside a container
-  Single threaded mode (GOMAXPROC=1)
