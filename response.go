package zabbix

import (
	"reflect"
)

type ZabbixHost map[string]interface{}
type ZabbixGraph map[string]interface{}
type ZabbixHostInterface map[string]interface{}
type ZabbixHostGroup map[string]interface{}
type ZabbixUser map[string]interface{}

type ZabbixHistoryItem struct {
	Clock  string `json:"clock"`
	Value  string `json:"value"`
	Itemid string `json:"itemid"`
}

var responseTypes map[string]reflect.Type = make(map[string]reflect.Type)

func CreateResponseTypeByActionType(actionType string) interface{} {
	return reflect.MakeSlice(responseTypes[actionType], 0, 0).Interface()
}

func init() {
	{
		var a ZabbixHost
		responseTypes[`host`] = reflect.TypeOf(a)
	}

	{
		var a ZabbixHostGroup
		responseTypes[`hostgroup`] = reflect.TypeOf(a)
	}

	{
		var a ZabbixHostInterface
		responseTypes[`hostinterface`] = reflect.TypeOf(a)
	}

	{
		var a ZabbixHistoryItem
		responseTypes[`history`] = reflect.TypeOf(a)
	}

	{
		var a ZabbixGraph
		responseTypes[`graph`] = reflect.TypeOf(a)
	}

	{
		var a ZabbixUser
		responseTypes[`user`] = reflect.TypeOf(a)
	}
}
