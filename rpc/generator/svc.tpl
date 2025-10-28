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

package svc

import {{.imports}}

type ServiceContext struct {
    Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
    return &ServiceContext{
        Config:c,
    }
}
