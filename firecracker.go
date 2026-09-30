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
	"context"
	"time"

	"github.com/go-openapi/strfmt"

	"github.com/sirupsen/logrus"

	"github.com/firecracker-microvm/firecracker-go-sdk/client"
	models "github.com/firecracker-microvm/firecracker-go-sdk/client/models"
	ops "github.com/firecracker-microvm/firecracker-go-sdk/client/operations"
)

const (
	// env name to make firecracker request timeout configurable
	firecrackerRequestTimeoutEnv = "FIRECRACKER_GO_SDK_REQUEST_TIMEOUT_MILLISECONDS"

	defaultFirecrackerRequestTimeout = 500
)

// newFirecrackerClient creates a FirecrackerClient
func newFirecrackerClient(socketPath string, logger *logrus.Entry, debug bool) *client.FirecrackerAPI {
	httpClient := client.NewHTTPClient(strfmt.NewFormats())

	transport := NewUnixSocketTransport(socketPath, logger, debug)
	httpClient.SetTransport(transport)

	return httpClient
}

// ClientOpt is a functional option used to modify the client after construction.
type ClientOpt func(*Client)

// WithOpsClient will return a functional option and replace the operations
// client. This is useful for mock and stub testing.
func WithOpsClient(opsClient ops.ClientService) ClientOpt {
	return func(c *Client) {
		c.client.Operations = opsClient
	}
}

// Client is a client for interacting with the Firecracker API
type Client struct {
	client                    *client.FirecrackerAPI
	firecrackerRequestTimeout int
	firecrackerInitTimeout    int
}

// NewClient creates a Client
func NewClient(socketPath string, logger *logrus.Entry, debug bool, opts ...ClientOpt) *Client {
	httpClient := newFirecrackerClient(socketPath, logger, debug)
	c := &Client{client: httpClient}
	c.firecrackerRequestTimeout = envValueOrDefaultInt(firecrackerRequestTimeoutEnv, defaultFirecrackerRequestTimeout)
	c.firecrackerInitTimeout = envValueOrDefaultInt(firecrackerInitTimeoutEnv, defaultFirecrackerInitTimeoutSeconds)

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// GetFirecrackerVersionOpt is a functional option to be used for the
// GetFirecrackerVersion API in setting any additional optional fields.
type GetFirecrackerVersionOpt func(*ops.GetFirecrackerVersionParams)

// GetFirecrackerVersion is a wrapper for the swagger generated client to make
// calling of the API easier.
func (f *Client) GetFirecrackerVersion(ctx context.Context, opts ...GetFirecrackerVersionOpt) (*ops.GetFirecrackerVersionOK, error) {
	params := ops.NewGetFirecrackerVersionParams()

	for _, opt := range opts {
		opt(params)
	}

	return f.client.Operations.GetFirecrackerVersionContext(ctx, params)
}

// PutLoggerOpt is a functional option to be used for the PutLogger API in
// setting any additional optional fields.
type PutLoggerOpt func(*ops.PutLoggerParams)

// PutLogger is a wrapper for the swagger generated client to make calling of
// the API easier.
func (f *Client) PutLogger(ctx context.Context, logger *models.Logger, opts ...PutLoggerOpt) (*ops.PutLoggerNoContent, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Duration(f.firecrackerRequestTimeout)*time.Millisecond)
	defer cancel()

	loggerParams := ops.NewPutLoggerParams()
	loggerParams.SetBody(logger)

	for _, opt := range opts {
		opt(loggerParams)
	}

	return f.client.Operations.PutLoggerContext(timeout, loggerParams)
}

// PutMetricsOpt is a functional option to be used for the PutMetrics API in
// setting any additional optional fields.
type PutMetricsOpt func(*ops.PutMetricsParams)

