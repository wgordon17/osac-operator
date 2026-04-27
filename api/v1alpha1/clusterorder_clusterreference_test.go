package v1alpha1_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "github.com/osac-project/osac-operator/api/v1alpha1"
)

var _ = Describe("ClusterOrder ClusterReference", func() {
	var co *v1alpha1.ClusterOrder

	BeforeEach(func() {
		co = &v1alpha1.ClusterOrder{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-order",
				Namespace: "default",
			},
		}
	})

	Describe("EnsureClusterReference", func() {
		It("should initialize ClusterReference when nil", func() {
			Expect(co.Status.ClusterReference).To(BeNil())
			co.EnsureClusterReference()
			Expect(co.Status.ClusterReference).NotTo(BeNil())
		})

		It("should not overwrite an existing ClusterReference", func() {
			co.EnsureClusterReference()
			co.Status.ClusterReference.Namespace = "existing-ns"

			co.EnsureClusterReference()
			Expect(co.Status.ClusterReference.Namespace).To(Equal("existing-ns"))
		})
	})

	Describe("Getters with nil ClusterReference", func() {
		BeforeEach(func() {
			Expect(co.Status.ClusterReference).To(BeNil())
		})

		It("GetClusterReferenceNamespace should return empty string", func() {
			Expect(co.GetClusterReferenceNamespace()).To(Equal(""))
		})

		It("GetClusterReferenceServiceAccountName should return empty string", func() {
			Expect(co.GetClusterReferenceServiceAccountName()).To(Equal(""))
		})

		It("GetClusterReferenceRoleBindingName should return empty string", func() {
			Expect(co.GetClusterReferenceRoleBindingName()).To(Equal(""))
		})

		It("GetClusterReferenceHostedClusterName should return empty string", func() {
			Expect(co.GetClusterReferenceHostedClusterName()).To(Equal(""))
		})
	})

	Describe("SetClusterReferenceNamespace", func() {
		It("should set the namespace and initialize ClusterReference if nil", func() {
			co.SetClusterReferenceNamespace("my-namespace")
			Expect(co.Status.ClusterReference).NotTo(BeNil())
			Expect(co.Status.ClusterReference.Namespace).To(Equal("my-namespace"))
		})

		It("should update the namespace on an existing ClusterReference", func() {
			co.EnsureClusterReference()
			co.SetClusterReferenceNamespace("first")
			co.SetClusterReferenceNamespace("second")
			Expect(co.Status.ClusterReference.Namespace).To(Equal("second"))
		})
	})

	Describe("GetClusterReferenceNamespace", func() {
		It("should return the namespace when set", func() {
			co.SetClusterReferenceNamespace("my-namespace")
			Expect(co.GetClusterReferenceNamespace()).To(Equal("my-namespace"))
		})
	})

	Describe("SetClusterReferenceServiceAccountName", func() {
		It("should set the service account name and initialize ClusterReference if nil", func() {
			co.SetClusterReferenceServiceAccountName("my-sa")
			Expect(co.Status.ClusterReference).NotTo(BeNil())
			Expect(co.Status.ClusterReference.ServiceAccountName).To(Equal("my-sa"))
		})

		It("should update the service account name on an existing ClusterReference", func() {
			co.EnsureClusterReference()
			co.SetClusterReferenceServiceAccountName("sa-1")
			co.SetClusterReferenceServiceAccountName("sa-2")
			Expect(co.Status.ClusterReference.ServiceAccountName).To(Equal("sa-2"))
		})
	})

	Describe("GetClusterReferenceServiceAccountName", func() {
		It("should return the service account name when set", func() {
			co.SetClusterReferenceServiceAccountName("my-sa")
			Expect(co.GetClusterReferenceServiceAccountName()).To(Equal("my-sa"))
		})
	})

	Describe("SetClusterReferenceRoleBindingName", func() {
		It("should set the role binding name and initialize ClusterReference if nil", func() {
			co.SetClusterReferenceRoleBindingName("my-rb")
			Expect(co.Status.ClusterReference).NotTo(BeNil())
			Expect(co.Status.ClusterReference.RoleBindingName).To(Equal("my-rb"))
		})

		It("should update the role binding name on an existing ClusterReference", func() {
			co.EnsureClusterReference()
			co.SetClusterReferenceRoleBindingName("rb-1")
			co.SetClusterReferenceRoleBindingName("rb-2")
			Expect(co.Status.ClusterReference.RoleBindingName).To(Equal("rb-2"))
		})
	})

	Describe("GetClusterReferenceRoleBindingName", func() {
		It("should return the role binding name when set", func() {
			co.SetClusterReferenceRoleBindingName("my-rb")
			Expect(co.GetClusterReferenceRoleBindingName()).To(Equal("my-rb"))
		})
	})

	Describe("SetClusterReferenceHostedClusterName", func() {
		It("should set the hosted cluster name and initialize ClusterReference if nil", func() {
			co.SetClusterReferenceHostedClusterName("my-hc")
			Expect(co.Status.ClusterReference).NotTo(BeNil())
			Expect(co.Status.ClusterReference.HostedClusterName).To(Equal("my-hc"))
		})

		It("should update the hosted cluster name on an existing ClusterReference", func() {
			co.EnsureClusterReference()
			co.SetClusterReferenceHostedClusterName("hc-1")
			co.SetClusterReferenceHostedClusterName("hc-2")
			Expect(co.Status.ClusterReference.HostedClusterName).To(Equal("hc-2"))
		})
	})

	Describe("GetClusterReferenceHostedClusterName", func() {
		It("should return the hosted cluster name when set", func() {
			co.SetClusterReferenceHostedClusterName("my-hc")
			Expect(co.GetClusterReferenceHostedClusterName()).To(Equal("my-hc"))
		})
	})

	Describe("Combined setter/getter workflow", func() {
		It("should set and retrieve all fields independently", func() {
			co.SetClusterReferenceNamespace("ns-1")
			co.SetClusterReferenceServiceAccountName("sa-1")
			co.SetClusterReferenceRoleBindingName("rb-1")
			co.SetClusterReferenceHostedClusterName("hc-1")

			Expect(co.GetClusterReferenceNamespace()).To(Equal("ns-1"))
			Expect(co.GetClusterReferenceServiceAccountName()).To(Equal("sa-1"))
			Expect(co.GetClusterReferenceRoleBindingName()).To(Equal("rb-1"))
			Expect(co.GetClusterReferenceHostedClusterName()).To(Equal("hc-1"))
		})
	})
})
