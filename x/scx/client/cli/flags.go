package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagOrganizationDescription = "organization-description"
	FlagProductDescription      = "product-description"
	FlagUnitDetails             = "unit-details"
	FlagUnitComponents          = "unit-components"
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

func FlagSetUnitDetailsCreate() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagUnitDetails, "", "The details of an unit")

	return fs
}

func FlagSetUnitComponentsCreate() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagUnitComponents, "", "A comma separated list of the reference of the components of the unit")

	return fs
}
