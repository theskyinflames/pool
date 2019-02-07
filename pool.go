/*
Copyright 2015 - Jaume Arús9

Author Jaume Arús - jaumearus@gmail.com

Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package pool

type (
	Poolable struct {
		Item interface{}
	}

	Pool struct {
		size int32
		pool chan (*Poolable)
	}
)

func NewPool(size int32) *Pool {
	return &Pool{size, make(chan (*Poolable), size)}
}

func (p *Pool) GetSize() int32 {
	return p.size
}

func (p *Pool) GetPool() chan (*Poolable) {
	return p.pool
}
