package pscan

import "testing"

func TestPscan(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			in:  "scanme.nmap.org",
			out: "22 80",
		},
	}

	for _, test := range tests {
		if val := Scan(test.in); val != test.out {
			t.Errorf("expected %v \ngot %v", test.in, val)
		}
	}
}
