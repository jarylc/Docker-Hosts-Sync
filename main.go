package main

import (
	"fmt"
	docker "github.com/fsouza/go-dockerclient"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
)

var debug = os.Getenv("DEBUG") == "1"

var path = "/etc/hosts"
var separator = "# DOCKER-HOST-SYNC - AUTO GENERATED - DO NOT REMOVE/EDIT #"
var regex = regexp.MustCompile("\\n\\n" + separator + "(.|\\n)*" + separator)

type host struct {
	name string
	ip   string
}

var hosts []host

func add(name string, ip string) {
	hosts = append(hosts, host{
		name: name,
		ip:   ip,
	})
}
func clear() {
	hosts = []host{}
}

//nolint:funlen
func main() {
	var interrupt = make(chan os.Signal, 1)
	defer close(interrupt)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	client, err := docker.NewClientFromEnv()
	if err != nil {
		log.Panic(err)
	}

	change := make(chan *docker.APIEvents)
	err = client.AddEventListenerWithOptions(docker.EventsOptions{
		Filters: map[string][]string{
			"type":  {"container"},
			"event": {"start", "destroy"},
		},
	}, change)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Started.")
	for {
		containers, err := client.ListContainers(docker.ListContainersOptions{All: true})
		if err != nil {
			log.Panic(err)
		}
		clear()
		for _, container := range containers {
			if strings.Contains(container.Image, "docker-hosts-sync") { // ignore this container
				continue
			}
			for _, name := range container.Names {
				for _, network := range container.Networks.Networks {
					if network.IPAddress != "" {
						if debug {
							log.Printf("%s - %s", name[1:], network.IPAddress)
						}
						add(name[1:], network.IPAddress)
					}
				}
			}
		}
		err = update()
		if err != nil {
			log.Panic(err)
			return
		}
		select {
		case <-interrupt:
			if os.Getenv("EXIT_RESET") == "1" {
				log.Println("Resetting...")
				err = reset()
				if err != nil {
					log.Panic(err)
					return
				}
			}
			return
		case <-change:
			if debug {
				log.Println("Containers changed.")
			}
		}
	}
}

func update() error {
	existing, err := read()
	if err != nil {
		return err
	}

	if !regex.MatchString(existing) { // if does not exist, create separators
		existing += "\n\n" + separator + "\n" + separator
	}

	out := "\n\n" + separator
	for _, h := range hosts {
		out += fmt.Sprintf("\n%s\t%s", h.ip, h.name)
	}
	out += "\n" + separator

	err = write(existing, regex.ReplaceAllString(existing, out))
	if err != nil {
		return err
	}

	return nil
}
func reset() error {
	existing, err := read()
	if err != nil {
		return err
	}
	err = write(existing, regex.ReplaceAllString(existing, ""))
	if err != nil {
		return err
	}
	return nil
}
func read() (string, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil { // create hosts file if doesn't exist
		err = ioutil.WriteFile(path, []byte{}, 0644) //nolint:gosec
		if err != nil {
			return "", err
		}
		raw, err = ioutil.ReadFile(path)
		if err != nil {
			return "", err
		}
	}
	return string(raw), nil
}
func write(existing string, out string) error {
	if existing == out {
		return nil
	}
	err := ioutil.WriteFile(path, []byte(out), 0644) //nolint:gosec
	if err != nil {
		return err
	}
	return nil
}
