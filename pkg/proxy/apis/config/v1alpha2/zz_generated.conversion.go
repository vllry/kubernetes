// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha2

import (
	unsafe "unsafe"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1alpha2 "k8s.io/kube-proxy/config/v1alpha2"
	config "k8s.io/kubernetes/pkg/proxy/apis/config"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1alpha2.KubeProxyConfiguration)(nil), (*config.KubeProxyConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeProxyConfiguration_To_config_KubeProxyConfiguration(a.(*v1alpha2.KubeProxyConfiguration), b.(*config.KubeProxyConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeProxyConfiguration)(nil), (*v1alpha2.KubeProxyConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeProxyConfiguration_To_v1alpha2_KubeProxyConfiguration(a.(*config.KubeProxyConfiguration), b.(*v1alpha2.KubeProxyConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha2.KubeProxyConntrackConfiguration)(nil), (*config.KubeProxyConntrackConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeProxyConntrackConfiguration_To_config_KubeProxyConntrackConfiguration(a.(*v1alpha2.KubeProxyConntrackConfiguration), b.(*config.KubeProxyConntrackConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeProxyConntrackConfiguration)(nil), (*v1alpha2.KubeProxyConntrackConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeProxyConntrackConfiguration_To_v1alpha2_KubeProxyConntrackConfiguration(a.(*config.KubeProxyConntrackConfiguration), b.(*v1alpha2.KubeProxyConntrackConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha2.KubeProxyIPTablesConfiguration)(nil), (*config.KubeProxyIPTablesConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeProxyIPTablesConfiguration_To_config_KubeProxyIPTablesConfiguration(a.(*v1alpha2.KubeProxyIPTablesConfiguration), b.(*config.KubeProxyIPTablesConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeProxyIPTablesConfiguration)(nil), (*v1alpha2.KubeProxyIPTablesConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeProxyIPTablesConfiguration_To_v1alpha2_KubeProxyIPTablesConfiguration(a.(*config.KubeProxyIPTablesConfiguration), b.(*v1alpha2.KubeProxyIPTablesConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha2.KubeProxyIPVSConfiguration)(nil), (*config.KubeProxyIPVSConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeProxyIPVSConfiguration_To_config_KubeProxyIPVSConfiguration(a.(*v1alpha2.KubeProxyIPVSConfiguration), b.(*config.KubeProxyIPVSConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeProxyIPVSConfiguration)(nil), (*v1alpha2.KubeProxyIPVSConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeProxyIPVSConfiguration_To_v1alpha2_KubeProxyIPVSConfiguration(a.(*config.KubeProxyIPVSConfiguration), b.(*v1alpha2.KubeProxyIPVSConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha2.KubeProxyWinkernelConfiguration)(nil), (*config.KubeProxyWinkernelConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeProxyWinkernelConfiguration_To_config_KubeProxyWinkernelConfiguration(a.(*v1alpha2.KubeProxyWinkernelConfiguration), b.(*config.KubeProxyWinkernelConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeProxyWinkernelConfiguration)(nil), (*v1alpha2.KubeProxyWinkernelConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeProxyWinkernelConfiguration_To_v1alpha2_KubeProxyWinkernelConfiguration(a.(*config.KubeProxyWinkernelConfiguration), b.(*v1alpha2.KubeProxyWinkernelConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha2_KubeProxyConfiguration_To_config_KubeProxyConfiguration(in *v1alpha2.KubeProxyConfiguration, out *config.KubeProxyConfiguration, s conversion.Scope) error {
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	out.BindAddress = *(*[]string)(unsafe.Pointer(&in.BindAddress))
	out.HealthzBindAddress = in.HealthzBindAddress
	out.MetricsBindAddress = in.MetricsBindAddress
	out.EnableProfiling = in.EnableProfiling
	out.ClusterCIDR = *(*[]string)(unsafe.Pointer(&in.ClusterCIDR))
	out.HostnameOverride = in.HostnameOverride
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ClientConnection, &out.ClientConnection, 0); err != nil {
		return err
	}
	if err := Convert_v1alpha2_KubeProxyIPTablesConfiguration_To_config_KubeProxyIPTablesConfiguration(&in.IPTables, &out.IPTables, s); err != nil {
		return err
	}
	if err := Convert_v1alpha2_KubeProxyIPVSConfiguration_To_config_KubeProxyIPVSConfiguration(&in.IPVS, &out.IPVS, s); err != nil {
		return err
	}
	out.OOMScoreAdj = (*int32)(unsafe.Pointer(in.OOMScoreAdj))
	out.Mode = config.ProxyMode(in.Mode)
	out.PortRange = in.PortRange
	out.UDPIdleTimeout = in.UDPIdleTimeout
	if err := Convert_v1alpha2_KubeProxyConntrackConfiguration_To_config_KubeProxyConntrackConfiguration(&in.Conntrack, &out.Conntrack, s); err != nil {
		return err
	}
	out.ConfigSyncPeriod = in.ConfigSyncPeriod
	out.NodePortAddresses = *(*[]string)(unsafe.Pointer(&in.NodePortAddresses))
	if err := Convert_v1alpha2_KubeProxyWinkernelConfiguration_To_config_KubeProxyWinkernelConfiguration(&in.Winkernel, &out.Winkernel, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_KubeProxyConfiguration_To_config_KubeProxyConfiguration is an autogenerated conversion function.
func Convert_v1alpha2_KubeProxyConfiguration_To_config_KubeProxyConfiguration(in *v1alpha2.KubeProxyConfiguration, out *config.KubeProxyConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha2_KubeProxyConfiguration_To_config_KubeProxyConfiguration(in, out, s)
}

func autoConvert_config_KubeProxyConfiguration_To_v1alpha2_KubeProxyConfiguration(in *config.KubeProxyConfiguration, out *v1alpha2.KubeProxyConfiguration, s conversion.Scope) error {
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	out.BindAddress = *(*[]string)(unsafe.Pointer(&in.BindAddress))
	out.HealthzBindAddress = in.HealthzBindAddress
	out.MetricsBindAddress = in.MetricsBindAddress
	out.EnableProfiling = in.EnableProfiling
	out.ClusterCIDR = *(*[]string)(unsafe.Pointer(&in.ClusterCIDR))
	out.HostnameOverride = in.HostnameOverride
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ClientConnection, &out.ClientConnection, 0); err != nil {
		return err
	}
	if err := Convert_config_KubeProxyIPTablesConfiguration_To_v1alpha2_KubeProxyIPTablesConfiguration(&in.IPTables, &out.IPTables, s); err != nil {
		return err
	}
	if err := Convert_config_KubeProxyIPVSConfiguration_To_v1alpha2_KubeProxyIPVSConfiguration(&in.IPVS, &out.IPVS, s); err != nil {
		return err
	}
	out.OOMScoreAdj = (*int32)(unsafe.Pointer(in.OOMScoreAdj))
	out.Mode = v1alpha2.ProxyMode(in.Mode)
	out.PortRange = in.PortRange
	out.UDPIdleTimeout = in.UDPIdleTimeout
	if err := Convert_config_KubeProxyConntrackConfiguration_To_v1alpha2_KubeProxyConntrackConfiguration(&in.Conntrack, &out.Conntrack, s); err != nil {
		return err
	}
	out.ConfigSyncPeriod = in.ConfigSyncPeriod
	out.NodePortAddresses = *(*[]string)(unsafe.Pointer(&in.NodePortAddresses))
	if err := Convert_config_KubeProxyWinkernelConfiguration_To_v1alpha2_KubeProxyWinkernelConfiguration(&in.Winkernel, &out.Winkernel, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_KubeProxyConfiguration_To_v1alpha2_KubeProxyConfiguration is an autogenerated conversion function.
func Convert_config_KubeProxyConfiguration_To_v1alpha2_KubeProxyConfiguration(in *config.KubeProxyConfiguration, out *v1alpha2.KubeProxyConfiguration, s conversion.Scope) error {
	return autoConvert_config_KubeProxyConfiguration_To_v1alpha2_KubeProxyConfiguration(in, out, s)
}

func autoConvert_v1alpha2_KubeProxyConntrackConfiguration_To_config_KubeProxyConntrackConfiguration(in *v1alpha2.KubeProxyConntrackConfiguration, out *config.KubeProxyConntrackConfiguration, s conversion.Scope) error {
	out.MaxPerCore = (*int32)(unsafe.Pointer(in.MaxPerCore))
	out.Min = (*int32)(unsafe.Pointer(in.Min))
	out.TCPEstablishedTimeout = (*v1.Duration)(unsafe.Pointer(in.TCPEstablishedTimeout))
	out.TCPCloseWaitTimeout = (*v1.Duration)(unsafe.Pointer(in.TCPCloseWaitTimeout))
	return nil
}

// Convert_v1alpha2_KubeProxyConntrackConfiguration_To_config_KubeProxyConntrackConfiguration is an autogenerated conversion function.
func Convert_v1alpha2_KubeProxyConntrackConfiguration_To_config_KubeProxyConntrackConfiguration(in *v1alpha2.KubeProxyConntrackConfiguration, out *config.KubeProxyConntrackConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha2_KubeProxyConntrackConfiguration_To_config_KubeProxyConntrackConfiguration(in, out, s)
}

func autoConvert_config_KubeProxyConntrackConfiguration_To_v1alpha2_KubeProxyConntrackConfiguration(in *config.KubeProxyConntrackConfiguration, out *v1alpha2.KubeProxyConntrackConfiguration, s conversion.Scope) error {
	out.MaxPerCore = (*int32)(unsafe.Pointer(in.MaxPerCore))
	out.Min = (*int32)(unsafe.Pointer(in.Min))
	out.TCPEstablishedTimeout = (*v1.Duration)(unsafe.Pointer(in.TCPEstablishedTimeout))
	out.TCPCloseWaitTimeout = (*v1.Duration)(unsafe.Pointer(in.TCPCloseWaitTimeout))
	return nil
}

// Convert_config_KubeProxyConntrackConfiguration_To_v1alpha2_KubeProxyConntrackConfiguration is an autogenerated conversion function.
func Convert_config_KubeProxyConntrackConfiguration_To_v1alpha2_KubeProxyConntrackConfiguration(in *config.KubeProxyConntrackConfiguration, out *v1alpha2.KubeProxyConntrackConfiguration, s conversion.Scope) error {
	return autoConvert_config_KubeProxyConntrackConfiguration_To_v1alpha2_KubeProxyConntrackConfiguration(in, out, s)
}

func autoConvert_v1alpha2_KubeProxyIPTablesConfiguration_To_config_KubeProxyIPTablesConfiguration(in *v1alpha2.KubeProxyIPTablesConfiguration, out *config.KubeProxyIPTablesConfiguration, s conversion.Scope) error {
	out.MasqueradeBit = (*int32)(unsafe.Pointer(in.MasqueradeBit))
	out.MasqueradeAll = in.MasqueradeAll
	out.SyncPeriod = in.SyncPeriod
	out.MinSyncPeriod = in.MinSyncPeriod
	return nil
}

// Convert_v1alpha2_KubeProxyIPTablesConfiguration_To_config_KubeProxyIPTablesConfiguration is an autogenerated conversion function.
func Convert_v1alpha2_KubeProxyIPTablesConfiguration_To_config_KubeProxyIPTablesConfiguration(in *v1alpha2.KubeProxyIPTablesConfiguration, out *config.KubeProxyIPTablesConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha2_KubeProxyIPTablesConfiguration_To_config_KubeProxyIPTablesConfiguration(in, out, s)
}

func autoConvert_config_KubeProxyIPTablesConfiguration_To_v1alpha2_KubeProxyIPTablesConfiguration(in *config.KubeProxyIPTablesConfiguration, out *v1alpha2.KubeProxyIPTablesConfiguration, s conversion.Scope) error {
	out.MasqueradeBit = (*int32)(unsafe.Pointer(in.MasqueradeBit))
	out.MasqueradeAll = in.MasqueradeAll
	out.SyncPeriod = in.SyncPeriod
	out.MinSyncPeriod = in.MinSyncPeriod
	return nil
}

// Convert_config_KubeProxyIPTablesConfiguration_To_v1alpha2_KubeProxyIPTablesConfiguration is an autogenerated conversion function.
func Convert_config_KubeProxyIPTablesConfiguration_To_v1alpha2_KubeProxyIPTablesConfiguration(in *config.KubeProxyIPTablesConfiguration, out *v1alpha2.KubeProxyIPTablesConfiguration, s conversion.Scope) error {
	return autoConvert_config_KubeProxyIPTablesConfiguration_To_v1alpha2_KubeProxyIPTablesConfiguration(in, out, s)
}

func autoConvert_v1alpha2_KubeProxyIPVSConfiguration_To_config_KubeProxyIPVSConfiguration(in *v1alpha2.KubeProxyIPVSConfiguration, out *config.KubeProxyIPVSConfiguration, s conversion.Scope) error {
	out.SyncPeriod = in.SyncPeriod
	out.MinSyncPeriod = in.MinSyncPeriod
	out.Scheduler = in.Scheduler
	out.ExcludeCIDRs = *(*[]string)(unsafe.Pointer(&in.ExcludeCIDRs))
	out.StrictARP = in.StrictARP
	return nil
}

// Convert_v1alpha2_KubeProxyIPVSConfiguration_To_config_KubeProxyIPVSConfiguration is an autogenerated conversion function.
func Convert_v1alpha2_KubeProxyIPVSConfiguration_To_config_KubeProxyIPVSConfiguration(in *v1alpha2.KubeProxyIPVSConfiguration, out *config.KubeProxyIPVSConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha2_KubeProxyIPVSConfiguration_To_config_KubeProxyIPVSConfiguration(in, out, s)
}

func autoConvert_config_KubeProxyIPVSConfiguration_To_v1alpha2_KubeProxyIPVSConfiguration(in *config.KubeProxyIPVSConfiguration, out *v1alpha2.KubeProxyIPVSConfiguration, s conversion.Scope) error {
	out.SyncPeriod = in.SyncPeriod
	out.MinSyncPeriod = in.MinSyncPeriod
	out.Scheduler = in.Scheduler
	out.ExcludeCIDRs = *(*[]string)(unsafe.Pointer(&in.ExcludeCIDRs))
	out.StrictARP = in.StrictARP
	return nil
}

// Convert_config_KubeProxyIPVSConfiguration_To_v1alpha2_KubeProxyIPVSConfiguration is an autogenerated conversion function.
func Convert_config_KubeProxyIPVSConfiguration_To_v1alpha2_KubeProxyIPVSConfiguration(in *config.KubeProxyIPVSConfiguration, out *v1alpha2.KubeProxyIPVSConfiguration, s conversion.Scope) error {
	return autoConvert_config_KubeProxyIPVSConfiguration_To_v1alpha2_KubeProxyIPVSConfiguration(in, out, s)
}

func autoConvert_v1alpha2_KubeProxyWinkernelConfiguration_To_config_KubeProxyWinkernelConfiguration(in *v1alpha2.KubeProxyWinkernelConfiguration, out *config.KubeProxyWinkernelConfiguration, s conversion.Scope) error {
	out.NetworkName = in.NetworkName
	out.SourceVip = in.SourceVip
	out.EnableDSR = in.EnableDSR
	return nil
}

// Convert_v1alpha2_KubeProxyWinkernelConfiguration_To_config_KubeProxyWinkernelConfiguration is an autogenerated conversion function.
func Convert_v1alpha2_KubeProxyWinkernelConfiguration_To_config_KubeProxyWinkernelConfiguration(in *v1alpha2.KubeProxyWinkernelConfiguration, out *config.KubeProxyWinkernelConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha2_KubeProxyWinkernelConfiguration_To_config_KubeProxyWinkernelConfiguration(in, out, s)
}

func autoConvert_config_KubeProxyWinkernelConfiguration_To_v1alpha2_KubeProxyWinkernelConfiguration(in *config.KubeProxyWinkernelConfiguration, out *v1alpha2.KubeProxyWinkernelConfiguration, s conversion.Scope) error {
	out.NetworkName = in.NetworkName
	out.SourceVip = in.SourceVip
	out.EnableDSR = in.EnableDSR
	return nil
}

// Convert_config_KubeProxyWinkernelConfiguration_To_v1alpha2_KubeProxyWinkernelConfiguration is an autogenerated conversion function.
func Convert_config_KubeProxyWinkernelConfiguration_To_v1alpha2_KubeProxyWinkernelConfiguration(in *config.KubeProxyWinkernelConfiguration, out *v1alpha2.KubeProxyWinkernelConfiguration, s conversion.Scope) error {
	return autoConvert_config_KubeProxyWinkernelConfiguration_To_v1alpha2_KubeProxyWinkernelConfiguration(in, out, s)
}
