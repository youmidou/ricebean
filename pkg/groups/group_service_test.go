// Copyright (c) TFG Co. All Rights Reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package groups

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"ricebean/pkg/constants"
)

func testCreateDuplicatedGroup(gs GroupService, t *testing.T) {
	ctx := context.Background()
	//t.Parallel()
	err := gs.GroupCreate(ctx, "testCreateDuplicatedGroup")
	assert.NoError(t, err)
	err = gs.GroupCreate(ctx, "testCreateDuplicatedGroup")
	assert.Error(t, err)
	assert.Equal(t, constants.ErrGroupAlreadyExists, err)
}

func testCreateGroup(gs GroupService, t *testing.T) {
	ctx := context.Background()
	err := gs.GroupCreate(ctx, "testCreateGroup")
	assert.NoError(t, err)
	count, err := gs.GroupCountMembers(ctx, "testCreateGroup")
	assert.NoError(t, err)
	assert.Equal(t, 0, count)
	err = gs.GroupRenewTTL(ctx, "testCreateGroup")
	assert.Error(t, err)
}

func testCreateGroupWithTTL(gs GroupService, t *testing.T) {
	ctx := context.Background()
	err := gs.GroupCreateWithTTL(ctx, "testCreateGroupWithTTL", 10)
	assert.NoError(t, err)
	count, err := gs.GroupCountMembers(ctx, "testCreateGroupWithTTL")
	assert.NoError(t, err)
	assert.Equal(t, 0, count)
	err = gs.GroupRenewTTL(ctx, "testCreateGroupWithTTL")
	assert.NoError(t, err)
}

func testGroupAddMember(gs GroupService, t *testing.T) {
	ctx := context.Background()
	tables := []struct {
		name     string
		frontend bool
		UID      string
	}{
		{"frontend_uid", true, "someuid1"},
		{"backend_uid", false, "ola1"},
	}

	err := gs.GroupCreate(ctx, "testGroupAddMember")
	assert.NoError(t, err)

	for _, table := range tables {
		t.Run(table.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			err = gs.GroupAddMember(ctx, "testGroupAddMember", table.UID)
			assert.NoError(t, err)
			_, err := gs.GroupContainsMember(ctx, "testGroupAddMember", table.UID)
			assert.NoError(t, err)
		})
	}
}

func testGroupAddDuplicatedMember(gs GroupService, t *testing.T) {
	ctx := context.Background()
	err := gs.GroupCreate(ctx, "testGroupAddDuplicatedMember")
	assert.NoError(t, err)
	err = gs.GroupAddMember(ctx, "testGroupAddDuplicatedMember", "duplicatedUid")
	assert.NoError(t, err)
	err = gs.GroupAddMember(ctx, "testGroupAddDuplicatedMember", "duplicatedUid")
	assert.Error(t, err)
	assert.Equal(t, constants.ErrMemberAlreadyExists, err)
}

func testGroupContainsMember(gs GroupService, t *testing.T) {
	ctx := context.Background()
	tables := []struct {
		name     string
		frontend bool
		UID      string
	}{
		{"frontend_uid", true, "someuid2"},
		{"backend_uid", false, "ola2"},
	}

	err := gs.GroupCreate(ctx, "testGroupContainsMember")
	assert.NoError(t, err)

	for _, table := range tables {
		t.Run(table.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			err = gs.GroupAddMember(ctx, "testGroupContainsMember", table.UID)
			assert.NoError(t, err)
			b, err := gs.GroupContainsMember(ctx, "testGroupContainsMember", table.UID)
			assert.True(t, b)
			assert.NoError(t, err)
		})
	}
}

func testRemove(gs GroupService, t *testing.T) {
	ctx := context.Background()
	tables := []struct {
		name     string
		frontend bool
		UID      string
		err      error
	}{
		{"frontend_uid", true, "someuid2", nil},
		{"backend_uid", false, "ola2", nil},
	}

	err := gs.GroupCreate(ctx, "testRemove")
	assert.NoError(t, err)

	for _, table := range tables {
		t.Run(table.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			err = gs.GroupAddMember(ctx, "testRemove", table.UID)
			assert.NoError(t, err)
			err = gs.GroupRemoveMember(ctx, "testRemove", table.UID)
			assert.NoError(t, err)
			res, err := gs.GroupContainsMember(ctx, "testRemove", table.UID)
			assert.NoError(t, err)
			assert.False(t, res)
		})
	}
}

