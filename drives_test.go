// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
package firecracker

import (
	"reflect"
	"testing"

	models "github.com/firecracker-microvm/firecracker-go-sdk/client/models"
)

func TestDrivesBuilder(t *testing.T) {
	expectedPath := "/path/to/rootfs"
	expectedDrives := []models.Drive{
		{
			DriveID:      new(rootDriveName),
			PathOnHost:   expectedPath,
			IsRootDevice: new(true),
			IsReadOnly:   false,
		},
	}

	drives := NewDrivesBuilder(expectedPath).Build()
	if e, a := expectedDrives, drives; !reflect.DeepEqual(e, a) {
		t.Errorf("expected drives %+v, but received %+v", e, a)
	}
}

func TestDrivesBuilderWithRootDrive(t *testing.T) {
	expectedPath := "/path/to/rootfs"
	expectedDrives := []models.Drive{
		{
			DriveID:      new("foo"),
			PathOnHost:   expectedPath,
			IsRootDevice: new(true),
			IsReadOnly:   false,
		},
	}

	b := NewDrivesBuilder(expectedPath)
	drives := b.WithRootDrive(expectedPath, WithDriveID("foo")).Build()

	if e, a := expectedDrives, drives; !reflect.DeepEqual(e, a) {
		t.Errorf("expected drives %+v, but received %+v", e, a)
	}
}

func TestDrivesBuilderWithCacheType(t *testing.T) {
	expectedPath := "/path/to/rootfs"
	expectedCacheType := models.DriveCacheTypeWriteback
	expectedDrives := []models.Drive{
		{
			DriveID:      new("root_drive"),
			PathOnHost:   expectedPath,
			IsRootDevice: new(true),
			IsReadOnly:   false,
			CacheType:    new(expectedCacheType),
		},
	}

	b := NewDrivesBuilder(expectedPath)
	drives := b.WithRootDrive(expectedPath, WithDriveID("root_drive"), WithCacheType("Writeback")).Build()

	if e, a := expectedDrives, drives; !reflect.DeepEqual(e, a) {
		t.Errorf("expected drives %+v, but received %+v", e, a)
	}
}

func TestDrivesBuilderAddDrive(t *testing.T) {
	rootPath := "/root/path"
	drivesToAdd := []struct {
		Path     string
		ReadOnly bool
		Opt      func(drive *models.Drive)
	}{
		{
			Path:     "/2",
			ReadOnly: true,
		},
		{
			Path:     "/3",
			ReadOnly: false,
		},
		{
			Path:     "/4",
			ReadOnly: false,
			Opt: func(drive *models.Drive) {
				drive.Partuuid = "uuid"
			},
		},
		{
			Path:     "/5",
			ReadOnly: true,
			Opt: func(drive *models.Drive) {
				drive.CacheType = new(models.DriveCacheTypeWriteback)
			},
		},
	}
	expectedDrives := []models.Drive{
		{
			DriveID:      new("0"),
			PathOnHost:   "/2",
			IsRootDevice: new(false),
			IsReadOnly:   true,
		},
		{
			DriveID:      new("1"),
			PathOnHost:   "/3",
			IsRootDevice: new(false),
			IsReadOnly:   false,
		},
		{
			DriveID:      new("2"),
			PathOnHost:   "/4",
			IsRootDevice: new(false),
			IsReadOnly:   false,
			Partuuid:     "uuid",
		},
		{
			DriveID:      new("3"),
			PathOnHost:   "/5",
			IsRootDevice: new(false),
			IsReadOnly:   true,
			CacheType:    new(models.DriveCacheTypeWriteback),
		},
		{
			DriveID:      new(rootDriveName),
			PathOnHost:   rootPath,
			IsRootDevice: new(true),
			IsReadOnly:   false,
		},
	}

	b := NewDrivesBuilder(rootPath)

	for _, drive := range drivesToAdd {
		if drive.Opt != nil {
			b = b.AddDrive(drive.Path, drive.ReadOnly, drive.Opt)
		} else {
			b = b.AddDrive(drive.Path, drive.ReadOnly)
		}
	}

	drives := b.Build()
	if e, a := expectedDrives, drives; !reflect.DeepEqual(e, a) {
		t.Errorf("expected drives %+v\n, but received %+v", e, a)
	}
}

func TestDrivesBuilderWithIoEngine(t *testing.T) {
	expectedPath := "/path/to/rootfs"
	expectedVal := "Async"
	expectedDrives := []models.Drive{
		{
			DriveID:      new(rootDriveName),
			PathOnHost:   expectedPath,
			IsRootDevice: new(true),
			IsReadOnly:   false,
			IoEngine:     &expectedVal,
		},
	}

	drives := NewDrivesBuilder(expectedPath).WithRootDrive(expectedPath,
		WithDriveID(string(rootDriveName)), WithIoEngine(expectedVal)).Build()
	if e, a := expectedDrives, drives; !reflect.DeepEqual(e, a) {
		t.Errorf("expected drives %+v, but received %+v", e, a)
	}
}
