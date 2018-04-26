package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {

	log.Info("Starting Wrapper")

	//load config
	config := Config{}
	if _, err := os.Stat("/usr/share/ansible/summit/config.yaml"); os.IsNotExist(err) {
		configor.Load(&config, "config.yaml")
	} else {
		configor.Load(&config, "/usr/share/ansible/summit/config.yaml")
	}

	app := cli.NewApp()

	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		b, err := ioutil.ReadFile("/etc/motd") // just pass the file name
		if err != nil {
			fmt.Print(err)
		}
		str := string(b)
		fmt.Println(str)
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "scenario, S, s", Usage: "The Name of the scenario"},
		cli.StringFlag{Name: "action, A, a", Usage: "action", Value: "init"},
		cli.BoolFlag{Name: "list, L, l", Usage: "List all scenarios"}}

	app.Name = "RH Summit Labs"
	app.Version = "0.0.2"
	app.Usage = "Init Labs Scenarios"
	app.Authors = []cli.Author{
		cli.Author{
			Name: "Dream Team",
		},
	}
	app.Action = func(c *cli.Context) error {

		if c.IsSet("list") {
			for key, item := range config.Scripts {
				fmt.Println("---------------------------------------------------------------------")
				fmt.Println(fmt.Sprintf("Scenario %d, %s", key, item.Name))
				fmt.Println("Description:", item.Description)
				fmt.Println("Actions:", item.Actions)
				fmt.Println("To init this scenario execute:")
				fmt.Println(fmt.Sprintf("  lab -s %d -a init", key))
				fmt.Println("")

			}
			return nil
		}

		if c.IsSet("action") || c.IsSet("scenario") {
			log.Infof("Execute scenario %v and phase %v", c.String("scenario"), c.String("action"))
			fmt.Println()
			key, _ := strconv.Atoi(c.String("scenario"))
			err := execCommand(config, key, c.String("action"))
			if err != nil {
				log.Errorf("Error while executing scenario [%s] and action [%s]. error: %s", config.Scripts[key].Name, c.String("action"), err.Error())
			} else {
				log.Infof("Scenario [%s] activated ", config.Scripts[key].Name)
			}

		}

		return nil
	}

	app.Run(os.Args)
}

func execCommand(config Config, key int, action string) error {

	cmd := config.Scripts[key].Executor.Script
	args1 := strings.Split(config.Scripts[key].Executor.ActionMapper[action], " ")
	executable := exec.Command(cmd, args1...)
	executable.Dir = config.Scripts[key].Executor.Path
	if out, err := executable.Output(); err != nil {
		log.Errorf("error: %v and %v", err.Error(), string(out))
		return err
	}
	return nil
}
