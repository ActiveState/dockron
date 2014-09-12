package main

import "strings"
import "os/signal"
import "syscall"
import "os"
import "os/exec"
import "sync"
import "github.com/robfig/cron"

type Cron struct {
	*cron.Cron
	// WaitGroup is used to ensure that at most a single instance
	// of the command is running.
	wg *sync.WaitGroup
}

func NewCron(schedule string, command string, args []string) *Cron {
	println("New cron:", schedule)
	c := &Cron{cron.New(), &sync.WaitGroup{}}

	c.AddFunc(schedule, func() {
		c.wg.Add(1)

		println("Executing:", command, strings.Join(args, " "))
		err := execute(command, args)
		if err != nil {
			println("Failed:", err)
		}else{
			print("Succeeded")
		}
		c.wg.Done()
	})
	return c
}

func (c *Cron) Stop() {
	println("Stopping cron")
	c.Cron.Stop()
	println("Waiting")
	c.wg.Wait()
	print("Exiting")
}

// execute executes the given command
func execute(command string, args []string) error {
	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}
	return cmd.Wait()
}

func main() {
	c := NewCron(os.Args[1], os.Args[2], os.Args[3:len(os.Args)])
	go c.Start()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	println(<-ch)

	c.Stop()
	os.Exit(1)
}
