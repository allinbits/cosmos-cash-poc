package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/stretchr/testify/assert"
)

var opts = godog.Options{Output: colors.Colored(os.Stdout)}

func iAmAnActor(arg1 string, arg2 string) error {
	output, err := exec.Command("poacli", "keys", "show", arg2, "-a").Output()
	if len(output) == 0 {
		return fmt.Errorf("key is unavailable")
	}
	if err != nil {
		return err
	}
	return nil
}

func iCreateAnAccountWithName(arg1 string) error {
	output, err := exec.Command("bash", "-c",
		"echo 'y' | poacli keys add"+arg1,
	).Output()
	if len(output) > 0 {
		return nil
	}
	if err != nil {
		return fmt.Errorf("error generating issuer key")
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
	_, err = exec.Command("poacli", "tx", "issuer", "create-issuer",
		arg1, issuerAddress, arg3, arg2,
		"--trust-node", "-y", "--from", "validator", "--chain-id", "cash", "--home", "../../../build/.poad",
	).Output()
	if err != nil {
		return err
	}
	return nil
}

func myAccountHasTokens(arg1 string, arg2 string) error {
	output, err := exec.Command("poacli", "keys", "show", arg1, "-a").Output()
	if err != nil {
		return err
	}
	s := string(output)
	address := s[:len(s)-1]
	output, err = exec.Command("bash", "-c",
		"poacli query auth account "+address+" --output json | jq '.value.coins[0].amount'",
	).Output()
	if err != nil {
		return fmt.Errorf("error querying account")
	}

	s = string(output)
	amount := s[1 : len(s)-2]
	var t assert.TestingT
	if !assert.Equal(t, amount, arg2,
		"Expected account balance to be %s, but there is %s", arg2, amount,
	) {
		return fmt.Errorf("error in account balance")
	}

	return nil
}

func iCreateTokens(arg1 string, arg2 string) error {
	_, err := exec.Command("bash", "-c",
		"poacli tx issuer mint-token "+arg2+" "+arg1+" --trust-node --from token-issuer --chain-id cash -y",
	).Output()

	if err != nil {
		return fmt.Errorf("error minting tokens")
	}
	return nil
}

func iBurnTokens(arg1 string, arg2 string) error {
	_, err := exec.Command("bash", "-c",
		"poacli tx issuer burn-token "+arg2+" "+arg1+" --trust-node --from token-issuer --chain-id cash -y",
	).Output()

	if err != nil {
		return fmt.Errorf("error burning tokens")
	}
	return nil
}

func iWaitSecondsForMyTransactionToBeCompleted() error {
	time.Sleep(5 * time.Second)
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I am an "([^"]*)" called "([^"]*)"$`, iAmAnActor)
	ctx.Step(`^I create an account with name "([^"]*)"$`, iCreateAnAccountWithName)
	ctx.Step(`^I create an issuer "([^"]*)" with "([^"]*)" "([^"]*)"$`, iCreateAnIssuerWithTokens)
	ctx.Step(`^my account "([^"]*)" has "([^"]*)" tokens$`, myAccountHasTokens)
	ctx.Step(`^I create "([^"]*)" "([^"]*)" tokens$`, iCreateTokens)
	ctx.Step(`^I wait for my transaction to be completed$`, iWaitSecondsForMyTransactionToBeCompleted)
	ctx.Step(`^I burn "([^"]*)" "([^"]*)" tokens$`, iBurnTokens)
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
