package stub

import "os"

// Env stubs environmental variable
func Env(k, v string) Stubber {
	return newStubs().Env(k, v)
}

func (s *stubs) checkEnvKey(k string) {
	if _, ok := s.envs[k]; !ok {
		v, ok := os.LookupEnv(k)
		s.envs[k] = env{v, ok}
	}
}

// Env the specified environent variable to the specified value.
func (s *stubs) Env(k, v string) Stubber {
	s.checkEnvKey(k)
	os.Setenv(k, v)
	return s
}

// UnsetEnv unset the specified environent variable.
func (s *stubs) UnsetEnv(k string) *stubs {
	s.checkEnvKey(k)

	os.Unsetenv(k)
	return s
}
