package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagOrganizationDescription = "organization-description"
	FlagProductDescription      = "product-description"
)

func FlagSetOrganizationDescriptionCreate() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagOrganizationDescription, "", "The description of the organization")

	return fs
}

func FlagSetProductDescriptionCreate() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagProductDescription, "", "The description of the product")

	return fs
}
