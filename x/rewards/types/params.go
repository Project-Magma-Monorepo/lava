package types

import (
	"fmt"

	"cosmossdk.io/math"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMinBondedTarget                    = []byte("MinBondedTarget")
	DefaultMinBondedTarget math.LegacyDec = math.LegacyNewDecWithPrec(6, 1) // 0.6
)

var (
	KeyMaxBondedTarget                    = []byte("MaxBondedTarget")
	DefaultMaxBondedTarget math.LegacyDec = math.LegacyNewDecWithPrec(8, 1) // 0.8
)

var (
	KeyLowFactor                    = []byte("LowFactor")
	DefaultLowFactor math.LegacyDec = math.LegacyNewDecWithPrec(5, 1) // 0.5
)

var (
	KeyLeftoverBurnRate                    = []byte("LeftoverBurnRate")
	DefaultLeftOverBurnRate math.LegacyDec = math.LegacyOneDec()
)

var (
	KeyMaxRewardBoost     = []byte("MaxRewardBoost")
	DefaultMaxRewardBoost = uint64(5)
)

var (
	KeyValidatorsSubscriptionParticipation                    = []byte("ValidatorsSubscriptionParticipation")
	DefaultValidatorsSubscriptionParticipation math.LegacyDec = math.LegacyNewDecWithPrec(5, 2) // 0.05
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	minBondedTarget math.LegacyDec,
	maxBondedTarget math.LegacyDec,
	lowFactor math.LegacyDec,
	leftoverBurnRate math.LegacyDec,
	maxRewardBoost uint64,
	validatorsSubscriptionParticipation math.LegacyDec,
) Params {
	return Params{
		MinBondedTarget:                     minBondedTarget,
		MaxBondedTarget:                     maxBondedTarget,
		LowFactor:                           lowFactor,
		LeftoverBurnRate:                    leftoverBurnRate,
		MaxRewardBoost:                      maxRewardBoost,
		ValidatorsSubscriptionParticipation: validatorsSubscriptionParticipation,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMinBondedTarget,
		DefaultMaxBondedTarget,
		DefaultLowFactor,
		DefaultLeftOverBurnRate,
		DefaultMaxRewardBoost,
		DefaultValidatorsSubscriptionParticipation,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMinBondedTarget, &p.MinBondedTarget, validateDec),
		paramtypes.NewParamSetPair(KeyMaxBondedTarget, &p.MaxBondedTarget, validateDec),
		paramtypes.NewParamSetPair(KeyLowFactor, &p.LowFactor, validateDec),
		paramtypes.NewParamSetPair(KeyLeftoverBurnRate, &p.LeftoverBurnRate, validateDec),
		paramtypes.NewParamSetPair(KeyMaxRewardBoost, &p.MaxRewardBoost, validateuint64),
		paramtypes.NewParamSetPair(KeyValidatorsSubscriptionParticipation, &p.ValidatorsSubscriptionParticipation, validateDec),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateDec(p.MinBondedTarget); err != nil {
		return fmt.Errorf("invalid MinBondedTarget. Error: %s", err.Error())
	}

	if err := validateDec(p.MaxBondedTarget); err != nil {
		return fmt.Errorf("invalid MaxBondedTarget. Error: %s", err.Error())
	}

	if p.MinBondedTarget.GTE(p.MaxBondedTarget) {
		return fmt.Errorf("min_bonded_target cannot be greater or equal to max_bonded_target")
	}

	if err := validateDec(p.LowFactor); err != nil {
		return fmt.Errorf("invalid LowFactor. Error: %s", err.Error())
	}

	if err := validateDec(p.LeftoverBurnRate); err != nil {
		return fmt.Errorf("invalid LeftoverBurnRate. Error: %s", err.Error())
	}

	if p.MaxRewardBoost == 0 {
		return fmt.Errorf("MaxRewardBoost cannot be 0")
	}

	if err := validateDec(p.ValidatorsSubscriptionParticipation); err != nil {
		return fmt.Errorf("invalid ValidatorsSubscriptionParticipation. Error: %s", err.Error())
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateDec validates the Dec param is between 0 and 1
func validateDec(v interface{}) error {
	param, ok := v.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if param.GT(math.LegacyOneDec()) || param.LT(math.LegacyZeroDec()) {
		return fmt.Errorf("invalid dec parameter")
	}

	return nil
}

func validateuint64(v interface{}) error {
	_, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	return nil
}
