// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NS112B57B2_B3D2_44D2_80CD_FAA49D494CC6
//////////////////////////////////////////////
package ns112b57b2_b3d2_44d2_80cd_faa49d494cc6
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_InstCreation struct
type CIM_InstCreation struct { 
	*CIM_InstIndication
}

	func NewCIM_InstCreationEx1(instance *cim.WmiInstance) (newInstance *CIM_InstCreation, err error) {tmp, err := NewCIM_InstIndicationEx1(instance)
		
	if err != nil { return }
	newInstance = &CIM_InstCreation {
	CIM_InstIndication: tmp,
	}
	return
	}
	

	func NewCIM_InstCreationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CIM_InstCreation, err error) {tmp, err := NewCIM_InstIndicationEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CIM_InstCreation {
	CIM_InstIndication: tmp,
	}
	return
	}
	

