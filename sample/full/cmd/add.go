package cmd

import (
	"fmt"
	"github.com/desertbit/grumble"
)

func init() {
	adminCommand := &grumble.Command{
		Name:     "add",
		Help:     "add command",
		LongHelp: "add command",
		Run: func(c *grumble.Context) error {
			App.AddCommand(&grumble.Command{
				Name:     "test1",
				LongHelp: "test1",
				Help:     "test1",
				Run: func(c *grumble.Context) error {
					fmt.Println(c.Flags.String("directory"))
					return fmt.Errorf("failed")
				},
				HelpGroup: "aaa",
			})
			return nil
		},
	}
	App.AddCommand(adminCommand)
	App.SetInterruptHandler(func(a *grumble.App, count int) {
		fmt.Println("1")
	})
	App.OnInit(func(a *grumble.App, flags grumble.FlagMap) error {
		fmt.Println("oninit")
		return nil
	})
	App.OnShell(func(a *grumble.App) error {
		fmt.Println("onshell")
		return nil
	})

}
