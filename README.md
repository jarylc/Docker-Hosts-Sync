[![Docker Hosts Sync Logo](docker-hosts-sync.png)](https://gitlab.com/jarylc/docker-hosts-sync)

# Docker Hosts Sync
Small and configurable Docker image to periodically synchronize Docker host's /etc/hosts file with Docker

[**Docker Hub Image »**](https://hub.docker.com/r/jarylc/docker-hosts-sync)

[**Explore the docs »**](https://gitlab.com/jarylc/docker-hosts-sync)

[Report Bugs](https://gitlab.com/jarylc/docker-hosts-sync/-/issues/new?issuable_template=Bug)
· [Request Features](https://gitlab.com/jarylc/docker-hosts-sync/-/issues/new?issuable_template=Feature%20Request)


## About The Project
This application makes possible to communicate via network with Docker images using container names on the Docker host.
### Environment Variables
| Environment | Default value | Description
|---|---|---|
| INTERVAL | 60 | Interval in seconds to update /etc/hosts |
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
  -e INTERVAL=60 \
  -e EXIT_RESET=1 \
  -v /var/run/docker.sock:/var/run/docker.sock:ro \
  -v /etc/hosts:/etc/hosts \
  jarylc/docker-hosts-sync
```
### Docker-compose
```docker-compose
docker-hosts-sync:
    image: jarylc/docker-hosts-sync
    environment:
      - INTERVAL=60
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
Distributed under the GNU GENERAL PUBLIC LICENSE V3. See `LICENSE` for more information.


## Contact
Jaryl Chng - git@jarylchng.com

https://jarylchng.com

Project Link: [https://gitlab.com/jarylc/docker-hosts-sync/](https://gitlab.com/jarylc/docker-hosts-sync/)
