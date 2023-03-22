package services

import (
	"bytes"
	"encoding/json"
	discordstructs "klineService/services/discordStructs"
	"log"
	"net/http"
)

// curl -X POST https://discord.com/api/webhooks/1057831941514207282/AlQM4bvDA3DX1nFTiYwzMcZ0hKfgjWwlBndZd1g-FUblfglaLSFJmK9kZudQYiNJi2T9 -d '{"content":"testando"}' -H "Content-Type: application/json
// curl -X POST https://discord.com/api/webhooks/1059954298458488893/4HA2jqowyFPmMuQJx1TnDWK8P2GrqUblpjuggX4pQt-dDxcXo_DDttgtQko-vdDa-rdo -d '{"content":"testando"}' -H "Content-Type: application/json
// curl -X POST https://discord.com/api/webhooks/1059973551874129960/lNgjmyt3_Mt565GY6UQaKHquW7EvdzOVnRMKy9_HtTDSIf-oZoZNma70z9Jy67oOrt_X -d '{"content":"testando"}' -H "Content-Type: application/json
// curl -X POST https://discord.com/api/webhooks/1060354322799546448/3Tfq9JvqYZ-DDHCgsMWX9BAkA1stSkcXayD9sG_mXhB7zfSCRz-OtGhyJe5CmtctuMxI -d '{"content":"testando"}' -H "Content-Type: application/json
// curl -X POST https://discord.com/api/webhooks/1062894822710575105/gqgTByNIkh66frvpEAkKQDRl5Sql9qQRCL192HxXsJ_OZhmAmBhITPKLHjj5y9AG5yw- -d '{"content":"testando"}' -H "Content-Type: application/json
// curl -X POST https://discord.com/api/webhooks/1063592357712376019/kJH1-BDZ3kAsmk5b1IpYrJsWFlPQIHr-bu_VRCN9i31qRi8S8Safhx02DF3jeFt4gLba -d '{"content":"testando"}' -H "Content-Type: application/json
// curl -X POST https://discord.com/api/webhooks/1065283568395362365/qN_z6acUQQU5-WxgzxmW7yiTVJULHzXPX6G8HH4KDggD_B_as4EJtg8zmFg-kYlg5NQp -d '{"content":"testando"}' -H "Content-Type: application/json

const discordBaseURL = "https://discord.com/api/webhooks"
const OrderNotifyURL = "/1057831941514207282/AlQM4bvDA3DX1nFTiYwzMcZ0hKfgjWwlBndZd1g-FUblfglaLSFJmK9kZudQYiNJi2T9"
const HealthStatusURL = "/1059954298458488893/4HA2jqowyFPmMuQJx1TnDWK8P2GrqUblpjuggX4pQt-dDxcXo_DDttgtQko-vdDa-rdo"
const OpenOperationURL = "/1059973551874129960/lNgjmyt3_Mt565GY6UQaKHquW7EvdzOVnRMKy9_HtTDSIf-oZoZNma70z9Jy67oOrt_X"
const DebuggerURL = "/1060354322799546448/3Tfq9JvqYZ-DDHCgsMWX9BAkA1stSkcXayD9sG_mXhB7zfSCRz-OtGhyJe5CmtctuMxI"
const NewOperationURL = "/1062894822710575105/gqgTByNIkh66frvpEAkKQDRl5Sql9qQRCL192HxXsJ_OZhmAmBhITPKLHjj5y9AG5yw-"
const ProfitURL = "/1063592357712376019/kJH1-BDZ3kAsmk5b1IpYrJsWFlPQIHr-bu_VRCN9i31qRi8S8Safhx02DF3jeFt4gLba"
const ErrorsURL = "/1065283568395362365/qN_z6acUQQU5-WxgzxmW7yiTVJULHzXPX6G8HH4KDggD_B_as4EJtg8zmFg-kYlg5NQp"

func (s *Discord) SendNotification(params *discordstructs.Notification) {
	client := &http.Client{}

	jsonstr, _ := json.Marshal(params)

	var req *http.Request
	var err error
	switch params.Channel {
	case "Order":
		req, err = http.NewRequest("POST", discordBaseURL+OrderNotifyURL, bytes.NewBuffer(jsonstr))
	case "Health":
		req, err = http.NewRequest("POST", discordBaseURL+HealthStatusURL, bytes.NewBuffer(jsonstr))
	case "OpenOp":
		req, err = http.NewRequest("POST", discordBaseURL+OpenOperationURL, bytes.NewBuffer(jsonstr))
	case "Debug":
		req, err = http.NewRequest("POST", discordBaseURL+DebuggerURL, bytes.NewBuffer(jsonstr))
	case "NewOrder":
		req, err = http.NewRequest("POST", discordBaseURL+NewOperationURL, bytes.NewBuffer(jsonstr))
	case "Profit":
		req, err = http.NewRequest("POST", discordBaseURL+ProfitURL, bytes.NewBuffer(jsonstr))
	case "Errors":
		req, err = http.NewRequest("POST", discordBaseURL+ErrorsURL, bytes.NewBuffer(jsonstr))
	}

	if err != nil {
		log.Println("DSN 01: ", err)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Println("Discord req exec: ", err)
	}

	defer resp.Body.Close()
}
