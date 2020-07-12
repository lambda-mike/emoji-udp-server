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
		t.Log("Given negative n")
		{
			isRaw := true
			n := -3
			sep := ","
			_, err := Create(isRaw, n, sep)
			if err == nil {
				t.Fatal("It should return err, got: nil")
			}
			t.Log("It should return err", err)
		}
	}
}

func TestParsePort(t *testing.T) {
	t.Log("ParsePort")
	{
		t.Log("Given correct port string")
		{
			// 49152–65535
			str := "54321"
			port, err := ParsePort(str)
			if err != nil {
				t.Fatal("It should not return err, got: ", err)
			}
			t.Log("It should not return err")
			expected := 54321
			if port != expected {
				t.Fatalf("It should return valid port %d, got: %d", expected, port)
			}
			t.Log("It should return valid port", expected)
		}
		t.Log("Given too small port")
		{
			// 49152–65535
			str := "49151"
			_, err := ParsePort(str)
			if err == nil {
				t.Fatal("It should return err, got nil")
			}
			t.Log("It should return err", err)
		}
		t.Log("Given too big port")
		{
			// 49152–65535
			str := "65536"
			_, err := ParsePort(str)
			if err == nil {
				t.Fatal("It should return err, got nil")
			}
			t.Log("It should return err", err)
		}
		t.Log("Given incorrect port")
		{
			str := "nope"
			_, err := ParsePort(str)
			if err == nil {
				t.Fatal("It should return err, got nil")
			}
			t.Log("It should return err", err)
		}
	}
}
