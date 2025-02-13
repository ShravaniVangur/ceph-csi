/*
Copyright 2021 The Ceph-CSI Authors.

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

package core

import (
	"testing"

	cerrors "github.com/ceph/ceph-csi/internal/cephfs/errors"

	fsa "github.com/ceph/go-ceph/cephfs/admin"
	"github.com/stretchr/testify/require"
)

func TestCloneStateToError(t *testing.T) {
	t.Parallel()
	errorState := make(map[cephFSCloneState]error)
	errorState[cephFSCloneState{fsa.CloneComplete, fsa.CloneProgressReport{}, "", ""}] = nil
	errorState[*CephFSCloneError] = cerrors.ErrInvalidClone
	errorState[cephFSCloneState{fsa.CloneInProgress, fsa.CloneProgressReport{}, "", ""}] = cerrors.ErrCloneInProgress
	errorState[cephFSCloneState{fsa.ClonePending, fsa.CloneProgressReport{}, "", ""}] = cerrors.ErrClonePending
	errorState[cephFSCloneState{fsa.CloneFailed, fsa.CloneProgressReport{}, "", ""}] = cerrors.ErrCloneFailed

	for state, err := range errorState {
		require.ErrorIs(t, state.ToError(), err)
	}
}
