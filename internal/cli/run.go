package cli

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
	c "github.com/urfave/cli/v3"
)

func Run() {
	cmd := &c.Command{
		Name:    "docpst",
		Version: "v0.1.0",
		// Authors: []any{
		// 	&c.Author{
		// 		Name:  "Alejandro Benitez",
		// 		Email: "alexxbez@proton.me",
		// 	},
		// },
		Copyright: "(c) 2025 Alejandro Benitez",
		Usage:     "Manage tyspt documents",
		UsageText: "docpst <command> [flags] [args]",
		Commands: []*c.Command{
			&cli.Command{
				Name:        "new",
				Aliases:     []string{"n"},
				Usage:       "Create a new project",
				UsageText:   "docpst new NAME",
				Description: "Creates a new directory NAME the selected or default template.",
				ArgsUsage:   "NAME",
				Action: func(ctx context.Context, cmd *c.Command) error {
					if cmd.Args().Len() != 1 {
						fmt.Printf("Usage: %v\n", cmd.UsageText)
						return c.Exit("Incorrect usage", 1)
					}
					create(cmd.Args().Get(0))
					return nil
				},
			},
			&cli.Command{
				Name:        "watch",
				Aliases:     []string{"w"},
				Usage:       "Watches and opens pdf viewer",
				UsageText:   "docpst watch",
				Description: "Executes the typst watch command and opens the pdf viewer.",
				Action: func(ctx context.Context, cmd *c.Command) error {
					if cmd.Args().Len() != 0 {
						fmt.Printf("Usage: %v\n", cmd.UsageText)
						return c.Exit("Incorrect usage", 1)
					}
					watch()
					return nil
				},
			},
			&cli.Command{
				Name:        "set",
				Aliases:     []string{"s"},
				Usage:       "Sets up different settings",
				UsageText:   "docpst set [ARGS]",
				Description: "Allows to setup different parts of the configuration.",
				Commands: []*c.Command{
					&cli.Command{
						Name:        "config",
						Aliases:     []string{"configuration", "conf"},
						Usage:       "Sets up the configuration directory.",
						UsageText:   "docpst set config",
						Description: "Sets up the docpst config file. In UNIX this is $HOME/.config. In Windows this is %AppData%.",
						Action: func(ctx context.Context, cmd *c.Command) error {
							if cmd.Args().Len() != 0 {
								fmt.Printf("Usage: %v\n", cmd.UsageText)
								return c.Exit("Incorrect usage", 1)
							}
							setup()
							return nil
						},
					},
				},
			},
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
