package quickstart

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/GlennLiu0607/glennctl/api/gogen"
	"github.com/GlennLiu0607/glennctl/pkg/golang"
	"github.com/GlennLiu0607/glennctl/util"
	"github.com/GlennLiu0607/glennctl/util/pathx"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	//go:embed idl/greet.api
	apiContent string
	//go:embed idl/svc.tpl
	svcContent string
	//go:embed idl/apilogic.tpl
	apiLogicContent string
	//go:embed idl/api.yaml
	apiEtcContent string

	apiWorkDir string
	rpcWorkDir string
)

func initAPIFlags() error {
	rpcWorkDir = filepath.Join(projectDir, "rpc")
	apiWorkDir = filepath.Join(projectDir, "api")
	if err := pathx.MkdirIfNotExist(apiWorkDir); err != nil {
		return err
	}

	apiFilename := filepath.Join(apiWorkDir, "greet.api")
	// prepend dynamic header to embedded api content
	now := time.Now().Format("2006-01-02 15:04:05")
	winPath := strings.ReplaceAll(filepath.Clean(apiFilename), "/", "\\")
	header := fmt.Sprintf("/*\n * @Author: GlennLiu <glennliu0607@gmail.com>\n * @Date: %s\n * @LastEditors: Glenn 18322653727@163.com\n * @LastEditTime: %s\n * @FilePath: %s\n * @Description:\n *\n * Copyright (c) 2025 by 天津晟源士兴科技有限公司, All Rights Reserved.\n */\n\n", now, now, winPath)
	apiBytes := []byte(header + apiContent)
	if err := os.WriteFile(apiFilename, apiBytes, 0o666); err != nil {
		return err
	}

	gogen.VarStringDir = apiWorkDir
	gogen.VarStringAPI = apiFilename
	return nil
}

type mono struct {
	callRPC bool
}

func newMonoService(callRPC bool) mono {
	m := mono{callRPC}
	m.createAPIProject()
	return m
}

func (m mono) createAPIProject() {
	logx.Must(initAPIFlags())
	log.Debug(">> Generating quickstart api project...")
	logx.Must(gogen.GoCommand(nil, nil))
	etcFile := filepath.Join(apiWorkDir, "etc", "greet.yaml")
	logx.Must(os.WriteFile(etcFile, []byte(apiEtcContent), 0o666))
	logicFile := filepath.Join(apiWorkDir, "internal", "logic", "pinglogic.go")
	svcFile := filepath.Join(apiWorkDir, "internal", "svc", "servicecontext.go")
	configPath := filepath.Join(apiWorkDir, "internal", "config")
	svcPath := filepath.Join(apiWorkDir, "internal", "svc")
	typesPath := filepath.Join(apiWorkDir, "internal", "types")
	svcPkg, _, err := golang.GetParentPackage(svcPath)
	logx.Must(err)
	typesPkg, _, err := golang.GetParentPackage(typesPath)
	logx.Must(err)
	configPkg, _, err := golang.GetParentPackage(configPath)
	logx.Must(err)

	var rpcClientPkg string
	if m.callRPC {
		rpcClientPath := filepath.Join(rpcWorkDir, "greet")
		rpcClientPkg, _, err = golang.GetParentPackage(rpcClientPath)
		logx.Must(err)
	}

	logx.Must(util.With("logic").Parse(apiLogicContent).SaveTo(map[string]any{
		"svcPkg":       svcPkg,
		"typesPkg":     typesPkg,
		"rpcClientPkg": rpcClientPkg,
		"callRPC":      m.callRPC,
	}, logicFile, true))

	logx.Must(util.With("svc").Parse(svcContent).SaveTo(map[string]any{
		"rpcClientPkg": rpcClientPkg,
		"configPkg":    configPkg,
		"callRPC":      m.callRPC,
	}, svcFile, true))
}

func (m mono) start() {
	if !m.callRPC {
		goModTidy(projectDir)
	}
	log.Debug(">> Ready to start an API server...")
	log.Debug(">> Run 'curl http://127.0.0.1:8888/ping' after service startup...")
	goStart(apiWorkDir)
}
