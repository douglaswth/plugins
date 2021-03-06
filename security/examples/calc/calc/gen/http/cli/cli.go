// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// calc HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/plugins/security/examples/calc/calc/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
	calcsvcc "goa.design/plugins/security/examples/calc/calc/gen/http/calc/client"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `calc (login|add)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` calc login --body '{
      "password": "password",
      "user": "username"
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		calcFlags = flag.NewFlagSet("calc", flag.ContinueOnError)

		calcLoginFlags    = flag.NewFlagSet("login", flag.ExitOnError)
		calcLoginBodyFlag = calcLoginFlags.String("body", "REQUIRED", "")

		calcAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		calcAddBodyFlag = calcAddFlags.String("body", "REQUIRED", "")
		calcAddAFlag    = calcAddFlags.String("a", "REQUIRED", "Left operand")
		calcAddBFlag    = calcAddFlags.String("b", "REQUIRED", "Right operand")
	)
	calcFlags.Usage = calcUsage
	calcLoginFlags.Usage = calcLoginUsage
	calcAddFlags.Usage = calcAddUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if len(os.Args) < flag.NFlag()+3 {
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = os.Args[1+flag.NFlag()]
		switch svcn {
		case "calc":
			svcf = calcFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(os.Args[2+flag.NFlag():]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = os.Args[2+flag.NFlag()+svcf.NFlag()]
		switch svcn {
		case "calc":
			switch epn {
			case "login":
				epf = calcLoginFlags

			case "add":
				epf = calcAddFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if len(os.Args) > 2+flag.NFlag()+svcf.NFlag() {
		if err := epf.Parse(os.Args[3+flag.NFlag()+svcf.NFlag():]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "calc":
			c := calcsvcc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "login":
				endpoint = c.Login()
				data, err = calcsvcc.BuildLoginPayload(*calcLoginBodyFlag)
			case "add":
				endpoint = c.Add()
				data, err = calcsvcc.BuildAddPayload(*calcAddBodyFlag, *calcAddAFlag, *calcAddBFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// calcUsage displays the usage of the calc command and its subcommands.
func calcUsage() {
	fmt.Fprintf(os.Stderr, `The calc service exposes public endpoints that require valid authorization credentials.
Usage:
    %s [globalflags] calc COMMAND [flags]

COMMAND:
    login: Creates a valid JWT
    add: Add adds up the two integer parameters and returns the results. This endpoint is secured with the JWT scheme

Additional help:
    %s calc COMMAND --help
`, os.Args[0], os.Args[0])
}
func calcLoginUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] calc login -body JSON

Creates a valid JWT
    -body JSON: 

Example:
    `+os.Args[0]+` calc login --body '{
      "password": "password",
      "user": "username"
   }'
`, os.Args[0])
}

func calcAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] calc add -body JSON -a INT -b INT

Add adds up the two integer parameters and returns the results. This endpoint is secured with the JWT scheme
    -body JSON: 
    -a INT: Left operand
    -b INT: Right operand

Example:
    `+os.Args[0]+` calc add --body '{
      "token": "Ullam eius odio minima ipsam voluptatem mollitia."
   }' --a 1 --b 2
`, os.Args[0])
}
