package zabbix

import (
	"reflect"
)

type ZabbixHost map[string]interface{}
type ZabbixHosts []ZabbixHost

type ZabbixGraph map[string]interface{}
type ZabbixGraphs []ZabbixGraph

type ZabbixHostInterface map[string]interface{}
type ZabbixHostInterfaces []ZabbixHostInterface

type ZabbixHostGroup map[string]interface{}
type ZabbixHostGroups []ZabbixHostGroup

type ZabbixUser map[string]interface{}
type ZabbixUsers []ZabbixUser

type ZabbixHistoryItem struct {
	Clock  string `json:"clock"`
	Value  string `json:"value"`
	Itemid string `json:"itemid"`
}
type ZabbixHistoryItems []ZabbixHistoryItem

var responseTypes map[string]reflect.Type = make(map[string]reflect.Type)

func CreateResponseTypeByActionType(actionType string) interface{} {
	return reflect.MakeSlice(responseTypes[actionType], 0, 0).Interface()
}

func init() {
	{
		var a ZabbixHosts
		responseTypes[`host`] = reflect.TypeOf(a)
	}

	{
		var a ZabbixHostGroups
		responseTypes[`hostgroup`] = reflect.TypeOf(a)
	}

	{
		var a ZabbixHostInterfaces
		responseTypes[`hostinterface`] = reflect.TypeOf(a)
	}

	{
		var a ZabbixHistoryItems
		responseTypes[`history`] = reflect.TypeOf(a)
	}

	{
		var a ZabbixGraphs
		responseTypes[`graph`] = reflect.TypeOf(a)
	}

	{
		var a ZabbixUsers
		responseTypes[`user`] = reflect.TypeOf(a)
	}
}
