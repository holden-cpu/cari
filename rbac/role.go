/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rbac

type RoleResponse struct {
	Total int64   `json:"total,omitempty"`
	Roles []*Role `json:"data,omitempty"`
}

type Role struct {
	ID         string        `json:"id,omitempty"`
	Name       string        `json:"name,omitempty"`
	Perms      []*Permission `json:"perms,omitempty"`
	CreateTime string        `json:"createTime,omitempty"`
	UpdateTime string        `json:"updateTime,omitempty"`
}

type Permission struct {
	Resources []*Resource `json:"resources,omitempty"`
	Verbs     []string    `json:"verbs,omitempty"`
}
type Resource struct {
	Type   string            `json:"type,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
}
