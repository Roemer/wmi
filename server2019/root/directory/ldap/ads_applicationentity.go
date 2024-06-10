// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.directory.LDAP
//
// ////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_applicationentity struct
type ads_applicationentity struct {
	*ds_top

	//
	DS_l string

	//
	DS_o []string

	//
	DS_ou []string

	//
	DS_presentationAddress Uint8Array

	//
	DS_seeAlso []string

	//
	DS_supportedApplicationContext []Uint8Array
}

func Newads_applicationentityEx1(instance *cim.WmiInstance) (newInstance *ads_applicationentity, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_applicationentity{
		ds_top: tmp,
	}
	return
}

func Newads_applicationentityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_applicationentity, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_applicationentity{
		ds_top: tmp,
	}
	return
}

// SetDS_l sets the value of DS_l for the instance
func (instance *ads_applicationentity) SetPropertyDS_l(value string) (err error) {
	return instance.SetProperty("DS_l", (value))
}

// GetDS_l gets the value of DS_l for the instance
func (instance *ads_applicationentity) GetPropertyDS_l() (value string, err error) {
	retValue, err := instance.GetProperty("DS_l")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDS_o sets the value of DS_o for the instance
func (instance *ads_applicationentity) SetPropertyDS_o(value []string) (err error) {
	return instance.SetProperty("DS_o", (value))
}

// GetDS_o gets the value of DS_o for the instance
func (instance *ads_applicationentity) GetPropertyDS_o() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_o")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetDS_ou sets the value of DS_ou for the instance
func (instance *ads_applicationentity) SetPropertyDS_ou(value []string) (err error) {
	return instance.SetProperty("DS_ou", (value))
}

// GetDS_ou gets the value of DS_ou for the instance
func (instance *ads_applicationentity) GetPropertyDS_ou() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_ou")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetDS_presentationAddress sets the value of DS_presentationAddress for the instance
func (instance *ads_applicationentity) SetPropertyDS_presentationAddress(value Uint8Array) (err error) {
	return instance.SetProperty("DS_presentationAddress", (value))
}

// GetDS_presentationAddress gets the value of DS_presentationAddress for the instance
func (instance *ads_applicationentity) GetPropertyDS_presentationAddress() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_presentationAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}

// SetDS_seeAlso sets the value of DS_seeAlso for the instance
func (instance *ads_applicationentity) SetPropertyDS_seeAlso(value []string) (err error) {
	return instance.SetProperty("DS_seeAlso", (value))
}

// GetDS_seeAlso gets the value of DS_seeAlso for the instance
func (instance *ads_applicationentity) GetPropertyDS_seeAlso() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_seeAlso")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetDS_supportedApplicationContext sets the value of DS_supportedApplicationContext for the instance
func (instance *ads_applicationentity) SetPropertyDS_supportedApplicationContext(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_supportedApplicationContext", (value))
}

// GetDS_supportedApplicationContext gets the value of DS_supportedApplicationContext for the instance
func (instance *ads_applicationentity) GetPropertyDS_supportedApplicationContext() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_supportedApplicationContext")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Uint8Array)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Uint8Array(valuetmp))
	}

	return
}
