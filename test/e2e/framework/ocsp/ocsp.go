/*
Copyright 2020 The Kubernetes Authors.

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

package ocsp

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/onsi/ginkgo"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/ingress-nginx/test/e2e/framework"
)

const (
	// ocsp response code are documented in RFC5280
	// https://tools.ietf.org/html/rfc5280#section-5.3.1
	ocspReasonWithdrawn int    = 9
	configTemplate      string = `
	{
		"signing": {
			"default": {
				"ocsp_url": "http://ocspserve.%v.svc.cluster.local",
				"expiry": "219000h",
				"usages": [
					"signing",
					"key encipherment",
					"client auth",
					"ocsp signing" 
				]
			},
			"profiles": {
				"ocsp": {
					"usages": ["digital signature", "ocsp signing"],
					"expiry": "8760h"
				},
				"intermediate": {
					"usages": ["cert sign", "crl sign"],
					"expiry": "219000h",
					"ca_constraint": {
						"is_ca": true
					}
				},
				"server": {
					"usages": ["signing", "key encipherment", "server auth"],
					"expiry": "8760h"
				},
				"client": {
					"usages": ["signing", "key encipherment", "client auth"],
					"expiry": "8760h"
				}
			}
		}
	}
	`
)

// OCSP Framework support common setup operations for OCSP and OSCP Responsder
type OcspFramework struct {
	framework *framework.Framework
}

func NewFramework(f *framework.Framework) *OcspFramework {
	return &OcspFramework{framework: f}
}

// CreateIngressOcspSecret creates or updates a Secret containing a PKI
// certificate chain and TLS certificates for a Ingress.
func (o *OcspFramework) CreateIngressOcspSecret(host string, secretName string, namespace string) error {
	err := o.prepareCertificates(namespace)
	if err != nil {
		return err
	}

	caCert, err := ioutil.ReadFile("ca.pem")
	if err != nil {
		return err
	}

	leafCert, err := ioutil.ReadFile("leaf.pem")
	if err != nil {
		return err
	}

	leafKey, err := ioutil.ReadFile("leaf-key.pem")
	if err != nil {
		return err
	}

	var pemCertBuffer bytes.Buffer
	pemCertBuffer.Write(leafCert)
	pemCertBuffer.Write([]byte("\n"))
	pemCertBuffer.Write(caCert)

	newSecret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: secretName,
		},
		Data: map[string][]byte{
			corev1.TLSCertKey:       pemCertBuffer.Bytes(),
			corev1.TLSPrivateKeyKey: leafKey,
			"ca.crt":                caCert,
		},
	}

	var apierr error
	curSecret, err := o.framework.KubeClientSet.CoreV1().Secrets(namespace).Get(context.TODO(), secretName, metav1.GetOptions{})
	if err == nil && curSecret != nil {
		curSecret.Data = newSecret.Data
		_, apierr = o.framework.KubeClientSet.CoreV1().Secrets(namespace).Update(context.TODO(), curSecret, metav1.UpdateOptions{})
	} else {
		_, apierr = o.framework.KubeClientSet.CoreV1().Secrets(namespace).Create(context.TODO(), newSecret, metav1.CreateOptions{})
	}
	if apierr != nil {
		return apierr
	}

	return nil
}

// prepareCertificates generates a PKI certificate chain and stores
// valid OCSP respones in a file.
func (o *OcspFramework) prepareCertificates(namespace string) error {
	config := fmt.Sprintf(configTemplate, namespace)
	err := ioutil.WriteFile("cfssl_config.json", []byte(config), 0644)
	if err != nil {
		return fmt.Errorf("creating cfssl_config.json file: %v", err)
	}

	commands := []string{
		"cfssl gencert -initca ca_csr.json | cfssljson -bare ca",
		"cfssl gencert -ca ca.pem -ca-key ca-key.pem -config=cfssl_config.json -profile=server leaf_csr.json | cfssljson -bare leaf",
		"cfssl gencert -ca ca.pem -ca-key ca-key.pem -config=cfssl_config.json -profile=ocsp ocsp_csr.json | cfssljson -bare ocsp",
		"cfssl gencert -ca ca.pem -ca-key ca-key.pem -config=cfssl_config.json -profile=client smime_csr.json | cfssljson -bare smime",
	}

	for _, command := range commands {
		ginkgo.By(fmt.Sprintf("running %v", command))
		out, err := exec.Command("bash", "-c", command).CombinedOutput()
		if err != nil {
			framework.Logf("Command error: %v\n%v\n%v", command, err, string(out))
			return err
		}
	}

	err = o.OcspSignCertificates(true, "ca.pem", "leaf.pem", "smime.pem")
	if err != nil {
		return fmt.Errorf("error signing certificates: %v", err)
	}

	return nil
}

// OcspSignCertificates stores the OCSP signature responses in a file.
// Validity is optional and the response hardcoded to Priveledge Withdrawn (9)
func (o *OcspFramework) OcspSignCertificates(valid bool, certs ...string) error {
	// We always want to have a blank file to write new ocsp responses too
	err := ioutil.WriteFile("responses", []byte{}, 0644)
	if err != nil {
		return fmt.Errorf("error writing blank responses file: %v", err)
	}

	signCommand := "cfssl ocspsign -ca ca.pem -responder ocsp.pem -responder-key ocsp-key.pem -cert %s | cfssljson -bare -stdout >> responses"
	for _, cert := range certs {
		if !valid {
			cert += fmt.Sprintf(" -status revoked -reason %v", ocspReasonWithdrawn)
		}
		command := fmt.Sprintf(signCommand, cert)
		ginkgo.By(fmt.Sprintf("running %v", command))
		out, err := exec.Command("bash", "-c", command).CombinedOutput()
		if err != nil {
			framework.Logf("Command error: %v\n%v\n%v", command, err, string(out))
			return err
		}
	}

	return nil
}

// TlsConfig returns a TLS Configurations suitable for client certificate HTTP Clients
func (o *OcspFramework) TlsConfig(host string) (*tls.Config, error) {
	clientPair, err := tls.LoadX509KeyPair("smime.pem", "smime-key.pem")
	if err != nil {
		return nil, err
	}

	caCert, err := ioutil.ReadFile("ca.pem")
	if err != nil {
		return nil, err
	}

	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCert)

	return &tls.Config{
		ServerName:         host,
		Certificates:       []tls.Certificate{clientPair},
		RootCAs:            caPool,
		InsecureSkipVerify: true,
	}, nil
}

// EnsureOCSPResponderDeployment deploysa a OCSP Responder. Requires a OCSP signature file to be created.
func (o *OcspFramework) EnsureOCSPResponderDeployment(nameSpace, name string) error {
	c, err := ocspSecret(o.framework.KubeClientSet, name, nameSpace)
	if err != nil {
		return fmt.Errorf("error reading OCSP signature file (did you generate the ocsp signatures?): %v", err)
	}

	_, err = o.framework.EnsureConfigMap(c)
	if err != nil {
		return err
	}

	d, s := ocspServeDeployment(nameSpace, name)
	o.framework.EnsureDeployment(d)
	o.framework.EnsureService(s)

	return nil
}

func ocspSecret(client kubernetes.Interface, configName string, nameSpace string) (*corev1.ConfigMap, error) {
	responses, err := ioutil.ReadFile("responses")
	if err != nil {
		return nil, err
	}

	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: configName,
		},
		BinaryData: map[string][]byte{
			"responses": responses,
		},
	}, nil
}

func ocspServeDeployment(namespace, name string) (*appsv1.Deployment, *corev1.Service) {
	return &appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: namespace,
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: framework.NewInt32(1),
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						"app": name,
					},
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"app":     name,
							"service": "ocspserve",
						},
					},
					Spec: corev1.PodSpec{
						TerminationGracePeriodSeconds: framework.NewInt64(0),
						Containers: []corev1.Container{
							{
								Name:  name,
								Image: "k8s.gcr.io/ingress-nginx/e2e-test-cfssl@sha256:be2f69024f7b7053f35b86677de16bdaa5d3ff0f81b17581ef0b0c6804188b03",
								Command: []string{
									"/bin/bash",
									"-c",
									"cfssl ocspserve -port=80 -address=0.0.0.0 -responses=/data/responses -loglevel=0",
								},
								Ports: []corev1.ContainerPort{
									{
										Name:          "http",
										ContainerPort: 80,
									},
								},
								VolumeMounts: []corev1.VolumeMount{
									{
										Name:      name,
										MountPath: "/data",
										ReadOnly:  true,
									},
								},
							},
						},
						Volumes: []corev1.Volume{
							{
								Name: name,
								VolumeSource: corev1.VolumeSource{
									ConfigMap: &corev1.ConfigMapVolumeSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: name,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		&corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: namespace,
			},
			Spec: corev1.ServiceSpec{
				Ports: []corev1.ServicePort{
					{
						Name:       "http",
						Port:       80,
						TargetPort: intstr.FromInt(80),
						Protocol:   corev1.ProtocolTCP,
					},
				},
				Selector: map[string]string{
					"app": name,
				},
			},
		}
}
