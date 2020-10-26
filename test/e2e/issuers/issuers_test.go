package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

var opts = godog.Options{Output: colors.Colored(os.Stdout)}

func iAmAnAdmin(arg1 string) error {
	output, err := exec.Command("poacli", "keys", "show", arg1, "-a").Output()
	if len(output) == 0 {
		return fmt.Errorf("admin key is unavailable")
	}
	if err != nil {
		return err
	}
	return nil
}

func iCreateAnIssuerWithTokens(arg1 string, arg2 string, arg3 string) error {
	output, err := exec.Command("poacli", "keys", "show", arg1, "-a").Output()
	if err != nil {
		return err
	}
	s := string(output)
	issuerAddress := s[:len(s)-1]
	txOutput, err := exec.Command("poacli", "tx", "issuer", "create-issuer", arg1, issuerAddress, arg3, arg2, "--trust-node", "-y", "--from", "validator", "--chain-id", "cash", "--home", "../../../build/.poad").Output()
	fmt.Println(string(txOutput))
	if err != nil {
		return err
	}
	return nil
}

func iCreateAndIssuerAccountWithName(arg1 string) error {
	_, err := exec.Command("poacli", "keys", "add", arg1).Output()
	if err != nil {
		return err
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I am an admin called "([^"]*)"$`, iAmAnAdmin)
	ctx.Step(`^I create an account with name "([^"]*)"$`, iCreateAndIssuerAccountWithName)
	ctx.Step(`^I create an issuer "([^"]*)" with "([^"]*)" "([^"]*)"$`, iCreateAnIssuerWithTokens)

	// TODO: add in steps
	//	ctx.Step(`^I am an issuer$`, iAmAnIssuer)
	//	ctx.Step(`^my account has (\d+) tokens$`, myAccountHasTokens)
	//	ctx.Step(`^I create (\d+) tokens$`, iCreateTokens)
	//	ctx.Step(`^my account will have (\d+) tokens$`, myAccountWillHaveTokens)
}

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

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opts.Paths = flag.Args()

	status := godog.TestSuite{
		Name:                "godogs",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
