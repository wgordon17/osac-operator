package v1alpha1_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "github.com/osac-project/osac-operator/api/v1alpha1"
)

var _ = Describe("ClusterOrder Conditions", func() {
	var co *v1alpha1.ClusterOrder

	BeforeEach(func() {
		co = &v1alpha1.ClusterOrder{
			ObjectMeta: metav1.ObjectMeta{
				Name:       "test-order",
				Namespace:  "default",
				Generation: 1,
			},
		}
	})

	Describe("SetStatusCondition", func() {
		It("should add a new condition when none exist", func() {
			changed := co.SetStatusCondition("Ready", metav1.ConditionTrue, "all good", "AllReady")
			Expect(changed).To(BeTrue())
			Expect(co.Status.Conditions).To(HaveLen(1))
			Expect(co.Status.Conditions[0].Type).To(Equal("Ready"))
			Expect(co.Status.Conditions[0].Status).To(Equal(metav1.ConditionTrue))
			Expect(co.Status.Conditions[0].Message).To(Equal("all good"))
			Expect(co.Status.Conditions[0].Reason).To(Equal("AllReady"))
		})

		It("should initialize the Conditions slice when it is nil", func() {
			Expect(co.Status.Conditions).To(BeNil())
			co.SetStatusCondition("Accepted", metav1.ConditionTrue, "accepted", "OrderAccepted")
			Expect(co.Status.Conditions).NotTo(BeNil())
			Expect(co.Status.Conditions).To(HaveLen(1))
		})

		It("should update an existing condition", func() {
			co.SetStatusCondition("Ready", metav1.ConditionFalse, "not ready", "Provisioning")
			co.SetStatusCondition("Ready", metav1.ConditionTrue, "all good", "AllReady")

			Expect(co.Status.Conditions).To(HaveLen(1))
			Expect(co.Status.Conditions[0].Status).To(Equal(metav1.ConditionTrue))
			Expect(co.Status.Conditions[0].Message).To(Equal("all good"))
			Expect(co.Status.Conditions[0].Reason).To(Equal("AllReady"))
		})

		It("should add multiple conditions with different types", func() {
			co.SetStatusCondition("Accepted", metav1.ConditionTrue, "accepted", "OrderAccepted")
			co.SetStatusCondition("Progressing", metav1.ConditionTrue, "working", "InProgress")
			co.SetStatusCondition("Available", metav1.ConditionFalse, "not yet", "Waiting")

			Expect(co.Status.Conditions).To(HaveLen(3))
		})

		It("should return false when setting to the same value", func() {
			co.SetStatusCondition("Ready", metav1.ConditionTrue, "all good", "AllReady")
			changed := co.SetStatusCondition("Ready", metav1.ConditionTrue, "all good", "AllReady")
			Expect(changed).To(BeFalse())
		})
	})

	Describe("RemoveStatusCondition", func() {
		It("should remove an existing condition", func() {
			co.SetStatusCondition("Ready", metav1.ConditionTrue, "ready", "AllReady")
			co.SetStatusCondition("Accepted", metav1.ConditionTrue, "accepted", "Accepted")
			Expect(co.Status.Conditions).To(HaveLen(2))

			removed := co.RemoveStatusCondition("Ready")
			Expect(removed).To(BeTrue())
			Expect(co.Status.Conditions).To(HaveLen(1))
			Expect(co.Status.Conditions[0].Type).To(Equal("Accepted"))
		})

		It("should return false when removing a non-existent condition", func() {
			co.SetStatusCondition("Ready", metav1.ConditionTrue, "ready", "AllReady")

			removed := co.RemoveStatusCondition("NonExistent")
			Expect(removed).To(BeFalse())
			Expect(co.Status.Conditions).To(HaveLen(1))
		})

		It("should handle removing from nil conditions slice", func() {
			Expect(co.Status.Conditions).To(BeNil())
			removed := co.RemoveStatusCondition("Ready")
			Expect(removed).To(BeFalse())
		})

		It("should handle removing from empty conditions slice", func() {
			co.Status.Conditions = []metav1.Condition{}
			removed := co.RemoveStatusCondition("Ready")
			Expect(removed).To(BeFalse())
		})
	})

	Describe("IsStatusConditionTrue", func() {
		It("should return true when condition is True", func() {
			co.SetStatusCondition("Ready", metav1.ConditionTrue, "ready", "AllReady")
			Expect(co.IsStatusConditionTrue("Ready")).To(BeTrue())
		})

		It("should return false when condition is False", func() {
			co.SetStatusCondition("Ready", metav1.ConditionFalse, "not ready", "Provisioning")
			Expect(co.IsStatusConditionTrue("Ready")).To(BeFalse())
		})

		It("should return false when condition is Unknown", func() {
			co.SetStatusCondition("Ready", metav1.ConditionUnknown, "unknown", "Unknown")
			Expect(co.IsStatusConditionTrue("Ready")).To(BeFalse())
		})

		It("should return false when condition does not exist", func() {
			Expect(co.IsStatusConditionTrue("NonExistent")).To(BeFalse())
		})

		It("should return false when conditions slice is nil", func() {
			Expect(co.Status.Conditions).To(BeNil())
			Expect(co.IsStatusConditionTrue("Ready")).To(BeFalse())
		})
	})

	Describe("IsStatusConditionFalse", func() {
		It("should return true when condition is False", func() {
			co.SetStatusCondition("Ready", metav1.ConditionFalse, "not ready", "Provisioning")
			Expect(co.IsStatusConditionFalse("Ready")).To(BeTrue())
		})

		It("should return false when condition is True", func() {
			co.SetStatusCondition("Ready", metav1.ConditionTrue, "ready", "AllReady")
			Expect(co.IsStatusConditionFalse("Ready")).To(BeFalse())
		})

		It("should return false when condition is Unknown", func() {
			co.SetStatusCondition("Ready", metav1.ConditionUnknown, "unknown", "Unknown")
			Expect(co.IsStatusConditionFalse("Ready")).To(BeFalse())
		})

		It("should return false when condition does not exist", func() {
			Expect(co.IsStatusConditionFalse("NonExistent")).To(BeFalse())
		})

		It("should return false when conditions slice is nil", func() {
			Expect(co.Status.Conditions).To(BeNil())
			Expect(co.IsStatusConditionFalse("Ready")).To(BeFalse())
		})
	})

	Describe("IsStatusConditionPresentAndEqual", func() {
		It("should return true when condition matches the given status", func() {
			co.SetStatusCondition("Ready", metav1.ConditionTrue, "ready", "AllReady")
			Expect(co.IsStatusConditionPresentAndEqual("Ready", metav1.ConditionTrue)).To(BeTrue())
		})

		It("should return false when condition exists but does not match", func() {
			co.SetStatusCondition("Ready", metav1.ConditionFalse, "not ready", "Provisioning")
			Expect(co.IsStatusConditionPresentAndEqual("Ready", metav1.ConditionTrue)).To(BeFalse())
		})

		It("should return false when condition does not exist", func() {
			Expect(co.IsStatusConditionPresentAndEqual("NonExistent", metav1.ConditionTrue)).To(BeFalse())
		})

		It("should return false when conditions slice is nil", func() {
			Expect(co.Status.Conditions).To(BeNil())
			Expect(co.IsStatusConditionPresentAndEqual("Ready", metav1.ConditionTrue)).To(BeFalse())
		})

		It("should match Unknown status", func() {
			co.SetStatusCondition("Ready", metav1.ConditionUnknown, "unknown", "Unknown")
			Expect(co.IsStatusConditionPresentAndEqual("Ready", metav1.ConditionUnknown)).To(BeTrue())
		})
	})
})
