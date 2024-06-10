// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.RSOP.NSADDAA4BD_2BE8_4395_88D3_4679F8FF4561
//
// ////////////////////////////////////////////
package nsaddaa4bd_2be8_4395_88d3_4679f8ff4561

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// __InstanceOperationEvent struct
type __InstanceOperationEvent struct {
	*__Event

	//
	TargetInstance interface{}
}

func New__InstanceOperationEventEx1(instance *cim.WmiInstance) (newInstance *__InstanceOperationEvent, err error) {
	tmp, err := New__EventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__InstanceOperationEvent{
		__Event: tmp,
	}
	return
}

func New__InstanceOperationEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__InstanceOperationEvent, err error) {
	tmp, err := New__EventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__InstanceOperationEvent{
		__Event: tmp,
	}
	return
}

// SetTargetInstance sets the value of TargetInstance for the instance
func (instance *__InstanceOperationEvent) SetPropertyTargetInstance(value interface{}) (err error) {
	return instance.SetProperty("TargetInstance", (value))
}

// GetTargetInstance gets the value of TargetInstance for the instance
func (instance *__InstanceOperationEvent) GetPropertyTargetInstance() (value interface{}, err error) {
	retValue, err := instance.GetProperty("TargetInstance")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}
