// Code scaffolded by goctl. Safe to edit.
// glennctl {{.version}}

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

package config

import {{.authImport}}

type Config struct {
	rest.RestConf
	{{.auth}}
	{{.jwtTrans}}
}
