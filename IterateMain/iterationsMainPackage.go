package main

import "github.com/robhitsaxena/Atom-GoLang-Learn-Proj-Repo/GoIterations"

func main() {
	//GoIterations.ForLoopmain()
	//GoIterations.SwitchCaseMain()
	GoIterations.SwitchNoExp()
	GoIterations.GoOnTypes(7)
	GoIterations.GoOnTypes("Robhit")
	var t = GoIterations.Contact{"Good to see you,", "Tim"}
	GoIterations.GoOnTypes(t)
	GoIterations.GoOnTypes(t.Greeting)
	GoIterations.GoOnTypes(t.Name)

}
