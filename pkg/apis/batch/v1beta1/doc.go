// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package
// +k8s:protobuf-gen=package
// +k8s:conversion-gen=github.com/onexstack/onex/pkg/apis/batch
// +k8s:conversion-gen-external-types=github.com/onexstack/onex/pkg/apis/batch/v1beta1
// +k8s:defaulter-gen=TypeMeta
// +k8s:defaulter-gen-input=github.com/onexstack/onex/pkg/apis/batch/v1beta1
// +groupName=batch.onex.io

// Package v1beta1 is the v1beta1 version of the API.
package v1beta1 // import "github.com/onexstack/onex/pkg/apis/batch/v1beta1"
