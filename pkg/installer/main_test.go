/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package installer

import "testing"

func Test_getBucket(t *testing.T) {
	type args struct {
		bucket string
		elem   []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "gs://ace",
			args: args{
				bucket: "gs://ace",
				elem: []string{
					"avatars",
				},
			},
			want:    "gs://ace/avatars",
			wantErr: false,
		},
		{
			name: "gs://ace/sub?q=x",
			args: args{
				bucket: "gs://ace/sub?q=x",
				elem: []string{
					"avatars",
				},
			},
			want:    "gs://ace/sub/avatars?q=x",
			wantErr: false,
		},
		{
			name: "gs://ace?q=x",
			args: args{
				bucket: "gs://ace?q=x",
				elem: []string{
					"sub",
					"avatars",
				},
			},
			want:    "gs://ace/sub/avatars?q=x",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getBucket(tt.args.bucket, tt.args.elem...)
			if (err != nil) != tt.wantErr {
				t.Errorf("getBucket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getBucket() got = %v, want %v", got, tt.want)
			}
		})
	}
}
