package controllers_test

import (
	"context"
	"github.com/numary/ledger/pkg/api"
	"github.com/numary/ledger/pkg/api/controllers"
	"github.com/numary/ledger/pkg/api/internal"
	"github.com/numary/ledger/pkg/storage"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"net/http"
	"testing"
)

func TestGetInfo(t *testing.T) {
	internal.RunTest(t, fx.Invoke(func(lc fx.Lifecycle, h *api.API, driver storage.Driver) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				rsp := internal.GetInfo(h)
				assert.Equal(t, http.StatusOK, rsp.Result().StatusCode)

				info := controllers.ConfigInfo{}
				internal.DecodeSingleResponse(t, rsp.Body, &info)
				info.Config.LedgerStorage.Ledgers = []string{}
				assert.EqualValues(t, controllers.ConfigInfo{
					Server:  "numary-ledger",
					Version: "latest",
					Config: &controllers.Config{
						LedgerStorage: &controllers.LedgerStorage{
							Driver:  driver.Name(),
							Ledgers: []string{},
						},
					},
				}, info)
				return nil
			},
		})
	}))
}
