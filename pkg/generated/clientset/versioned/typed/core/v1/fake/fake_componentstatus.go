// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	corev1 "github.com/onexstack/onex/pkg/generated/applyconfigurations/core/v1"
	typedcorev1 "github.com/onexstack/onex/pkg/generated/clientset/versioned/typed/core/v1"
	v1 "k8s.io/api/core/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeComponentStatuses implements ComponentStatusInterface
type fakeComponentStatuses struct {
	*gentype.FakeClientWithListAndApply[*v1.ComponentStatus, *v1.ComponentStatusList, *corev1.ComponentStatusApplyConfiguration]
	Fake *FakeCoreV1
}

func newFakeComponentStatuses(fake *FakeCoreV1) typedcorev1.ComponentStatusInterface {
	return &fakeComponentStatuses{
		gentype.NewFakeClientWithListAndApply[*v1.ComponentStatus, *v1.ComponentStatusList, *corev1.ComponentStatusApplyConfiguration](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("componentstatuses"),
			v1.SchemeGroupVersion.WithKind("ComponentStatus"),
			func() *v1.ComponentStatus { return &v1.ComponentStatus{} },
			func() *v1.ComponentStatusList { return &v1.ComponentStatusList{} },
			func(dst, src *v1.ComponentStatusList) { dst.ListMeta = src.ListMeta },
			func(list *v1.ComponentStatusList) []*v1.ComponentStatus { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.ComponentStatusList, items []*v1.ComponentStatus) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
