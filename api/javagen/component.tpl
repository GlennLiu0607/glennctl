/*
 * @Author: GlennLiu <glennliu0607@gmail.com>
 * @Date: {{.Date}}
 * @LastEditors: Glenn 18322653727@163.com
 * @LastEditTime: 2025-10-28 14:32:49
 * @FilePath: \glennctl\rpc\generator\rpc.tpl
 * @Description:
 *		glennctl {{.version}}
 * Copyright (c) 2025 by 天津晟源士兴科技有限公司, All Rights Reserved.
 */

package com.xhb.logic.http.packet.{{.packet}}.model;

import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;
{{.imports}}

public class {{.className}} extends {{.superClassName}} {

{{.properties}}
{{if .HasProperty}}

	public {{.className}}() {
	}

	public {{.className}}({{.params}}) {
{{.constructorSetter}}
	}
{{end}}

{{.getSet}}
}
