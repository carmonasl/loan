package keeper_test

import (
	"testing"

	testkeeper "github.com/carmonasl/loan/testutil/keeper"
	"github.com/carmonasl/loan/x/loan/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LoanKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
