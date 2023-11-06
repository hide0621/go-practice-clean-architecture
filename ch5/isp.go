package ch5

type SharedPreferenceManager struct {
	sp map[string]interface{}
}

func NewSharedPreferenceManager() *SharedPreferenceManager {
	return &SharedPreferenceManager{
		sp: make(map[string]interface{}),
	}
}

func (spm *SharedPreferenceManager) SetBoolean(key string, value interface{}) {
	spm.sp[key] = value
}

func (spm *SharedPreferenceManager) IsResPrefix() bool {
	val, exists := spm.sp["res"].(bool)
	return exists && val
}

func (spm *SharedPreferenceManager) IsEnableNotifySound() bool {
	val, exists := spm.sp["sound"].(bool)
	return exists && val
}

func (spm *SharedPreferenceManager) IsEnableNotifyVib() bool {
	val, exists := spm.sp["vibration"].(bool)
	return exists && val
}
