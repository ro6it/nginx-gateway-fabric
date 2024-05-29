package policies_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nginxinc/nginx-gateway-fabric/internal/framework/conditions"
	"github.com/nginxinc/nginx-gateway-fabric/internal/mode/static/policies"
	"github.com/nginxinc/nginx-gateway-fabric/internal/mode/static/policies/policiesfakes"
	staticConds "github.com/nginxinc/nginx-gateway-fabric/internal/mode/static/state/conditions"
)

var _ = Describe("Policy Manager", func() {
	orangeGVK := schema.GroupVersionKind{Group: "fruit", Version: "1", Kind: "orange"}
	orangePolicy := &policiesfakes.FakePolicy{
		GetNameStub: func() string {
			return "orange"
		},
	}

	appleGVK := schema.GroupVersionKind{Group: "fruit", Version: "1", Kind: "apple"}
	applePolicy := &policiesfakes.FakePolicy{
		GetNameStub: func() string {
			return "apple"
		},
	}

	mustExtractGVK := func(object client.Object) schema.GroupVersionKind {
		switch object.GetName() {
		case "apple":
			return appleGVK
		case "orange":
			return orangeGVK
		default:
			return schema.GroupVersionKind{}
		}
	}

	mgr := policies.NewManager(
		mustExtractGVK,
		policies.ManagerConfig{
			Validator: &policiesfakes.FakeValidator{
				ValidateStub: func(_ policies.Policy, _ *policies.GlobalSettings) []conditions.Condition {
					return []conditions.Condition{staticConds.NewPolicyInvalid("apple error")}
				},
				ConflictsStub: func(_ policies.Policy, _ policies.Policy) bool { return true },
			},
			Generator: func(_ policies.Policy, _ *policies.GlobalSettings) []byte {
				return []byte("apple")
			},
			GVK: appleGVK,
		},
		policies.ManagerConfig{
			Validator: &policiesfakes.FakeValidator{
				ValidateStub: func(_ policies.Policy, _ *policies.GlobalSettings) []conditions.Condition {
					return []conditions.Condition{staticConds.NewPolicyInvalid("orange error")}
				},
				ConflictsStub: func(_ policies.Policy, _ policies.Policy) bool { return false },
			},
			Generator: func(_ policies.Policy, _ *policies.GlobalSettings) []byte {
				return []byte("orange")
			},
			GVK: orangeGVK,
		},
	)

	Context("Validation", func() {
		When("Policy is registered with manager", func() {
			It("Validates the policy", func() {
				conds := mgr.Validate(applePolicy, nil)
				Expect(conds).To(HaveLen(1))
				Expect(conds[0].Message).To(Equal("apple error"))

				conds = mgr.Validate(orangePolicy, nil)
				Expect(conds).To(HaveLen(1))
				Expect(conds[0].Message).To(Equal("orange error"))
			})
			It("Returns whether the policies conflict", func() {
				Expect(mgr.Conflicts(applePolicy, applePolicy)).To(BeTrue())
				Expect(mgr.Conflicts(orangePolicy, orangePolicy)).To(BeFalse())
			})
		})
		When("Policy is not registered with manager", func() {
			It("Panics on call to validate", func() {
				validate := func() {
					_ = mgr.Validate(&policiesfakes.FakePolicy{}, nil)
				}

				Expect(validate).To(Panic())
			})
			It("panics on call to conflicts", func() {
				conflict := func() {
					_ = mgr.Conflicts(&policiesfakes.FakePolicy{}, &policiesfakes.FakePolicy{})
				}

				Expect(conflict).To(Panic())
			})
		})
	})
	Context("Generation", func() {
		When("Policy is registered with manager", func() {
			It("Generates the configuration for the policy", func() {
				Expect(mgr.Generate(applePolicy, nil)).To(Equal([]byte("apple")))
				Expect(mgr.Generate(orangePolicy, nil)).To(Equal([]byte("orange")))
			})
		})
		When("Policy is not registered with manager", func() {
			It("Panics on generate", func() {
				generate := func() {
					_ = mgr.Generate(&policiesfakes.FakePolicy{}, nil)
				}

				Expect(generate).To(Panic())
			})
		})
	})
})
