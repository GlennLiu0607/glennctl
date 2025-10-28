/*
 * @Author: GlennLiu <glennliu0607@gmail.com>
 * @Date: {{.Date}}
 * @LastEditors: Glenn 18322653727@163.com
 * @LastEditTime: {{.LastEditTime}}
 * @FilePath: {{.FilePath}}
 * @Description:
 *		glennctl {{.version}}
 * Copyright (c) 2025 by 天津晟源士兴科技有限公司, All Rights Reserved.
 */

package middleware

import "net/http"

type {{.name}} struct {
}

func New{{.name}}() *{{.name}} {
	return &{{.name}}{}
}

func (m *{{.name}})Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
