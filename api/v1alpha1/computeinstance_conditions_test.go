package v1alpha1_test

import (
	. "github.com/onsi/ginkgo/v2" //nolint:revive,staticcheck
	. "github.com/onsi/gomega"    //nolint:revive,staticcheck
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "github.com/osac-project/osac-operator/api/v1alpha1"
)

var _ = Describe("ComputeInstance Conditions", func() {
	var ci *v1alpha1.ComputeInstance

	BeforeEach(func() {
		ci = &v1alpha1.ComputeInstance{}
	})

	Describe("SetStatusCondition", func() {
		It("should add a condition to an empty conditions slice", func() {
			changed := ci.SetStatusCondition(
				v1alpha1.ComputeInstanceConditionAvailable,
				metav1.ConditionTrue,
				"instance is available",
				"Available",
			)

			Expect(changed).To(BeTrue())
			Expect(ci.Status.Conditions).To(HaveLen(1))
			Expect(ci.Status.Conditions[0].Type).To(Equal(string(v1alpha1.ComputeInstanceConditionAvailable)))
			Expect(ci.Status.Conditions[0].Status).To(Equal(metav1.ConditionTrue))
			Expect(ci.Status.Conditions[0].Message).To(Equal("instance is available"))
			Expect(ci.Status.Conditions[0].Reason).To(Equal("Available"))
		})

		It("should initialize nil conditions slice", func() {
			ci.Status.Conditions = nil

			ci.SetStatusCondition(
				v1alpha1.ComputeInstanceConditionProvisioned,
				metav1.ConditionFalse,
				"not yet provisioned",
				"Pending",
			)

			Expect(ci.Status.Conditions).ToNot(BeNil())
			Expect(ci.Status.Conditions).To(HaveLen(1))
		})

		It("should update an existing condition", func() {
			ci.SetStatusCondition(
				v1alpha1.ComputeInstanceConditionAvailable,
				metav1.ConditionFalse,
				"starting up",
				"Starting",
			)

			changed := ci.SetStatusCondition(
				v1alpha1.ComputeInstanceConditionAvailable,
				metav1.ConditionTrue,
				"instance is available",
				"Available",
			)

			Expect(changed).To(BeTrue())
			Expect(ci.Status.Conditions).To(HaveLen(1))
			Expect(ci.Status.Conditions[0].Status).To(Equal(metav1.ConditionTrue))
			Expect(ci.Status.Conditions[0].Message).To(Equal("instance is available"))
		})

		It("should add multiple different conditions", func() {
			ci.SetStatusCondition(
				v1alpha1.ComputeInstanceConditionAvailable,
				metav1.ConditionTrue,
				"available",
				"Available",
			)
			ci.SetStatusCondition(
				v1alpha1.ComputeInstanceConditionProvisioned,
				metav1.ConditionTrue,
				"provisioned",
				"Provisioned",
			)

			Expect(ci.Status.Conditions).To(HaveLen(2))
		})
	})

	Describe("GetStatusCondition", func() {
		It("should return the matching condition", func() {
			ci.SetStatusCondition(
				v1alpha1.ComputeInstanceConditionAvailable,
				metav1.ConditionTrue,
				"available",
				"Available",
			)

			cond := ci.GetStatusCondition(v1alpha1.ComputeInstanceConditionAvailable)
			Expect(cond).ToNot(BeNil())
			Expect(cond.Type).To(Equal(string(v1alpha1.ComputeInstanceConditionAvailable)))
			Expect(cond.Status).To(Equal(metav1.ConditionTrue))
		})

		It("should return nil when condition is not found", func() {
			ci.SetStatusCondition(
				v1alpha1.ComputeInstanceConditionAvailable,
				metav1.ConditionTrue,
				"available",
				"Available",
			)

			cond := ci.GetStatusCondition(v1alpha1.ComputeInstanceConditionProvisioned)
			Expect(cond).To(BeNil())
		})

		It("should return nil when conditions slice is nil", func() {
			ci.Status.Conditions = nil

			cond := ci.GetStatusCondition(v1alpha1.ComputeInstanceConditionAvailable)
			Expect(cond).To(BeNil())
		})

		It("should return nil when conditions slice is empty", func() {
			ci.Status.Conditions = []metav1.Condition{}

			cond := ci.GetStatusCondition(v1alpha1.ComputeInstanceConditionAvailable)
			Expect(cond).To(BeNil())
		})
	})

	Describe("IsStatusConditionTrue", func() {
		It("should return true when condition status is True", func() {
			ci.SetStatusCondition(
				v1alpha1.ComputeInstanceConditionAvailable,
				metav1.ConditionTrue,
				"available",
				"Available",
			)

			Expect(ci.IsStatusConditionTrue(v1alpha1.ComputeInstanceConditionAvailable)).To(BeTrue())
		})

		It("should return false when condition status is False", func() {
			ci.SetStatusCondition(
				v1alpha1.ComputeInstanceConditionAvailable,
				metav1.ConditionFalse,
				"not available",
				"Unavailable",
			)

			Expect(ci.IsStatusConditionTrue(v1alpha1.ComputeInstanceConditionAvailable)).To(BeFalse())
		})

		It("should return false when condition status is Unknown", func() {
			ci.SetStatusCondition(
				v1alpha1.ComputeInstanceConditionAvailable,
				metav1.ConditionUnknown,
				"unknown",
				"Unknown",
			)

			Expect(ci.IsStatusConditionTrue(v1alpha1.ComputeInstanceConditionAvailable)).To(BeFalse())
		})

		It("should return false when condition does not exist", func() {
			Expect(ci.IsStatusConditionTrue(v1alpha1.ComputeInstanceConditionAvailable)).To(BeFalse())
		})
	})

	Describe("IsStatusConditionFalse", func() {
		It("should return true when condition status is False", func() {
			ci.SetStatusCondition(
				v1alpha1.ComputeInstanceConditionAvailable,
				metav1.ConditionFalse,
				"not available",
				"Unavailable",
			)

			Expect(ci.IsStatusConditionFalse(v1alpha1.ComputeInstanceConditionAvailable)).To(BeTrue())
		})

		It("should return false when condition status is True", func() {
			ci.SetStatusCondition(
				v1alpha1.ComputeInstanceConditionAvailable,
				metav1.ConditionTrue,
				"available",
				"Available",
			)

			Expect(ci.IsStatusConditionFalse(v1alpha1.ComputeInstanceConditionAvailable)).To(BeFalse())
		})

		It("should return false when condition status is Unknown", func() {
			ci.SetStatusCondition(
				v1alpha1.ComputeInstanceConditionAvailable,
				metav1.ConditionUnknown,
				"unknown",
				"Unknown",
			)

			Expect(ci.IsStatusConditionFalse(v1alpha1.ComputeInstanceConditionAvailable)).To(BeFalse())
		})

		It("should return false when condition does not exist", func() {
			Expect(ci.IsStatusConditionFalse(v1alpha1.ComputeInstanceConditionAvailable)).To(BeFalse())
		})
	})

	Describe("IsStatusConditionUnknown", func() {
		It("should return true when condition status is Unknown", func() {
			ci.SetStatusCondition(
				v1alpha1.ComputeInstanceConditionAvailable,
				metav1.ConditionUnknown,
				"unknown",
				"Unknown",
			)

			Expect(ci.IsStatusConditionUnknown(v1alpha1.ComputeInstanceConditionAvailable)).To(BeTrue())
		})

		It("should return true when condition does not exist", func() {
			Expect(ci.IsStatusConditionUnknown(v1alpha1.ComputeInstanceConditionAvailable)).To(BeTrue())
		})

		It("should return false when condition status is True", func() {
			ci.SetStatusCondition(
				v1alpha1.ComputeInstanceConditionAvailable,
				metav1.ConditionTrue,
				"available",
				"Available",
			)

			Expect(ci.IsStatusConditionUnknown(v1alpha1.ComputeInstanceConditionAvailable)).To(BeFalse())
		})

		It("should return false when condition status is False", func() {
			ci.SetStatusCondition(
				v1alpha1.ComputeInstanceConditionAvailable,
				metav1.ConditionFalse,
				"not available",
				"Unavailable",
			)

			Expect(ci.IsStatusConditionUnknown(v1alpha1.ComputeInstanceConditionAvailable)).To(BeFalse())
		})
	})
})
