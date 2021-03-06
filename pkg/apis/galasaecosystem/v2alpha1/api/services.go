/*
 * Copyright contributors to the Galasa Project
 */
package api

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (c *Api) getExposedService() *corev1.Service {
	s := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:            c.Name + "-external-service",
			Namespace:       c.Namespace,
			OwnerReferences: c.Owner,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceType(corev1.ServiceTypeNodePort),
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					TargetPort: intstr.FromInt(8080),
					Port:       8080,
				},
			},
			Selector: map[string]string{
				"app": c.Name,
			},
		},
	}
	s.SetGroupVersionKind(schema.FromAPIVersionAndKind("v1", "Service"))
	return s
}

func (c *Api) getInternalService() *corev1.Service {
	s := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				"service.alpha.kubernetes.io/tolerate-unready-endpoints": "true",
			},
			Name:      c.Name + "-internal-service",
			Namespace: c.Namespace,
			Labels: map[string]string{
				"app": c.Name,
			},
			OwnerReferences: c.Owner,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					TargetPort: intstr.FromInt(8080),
					Port:       8080,
				},
				{
					Name:       "metrics",
					TargetPort: intstr.FromInt(9010),
					Port:       9010,
				},
				{
					Name:       "health",
					TargetPort: intstr.FromInt(9011),
					Port:       9011,
				},
			},
			Selector: map[string]string{
				"app": c.Name,
			},
		},
	}
	s.SetGroupVersionKind(schema.FromAPIVersionAndKind("v1", "Service"))
	return s
}
