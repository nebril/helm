/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package environment

import (
	"bytes"
	"testing"
	"time"

	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/apimachinery/announced"
	//appsbeta1 "k8s.io/kubernetes/pkg/apis/apps/v1beta1"
	//batchv1 "k8s.io/kubernetes/pkg/apis/batch/v1"
	//"k8s.io/kubernetes/pkg/apis/extensions/v1beta1"
	"k8s.io/kubernetes/pkg/runtime"
	//"github.com/kr/pretty"
	//"k8s.io/kubernetes/pkg/api"
)

const (
	GroupName string = "appcontroller.k8s"
	Version   string = "v1alpha1"
)

var (
	SchemeGroupVersion = unversioned.GroupVersion{Group: GroupName, Version: Version}
	SchemeBuilder      = runtime.NewSchemeBuilder(addKnownTypes)
)

/*
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemeGroupVersion,
		&Dependency{},
	)
	return nil
}
*/

func kra() {
	if err := announced.NewGroupMetaFactory(
		&announced.GroupMetaFactoryArgs{
			GroupName:                  GroupName,
			VersionPreferenceOrder:     []string{SchemeGroupVersion.Version},
			AddInternalObjectsToScheme: SchemeBuilder.AddToScheme,
		},
		announced.VersionToSchemeFunc{
			SchemeGroupVersion.Version: SchemeBuilder.AddToScheme,
		},
	).Announce().RegisterAndEnable(); err != nil {
		panic(err)
	}
}

/*
type ResourceDefinition struct {
	unversioned.TypeMeta `json:",inline"`

	// Standard object metadata
	api.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Meta map[string]string `json:"meta,omitempty"`

	//TODO: add other object types
	Pod                   *v1.Pod                   `json:"pod,omitempty"`
	Job                   *batchv1.Job              `json:"job,omitempty"`
	Service               *v1.Service               `json:"service,omitempty"`
	ReplicaSet            *v1beta1.ReplicaSet       `json:"replicaset,omitempty"`
	StatefulSet           *appsbeta1.StatefulSet    `json:"statefulset,omitempty"`
	DaemonSet             *v1beta1.DaemonSet        `json:"daemonset,omitempty"`
	ConfigMap             *v1.ConfigMap             `json:"configmap,omitempty"`
	Secret                *v1.Secret                `json:"secret,omitempty"`
	Deployment            *v1beta1.Deployment       `json:"deployment, omitempty"`
	PersistentVolumeClaim *v1.PersistentVolumeClaim `json:"persistentvolumeclaim, omitempty"`
}


type Dependency struct {
	unversioned.TypeMeta `json:",inline"`

	// Standard object metadata
	api.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Parent string            `json:"parent"`
	Child  string            `json:"child"`
	Meta   map[string]string `json:"meta,omitempty"`
}
*/

func TestKubeClientCreateThirdPartyResource(t *testing.T) {
	env := New()

	manifests := map[string]string{
		"foo": `apiVersion: appcontroller.k8s/v1alpha1
kind: Dependency
metadata: {"name": "dep1"}
parent: pod/pod1
child: pod/pod2
`,
	}

	//pretty.Println(api.Scheme)

	b := bytes.NewBuffer(nil)
	for _, content := range manifests {
		//b.WriteString("\n---\n")
		b.WriteString(content)
	}

	now := time.Now()
	if err := env.KubeClient.Create("sharry-bobbins", b); err != nil {
		t.Errorf("Kubeclient failed: %s at %v", err, now)
	}
}
