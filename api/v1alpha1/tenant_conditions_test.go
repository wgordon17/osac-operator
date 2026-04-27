package v1alpha1_test

import (
	. "github.com/onsi/ginkgo/v2" //nolint:revive,staticcheck
	. "github.com/onsi/gomega"    //nolint:revive,staticcheck
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "github.com/osac-project/osac-operator/api/v1alpha1"
)

var _ = Describe("Tenant Conditions", func() {
	var tenant *v1alpha1.Tenant

	BeforeEach(func() {
		tenant = &v1alpha1.Tenant{}
	})

	Describe("SetStatusCondition", func() {
		It("should add a condition to an empty conditions slice", func() {
			changed := tenant.SetStatusCondition(
				v1alpha1.TenantConditionNamespaceReady,
				metav1.ConditionTrue,
				v1alpha1.TenantReasonFound,
				"namespace exists",
			)

			Expect(changed).To(BeTrue())
			Expect(tenant.Status.Conditions).To(HaveLen(1))
			Expect(tenant.Status.Conditions[0].Type).To(Equal(string(v1alpha1.TenantConditionNamespaceReady)))
			Expect(tenant.Status.Conditions[0].Status).To(Equal(metav1.ConditionTrue))
			Expect(tenant.Status.Conditions[0].Reason).To(Equal(v1alpha1.TenantReasonFound))
			Expect(tenant.Status.Conditions[0].Message).To(Equal("namespace exists"))
		})

		It("should initialize nil conditions slice", func() {
			tenant.Status.Conditions = nil

			tenant.SetStatusCondition(
				v1alpha1.TenantConditionNamespaceReady,
				metav1.ConditionFalse,
				v1alpha1.TenantReasonNotFound,
				"namespace not found",
			)

			Expect(tenant.Status.Conditions).ToNot(BeNil())
			Expect(tenant.Status.Conditions).To(HaveLen(1))
		})

		It("should update an existing condition", func() {
			tenant.SetStatusCondition(
				v1alpha1.TenantConditionNamespaceReady,
				metav1.ConditionFalse,
				v1alpha1.TenantReasonNotFound,
				"namespace not found",
			)

			changed := tenant.SetStatusCondition(
				v1alpha1.TenantConditionNamespaceReady,
				metav1.ConditionTrue,
				v1alpha1.TenantReasonFound,
				"namespace exists",
			)

			Expect(changed).To(BeTrue())
			Expect(tenant.Status.Conditions).To(HaveLen(1))
			Expect(tenant.Status.Conditions[0].Status).To(Equal(metav1.ConditionTrue))
			Expect(tenant.Status.Conditions[0].Reason).To(Equal(v1alpha1.TenantReasonFound))
		})

		It("should add multiple different conditions", func() {
			tenant.SetStatusCondition(
				v1alpha1.TenantConditionNamespaceReady,
				metav1.ConditionTrue,
				v1alpha1.TenantReasonFound,
				"namespace exists",
			)
			tenant.SetStatusCondition(
				v1alpha1.TenantConditionStorageClassReady,
				metav1.ConditionTrue,
				v1alpha1.TenantReasonFound,
				"storage class found",
			)

			Expect(tenant.Status.Conditions).To(HaveLen(2))
		})
	})

	Describe("GetStatusCondition", func() {
		It("should return the matching condition", func() {
			tenant.SetStatusCondition(
				v1alpha1.TenantConditionNamespaceReady,
				metav1.ConditionTrue,
				v1alpha1.TenantReasonFound,
				"namespace exists",
			)

			cond := tenant.GetStatusCondition(v1alpha1.TenantConditionNamespaceReady)
			Expect(cond).ToNot(BeNil())
			Expect(cond.Type).To(Equal(string(v1alpha1.TenantConditionNamespaceReady)))
			Expect(cond.Status).To(Equal(metav1.ConditionTrue))
		})

		It("should return nil when condition is not found", func() {
			tenant.SetStatusCondition(
				v1alpha1.TenantConditionNamespaceReady,
				metav1.ConditionTrue,
				v1alpha1.TenantReasonFound,
				"namespace exists",
			)

			cond := tenant.GetStatusCondition(v1alpha1.TenantConditionStorageClassReady)
			Expect(cond).To(BeNil())
		})

		It("should return nil when conditions slice is nil", func() {
			tenant.Status.Conditions = nil

			cond := tenant.GetStatusCondition(v1alpha1.TenantConditionNamespaceReady)
			Expect(cond).To(BeNil())
		})

		It("should return nil when conditions slice is empty", func() {
			tenant.Status.Conditions = []metav1.Condition{}

			cond := tenant.GetStatusCondition(v1alpha1.TenantConditionNamespaceReady)
			Expect(cond).To(BeNil())
		})

		It("should return the correct condition when multiple exist", func() {
			tenant.SetStatusCondition(
				v1alpha1.TenantConditionNamespaceReady,
				metav1.ConditionTrue,
				v1alpha1.TenantReasonFound,
				"namespace exists",
			)
			tenant.SetStatusCondition(
				v1alpha1.TenantConditionStorageClassReady,
				metav1.ConditionFalse,
				v1alpha1.TenantReasonNotFound,
				"storage class not found",
			)

			cond := tenant.GetStatusCondition(v1alpha1.TenantConditionStorageClassReady)
			Expect(cond).ToNot(BeNil())
			Expect(cond.Type).To(Equal(string(v1alpha1.TenantConditionStorageClassReady)))
			Expect(cond.Status).To(Equal(metav1.ConditionFalse))
			Expect(cond.Reason).To(Equal(v1alpha1.TenantReasonNotFound))
		})
	})
})
