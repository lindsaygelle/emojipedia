package program

import "strings"

type Runner map[string]*Function

func (runner *Runner) Add(program *Program) *Runner {
	(*runner)[strings.ToUpper(program.Key)] = program.Function
	return runner
}

func (runner *Runner) Get(key string) (function *Function, ok bool) {
	function, ok = (*runner)[strings.ToUpper(key)]
	return function, ok
}

func (runner *Runner) New(programs ...*Program) *Runner {
	for i := range programs {
		runner.Add(programs[i])
	}
	return runner
}
