package types

// GenesisState - all scx state that must be provided at genesis
type GenesisState struct {
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState() GenesisState {
	return GenesisState{}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

// ValidateGenesis validates the scx genesis parameters
func ValidateGenesis(data GenesisState) error {
	return nil
}
