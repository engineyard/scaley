package fs

// ScalingScriptService is a service that knows how to interact with scaling
// scripts via the file system.
type ScalingScriptService struct{}

// NewScalingScriptService returns a new ScalingScriptService.
func NewScalingScriptService() *ScalingScriptService {
	return &ScalingScriptService{}
}

// Exists takes a path and returns a boolean that expresses whether or not the
// scaling script at that location exists.
func (service *ScalingScriptService) Exists(path string) bool {
	return FileExists(path)
}
