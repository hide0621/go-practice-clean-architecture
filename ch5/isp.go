package ch5

type BooleanProvider interface {
	IsEnabled() bool
}

type SharedPreferenceManager struct {
	sp map[string]BooleanProvider
}

func NewSharedPreferenceManager() *SharedPreferenceManager {
	return &SharedPreferenceManager{
		sp: make(map[string]BooleanProvider),
	}
}

func (spm *SharedPreferenceManager) AddBooleanProvider(key string, provider BooleanProvider) {
	spm.sp[key] = provider
}

type ResProvider struct {
	value bool
}

func (rp ResProvider) IsEnabled() bool {
	return rp.value
}

type SoundProvider struct {
	value bool
}

func (sp SoundProvider) IsEnabled() bool {
	return sp.value
}

type VibrationProvider struct {
	value bool
}

func (vp VibrationProvider) IsEnabled() bool {
	return vp.value
}
