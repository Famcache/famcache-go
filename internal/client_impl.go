package internal

import (
	"errors"
	"net"
	"strconv"

	"github.com/famcache/famcache-go/domain"
	"github.com/google/uuid"
)

func (c *FamcacheClient) Connect() error {
	url := c.options.GetHost() + ":" + strconv.Itoa(c.options.GetPort())

	socket, err := net.Dial("tcp", url)

	if err != nil {
		return err
	}

	c.socket = socket
	c.messaging.(*Messaging).socket = &socket

	go c.Listen()

	return nil
}

func (c *FamcacheClient) Close() error {
	return c.socket.Close()
}

func (c *FamcacheClient) Set(key string, value string, ttl *uint64) error {
	uuid := uuid.New().String()
	command := c.commands.Set(uuid, key, value, ttl)

	_, err := c.socket.Write([]byte(command))

	if err != nil {
		return err
	}

	task := NewTask()

	c.taskRegistry.Set(uuid, task)

	task.Wait()

	defer c.taskRegistry.Free(uuid)

	if task.IsError() {
		return errors.New("error setting value")
	}

	return nil
}

func (c *FamcacheClient) Get(key string) (string, error) {
	uuid := uuid.New().String()

	command := c.commands.Get(uuid, key)

	_, err := c.socket.Write([]byte(command))

	if err != nil {
		return "", err
	}

	task := NewTask()
	defer c.taskRegistry.Free(uuid)

	c.taskRegistry.Set(uuid, task)

	task.Wait()

	if task.IsError() {
		return "", errors.New("error getting value")
	}

	data, _ := task.GetResult()

	return data, nil
}

func (c *FamcacheClient) Delete(key string) error {
	uuid := uuid.New().String()

	command := c.commands.Delete(uuid, key)

	_, err := c.socket.Write([]byte(command))

	if err != nil {
		return err
	}

	task := NewTask()
	defer c.taskRegistry.Free(uuid)

	c.taskRegistry.Set(uuid, task)

	task.Wait()

	if task.IsError() {
		return errors.New("error deleting value")
	}

	return nil
}

func (c *FamcacheClient) Messaging() domain.Messaging {
	return c.messaging
}
