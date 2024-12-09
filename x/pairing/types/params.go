package types

import (
	"fmt"

	"cosmossdk.io/math"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyEpochBlocksOverlap = []byte("EpochBlocksOverlap")
	// TODO: Determine the default value
	DefaultEpochBlocksOverlap uint64 = 4
)

var (
	KeyQoSWeight                    = []byte("QoSWeight")
	DefaultQoSWeight math.LegacyDec = math.LegacyNewDecWithPrec(5, 1) // 0.5
)

var (
	KeyRecommendedEpochNumToCollectPayment            = []byte("RecommendedEpochNumToCollectPayment") // the recommended amount of max epochs that a provider should wait before collecting its payment (if he'll collect later, there's a higher chance to get punished)
	DefaultRecommendedEpochNumToCollectPayment uint64 = 3
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	epochBlocksOverlap uint64,
	qoSWeight math.LegacyDec,
	recommendedEpochNumToCollectPayment uint64,
) Params {
	return Params{
		EpochBlocksOverlap:                  epochBlocksOverlap,
		QoSWeight:                           qoSWeight,
		RecommendedEpochNumToCollectPayment: recommendedEpochNumToCollectPayment,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultEpochBlocksOverlap,
		DefaultQoSWeight,
		DefaultRecommendedEpochNumToCollectPayment,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyEpochBlocksOverlap, &p.EpochBlocksOverlap, validateEpochBlocksOverlap),
		paramtypes.NewParamSetPair(KeyQoSWeight, &p.QoSWeight, validateQoSWeight),
		paramtypes.NewParamSetPair(KeyRecommendedEpochNumToCollectPayment, &p.RecommendedEpochNumToCollectPayment, validateRecommendedEpochNumToCollectPayment),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateEpochBlocksOverlap(p.EpochBlocksOverlap); err != nil {
		return err
	}

	if err := validateRecommendedEpochNumToCollectPayment(p.RecommendedEpochNumToCollectPayment); err != nil {
		return err
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateEpochBlocksOverlap validates the EpochBlocksOverlap param
func validateEpochBlocksOverlap(v interface{}) error {
	epochBlocksOverlap, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = epochBlocksOverlap

	return nil
}

// validateDataReliabilityReward validates the param
func validateQoSWeight(v interface{}) error {
	QoSWeight, ok := v.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if QoSWeight.GT(math.LegacyOneDec()) || QoSWeight.LT(math.LegacyZeroDec()) {
		return fmt.Errorf("invalid parameter QoSWeight")
	}

	return nil
}

// validateRecommendedEpochNumToCollectPayment validates the RecommendedEpochNumToCollectPayment param
func validateRecommendedEpochNumToCollectPayment(v interface{}) error {
	recommendedEpochNumToCollectPayment, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = recommendedEpochNumToCollectPayment

	return nil
}
