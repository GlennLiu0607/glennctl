/*
 * @Author: GlennLiu <glennliu0607@gmail.com>
 * @Date: {{.Date}}
 * @LastEditors: Glenn 18322653727@163.com
 * @LastEditTime: {{.LastEditTime}}
 * @FilePath: {{.FilePath}}
 * @Description:
 *
 * Copyright (c) 2025 by 天津晟源士兴科技有限公司, All Rights Reserved.
 */

package model

import (
    "context"
)

// TODO: Replace with actual TDengine client and logic
type {{.Type}}Model interface{
    Insert(ctx context.Context, data *{{.Type}}) error
    FindOne(ctx context.Context, id int64) (*{{.Type}}, error)
    Update(ctx context.Context, data *{{.Type}}) error
    Delete(ctx context.Context, id int64) error
}

type default{{.Type}}Model struct{}

func newDefault{{.Type}}Model() *default{{.Type}}Model {
    return &default{{.Type}}Model{}
}