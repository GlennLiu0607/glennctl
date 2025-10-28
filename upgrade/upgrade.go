/*
 * @Author: GlennLiu <glennliu0607@gmail.com>
 * @Date: 2025-10-28 13:55:09
 * @LastEditors: Glenn 18322653727@163.com
 * @LastEditTime: 2025-10-28 14:06:08
 * @FilePath: \glennctl\upgrade\upgrade.go
 * @Description:
 *
 * Copyright (c) 2025 by 天津晟源士兴科技有限公司, All Rights Reserved.
 */
package upgrade

import (
	"fmt"
	"runtime"

	"github.com/GlennLiu0607/glennctl/rpc/execx"
	"github.com/spf13/cobra"
)

// upgrade gets the latest goctl by
// go install github.com/GlennLiu0607/glennctl@latest
func upgrade(_ *cobra.Command, _ []string) error {
	cmd := `go install github.com/GlennLiu0607/glennctl@latest`
	if runtime.GOOS == "windows" {
		cmd = `go install github.com/GlennLiu0607/glennctl@latest`
	}
	info, err := execx.Run(cmd, "")
	if err != nil {
		return err
	}

	fmt.Print(info)
	return nil
}
