package v1alpha1_test

import (
	. "github.com/onsi/ginkgo/v2" //nolint:revive,staticcheck
	. "github.com/onsi/gomega"    //nolint:revive,staticcheck
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "github.com/osac-project/osac-operator/api/v1alpha1"
)

var _ = Describe("Subnet Conditions", func() {
	var subnet *v1alpha1.Subnet

	BeforeEach(func() {
		subnet = &v1alpha1.Subnet{}
	})

	Describe("SetStatusCondition", func() {
		It("should add a condition to an empty conditions slice", func() {
			condition := metav1.Condition{
				Type:               string(v1alpha1.SubnetConditionReady),
				Status:             metav1.ConditionTrue,
				Reason:             v1alpha1.SubnetProvisioningSucceeded,
				Message:            "subnet is ready",
				LastTransitionTime: metav1.Now(),
			}

			v1alpha1.SetStatusCondition(subnet, condition)

			Expect(subnet.Status.Conditions).To(HaveLen(1))
			Expect(subnet.Status.Conditions[0].Type).To(Equal(string(v1alpha1.SubnetConditionReady)))
			Expect(subnet.Status.Conditions[0].Status).To(Equal(metav1.ConditionTrue))
			Expect(subnet.Status.Conditions[0].Reason).To(Equal(v1alpha1.SubnetProvisioningSucceeded))
			Expect(subnet.Status.Conditions[0].Message).To(Equal("subnet is ready"))
		})

		It("should update an existing condition", func() {
			initial := metav1.Condition{
				Type:               string(v1alpha1.SubnetConditionReady),
				Status:             metav1.ConditionFalse,
				Reason:             v1alpha1.SubnetProvisioningFailed,
				Message:            "provisioning failed",
				LastTransitionTime: metav1.Now(),
			}
			v1alpha1.SetStatusCondition(subnet, initial)

			updated := metav1.Condition{
				Type:               string(v1alpha1.SubnetConditionReady),
				Status:             metav1.ConditionTrue,
				Reason:             v1alpha1.SubnetProvisioningSucceeded,
				Message:            "subnet is ready",
				LastTransitionTime: metav1.Now(),
			}
			v1alpha1.SetStatusCondition(subnet, updated)

			Expect(subnet.Status.Conditions).To(HaveLen(1))
			Expect(subnet.Status.Conditions[0].Status).To(Equal(metav1.ConditionTrue))
			Expect(subnet.Status.Conditions[0].Reason).To(Equal(v1alpha1.SubnetProvisioningSucceeded))
		})

		It("should add multiple different conditions", func() {
			ready := metav1.Condition{
				Type:               string(v1alpha1.SubnetConditionReady),
				Status:             metav1.ConditionTrue,
				Reason:             v1alpha1.SubnetProvisioningSucceeded,
				Message:            "subnet is ready",
				LastTransitionTime: metav1.Now(),
			}
			v1alpha1.SetStatusCondition(subnet, ready)

			networkProvisioned := metav1.Condition{
				Type:               string(v1alpha1.SubnetConditionNetworkProvisioned),
				Status:             metav1.ConditionTrue,
				Reason:             v1alpha1.SubnetProvisioningSucceeded,
				Message:            "network provisioned",
				LastTransitionTime: metav1.Now(),
			}
			v1alpha1.SetStatusCondition(subnet, networkProvisioned)

			Expect(subnet.Status.Conditions).To(HaveLen(2))
		})
	})

	Describe("GetStatusCondition", func() {
		It("should return the matching condition", func() {
			condition := metav1.Condition{
				Type:               string(v1alpha1.SubnetConditionReady),
				Status:             metav1.ConditionTrue,
				Reason:             v1alpha1.SubnetProvisioningSucceeded,
				Message:            "subnet is ready",
				LastTransitionTime: metav1.Now(),
			}
			v1alpha1.SetStatusCondition(subnet, condition)

			got := v1alpha1.GetStatusCondition(subnet, v1alpha1.SubnetConditionReady)
			Expect(got).ToNot(BeNil())
			Expect(got.Type).To(Equal(string(v1alpha1.SubnetConditionReady)))
			Expect(got.Status).To(Equal(metav1.ConditionTrue))
		})

		It("should return nil when condition is not found", func() {
			condition := metav1.Condition{
				Type:               string(v1alpha1.SubnetConditionReady),
				Status:             metav1.ConditionTrue,
				Reason:             v1alpha1.SubnetProvisioningSucceeded,
				Message:            "subnet is ready",
				LastTransitionTime: metav1.Now(),
			}
			v1alpha1.SetStatusCondition(subnet, condition)

			got := v1alpha1.GetStatusCondition(subnet, v1alpha1.SubnetConditionNetworkProvisioned)
			Expect(got).To(BeNil())
		})

		It("should return nil when conditions slice is empty", func() {
			got := v1alpha1.GetStatusCondition(subnet, v1alpha1.SubnetConditionReady)
			Expect(got).To(BeNil())
		})

		It("should return the correct condition when multiple exist", func() {
			ready := metav1.Condition{
				Type:               string(v1alpha1.SubnetConditionReady),
				Status:             metav1.ConditionTrue,
				Reason:             v1alpha1.SubnetProvisioningSucceeded,
				Message:            "subnet is ready",
				LastTransitionTime: metav1.Now(),
			}
			v1alpha1.SetStatusCondition(subnet, ready)

			networkReady := metav1.Condition{
				Type:               string(v1alpha1.SubnetConditionNetworkReady),
				Status:             metav1.ConditionFalse,
				Reason:             v1alpha1.SubnetNetworkProviderError,
				Message:            "network not ready",
				LastTransitionTime: metav1.Now(),
			}
			v1alpha1.SetStatusCondition(subnet, networkReady)

			got := v1alpha1.GetStatusCondition(subnet, v1alpha1.SubnetConditionNetworkReady)
			Expect(got).ToNot(BeNil())
			Expect(got.Type).To(Equal(string(v1alpha1.SubnetConditionNetworkReady)))
			Expect(got.Status).To(Equal(metav1.ConditionFalse))
			Expect(got.Reason).To(Equal(v1alpha1.SubnetNetworkProviderError))
		})
	})
})
