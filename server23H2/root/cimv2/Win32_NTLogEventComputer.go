// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_NTLogEventComputer struct
type Win32_NTLogEventComputer struct {
	*cim.WmiInstance

	//
	Computer Win32_ComputerSystem

	//
	Record Win32_NTLogEvent
}

func NewWin32_NTLogEventComputerEx1(instance *cim.WmiInstance) (newInstance *Win32_NTLogEventComputer, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Win32_NTLogEventComputer{
		WmiInstance: tmp,
	}
	return
}

func NewWin32_NTLogEventComputerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_NTLogEventComputer, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_NTLogEventComputer{
		WmiInstance: tmp,
	}
	return
}

// SetComputer sets the value of Computer for the instance
func (instance *Win32_NTLogEventComputer) SetPropertyComputer(value Win32_ComputerSystem) (err error) {
	return instance.SetProperty("Computer", (value))
}

// GetComputer gets the value of Computer for the instance
func (instance *Win32_NTLogEventComputer) GetPropertyComputer() (value Win32_ComputerSystem, err error) {
	retValue, err := instance.GetProperty("Computer")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Win32_ComputerSystem)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Win32_ComputerSystem is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Win32_ComputerSystem(valuetmp)

	return
}

// SetRecord sets the value of Record for the instance
func (instance *Win32_NTLogEventComputer) SetPropertyRecord(value Win32_NTLogEvent) (err error) {
	return instance.SetProperty("Record", (value))
}

// GetRecord gets the value of Record for the instance
func (instance *Win32_NTLogEventComputer) GetPropertyRecord() (value Win32_NTLogEvent, err error) {
	retValue, err := instance.GetProperty("Record")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Win32_NTLogEvent)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Win32_NTLogEvent is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Win32_NTLogEvent(valuetmp)

	return
}
