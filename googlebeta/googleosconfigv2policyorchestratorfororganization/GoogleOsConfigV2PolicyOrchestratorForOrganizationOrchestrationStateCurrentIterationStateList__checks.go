// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package googleosconfigv2policyorchestratorfororganization

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (g *jsiiProxy_GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewGoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

