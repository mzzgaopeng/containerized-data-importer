/*
Copyright 2018 The CDI Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "kubevirt.io/containerized-data-importer/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// CDIs returns a CDIInformer.
	CDIs() CDIInformer
	// CDIConfigs returns a CDIConfigInformer.
	CDIConfigs() CDIConfigInformer
	// DataImportCrons returns a DataImportCronInformer.
	DataImportCrons() DataImportCronInformer
	// DataSources returns a DataSourceInformer.
	DataSources() DataSourceInformer
	// DataVolumes returns a DataVolumeInformer.
	DataVolumes() DataVolumeInformer
	// ObjectTransfers returns a ObjectTransferInformer.
	ObjectTransfers() ObjectTransferInformer
	// StorageProfiles returns a StorageProfileInformer.
	StorageProfiles() StorageProfileInformer
	// VolumeCloneSources returns a VolumeCloneSourceInformer.
	VolumeCloneSources() VolumeCloneSourceInformer
	// VolumeImportSources returns a VolumeImportSourceInformer.
	VolumeImportSources() VolumeImportSourceInformer
	// VolumeUploadSources returns a VolumeUploadSourceInformer.
	VolumeUploadSources() VolumeUploadSourceInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// CDIs returns a CDIInformer.
func (v *version) CDIs() CDIInformer {
	return &cDIInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// CDIConfigs returns a CDIConfigInformer.
func (v *version) CDIConfigs() CDIConfigInformer {
	return &cDIConfigInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// DataImportCrons returns a DataImportCronInformer.
func (v *version) DataImportCrons() DataImportCronInformer {
	return &dataImportCronInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DataSources returns a DataSourceInformer.
func (v *version) DataSources() DataSourceInformer {
	return &dataSourceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DataVolumes returns a DataVolumeInformer.
func (v *version) DataVolumes() DataVolumeInformer {
	return &dataVolumeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ObjectTransfers returns a ObjectTransferInformer.
func (v *version) ObjectTransfers() ObjectTransferInformer {
	return &objectTransferInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// StorageProfiles returns a StorageProfileInformer.
func (v *version) StorageProfiles() StorageProfileInformer {
	return &storageProfileInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// VolumeCloneSources returns a VolumeCloneSourceInformer.
func (v *version) VolumeCloneSources() VolumeCloneSourceInformer {
	return &volumeCloneSourceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VolumeImportSources returns a VolumeImportSourceInformer.
func (v *version) VolumeImportSources() VolumeImportSourceInformer {
	return &volumeImportSourceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VolumeUploadSources returns a VolumeUploadSourceInformer.
func (v *version) VolumeUploadSources() VolumeUploadSourceInformer {
	return &volumeUploadSourceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
