package main

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/retry"
	"sigs.k8s.io/controller-runtime/pkg/client"

	adminv2 "github.com/solo-io/solo-apis/client-go/admin.gloo.solo.io/v2"
)

func main() {
	// Initialize the client to interact with the Kubernetes API
	cfg, err := clientcmd.NewDefaultClientConfigLoadingRules().Load()
	if err != nil {
		log.Fatal(err)
	}

	kubeConfig := clientcmd.NewDefaultClientConfig(*cfg, &clientcmd.ConfigOverrides{})
	clientConfig, err := kubeConfig.ClientConfig()
	if err != nil {
		log.Fatal(err)
	}

	adminv2Clientset, err := adminv2.NewClientsetFromConfig(clientConfig)
	if err != nil {
		log.Fatal(err)
	}

	wsClient := adminv2Clientset.Workspaces()

	// Create a new empty Workspace in the gloo-mesh namespace
	ctx := context.Background()

	newWorkspace := &adminv2.Workspace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "example",
			Namespace: "gloo-mesh",
		},
		Spec: adminv2.WorkspaceSpec{},
	}

	log.Printf("creating \"%s\" workspace...\n", newWorkspace.GetName())
	err = wsClient.CreateWorkspace(ctx, newWorkspace)
	if err != nil {
		log.Fatal(err)
	}

	// Add a new workload cluster to this workspace
	log.Printf("updating \"%s\" workspace with a new cluster...\n", newWorkspace.GetName())
	var updatedWorkspace = &adminv2.Workspace{}
	err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Workspace before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		updatedWorkspace, err = wsClient.GetWorkspace(ctx, client.ObjectKeyFromObject(newWorkspace))
		if err != nil {
			return err
		}

		updatedWorkspace.Spec.WorkloadClusters = []*adminv2.ClusterSelector{
			{
				Name: "cluster-1",
				Namespaces: []*adminv2.ClusterSelector_NamespaceSelector{
					{
						Name: "apps",
					},
				},
			},
		}
		err = wsClient.UpdateWorkspace(ctx, updatedWorkspace, &client.UpdateOptions{})
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	// Verify the new cluster has been added to the workspace
	log.Printf("clusters registered for this workspace: %+v\n", updatedWorkspace.Spec.WorkloadClusters)

	// Finally, clean up the environment by deleting the workspace
	log.Printf("deleting \"%s\" workspace...\n", updatedWorkspace.GetName())
	err = wsClient.DeleteWorkspace(ctx, client.ObjectKeyFromObject(updatedWorkspace))
	if err != nil {
		log.Fatal(err)
	}
}
