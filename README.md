[![Docker Hosts Sync Logo](docker-hosts-sync.png)](https://gitlab.com/jarylc/docker-hosts-sync)

# Docker Hosts Sync
Lightweight Docker image to synchronize Docker host's /etc/hosts file with Docker on container start and destruction.

[**Docker Hub Image »**](https://hub.docker.com/r/jarylc/docker-hosts-sync)

[**Explore the docs »**](https://gitlab.com/jarylc/docker-hosts-sync)

[Report Bugs](https://gitlab.com/jarylc/docker-hosts-sync/-/issues/new?issuable_template=Bug)
· [Request Features](https://gitlab.com/jarylc/docker-hosts-sync/-/issues/new?issuable_template=Feature%20Request)


## About The Project
This application makes possible to communicate via network with Docker containers via their container names on the Docker host. It's like as though you were communicating with another container within one!

### Features
- Super lightweight, final docker image based on `scratch`!
- Updates only on container start and destruction, doesn't waste resources by checking periodically (i.e. every 5 seconds).
- `EXIT_RESET` flag to set if changes to `/etc/hosts/` should be reset on exit.

### Environment Variables
| Environment | Default value | Description
|---|---|---|
| EXIT_RESET | 1 | Reset /etc/hosts on exit |

### Built With
* [golang](https://golang.org/)
* [go-dockerclient](https://github.com/fsouza/go-dockerclient)


## Getting Started
To get a local copy up and running follow these simple steps.

### Docker Run
```shell
docker run -it -d \
  --name docker-hosts-sync \
  --privileged \
  -e EXIT_RESET=1 \
  -v /var/run/docker.sock:/var/run/docker.sock:ro \
  -v /etc/hosts:/etc/hosts \
  jarylc/docker-hosts-sync
```

### Docker-compose
```docker-compose
docker-hosts-sync:
    image: jarylc/docker-hosts-sync
    privileged: true
    environment:
      - EXIT_RESET=1
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
      - /etc/hosts:/etc/hosts
```


## Development
### Building
```shell
$ cd /path/to/project/folder
$ go build -ldflags="-w -s"
```

### Docker build
```shell
$ cd /path/to/project/folder
$ docker build .
```


## Roadmap
See the [open issues](https://gitlab.com/jarylc/docker-hosts-sync/-/issues) for a list of proposed features (and known
issues).


## Contributing
Feel free to fork the repository and submit pull requests.


## License
Distributed under the MIT License. See `LICENSE` for more information.


## Contact
Jaryl Chng - git@jarylchng.com

https://jarylchng.com

Project Link: [https://gitlab.com/jarylc/docker-hosts-sync/](https://gitlab.com/jarylc/docker-hosts-sync/)
