package compute

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/compute/v1"
)

func ComputeDisks() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_disks",
		Description:  "Represents a Persistent Disk resource.",
		Resolver:     fetchComputeDisks,
		IgnoreError:  client.IgnoreErrorHandler,
		Multiplex:    client.ProjectMultiplex,
		DeleteFilter: client.DeleteProjectFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "creation_timestamp",
				Description: "Creation timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "An optional description of this resource Provide this property when you create the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "disk_encryption_key_kms_key_name",
				Description: "The name of the encryption key that is stored in Google Cloud KMS",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiskEncryptionKey.KmsKeyName"),
			},
			{
				Name:        "disk_encryption_key_kms_key_service_account",
				Description: "The service account being used for the encryption request for the given KMS key If absent, the Compute Engine default service account is used",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiskEncryptionKey.KmsKeyServiceAccount"),
			},
			{
				Name:        "disk_encryption_key_raw_key",
				Description: "Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiskEncryptionKey.RawKey"),
			},
			{
				Name:        "disk_encryption_key_sha256",
				Description: "The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiskEncryptionKey.Sha256"),
			},
			{
				Name:        "guest_os_features",
				Description: "A list of features to enable on the guest operating system Applicable only for bootable images Read  Enabling guest operating system features to see a list of available options",
				Type:        schema.TypeStringArray,
				Resolver:    resolveComputeDiskGuestOsFeatures,
			},
			{
				Name:        "id",
				Description: "The unique identifier for the resource This identifier is defined by the server",
				Type:        schema.TypeString,
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "kind",
				Description: "Type of the resource Always compute#disk for disks",
				Type:        schema.TypeString,
			},
			{
				Name:        "label_fingerprint",
				Description: "A fingerprint for the labels being applied to this disk, which is essentially a hash of the labels set used for optimistic locking The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet  To see the latest fingerprint, make a get() request to retrieve a disk",
				Type:        schema.TypeString,
			},
			{
				Name:        "labels",
				Description: "Labels to apply to this disk These can be later modified by the setLabels method",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "last_attach_timestamp",
				Description: "Last attach timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_detach_timestamp",
				Description: "Last detach timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "licenses",
				Description: "A list of publicly visible licenses Reserved for Google's use",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "location_hint",
				Description: "An opaque location hint used to place the disk close to other resources This field is for use by internal tools that use the public API",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "Name of the resource Provided by the client when the resource is created",
				Type:        schema.TypeString,
			},
			{
				Name:        "options",
				Description: "Internal use only",
				Type:        schema.TypeString,
			},
			{
				Name:        "physical_block_size_bytes",
				Description: "Physical block size of the persistent disk, in bytes If not present in a request, a default value is used The currently supported size is 4096, other sizes may be added in the future If an unsupported value is requested, the error message will list the supported values for the caller's project",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "provisioned_iops",
				Description: "Indicates how many IOPS must be provisioned for the disk",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "region",
				Description: "URL of the region where the disk resides Only applicable for regional resources You must specify this field as part of the HTTP request URL It is not settable as a field in the request body",
				Type:        schema.TypeString,
			},
			{
				Name:        "replica_zones",
				Description: "URLs of the zones where the disk should be replicated to Only applicable for regional resources",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "resource_policies",
				Description: "Resource policies applied to this disk for automatic snapshot creations",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "satisfies_pzs",
				Description: "Reserved for future use",
				Type:        schema.TypeBool,
			},
			{
				Name:        "self_link",
				Description: "Server-defined fully-qualified URL for this resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "size_gb",
				Description: "Size, in GB",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "source_disk",
				Description: "The source disk used to create this disk",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_disk_id",
				Description: "The unique ID of the disk used to create this disk This value identifies the exact disk that was used to create this persistent disk For example, if you created the persistent disk from a disk that was later deleted and recreated under the same name, the source disk ID would identify the exact version of the disk that was used",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_image",
				Description: "The source image used to create this disk",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_image_encryption_key_kms_key_name",
				Description: "The name of the encryption key that is stored in Google Cloud KMS",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceImageEncryptionKey.KmsKeyName"),
			},
			{
				Name:        "source_image_encryption_key_kms_key_service_account",
				Description: "The service account being used for the encryption request for the given KMS key If absent, the Compute Engine default service account is used",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceImageEncryptionKey.KmsKeyServiceAccount"),
			},
			{
				Name:        "source_image_encryption_key_raw_key",
				Description: "Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceImageEncryptionKey.RawKey"),
			},
			{
				Name:        "source_image_encryption_key_sha256",
				Description: "The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceImageEncryptionKey.Sha256"),
			},
			{
				Name:        "source_image_id",
				Description: "The ID value of the image used to create this disk This value identifies the exact image that was used to create this persistent disk For example, if you created the persistent disk from an image that was later deleted and recreated under the same name, the source image ID would identify the exact version of the image that was used",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_snapshot",
				Description: "The source snapshot used to create this disk You can provide this as a partial or full URL to the resource For example, the following are valid values: - https://wwwgoogleapis",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_snapshot_encryption_key_kms_key_name",
				Description: "The name of the encryption key that is stored in Google Cloud KMS",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceSnapshotEncryptionKey.KmsKeyName"),
			},
			{
				Name:        "source_snapshot_encryption_key_kms_key_service_account",
				Description: "The service account being used for the encryption request for the given KMS key If absent, the Compute Engine default service account is used",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceSnapshotEncryptionKey.KmsKeyServiceAccount"),
			},
			{
				Name:        "source_snapshot_encryption_key_raw_key",
				Description: "Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceSnapshotEncryptionKey.RawKey"),
			},
			{
				Name:        "source_snapshot_encryption_key_sha256",
				Description: "The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceSnapshotEncryptionKey.Sha256"),
			},
			{
				Name:        "source_snapshot_id",
				Description: "The unique ID of the snapshot used to create this disk This value identifies the exact snapshot that was used to create this persistent disk For example, if you created the persistent disk from a snapshot that was later deleted and recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_storage_object",
				Description: "The full Google Cloud Storage URI where the disk image is stored This file must be a gzip-compressed tarball whose name ends in targz or virtual machine disk whose name ends in vmdk Valid URIs may start with gs:// or https://storagegoogleapiscom/ This flag is not optimized for creating multiple disks from a source storage object To create many disks from a source storage object, use gcloud compute images import instead",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The status of disk creation - CREATING: Disk is provisioning - RESTORING: Source data is being copied into the disk - FAILED: Disk creation failed - READY: Disk is ready for use - DELETING: Disk is deleting",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "URL of the disk type resource describing which disk type to use to create the disk Provide this when creating the disk",
				Type:        schema.TypeString,
			},
			{
				Name: "users",
				Type: schema.TypeStringArray,
			},
			{
				Name:        "zone",
				Description: "URL of the zone where the disk resides You must specify this field as part of the HTTP request URL It is not settable as a field in the request body",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeDisks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Compute.Disks.AggregatedList(c.ProjectId).Context(ctx).PageToken(nextPageToken)
		output, err := client.Retryer(ctx, c, call.Do)
		if err != nil {
			return err
		}

		var diskTypes []*compute.Disk
		for _, items := range output.Items {
			diskTypes = append(diskTypes, items.Disks...)
		}
		res <- diskTypes

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func resolveComputeDiskGuestOsFeatures(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*compute.Disk)
	res := make([]string, len(r.GuestOsFeatures))
	for i, v := range r.GuestOsFeatures {
		res[i] = v.Type
	}
	return resource.Set("guest_os_features", res)
}
