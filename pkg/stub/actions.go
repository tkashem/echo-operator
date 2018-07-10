package stub

import (
	"reflect"

	echov1 "github.com/tkashem/echo-operator/pkg/apis/echo/v1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func deployment(m *echov1.EchoApp) *appsv1.Deployment {
	ls := echoLabels(m.Name)
	replicas := m.Spec.Size

	dep := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.Name,
			Namespace: m.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{{
						Image: m.Spec.Image,
						Name:  "echo",
						Ports: []v1.ContainerPort{{
							ContainerPort: 3000,
							Name:          "http",
						}},
					}},
				},
			},
		},
	}
	addOwnerRefToObject(dep, asOwner(m))
	return dep
}

func ensure(echo *echov1.EchoApp, deployment *appsv1.Deployment) error {
	if err := sdk.Get(deployment); err != nil {
		return err
	}

	size := echo.Spec.Size
	if size == *deployment.Spec.Replicas {
		return nil
	}

	deployment.Spec.Replicas = &size
	if err := sdk.Update(deployment); err != nil {
		return err
	}

	return nil
}

func update(echo *echov1.EchoApp) error {
	list := podList()
	selector := labels.SelectorFromSet(echoLabels(echo.Name)).String()
	listOps := &metav1.ListOptions{LabelSelector: selector}
	if err := sdk.List(echo.Namespace, list, sdk.WithListOptions(listOps)); err != nil {
		return nil
	}
	podNames := getPodNames(list.Items)

	if reflect.DeepEqual(podNames, echo.Status.Nodes) {
		return nil
	}

	echo.Status.Nodes = podNames
	if err := sdk.Update(echo); err != nil {
		return err
	}

	return nil
}

// labels returns the labels for selecting the resources belonging to the given echo CR name.
func echoLabels(name string) map[string]string {
	return map[string]string{"app": "echo", "echo_cr": name}
}

// addOwnerRefToObject appends the desired OwnerReference to the object
func addOwnerRefToObject(obj metav1.Object, ownerRef metav1.OwnerReference) {
	obj.SetOwnerReferences(append(obj.GetOwnerReferences(), ownerRef))
}

// asOwner returns an OwnerReference set as the given echo CR
func asOwner(m *echov1.EchoApp) metav1.OwnerReference {
	trueVar := true
	return metav1.OwnerReference{
		APIVersion: m.APIVersion,
		Kind:       m.Kind,
		Name:       m.Name,
		UID:        m.UID,
		Controller: &trueVar,
	}
}

// podList returns a v1.PodList object
func podList() *v1.PodList {
	return &v1.PodList{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
	}
}

// getPodNames returns the pod names of the array of pods passed in
func getPodNames(pods []v1.Pod) []string {
	var podNames []string
	for _, pod := range pods {
		podNames = append(podNames, pod.Name)
	}
	return podNames
}
