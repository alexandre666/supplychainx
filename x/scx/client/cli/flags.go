package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagOrganizationDescription = "organization-description"
)

func FlagSetOrganizationDescriptionCreate() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagOrganizationDescription, "", "The description of the organization")

	return fs
}
