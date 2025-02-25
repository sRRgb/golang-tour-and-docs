package main

import (
	"golang-tour-refresh/tasks/basics/flow"
	"golang-tour-refresh/tasks/basics/packages"
	"golang-tour-refresh/tasks/basics/types"
	"golang-tour-refresh/tasks/mi"
)

func main() {
	packages.ProcessPackagesTasks()
	flow.ProcessFlowTasks()
	types.ProcessAllTasks()
	mi.ProcessAllTasks()
}
