package base

type TransformContext struct {
	Version                 any
	InstanceID              string
	GetDatabaseMetadataFunc GetDatabaseMetadataFunc
}
