package main

import (
	"flag"
	"fmt"
	"log"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/restclient"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

func init() {
	flag.IntVar(&podNum, "pod", 0, "number of pods to create")
}

var podNum int

func main() {
	podID := 1234
	nsID := 5678

	c, err := createClient("localhost:8080")
	if err != nil {
		log.Fatalln("Can't connect to Kubernetes API:", err)
	}

	// First try a single pod creation request
	var podArgs []string
	pod := &api.Pod{
		ObjectMeta: api.ObjectMeta{
			Name: makePodName(podID),
		},
		Spec: api.PodSpec{
			Containers: []api.Container{
				{
					Name:  "none",
					Image: "none",
					Args:  podArgs,
				},
			},
		},
	}

	if _, err := c.Pods(makeNS(nsID)).Create(pod); err != nil { // ReplicationControllers(makeNS(nsID)).Create(rc); err != nil {
		log.Panicf("Pod creation failed: %v", err.Error())
	}
	fmt.Printf("created pod (namespace:%s'podID:%s)\n", makeNS(nsID), makePodName(podID))
}

func createClient(addr string) (*client.Client, error) {
	cfg := &restclient.Config{
		Host:  fmt.Sprintf("http://%s", addr),
		QPS:   1000,
		Burst: 1000,
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func makeNS(id int) string {
	return fmt.Sprintf("ksched-ns-%d", id)
}

func makePodName(id int) string {
	return fmt.Sprintf("ksched-pod-%d", id)
}
