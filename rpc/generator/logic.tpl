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

package {{.packageName}}

import (
	"context"

	{{.imports}}

	"github.com/zeromicro/go-zero/core/logx"
)

type {{.logicName}} struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func New{{.logicName}}(ctx context.Context,svcCtx *svc.ServiceContext) *{{.logicName}} {
	return &{{.logicName}}{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
{{.functions}}
