package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/Ak-Army/linereturn/pkg/linereturn"
)

func main() {
	singlechecker.Main(linereturn.NewAnalyzer())
}
