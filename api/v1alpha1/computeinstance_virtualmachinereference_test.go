package v1alpha1_test

import (
	. "github.com/onsi/ginkgo/v2" //nolint:revive,staticcheck
	. "github.com/onsi/gomega"    //nolint:revive,staticcheck

	v1alpha1 "github.com/osac-project/osac-operator/api/v1alpha1"
)

var _ = Describe("ComputeInstance VirtualMachineReference", func() {
	var ci *v1alpha1.ComputeInstance

	BeforeEach(func() {
		ci = &v1alpha1.ComputeInstance{}
	})

	Describe("VirtualMachineReference", func() {
		Describe("SetVirtualMachineReferenceNamespace", func() {
			It("should set the namespace", func() {
				ci.SetVirtualMachineReferenceNamespace("tenant-ns")

				Expect(ci.Status.VirtualMachineReference).ToNot(BeNil())
				Expect(ci.Status.VirtualMachineReference.Namespace).To(Equal("tenant-ns"))
			})

			It("should initialize nil reference before setting", func() {
				Expect(ci.Status.VirtualMachineReference).To(BeNil())

				ci.SetVirtualMachineReferenceNamespace("tenant-ns")

				Expect(ci.Status.VirtualMachineReference).ToNot(BeNil())
			})
		})

		Describe("SetVirtualMachineReferenceKubeVirtVirtualMachineName", func() {
			It("should set the KubeVirt VM name", func() {
				ci.SetVirtualMachineReferenceKubeVirtVirtualMachineName("my-vm")

				Expect(ci.Status.VirtualMachineReference).ToNot(BeNil())
				Expect(ci.Status.VirtualMachineReference.KubeVirtVirtualMachineName).To(Equal("my-vm"))
			})

			It("should initialize nil reference before setting", func() {
				Expect(ci.Status.VirtualMachineReference).To(BeNil())

				ci.SetVirtualMachineReferenceKubeVirtVirtualMachineName("my-vm")

				Expect(ci.Status.VirtualMachineReference).ToNot(BeNil())
			})
		})

		Describe("GetVirtualMachineReferenceNamespace", func() {
			It("should return the namespace when set", func() {
				ci.SetVirtualMachineReferenceNamespace("tenant-ns")

				Expect(ci.GetVirtualMachineReferenceNamespace()).To(Equal("tenant-ns"))
			})

			It("should return empty string when reference is nil", func() {
				Expect(ci.Status.VirtualMachineReference).To(BeNil())
				Expect(ci.GetVirtualMachineReferenceNamespace()).To(Equal(""))
			})
		})

		Describe("GetVirtualMachineReferenceKubeVirtVirtualMachineName", func() {
			It("should return the VM name when set", func() {
				ci.SetVirtualMachineReferenceKubeVirtVirtualMachineName("my-vm")

				Expect(ci.GetVirtualMachineReferenceKubeVirtVirtualMachineName()).To(Equal("my-vm"))
			})

			It("should return empty string when reference is nil", func() {
				Expect(ci.Status.VirtualMachineReference).To(BeNil())
				Expect(ci.GetVirtualMachineReferenceKubeVirtVirtualMachineName()).To(Equal(""))
			})
		})

		Describe("EnsureVirtualMachineReference", func() {
			It("should initialize a nil reference", func() {
				Expect(ci.Status.VirtualMachineReference).To(BeNil())

				ci.EnsureVirtualMachineReference()

				Expect(ci.Status.VirtualMachineReference).ToNot(BeNil())
			})

			It("should not overwrite an existing reference", func() {
				ci.SetVirtualMachineReferenceNamespace("existing-ns")
				ci.SetVirtualMachineReferenceKubeVirtVirtualMachineName("existing-vm")

				ci.EnsureVirtualMachineReference()

				Expect(ci.Status.VirtualMachineReference.Namespace).To(Equal("existing-ns"))
				Expect(ci.Status.VirtualMachineReference.KubeVirtVirtualMachineName).To(Equal("existing-vm"))
			})
		})
	})

	Describe("TenantReference", func() {
		Describe("SetTenantReferenceName", func() {
			It("should set the tenant name", func() {
				ci.SetTenantReferenceName("my-tenant")

				Expect(ci.Status.TenantReference).ToNot(BeNil())
				Expect(ci.Status.TenantReference.Name).To(Equal("my-tenant"))
			})

			It("should initialize nil reference before setting", func() {
				Expect(ci.Status.TenantReference).To(BeNil())

				ci.SetTenantReferenceName("my-tenant")

				Expect(ci.Status.TenantReference).ToNot(BeNil())
			})
		})

		Describe("SetTenantReferenceNamespace", func() {
			It("should set the tenant namespace", func() {
				ci.SetTenantReferenceNamespace("tenant-ns")

				Expect(ci.Status.TenantReference).ToNot(BeNil())
				Expect(ci.Status.TenantReference.Namespace).To(Equal("tenant-ns"))
			})

			It("should initialize nil reference before setting", func() {
				Expect(ci.Status.TenantReference).To(BeNil())

				ci.SetTenantReferenceNamespace("tenant-ns")

				Expect(ci.Status.TenantReference).ToNot(BeNil())
			})
		})

		Describe("GetTenantReferenceName", func() {
			It("should return the tenant name when set", func() {
				ci.SetTenantReferenceName("my-tenant")

				Expect(ci.GetTenantReferenceName()).To(Equal("my-tenant"))
			})

			It("should return empty string when reference is nil", func() {
				Expect(ci.Status.TenantReference).To(BeNil())
				Expect(ci.GetTenantReferenceName()).To(Equal(""))
			})
		})

		Describe("GetTenantReferenceNamespace", func() {
			It("should return the tenant namespace when set", func() {
				ci.SetTenantReferenceNamespace("tenant-ns")

				Expect(ci.GetTenantReferenceNamespace()).To(Equal("tenant-ns"))
			})

			It("should return empty string when reference is nil", func() {
				Expect(ci.Status.TenantReference).To(BeNil())
				Expect(ci.GetTenantReferenceNamespace()).To(Equal(""))
			})
		})

		Describe("EnsureTenantReference", func() {
			It("should initialize a nil reference", func() {
				Expect(ci.Status.TenantReference).To(BeNil())

				ci.EnsureTenantReference()

				Expect(ci.Status.TenantReference).ToNot(BeNil())
			})

			It("should not overwrite an existing reference", func() {
				ci.SetTenantReferenceName("existing-tenant")
				ci.SetTenantReferenceNamespace("existing-ns")

				ci.EnsureTenantReference()

				Expect(ci.Status.TenantReference.Name).To(Equal("existing-tenant"))
				Expect(ci.Status.TenantReference.Namespace).To(Equal("existing-ns"))
			})
		})
	})

	Describe("IPAddress", func() {
		Describe("SetIPAddress", func() {
			It("should set the IP address", func() {
				ci.SetIPAddress("10.0.0.5")

				Expect(ci.Status.IPAddress).To(Equal("10.0.0.5"))
			})
		})

		Describe("GetIPAddress", func() {
			It("should return the IP address when set", func() {
				ci.SetIPAddress("192.168.1.100")

				Expect(ci.GetIPAddress()).To(Equal("192.168.1.100"))
			})

			It("should return empty string when not set", func() {
				Expect(ci.GetIPAddress()).To(Equal(""))
			})
		})
	})
})