// PutMetrics is a wrapper for the swagger generated client to make calling of
// the API easier.
func (f *Client) PutMetrics(ctx context.Context, metrics *models.Metrics, opts ...PutMetricsOpt) (*ops.PutMetricsNoContent, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Duration(f.firecrackerRequestTimeout)*time.Millisecond)
	defer cancel()

	params := ops.NewPutMetricsParams()
	params.SetBody(metrics)

	for _, opt := range opts {
		opt(params)
	}

	return f.client.Operations.PutMetricsContext(timeout, params)
}

// PutMachineConfigurationOpt is a functional option to be used for the
// PutMachineConfiguration API in setting any additional optional fields.
type PutMachineConfigurationOpt func(*ops.PutMachineConfigurationParams)

// PutMachineConfiguration is a wrapper for the swagger generated client to
// make calling of the API easier.
func (f *Client) PutMachineConfiguration(ctx context.Context, cfg *models.MachineConfiguration, opts ...PutMachineConfigurationOpt) (*ops.PutMachineConfigurationNoContent, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Duration(f.firecrackerRequestTimeout)*time.Millisecond)
	defer cancel()

	mc := ops.NewPutMachineConfigurationParams()
	mc.SetBody(cfg)

	for _, opt := range opts {
		opt(mc)
	}

	return f.client.Operations.PutMachineConfigurationContext(timeout, mc)
}

// PutGuestBootSourceOpt is a functional option to be used for the
// PutGuestBootSource API in setting any additional optional fields.
type PutGuestBootSourceOpt func(*ops.PutGuestBootSourceParams)

// PutGuestBootSource is a wrapper for the swagger generated client to make
// calling of the API easier.
func (f *Client) PutGuestBootSource(ctx context.Context, source *models.BootSource, opts ...PutGuestBootSourceOpt) (*ops.PutGuestBootSourceNoContent, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Duration(f.firecrackerRequestTimeout)*time.Millisecond)
	defer cancel()

	bootSource := ops.NewPutGuestBootSourceParams()
	bootSource.SetBody(source)

	for _, opt := range opts {
		opt(bootSource)
	}

	return f.client.Operations.PutGuestBootSourceContext(timeout, bootSource)
}

// PutGuestNetworkInterfaceByIDOpt is a functional option to be used for the
// PutGuestNetworkInterfaceByID API in setting any additional optional fields.
type PutGuestNetworkInterfaceByIDOpt func(*ops.PutGuestNetworkInterfaceByIDParams)

// PutGuestNetworkInterfaceByID is a wrapper for the swagger generated client
// to make calling of the API easier.
func (f *Client) PutGuestNetworkInterfaceByID(ctx context.Context, ifaceID string, ifaceCfg *models.NetworkInterface, opts ...PutGuestNetworkInterfaceByIDOpt) (*ops.PutGuestNetworkInterfaceByIDNoContent, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Duration(f.firecrackerRequestTimeout)*time.Millisecond)
	defer cancel()

	cfg := ops.NewPutGuestNetworkInterfaceByIDParams()
	cfg.SetBody(ifaceCfg)
	cfg.SetIfaceID(ifaceID)

	for _, opt := range opts {
		opt(cfg)
	}

	return f.client.Operations.PutGuestNetworkInterfaceByIDContext(timeout, cfg)
}

// PatchGuestNetworkInterfaceByIDOpt is a functional option to be used for the
// PatchGuestNetworkInterfaceByID API in setting any additional optional fields.
type PatchGuestNetworkInterfaceByIDOpt func(*ops.PatchGuestNetworkInterfaceByIDParams)

// PatchGuestNetworkInterfaceByID is a wrapper for the swagger generated client to make calling of the
// API easier.
func (f *Client) PatchGuestNetworkInterfaceByID(ctx context.Context, ifaceID string, ifaceCfg *models.PartialNetworkInterface, opts ...PatchGuestNetworkInterfaceByIDOpt) (*ops.PatchGuestNetworkInterfaceByIDNoContent, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Duration(f.firecrackerRequestTimeout)*time.Millisecond)
	defer cancel()

	cfg := ops.NewPatchGuestNetworkInterfaceByIDParams()
	cfg.SetBody(ifaceCfg)
	cfg.SetIfaceID(ifaceID)

	for _, opt := range opts {
		opt(cfg)
	}

	return f.client.Operations.PatchGuestNetworkInterfaceByIDContext(timeout, cfg)
}

