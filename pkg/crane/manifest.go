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

package crane

import v1 "github.com/google/go-containerregistry/pkg/v1"

// Manifest returns the manifest for the remote image or index ref.
func Manifest(ref string, opt ...Option) ([]byte, error) {
	desc, err := getArtifact(ref, opt...)
	if err != nil {
		return nil, err
	}
	o := makeOptions(opt...)
	if idx, ok := desc.(v1.ImageIndex); ok && o.Platform != nil {
		img, err := childByPlatform(idx, *o.Platform)
		if err != nil {
			return nil, err
		}
		if img, ok := img.(v1.Image); ok {
			return img.RawManifest()
		}
	}
	return desc.RawManifest()
}
