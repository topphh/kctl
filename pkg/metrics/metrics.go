package metrics

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

type Service struct {
	Name     string
	PodCount int64
	Cpu      int64
	Memory   int64
}

func GetKubeServiceTops() ([]*Service, error) {
	kubeconfig := filepath.Join(GetHomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	metricsClient, err := versioned.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	podMetricsList, err := metricsClient.MetricsV1beta1().PodMetricses("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	serviceNameToServiceInfo := make(map[string]*Service)

	for _, podMetrics := range podMetricsList.Items {

		totalCPU := int64(0)
		totalMem := int64(0)

		for _, c := range podMetrics.Containers {
			cpu := c.Usage.Cpu().MilliValue()
			mem := c.Usage.Memory().Value()
			totalCPU += cpu
			totalMem += mem
		}

		podnameSplit := strings.Split(podMetrics.Name, "-")
		serviceName := podnameSplit[0]
		if serviceNameToServiceInfo[serviceName] == nil {
			serviceNameToServiceInfo[serviceName] = &Service{}
			serviceNameToServiceInfo[serviceName].Name = serviceName
		}

		serviceNameToServiceInfo[serviceName].PodCount++
		serviceNameToServiceInfo[serviceName].Cpu += totalCPU
		serviceNameToServiceInfo[serviceName].Memory += totalMem

	}

	res := make([]*Service, 0, len(serviceNameToServiceInfo))
	for _, serviceInfo := range serviceNameToServiceInfo {
		res = append(res, serviceInfo)
	}
	return res, nil
}

func GetHomeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