// PutGuestDriveByIDOpt is a functional option to be used for the
// PutGuestDriveByID API in setting any additional optional fields.
type PutGuestDriveByIDOpt func(*ops.PutGuestDriveByIDParams)

// PutGuestDriveByID is a wrapper for the swagger generated client to make
// calling of the API easier.
func (f *Client) PutGuestDriveByID(ctx context.Context, driveID string, drive *models.Drive, opts ...PutGuestDriveByIDOpt) (*ops.PutGuestDriveByIDNoContent, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Duration(f.firecrackerRequestTimeout)/2*time.Millisecond)
	defer cancel()

	params := ops.NewPutGuestDriveByIDParams()
	params.SetDriveID(driveID)
	params.SetBody(drive)

	for _, opt := range opts {
		opt(params)
	}

	return f.client.Operations.PutGuestDriveByIDContext(timeout, params)
}

// PutGuestVsockOpt is a functional option to be used for the
// PutGuestVsock API in setting any additional optional fields.
type PutGuestVsockOpt func(params *ops.PutGuestVsockParams)

// PutGuestVsock is a wrapper for the swagger generated client to make
// calling of the API easier.
func (f *Client) PutGuestVsock(ctx context.Context, vsock *models.Vsock, opts ...PutGuestVsockOpt) (*ops.PutGuestVsockNoContent, error) {
	params := ops.NewPutGuestVsockParams()
	params.SetBody(vsock)

	for _, opt := range opts {
		opt(params)
	}

	return f.client.Operations.PutGuestVsockContext(ctx, params)
}

// PatchVMOpt is a functional option to be used for the
// PatchVM API in setting any additional optional fields.
type PatchVMOpt func(*ops.PatchVMParams)

// PatchVM is a wrapper for the swagger generated client to make
// calling of the API easier.
func (f *Client) PatchVM(ctx context.Context, vm *models.VM, opts ...PatchVMOpt) (*ops.PatchVMNoContent, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Duration(f.firecrackerRequestTimeout)*time.Millisecond)
	defer cancel()

	params := ops.NewPatchVMParams()
	params.SetBody(vm)

	for _, opt := range opts {
		opt(params)
	}

	return f.client.Operations.PatchVMContext(timeout, params)
}

// CreateSnapshotOpt is a functional option to be used for the
// CreateSnapshot API in setting any additional optional fields.
type CreateSnapshotOpt func(*ops.CreateSnapshotParams)

// CreateSnapshot is a wrapper for the swagger generated client to make
// calling of the API easier.
func (f *Client) CreateSnapshot(ctx context.Context, snapshotParams *models.SnapshotCreateParams, opts ...CreateSnapshotOpt) (*ops.CreateSnapshotNoContent, error) {
	params := ops.NewCreateSnapshotParams()
	params.SetBody(snapshotParams)

	for _, opt := range opts {
		opt(params)
	}

	return f.client.Operations.CreateSnapshotContext(ctx, params)
}

// LoadSnapshotOpt is a functional option to be used for the
// LoadSnapshot API in setting any additional optional fields.
type LoadSnapshotOpt func(*ops.LoadSnapshotParams)

// LoadSnapshot is a wrapper for the swagger generated client to make
// calling of the API easier.
func (f *Client) LoadSnapshot(ctx context.Context, snapshotParams *models.SnapshotLoadParams, opts ...LoadSnapshotOpt) (*ops.LoadSnapshotNoContent, error) {
	params := ops.NewLoadSnapshotParams()
	params.SetBody(snapshotParams)

	for _, opt := range opts {
		opt(params)
	}

	return f.client.Operations.LoadSnapshotContext(ctx, params)
}

