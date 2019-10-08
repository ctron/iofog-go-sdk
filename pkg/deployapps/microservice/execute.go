/*
 *  *******************************************************************************
 *  * Copyright (c) 2019 Edgeworx, Inc.
 *  *
 *  * This program and the accompanying materials are made available under the
 *  * terms of the Eclipse Public License v. 2.0 which is available at
 *  * http://www.eclipse.org/legal/epl-2.0
 *  *
 *  * SPDX-License-Identifier: EPL-2.0
 *  *******************************************************************************
 *
 */

package deploymicroservice

import (
	types "github.com/eclipse-iofog/iofog-go-sdk/pkg/deployapps"
)

func Execute(controller types.IofogController, microservice types.Microservice) error {
	exe := newRemoteExecutor(controller, microservice)
	return exe.Execute()
}