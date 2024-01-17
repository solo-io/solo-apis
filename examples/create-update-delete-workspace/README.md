# Create, Update & Delete Workspace

This example program demonstrates the fundamental operations for managing on
**Workspace** resources, such as `Create`, `List`, `Update` and `Delete`.

You can adopt the source code from this example to write programs that manage
other types of resources through the Kubernetes API.

## Running this example

Make sure you have a Kubernetes cluster and `kubectl` is configured:

```
kubectl get nodes
```

You also need to install **Gloo Mesh Enterprise** first, follow the
[official documentation](https://docs.solo.io/gloo-mesh-enterprise/latest/setup)
to learn how to do this.

Finally, run this application on your workstation:

```
$ cd create-update-delete-workspace
$ go run .
2023/05/29 14:07:41 creating "example" workspace...
2023/05/29 14:07:41 updating "example" workspace with a new cluster...
2023/05/29 14:07:41 clusters registered for this workspace: [name:"cluster-1" namespaces:{name:"apps"}]
2023/05/29 14:07:41 deleting "example" workspace...
```

Running this command will execute the following operations on your cluster:

1. **Create Workspace:** This will create an empty Workspace (without workload
   clusters registered).
2. **Update Workspace:** This will update the Workspace resource created in
   previous step by registering a new workload cluster. You are encouraged to
   inspect the retry loop that handles conflicts.
3. **Delete Workspace:** This will delete the Workspace object.

## Cleanup

Successfully running this program will clean the created artifacts. If you
terminate the program without completing, you can clean up the created
workspace with:

```
kubectl delete workspace example -n gloo-mesh
```

## References

This example is inspired by the Kubernetes official client-go examples:
https://github.com/kubernetes/client-go/tree/master/examples/create-update-delete-deployment
