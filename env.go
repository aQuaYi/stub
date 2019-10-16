package stub

import "os"

func (s *Stubs) checkEnvKey(k string) {
	if _, ok := s.origEnv[k]; !ok {
		v, ok := os.LookupEnv(k)
		s.origEnv[k] = envVal{v, ok}
	}
}

// Env stubs environmental variable
func Env(k, v string) *Stubs {
	s := New()
	s.Env(k, v)
	return s
}

// Env the specified environent variable to the specified value.
func (s *Stubs) Env(k, v string) *Stubs {
	s.checkEnvKey(k)
	os.Setenv(k, v)
	return s
}

// SetEnv the specified environent variable to the specified value.
func (s *Stubs) SetEnv(k, v string) *Stubs {
	s.checkEnvKey(k)
	os.Setenv(k, v)
	return s
}

// UnsetEnv unsets the specified environent variable.
func (s *Stubs) UnsetEnv(k string) *Stubs {
	s.checkEnvKey(k)

	os.Unsetenv(k)
	return s
}

func (s *Stubs) resetEnv() {
	for k, v := range s.origEnv {
		if v.ok {
			os.Setenv(k, v.val)
		} else {
			os.Unsetenv(k)
		}
	}
}