// CreateSyncActionOpt is a functional option to be used for the
// CreateSyncAction API in setting any additional optional fields.
type CreateSyncActionOpt func(*ops.CreateSyncActionParams)

// CreateSyncAction is a wrapper for the swagger generated client to make
// calling of the API easier.
func (f *Client) CreateSyncAction(ctx context.Context, info *models.InstanceActionInfo, opts ...CreateSyncActionOpt) (*ops.CreateSyncActionNoContent, error) {
	params := ops.NewCreateSyncActionParams()
	params.SetInfo(info)

	for _, opt := range opts {
		opt(params)
	}

	return f.client.Operations.CreateSyncActionContext(ctx, params)
}

// PutMmdsOpt is a functional option to be used for the PutMmds API in setting
// any additional optional fields.
type PutMmdsOpt func(*ops.PutMmdsParams)

// PutMmds is a wrapper for the swagger generated client to make calling of the
// API easier.
func (f *Client) PutMmds(ctx context.Context, metadata any, opts ...PutMmdsOpt) (*ops.PutMmdsNoContent, error) {
	params := ops.NewPutMmdsParams()
	params.SetBody(metadata)

	for _, opt := range opts {
		opt(params)
	}

	return f.client.Operations.PutMmdsContext(ctx, params)
}

// GetMmdsOpt is a functional option to be used for the GetMmds API in setting
// any additional optional fields.
type GetMmdsOpt func(*ops.GetMmdsParams)

// GetMmds is a wrapper for the swagger generated client to make calling of the
// API easier.
func (f *Client) GetMmds(ctx context.Context, opts ...GetMmdsOpt) (*ops.GetMmdsOK, error) {
	params := ops.NewGetMmdsParams()

	for _, opt := range opts {
		opt(params)
	}

	return f.client.Operations.GetMmdsContext(ctx, params)
}

// PatchMmdsOpt is a functional option to be used for the GetMmds API in setting
// any additional optional fields.
type PatchMmdsOpt func(*ops.PatchMmdsParams)

// PatchMmds is a wrapper for the swagger generated client to make calling of the
// API easier.
func (f *Client) PatchMmds(ctx context.Context, metadata any, opts ...PatchMmdsOpt) (*ops.PatchMmdsNoContent, error) {
	params := ops.NewPatchMmdsParams()
	params.SetBody(metadata)

	for _, opt := range opts {
		opt(params)
	}

	return f.client.Operations.PatchMmdsContext(ctx, params)
}

// PutMmdsConfig is a wrapper for the swagger generated client to make calling of the
// API easier.
func (f *Client) PutMmdsConfig(ctx context.Context, config *models.MmdsConfig) (*ops.PutMmdsConfigNoContent, error) {
	params := ops.NewPutMmdsConfigParams()
	params.SetBody(config)

	return f.client.Operations.PutMmdsConfigContext(ctx, params)
}

// GetMachineConfigurationOpt  is a functional option to be used for the
// GetMachineConfiguration API in setting any additional optional fields.
type GetMachineConfigurationOpt func(*ops.GetMachineConfigurationParams)

// GetMachineConfiguration is a wrapper for the swagger generated client to make
// calling of the API easier.
func (f *Client) GetMachineConfiguration(opts ...GetMachineConfigurationOpt) (*ops.GetMachineConfigurationOK, error) {
	p := ops.NewGetMachineConfigurationParams()
	p.SetTimeout(time.Duration(f.firecrackerRequestTimeout) * time.Millisecond)

	for _, opt := range opts {
		opt(p)
	}

	return f.client.Operations.GetMachineConfiguration(p)
}

// DescribeInstanceOpt is a functional option to be used for the DescribeInstance API
// for any additional optional fields
type DescribeInstanceOpt func(*ops.DescribeInstanceParams)

