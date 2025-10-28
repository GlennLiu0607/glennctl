{{.head}}

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

package server

import (
	{{if .notStream}}"context"{{end}}

	{{.imports}}
)

type {{.server}}Server struct {
	svcCtx *svc.ServiceContext
	{{.unimplementedServer}}
}

func New{{.server}}Server(svcCtx *svc.ServiceContext) *{{.server}}Server {
	return &{{.server}}Server{
		svcCtx: svcCtx,
	}
}

{{.funcs}}
