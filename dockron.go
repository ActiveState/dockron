package main

import "github.com/ActiveState/log"
import "github.com/robfig/cron"
import "os"
import "os/exec"
import "os/signal"
import "strings"
import "sync"
import "syscall"

type Cron struct {
	*cron.Cron
	// WaitGroup is used to ensure that at most a single instance
	// of the command is running.
	wg *sync.WaitGroup
}

func NewCron(schedule string, command string, args []string) *Cron {
	log.Infof("Running per schedule: %v", schedule)
	c := &Cron{cron.New(), &sync.WaitGroup{}}

	c.AddFunc(schedule, func() {
		c.wg.Add(1)

		log.Infof("Executing: %v %v", command, strings.Join(args, " "))
		err := execute(command, args)
		if err != nil {
			log.Warnf("Failed: %v", err)
		} else {
			log.Info("Succeeded")
		}
		c.wg.Done()
	})
	return c
}

func (c *Cron) Stop() {
	log.Warn("Stopping cron")
	c.Cron.Stop()
	log.Info("Waiting")
	c.wg.Wait()
	log.Info("Exiting")
}

// execute executes the given command
func execute(command string, args []string) error {
	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func main() {
	c := NewCron(os.Args[1], os.Args[2], os.Args[3:len(os.Args)])
	go c.Start()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Warnf("%v", <-ch)

	c.Stop()
	os.Exit(1)
}