// GetInstanceInfo is a wrapper for the swagger generated client to make calling of
// the API easier
func (f *Client) GetInstanceInfo(ctx context.Context, opts ...DescribeInstanceOpt) (*ops.DescribeInstanceOK, error) {
	params := ops.NewDescribeInstanceParams()

	for _, opt := range opts {
		opt(params)
	}

	return f.client.Operations.DescribeInstanceContext(ctx, params)
}

// PatchGuestDriveByIDOpt is a functional option to be used for the PutMmds API in setting
// any additional optional fields.
type PatchGuestDriveByIDOpt func(*ops.PatchGuestDriveByIDParams)

// PatchGuestDriveByID is a wrapper for the swagger generated client to make calling of the
// API easier.
func (f *Client) PatchGuestDriveByID(ctx context.Context, driveID, pathOnHost string, opts ...PatchGuestDriveByIDOpt) (*ops.PatchGuestDriveByIDNoContent, error) {
	params := ops.NewPatchGuestDriveByIDParams()

	partialDrive := models.PartialDrive{
		DriveID:    &driveID,
		PathOnHost: pathOnHost,
	}
	params.SetBody(&partialDrive)
	params.DriveID = driveID

	for _, opt := range opts {
		opt(params)
	}

	return f.client.Operations.PatchGuestDriveByIDContext(ctx, params)
}

// PutBalloonOpt is a functional option to be used for the
// PutBalloon API in setting any additional optional fields.
type PutBalloonOpt func(*ops.PutBalloonParams)

// PutBalloon is a wrapper for the swagger generated client to make
// calling of the API easier.
func (f *Client) PutBalloon(ctx context.Context, balloon *models.Balloon, opts ...PutBalloonOpt) (*ops.PutBalloonNoContent, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Duration(f.firecrackerRequestTimeout)*time.Millisecond)
	defer cancel()

	params := ops.NewPutBalloonParams()
	params.SetBody(balloon)

	for _, opt := range opts {
		opt(params)
	}

	return f.client.Operations.PutBalloonContext(timeout, params)
}

// DescribeBalloonConfig is a wrapper for the swagger generated client to make
// calling of the API easier.
func (f *Client) DescribeBalloonConfig(ctx context.Context) (*ops.DescribeBalloonConfigOK, error) {
	params := ops.NewDescribeBalloonConfigParams()
	params.SetTimeout(time.Duration(f.firecrackerRequestTimeout) * time.Millisecond)

	return f.client.Operations.DescribeBalloonConfigContext(ctx, params)
}

// PatchBalloonOpt is a functional option to be used for the PatchBalloon API in setting
// any additional optional fields.
type PatchBalloonOpt func(*ops.PatchBalloonParams)

// PatchBalloon is a wrapper for the swagger generated client to make calling of the
// API easier.
func (f *Client) PatchBalloon(ctx context.Context, ballonUpdate *models.BalloonUpdate, opts ...PatchBalloonOpt) (*ops.PatchBalloonNoContent, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Duration(f.firecrackerRequestTimeout)*time.Millisecond)
	defer cancel()

	params := ops.NewPatchBalloonParams()
	params.SetBody(ballonUpdate)

	for _, opt := range opts {
		opt(params)
	}

	return f.client.Operations.PatchBalloonContext(timeout, params)
}

// DescribeBalloonStats is a wrapper for the swagger generated client to make calling of the
// API easier.
func (f *Client) DescribeBalloonStats(ctx context.Context) (*ops.DescribeBalloonStatsOK, error) {
	params := ops.NewDescribeBalloonStatsParams()
	params.SetTimeout(time.Duration(f.firecrackerRequestTimeout) * time.Millisecond)

	return f.client.Operations.DescribeBalloonStatsContext(ctx, params)
}

// PatchBalloonStatsIntervalOpt is a functional option to be used for the PatchBalloonStatsInterval API in setting
// any additional optional fields.
type PatchBalloonStatsIntervalOpt func(*ops.PatchBalloonStatsIntervalParams)

