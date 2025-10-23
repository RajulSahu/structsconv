package rules

import (
	"github.com/RajulSahu/structsconv"
	"github.com/RajulSahu/structsconv/example/domain"
	"github.com/RajulSahu/structsconv/example/dto"
)

// GetAddressDtoToAddressDomainRules dto.AddressDto{} -> domain.Address{}
func (d *RulesDefinitions) GetAddressDtoToAddressDomainRules() *structsconv.RulesDefinition {
	var rules = structsconv.RulesSet{}

	// Name association
	rules["Zip"] = "ZipCode"

	// Constant
	rules["Country"] = func() string { return "CL" }

	return &structsconv.RulesDefinition{
		Rules:  rules,
		Source: dto.AddressDto{},
		Target: domain.Address{},
	}
}
