package main

import (
	"fmt"
	osexec "os/exec"

	"github.com/cucumber/godog"
)

func iAmAnIssuer() error {
	return godog.ErrPending
}

func myAccountHasTokens(arg1 int) error {
	return godog.ErrPending
}

func iCreateTokens(arg1 int) error {
	return godog.ErrPending
}

func myAccountWillHaveTokens(arg1 int) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I am an issuer$`, iAmAnIssuer)
	ctx.Step(`^my account has (\d+) tokens$`, myAccountHasTokens)
	ctx.Step(`^I create (\d+) tokens$`, iCreateTokens)
	ctx.Step(`^my account will have (\d+) tokens$`, myAccountWillHaveTokens)
}

// execute executes a shell command.
func exec(args ...string) error {
	cmd := osexec.Command(args[0], args[1:]...)
	out, err := cmd.CombinedOutput()
	switch err := err.(type) {
	case nil:
		return nil
	case *osexec.ExitError:
		return fmt.Errorf("failed to run %q:\n%v", args, string(out))
	default:
		return err
	}
}
