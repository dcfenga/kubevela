/*
 Copyright 2022 The KubeVela Authors.

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

package version

import (
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
)

func TestIsOfficialKubeVelaVersion(t *testing.T) {
	assert.Equal(t, true, IsOfficialKubeVelaVersion("v1.2.3"))
	assert.Equal(t, true, IsOfficialKubeVelaVersion("1.2.3"))
	assert.Equal(t, true, IsOfficialKubeVelaVersion("v1.2"))
	assert.Equal(t, true, IsOfficialKubeVelaVersion("v1.2+myvela"))
	assert.Equal(t, false, IsOfficialKubeVelaVersion("v1.-"))
}

func TestGetVersion(t *testing.T) {
	version, err := GetOfficialKubeVelaVersion("v1.2.90")
	assert.Nil(t, err)
	assert.Equal(t, "1.2.90", version)
	version, err = GetOfficialKubeVelaVersion("1.2.90")
	assert.Nil(t, err)
	assert.Equal(t, "1.2.90", version)
	version, err = GetOfficialKubeVelaVersion("v1.2+myvela")
	assert.Nil(t, err)
	assert.Equal(t, "1.2.0", version)
}

func TestShouldUseLegacyHelmRepo(t *testing.T) {
	tests := []struct {
		ver  string
		want bool
	}{
		{
			ver:  "v1.2.0",
			want: true,
		},
		{
			ver:  "v1.9.0-beta.1",
			want: true,
		},
		{
			ver:  "v1.9.0-beta.1.post1",
			want: false,
		},
		{
			ver:  "v1.9.1",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.ver, func(t *testing.T) {
			ver := version.Must(version.NewVersion(tt.ver))
			assert.Equalf(t, tt.want, ShouldUseLegacyHelmRepo(ver), "ShouldUseLegacyHelmRepo(%v)", tt.ver)
		})
	}
}
