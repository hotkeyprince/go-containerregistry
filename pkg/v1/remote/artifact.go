// Copyright 2018 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package remote

import (
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/partial"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type descriptorArtifact struct {
	desc Descriptor
}

func (d *descriptorArtifact) Digest() (v1.Hash, error) {
	return d.desc.Digest, nil
}

func (d *descriptorArtifact) MediaType() (types.MediaType, error) {
	return d.desc.MediaType, nil
}

func (d *descriptorArtifact) RawManifest() ([]byte, error) {
	return d.desc.RawManifest()
}

func (d *descriptorArtifact) Size() (int64, error) {
	return d.desc.Size, nil
}

var _ partial.Artifact = (*descriptorArtifact)(nil)

// Get returns a partial.Artifact for the given reference.
//
// See Head if you don't need the response body.
func Artifact(ref name.Reference, options ...Option) (partial.Artifact, error) {
	o, err := makeOptions(options...)
	if err != nil {
		return nil, err
	}
	return newPuller(o).Artifact(o.context, ref)
}

// Handle options and fetch the manifest with the acceptable MediaTypes in the
// Accept header.
func artifact(ref name.Reference, acceptable []types.MediaType, options ...Option) (partial.Artifact, error) {
	o, err := makeOptions(append(options, WithAcceptableMediaTypes(acceptable))...)
	if err != nil {
		return nil, err
	}

	return newPuller(o).Artifact(o.context, ref)
}
