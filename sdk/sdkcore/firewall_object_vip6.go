package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
)

// JSONFirewallObjectVip contains the parameters for Create and Update API function
type JSONFirewallObjectVip6 struct {
	Name        string        `json:"name"`
	Comment     string        `json:"comment"`
	Extip       string        `json:"extip"`
	Mappedip    VIPMultValues `json:"mappedip"`
	Extintf     string        `json:"extintf,omitempty"`
	Portforward string        `json:"portforward,omitempty"`
	Protocol    string        `json:"protocol,omitempty"`
	Extport     string        `json:"extport,omitempty"`
	Mappedport  string        `json:"mappedport,omitempty"`
}

// JSONCreateFirewallObjectVipOutput contains the output results for Create API function
type JSONCreateFirewallObjectVip6Output struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateFirewallObjectVipOutput contains the output results for Update API function
// Attention: Considering scalability, the previous structure and the current structure may change differently
type JSONUpdateFirewallObjectVip6Output struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// ListFirewallObjectVip6 API operation for FortiOS lists the firewall virtual IPs
// available.
// Returns a list of firewall virtual IP value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vip chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ListFirewallObjectVip6() (output []*JSONFirewallObjectVip6, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/vip6"

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send(nil)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := io.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	slog.Debug("FOS-fortios reading response", slog.Any("body", string(body)))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	if err != nil {
		return nil, err
	}

	for _, v := range result["results"].([]interface{}) {
		mapTmp := v.(map[string]interface{})
		var vip = new(JSONFirewallObjectVip6)

		if mapTmp["name"] != nil {
			vip.Name = mapTmp["name"].(string)
		}
		if mapTmp["comment"] != nil {
			vip.Comment = mapTmp["comment"].(string)
		}
		if mapTmp["extip"] != nil {
			vip.Extip = mapTmp["extip"].(string)
		}
		if mapTmp["mappedip"] != nil {
			member := mapTmp["mappedip"].([]interface{})

			var members []VIPMultValue
			for _, v := range member {
				c := v.(map[string]interface{})

				members = append(members,
					VIPMultValue{
						Range: c["range"].(string),
					})
			}
			vip.Mappedip = members
		}
		if mapTmp["extintf"] != nil {
			vip.Extintf = mapTmp["extintf"].(string)
		}
		if mapTmp["portforward"] != nil {
			vip.Portforward = mapTmp["portforward"].(string)
		}
		if mapTmp["protocol"] != nil {
			vip.Protocol = mapTmp["protocol"].(string)
		}
		if mapTmp["extport"] != nil {
			vip.Extport = mapTmp["extport"].(string)
		}
		if mapTmp["mappedport"] != nil {
			vip.Mappedport = mapTmp["mappedport"].(string)
		}

		output = append(output, vip)
	}
	return
}

// CreateFirewallObjectVip API operation for FortiOS creates a new firewall virtual IP.
// Returns the index value of the firewall virtual IP and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vip chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallObjectVip6(params *JSONFirewallObjectVip6) (output *JSONCreateFirewallObjectVip6Output, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/vip6"
	output = &JSONCreateFirewallObjectVip6Output{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send(nil)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := io.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	log.Printf("FOS-fortios response: %v", result)

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}

		if result["mkey"] != nil {
			output.Mkey = result["mkey"].(string)
		} else {
			err = fmt.Errorf("Failed to get mkey")
			return
		}
	}

	return
}

// UpdateFirewallObjectVip API operation for FortiOS updates the specified firewall virtual IP.
// Returns the index value of the firewall virtual IP and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vip chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallObjectVip6(params *JSONFirewallObjectVip6, mkey string) (output *JSONUpdateFirewallObjectVip6Output, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/vip6"
	path += "/" + EscapeURLString(mkey)
	output = &JSONUpdateFirewallObjectVip6Output{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send(nil)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := io.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	slog.Debug("FOS-fortios response", slog.Any("body", string(body)))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}

		if result["mkey"] != nil {
			output.Mkey = result["mkey"].(string)
		}
	}

	return
}

// DeleteFirewallObjectVip API operation for FortiOS deletes the specified firewall virtual IP.
// Returns error for service API and SDK errors.
// See the firewall - vip chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallObjectVip6(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/vip6"
	path += "/" + EscapeURLString(mkey)

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send(nil)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := io.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	slog.Debug("FOS-fortios response", slog.Any("body", string(body)))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	return
}

// ReadFirewallObjectVip API operation for FortiOS gets the firewall virtual IP
// with the specified index value.
// Returns the requested firewall virtual IP value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vip chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallObjectVip6(mkey string) (output *JSONFirewallObjectVip6, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/vip6"
	path += "/" + EscapeURLString(mkey)

	output = &JSONFirewallObjectVip6{}

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send(nil)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := io.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	slog.Debug("FOS-fortios reading response", slog.Any("body", string(body)))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	if fortiAPIHttpStatus404Checking(result) == true {
		output = nil
		return
	}

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		mapTmp := (result["results"].([]interface{}))[0].(map[string]interface{})

		if mapTmp == nil {
			err = fmt.Errorf("cannot get the results from the response")
			return
		}

		if mapTmp["name"] != nil {
			output.Name = mapTmp["name"].(string)
		}
		if mapTmp["comment"] != nil {
			output.Comment = mapTmp["comment"].(string)
		}
		if mapTmp["extip"] != nil {
			output.Extip = mapTmp["extip"].(string)
		}
		if mapTmp["mappedip"] != nil {
			member := mapTmp["mappedip"].([]interface{})

			var members []VIPMultValue
			for _, v := range member {
				c := v.(map[string]interface{})

				members = append(members,
					VIPMultValue{
						Range: c["range"].(string),
					})
			}
			output.Mappedip = members
		}
		if mapTmp["extintf"] != nil {
			output.Extintf = mapTmp["extintf"].(string)
		}
		if mapTmp["portforward"] != nil {
			output.Portforward = mapTmp["portforward"].(string)
		}
		if mapTmp["protocol"] != nil {
			output.Protocol = mapTmp["protocol"].(string)
		}
		if mapTmp["extport"] != nil {
			output.Extport = mapTmp["extport"].(string)
		}
		if mapTmp["mappedport"] != nil {
			output.Mappedport = mapTmp["mappedport"].(string)
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
