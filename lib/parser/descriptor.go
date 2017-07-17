package parser

import (
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/lycoris0731/evans/lib/model"
)

type FileDescriptorSet struct {
	*descriptor.FileDescriptorSet
}

func (d *FileDescriptorSet) GetServices() model.Services {
	// TODO: Optimization
	var services []*model.Service
	for _, f := range d.GetFile() {
		for _, proto := range f.GetService() {
			services = append(services, model.NewService(proto))
		}
	}
	return services
}