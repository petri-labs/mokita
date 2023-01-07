package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	mokita "github.com/petri-labs/mokita/app"
	"github.com/petri-labs/mokita/app/params"
	"github.com/petri-labs/mokita/cmd/mokitad/cmd"
)

func main() {
	params.SetAddressPrefixes()
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, mokita.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