// PatchBalloonStatsInterval is a wrapper for the swagger generated client to make calling of the
// API easier.
func (f *Client) PatchBalloonStatsInterval(ctx context.Context, balloonStatsUpdate *models.BalloonStatsUpdate, opts ...PatchBalloonStatsIntervalOpt) (*ops.PatchBalloonStatsIntervalNoContent, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Duration(f.firecrackerRequestTimeout)*time.Millisecond)
	defer cancel()

	params := ops.NewPatchBalloonStatsIntervalParams()
	params.SetBody(balloonStatsUpdate)

	for _, opt := range opts {
		opt(params)
	}

	return f.client.Operations.PatchBalloonStatsIntervalContext(timeout, params)
}

type GetExportVMConfigOpt func(*ops.GetExportVMConfigParams)

func (f *Client) GetExportVMConfig(opts ...GetExportVMConfigOpt) (*ops.GetExportVMConfigOK, error) {
	p := ops.NewGetExportVMConfigParams()
	p.SetTimeout(time.Duration(f.firecrackerRequestTimeout) * time.Millisecond)

	for _, opt := range opts {
		opt(p)
	}

	return f.client.Operations.GetExportVMConfig(p)
}

// ConfigureMemoryHotplugOpt is a functional option to be used for the
// PutMemoryHotplugConfig API in setting any additional optional fields.
type ConfigureMemoryHotplugOpt func(*ops.PutMemoryHotplugParams)

// PutMemoryHotplug is a wrapper for the swagger generated client to make
// calling of the API easier.
func (f *Client) ConfigureMemoryHotplugging(ctx context.Context, memoryHotplugConfig *models.MemoryHotplugConfig, opts ...ConfigureMemoryHotplugOpt) (*ops.PutMemoryHotplugNoContent, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Duration(f.firecrackerRequestTimeout)*time.Millisecond)
	defer cancel()

	params := ops.NewPutMemoryHotplugParams()
	params.SetBody(memoryHotplugConfig)

	for _, opt := range opts {
		opt(params)
	}

	return f.client.Operations.PutMemoryHotplugContext(timeout, params)
}

// UpdateMemoryHotpluggingOpt is a functional option to be used for the
// UpdateMemoryHotplugging API in setting any additional optional fields.
type UpdateMemoryHotpluggingOpt func(*ops.PatchMemoryHotplugParams)

// UpdateMemoryHotplugging is a wrapper for the swagger generated client to make
// calling of the API easier.
func (f *Client) UpdateMemoryHotplugging(ctx context.Context, memoryHotplugConfig *models.MemoryHotplugSizeUpdate, opts ...UpdateMemoryHotpluggingOpt) (*ops.PatchMemoryHotplugNoContent, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Duration(f.firecrackerRequestTimeout)*time.Millisecond)
	defer cancel()

	params := ops.NewPatchMemoryHotplugParams()
	params.SetBody(memoryHotplugConfig)

	for _, opt := range opts {
		opt(params)
	}

	return f.client.Operations.PatchMemoryHotplugContext(timeout, params)
}

// GetMemoryHotpluggingStatusOpt is a functional option to be used for the
// GetMemoryHotpluggingStatus API in setting any additional optional fields.
type GetMemoryHotpluggingStatusOpt func(*ops.GetMemoryHotplugParams)

// GetMemoryHotpluggingStatus is a wrapper for the swagger generated client to make
// calling of the API easier.
func (f *Client) GetMemoryHotpluggingStatus(ctx context.Context, opts ...GetMemoryHotpluggingStatusOpt) (*ops.GetMemoryHotplugOK, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Duration(f.firecrackerRequestTimeout)*time.Millisecond)
	defer cancel()

	params := ops.NewGetMemoryHotplugParams()

	for _, opt := range opts {
		opt(params)
	}

	return f.client.Operations.GetMemoryHotplugContext(timeout, params)
}
