package config

func ReloadGeneration(r Runtime, err error) (Runtime, error) {
	if err != nil {
		return r, err
	}
	r.Next()
	return r, nil
}
