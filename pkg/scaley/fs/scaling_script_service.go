package fs

type ScalingScriptService struct{}

func NewScalingScriptService() *ScalingScriptService {
	return &ScalingScriptService{}
}

func (service *ScalingScriptService) Exists(path string) bool {
	return FileExists(path)
}
