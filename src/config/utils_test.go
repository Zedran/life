package config

import "testing"

/* A simple test of VerifyCfgFileExt. */
func TestVerifyCfgFileExt(t *testing.T) {
	const ext = ".json"

	cases := map[string]string{
		"file":      "file.json",
		"file.json": "file.json",
	}

	for c, eo := range cases {
		o := VerifyCfgFileExt(c)

		if o != eo {
			t.Errorf("Failed on %s. Expected %s, got %s.\n", c, eo, o)
		}
	}
}
