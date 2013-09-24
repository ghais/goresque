// Copyright 2013 Ghais Issa
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Package goresque provides a client for working with the Resuque
package goresque

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

type Client struct {
	conn redis.Conn
}

type job struct {
	Class string        `json:"class"`
	Args  []interface{} `json:"args"`
}

// Dial establishes a connection to the redis instance at url
func Dial(url string) (*Client, error) {
	c, err := redis.Dial("tcp", url)
	if err != nil {
		return nil, err
	}
	return &Client{c}, nil
}

func (c *Client) Enqueue(class, queue string, args ...interface{}) error {
	var j = &job{class, makeJobArgs(args)}

	job_json, _ := json.Marshal(j)

	if err := c.conn.Send("LPUSH", "resque:queue:"+queue, job_json); err != nil {
		return err
	}
	return c.conn.Flush()

}

func (c *Client) Close() error {
	return c.conn.Close()
}

//A trick to make [{}] json struct for empty args
func makeJobArgs(args []interface{}) []interface{} {
	if len(args) == 0 {
		// NOTE: Dirty hack to make a [{}] JSON struct
		return append(make([]interface{}, 0), make(map[string]interface{}, 0))
	}

	return args
}
