package config

import "testing"

func TestCreate(t *testing.T) {
	t.Log("Create")
	{
		t.Log("Given correct params")
		{
			isRaw := true
			n := 3
			sep := ","
			conf, err := Create(isRaw, n, sep)
			if err != nil {
				t.Fatal("It should not return err, got: ", err)
			}
			t.Log("It should not return err")
			if conf.Raw != isRaw {
				t.Errorf("It should return config with correct Raw value: %v, got: %v", isRaw, conf.Raw)
			}
			t.Log("It should return config with correct Raw value")
			if conf.N != n {
				t.Errorf("It should return config with correct N value: %v, got: %v", n, conf.N)
			}
			t.Log("It should return config with correct N value")
			if conf.Separator != sep {
				t.Errorf("It should return config with correct Separator value: %v, got: %v", sep, conf.Separator)
			}
			t.Log("It should return config with correct Separator value")
		}
	}
}
