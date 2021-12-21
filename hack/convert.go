package main

import (
	"os"
	"path/filepath"

	"github.com/solo-io/go-utils/log"
	"github.com/solo-io/skv2/codegen/util"
	"github.com/solo-io/solo-apis/hack/convert"
)

func main() {
	moduleRoot := util.GetModuleRoot()
	err := filepath.Walk(
		filepath.Join(moduleRoot, "api/gloo"),
		convert.ConvertToSkv2ProtoFile,
	)
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
}
