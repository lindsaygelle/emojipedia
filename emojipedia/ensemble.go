package emojipedia

type Ensemble map[string][]string

func (ensemble *Ensemble) Add(key string, value string) string {
	(*ensemble)[key] = append((*ensemble)[key], value)
	return value
}

func (ensemble *Ensemble) Set(key string) string {
	(*ensemble)[key] = []string{}
	return key
}