func testDelete(gs GroupService, t *testing.T) {
	ctx := context.Background()
	tables := []struct {
		name     string
		frontend bool
		UID      string
		err      error
	}{
		{"frontend_uid", true, "leaveSomeuid2", nil},
		{"backend_uid", false, "leaveOla2", nil},
	}

	err := gs.GroupCreate(ctx, "testDeleteSufix")
	assert.NoError(t, err)

	for _, table := range tables {
		t.Run(table.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			err = gs.GroupCreate(ctx, "testDelete")
			assert.NoError(t, err)
			err = gs.GroupAddMember(ctx, "testDeleteSufix", table.UID)
			assert.NoError(t, err)
			err = gs.GroupAddMember(ctx, "testDelete", table.UID)
			assert.NoError(t, err)
			err = gs.GroupDelete(ctx, "testDelete")
			assert.NoError(t, err)

			res, err := gs.GroupContainsMember(ctx, "testDeleteSufix", table.UID)
			assert.NoError(t, err)
			assert.True(t, res)
			_, err = gs.GroupContainsMember(ctx, "testDelete", table.UID)
			assert.Error(t, err)
			assert.Equal(t, constants.ErrGroupNotFound, err)
		})
	}
}

func testRemoveAll(gs GroupService, t *testing.T) {
	ctx := context.Background()
	tables := []struct {
		name     string
		frontend bool
		UID      string
		err      error
	}{
		{"frontend_uid", true, "removeSomeuid2", nil},
		{"backend_uid", false, "removeOla2", nil},
	}

	err := gs.GroupCreate(ctx, "testRemoveAllSufix")
	assert.NoError(t, err)
	err = gs.GroupCreate(ctx, "testRemoveAll")
	assert.NoError(t, err)

	for _, table := range tables {
		t.Run(table.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			err = gs.GroupAddMember(ctx, "testRemoveAllSufix", table.UID)
			assert.NoError(t, err)
			err = gs.GroupAddMember(ctx, "testRemoveAll", table.UID)
			assert.NoError(t, err)
			err = gs.GroupRemoveAll(ctx, "testRemoveAll")
			assert.NoError(t, err)

			res, err := gs.GroupContainsMember(ctx, "testRemoveAllSufix", table.UID)
			assert.NoError(t, err)
			assert.True(t, res)
			res, err = gs.GroupContainsMember(ctx, "testRemoveAll", table.UID)
			assert.NoError(t, err)
			assert.False(t, res)
		})
	}
}

func testCount(gs GroupService, t *testing.T) {
	ctx := context.Background()
	tables := []struct {
		name     string
		frontend bool
		UID      string
		err      error
	}{
		{"frontend_uid", true, "someuid2", nil},
		{"backend_uid", false, "ola2", nil},
	}

	err := gs.GroupCreate(ctx, "testCount")
	assert.NoError(t, err)

	for _, table := range tables {
		t.Run(table.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			err = gs.GroupAddMember(ctx, "testCount", table.UID)
			assert.NoError(t, err)
			res, err := gs.GroupCountMembers(ctx, "testCount")
			assert.NoError(t, err)
			assert.Equal(t, 1, res)

			err = gs.GroupRemoveAll(ctx, "testCount")
			assert.NoError(t, err)
		})
	}
}

func testMembers(gs GroupService, t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	err := gs.GroupCreate(ctx, "testGroupMembers")
	assert.NoError(t, err)
	uid1 := "someid1"
	uid2 := "someid2"
	err = gs.GroupAddMember(ctx, "testGroupMembers", uid1)
	assert.NoError(t, err)
	err = gs.GroupAddMember(ctx, "testGroupMembers", uid2)
	assert.NoError(t, err)

	res, err := gs.GroupMembers(ctx, "testGroupMembers")
	assert.NoError(t, err)
	assert.ElementsMatch(t, []string{"someid1", "someid2"}, res)
}
