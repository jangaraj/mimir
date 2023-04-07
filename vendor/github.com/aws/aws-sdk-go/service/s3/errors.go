// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

const (

	// ErrCodeBucketAlreadyExists for service response error code
	// "BucketAlreadyExists".
	//
	// The requested bucket name is not available. The bucket namespace is shared
	// by all users of the system. Select a different name and try again.
	ErrCodeBucketAlreadyExists = "BucketAlreadyExists"

	// ErrCodeBucketAlreadyOwnedByYou for service response error code
	// "BucketAlreadyOwnedByYou".
	//
	// The bucket you tried to create already exists, and you own it. Amazon S3
	// returns this error in all Amazon Web Services Regions except in the North
	// Virginia Region. For legacy compatibility, if you re-create an existing bucket
	// that you already own in the North Virginia Region, Amazon S3 returns 200
	// OK and resets the bucket access control lists (ACLs).
	ErrCodeBucketAlreadyOwnedByYou = "BucketAlreadyOwnedByYou"

	// ErrCodeInvalidObjectState for service response error code
	// "InvalidObjectState".
	//
	// Object is archived and inaccessible until restored.
	ErrCodeInvalidObjectState = "InvalidObjectState"

	// ErrCodeNoSuchBucket for service response error code
	// "NoSuchBucket".
	//
	// The specified bucket does not exist.
	ErrCodeNoSuchBucket = "NoSuchBucket"

	// ErrCodeNoSuchKey for service response error code
	// "NoSuchKey".
	//
	// The specified key does not exist.
	ErrCodeNoSuchKey = "NoSuchKey"

	// ErrCodeNoSuchUpload for service response error code
	// "NoSuchUpload".
	//
	// The specified multipart upload does not exist.
	ErrCodeNoSuchUpload = "NoSuchUpload"

	// ErrCodeObjectAlreadyInActiveTierError for service response error code
	// "ObjectAlreadyInActiveTierError".
	//
	// This action is not allowed against this storage tier.
	ErrCodeObjectAlreadyInActiveTierError = "ObjectAlreadyInActiveTierError"

	// ErrCodeObjectNotInActiveTierError for service response error code
	// "ObjectNotInActiveTierError".
	//
	// The source object of the COPY action is not in the active tier and is only
	// stored in Amazon S3 Glacier.
	ErrCodeObjectNotInActiveTierError = "ObjectNotInActiveTierError"
)
