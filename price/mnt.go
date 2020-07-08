package price

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	cfg "github.com/node-a-team/terra-oracle/config"
)

func (ps *PriceService) mntToKrw(logger log.Logger) {


	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					logger.Error("Unknown error", r)
				}

				time.Sleep(cfg.Config.Options.Interval * time.Second)
			}()

//			resp, err := http.Get("http://www.apilayer.net/api/live?access_key=f4f5c16e99a0f32baeab5be8ced1cd39")
			resp, err := http.Get(cfg.Config.APIs.MNT.Dunamu)
			if err != nil {
				logger.Error("Fail to fetch from freeforexapi", err.Error())
				return
			}
			defer func() {
				resp.Body.Close()
			}()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logger.Error("Fail to read body", err.Error())
				return
			}

			re, _ := regexp.Compile("\"basePrice\":[0-9.]+")
			str := re.FindString(string(body))
			re, _ = regexp.Compile("[0-9.]+")
			price := re.FindString(str)

			logger.Info(fmt.Sprintf("Recent mnt/krw: %s", price))

			decAmount, err := sdk.NewDecFromStr(price)
			if err != nil {
				logger.Error("Fail to parse price to Dec", err.Error())
				return
			}
			ps.SetPrice("mnt/krw", sdk.NewDecCoinFromDec("krw", decAmount))
		}()
	}
}

