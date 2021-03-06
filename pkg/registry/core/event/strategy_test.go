/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package event

import (
	"reflect"
	"testing"

	"k8s.io/kubernetes/pkg/api"
	apitesting "k8s.io/kubernetes/pkg/api/testing"
	"k8s.io/kubernetes/pkg/apimachinery/registered"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/util/diff"
)

func testEvent(name string) *api.Event {
	return &api.Event{
		ObjectMeta: api.ObjectMeta{
			Name:      name,
			Namespace: "default",
		},
		InvolvedObject: api.ObjectReference{
			Namespace: "default",
		},
		Reason: "forTesting",
	}
}

func TestGetAttrs(t *testing.T) {
	eventA := &api.Event{
		ObjectMeta: api.ObjectMeta{
			Name:      "f0118",
			Namespace: "default",
		},
		InvolvedObject: api.ObjectReference{
			Kind:            "Pod",
			Name:            "foo",
			Namespace:       "baz",
			UID:             "long uid string",
			APIVersion:      registered.GroupOrDie(api.GroupName).GroupVersion.String(),
			ResourceVersion: "0",
			FieldPath:       "",
		},
		Reason: "ForTesting",
		Source: api.EventSource{Component: "test"},
		Type:   api.EventTypeNormal,
	}
	field := EventToSelectableFields(eventA)
	expect := fields.Set{
		"metadata.name":                  "f0118",
		"metadata.namespace":             "default",
		"involvedObject.kind":            "Pod",
		"involvedObject.name":            "foo",
		"involvedObject.namespace":       "baz",
		"involvedObject.uid":             "long uid string",
		"involvedObject.apiVersion":      registered.GroupOrDie(api.GroupName).GroupVersion.String(),
		"involvedObject.resourceVersion": "0",
		"involvedObject.fieldPath":       "",
		"reason":                         "ForTesting",
		"source":                         "test",
		"type":                           api.EventTypeNormal,
	}
	if e, a := expect, field; !reflect.DeepEqual(e, a) {
		t.Errorf("diff: %s", diff.ObjectDiff(e, a))
	}
}

func TestSelectableFieldLabelConversions(t *testing.T) {
	fset := EventToSelectableFields(&api.Event{})
	apitesting.TestSelectableFieldLabelConversionsOfKind(t,
		registered.GroupOrDie(api.GroupName).GroupVersion.String(),
		"Event",
		fset,
		nil,
	)
}
