package config

import (
	"flag"

	"github.com/872409/gatom/log"
)

func Usage(appName string, appVersion string, appDesc string) {
	log.Infof(`Config Init

* ============================= *

%s [version: %s]
%s

* ============================= *

Options:
`, appName, appVersion, appDesc)
	flag.PrintDefaults()
}
