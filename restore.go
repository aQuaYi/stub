package stub

import "os"

// Restore resets all stubbed variables back to their original values.
func (s *stubs) Restore() {
	s.restoreVars()
	s.restoreEnv()
}

func (s *stubs) restoreVars() {
	for original, val := range s.vars {
		original.Elem().Set(val)
		delete(s.vars, original)
	}
}

func (s *stubs) restoreEnv() {
	for k, v := range s.envs {
		if v.isExisted {
			os.Setenv(k, v.val)
		} else {
			os.Unsetenv(k)
		}
	}
}
