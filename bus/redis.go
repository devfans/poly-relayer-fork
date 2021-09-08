/*
 * Copyright (C) 2021 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package bus

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/polynetwork/bridge-common/util"
)

type RedisConn struct {
	Conn    *redis.Client
	options *redis.Options
}

func (c *RedisConn) Key() string {
	return fmt.Sprintf("%s:%d", c.options.Addr, c.options.DB)
}

func (c *RedisConn) Create() interface{} {
	c.Conn = redis.NewClient(c.options)
	return c
}

func New(options *redis.Options) *redis.Client {
	return util.Single(&RedisConn{
		options: options,
	}).(*RedisConn).Conn
}