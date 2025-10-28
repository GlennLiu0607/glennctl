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

Name: {{.serviceName}}.rpc
ListenOn: 0.0.0.0:8080
Etcd:
  Hosts:
  - 127.0.0.1:2379
  Key: {{.serviceName}}.rpc
