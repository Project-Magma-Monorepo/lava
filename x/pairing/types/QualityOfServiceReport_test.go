package types

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/require"
)

func TestQosReport(t *testing.T) {
	qos1 := &QualityOfServiceReport{
		Latency:      math.LegacyMustNewDecFromStr("1.5"),
		Availability: math.LegacyMustNewDecFromStr("1"),
		Sync:         math.LegacyMustNewDecFromStr("0.1"),
	}
	qos2 := &QualityOfServiceReport{
		Latency:      math.LegacyMustNewDecFromStr("0.2"),
		Availability: math.LegacyMustNewDecFromStr("1"),
		Sync:         math.LegacyMustNewDecFromStr("0.1"),
	}
	qos3 := &QualityOfServiceReport{
		Latency:      math.LegacyMustNewDecFromStr("0.1"),
		Availability: math.LegacyMustNewDecFromStr("1"),
		Sync:         math.LegacyMustNewDecFromStr("0.5"),
	}
	qos4 := &QualityOfServiceReport{
		Latency:      math.LegacyMustNewDecFromStr("0.1"),
		Availability: math.LegacyMustNewDecFromStr("0.5"),
		Sync:         math.LegacyMustNewDecFromStr("0.5"),
	}

	qos1Res, errQos1 := qos1.ComputeQoSExcellence()
	qos2Res, errQos2 := qos2.ComputeQoSExcellence()
	qos3Res, errQos3 := qos3.ComputeQoSExcellence()
	qos4Res, errQos4 := qos4.ComputeQoSExcellence()
	require.NoError(t, errQos1)
	require.NoError(t, errQos2)
	require.NoError(t, errQos3)
	require.NoError(t, errQos4)
	require.True(t, qos1Res.LT(qos2Res))
	require.True(t, qos1Res.LT(qos3Res))
	require.True(t, qos1Res.LT(qos4Res))

	require.True(t, qos2Res.GT(qos3Res))
	require.True(t, qos2Res.GT(qos4Res))

	require.True(t, qos4Res.LT(qos3Res))
}
