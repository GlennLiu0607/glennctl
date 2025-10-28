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

package {{.pkgName}}

import (
	{{.imports}}
)

type {{.logic}} struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

{{if .hasDoc}}{{.doc}}{{end}}
func New{{.logic}}(ctx context.Context, svcCtx *svc.ServiceContext) *{{.logic}} {
	return &{{.logic}}{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *{{.logic}}) {{.function}}({{.request}}) {{.responseType}} {
    // todo: add your logic here and delete this line

	{{.returnString}}
}
