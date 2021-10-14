package gloo_fed_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/resource"
	mock_resource "github.com/solo-io/skv2/pkg/resource/mocks"
	"github.com/solo-io/solo-apis/pkg/api/fed.gloo.solo.io/v1"
	"github.com/solo-io/solo-apis/pkg/api/fed.gloo.solo.io/v1/types"
	"github.com/solo-io/solo-apis/pkg/api/multicluster.solo.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	client2 "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Controller", func() {

	var (
		ctl     *gomock.Controller
		client  *mock_resource.MockClient
		ctx     = context.Background()
	)

	JustBeforeEach(func() {
		ctl = gomock.NewController(GinkgoT())
		client = mock_resource.NewMockClient(ctl)
	})

	JustAfterEach(func() {
		ctl.Finish()
	})

	Context("FederatedUpstream", func() {

		var desired *v1.FederatedUpstream

		BeforeEach(func() {
			desired = &v1.FederatedUpstream{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "ns",
					Name:      "name",
				},
				Spec: types.FederatedUpstreamSpec{
					Placement: &v1alpha1.Placement{
						Namespaces: []string{"ns-desired"},
					},
				},
			}
		})

		It("Can create resource", func() {
			client.EXPECT().Get(ctx, client2.ObjectKeyFromObject(desired), gomock.Any()).Return(
				&errors.StatusError{
					ErrStatus: metav1.Status{Reason: metav1.StatusReasonNotFound},
				})
			client.EXPECT().Create(ctx, desired).Return(nil)

			result, err := controllerutils.Upsert(ctx, client, desired)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(controllerutil.OperationResultCreated))
		})

		It("Can transition resource", func() {
			var called bool

			client.EXPECT().Get(ctx, resource.ToClientKey(desired), gomock.Any()).Return(nil)
			client.EXPECT().Update(ctx, desired).Return(nil)

			existingTest := desired.DeepCopyObject().(*v1.FederatedUpstream)

			result, err := controllerutils.Upsert(ctx, client, desired, func(existing, desired runtime.Object) error {
				called = true

				// necessary to ensure there is a diff between existing and desired
				desired.(*v1.FederatedUpstream).Spec = types.FederatedUpstreamSpec{
					Placement: &v1alpha1.Placement{
						Namespaces: []string{"ns-other"},
					},
				}
				return nil
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(controllerutil.OperationResultUpdated))
			Expect(called).To(BeTrue())

			// object gets updated by transition function correctly
			Expect(existingTest).ToNot(Equal(desired))
			Expect(desired.Spec.Placement.Namespaces).To(Equal([]string{"ns-other"}))
		})

	})

})
