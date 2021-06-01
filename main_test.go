package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestUpdateAndReset(t *testing.T) {
	err := create()
	if err != nil {
		t.Error(err)
	}
	defer func() {
		err = remove()
		if err != nil {
			t.Error(err)
		}
	}()

	hosts = append(hosts, host{
		name: "test",
		ip:   "172.16.1.2",
	})

	err = update()
	if err != nil {
		t.Error(err)
	}
	current, err := read()
	if current != "127.0.0.1 localhost\n\n# DOCKER-HOST-SYNC - AUTO GENERATED - DO NOT REMOVE/EDIT #\n172.16.1.2\ttest\n# DOCKER-HOST-SYNC - AUTO GENERATED - DO NOT REMOVE/EDIT #" {
		t.Error("update mismatch")
	}

	err = reset()
	if err != nil {
		t.Error(err)
	}
	read, err := read()
	if read != "127.0.0.1 localhost" {
		t.Error("reset failed")
	}
}

// utilities

func create() error {
	tmp, err := ioutil.TempFile("", "hosts")
	if err != nil {
		return err
	}
	path = tmp.Name()
	err = write("", "127.0.0.1 localhost")
	if err != nil {
		return err
	}
	return nil
}
func remove() error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
