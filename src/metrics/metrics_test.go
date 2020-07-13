package metrics

import (
	"testing"
)

func TestIncValid(t *testing.T) {
	t.Log("IncValid")
	{
		t.Log("Given MetricsService")
		{
			var i, n uint
			for n = 1; n <= 3; n++ {
				t.Log("When called", n, "time(s)")
				{
					sut := Create()
					for i = 0; i < n; i++ {
						sut.IncValid()
					}
					expected := n
					if sut.valid != expected {
						t.Errorf("It should increment valid field to value: %v, got: %v", expected, sut.valid)
					} else {
						t.Log("It should increment valid field to value", expected)
					}
					if sut.invalid != 0 {
						t.Errorf("It should not increment invalid field, got: %v", sut.invalid)
					} else {
						t.Log("It should not increment invalid field")
					}
				}
			}
		}
	}
}

func TestIncInvalid(t *testing.T) {
	t.Log("IncInvalid")
	{
		t.Log("Given MetricsService")
		{
			var i, n uint
			for n = 1; n <= 3; n++ {
				t.Log("When called", n, "time(s)")
				{
					sut := Create()
					for i = 0; i < n; i++ {
						sut.IncInvalid()
					}
					expected := n
					if sut.invalid != expected {
						t.Errorf("It should increment invalid field to value: %v, got: %v", expected, sut.invalid)
					} else {
						t.Log("It should increment invalid field to value", expected)
					}
					if sut.valid != 0 {
						t.Errorf("It should not increment valid field, got: %v", sut.valid)
					} else {
						t.Log("It should not increment valid field")
					}
				}
			}
		}
	}
}

func TestGetReport(t *testing.T) {
	t.Log("GetReport")
	{
		t.Log("Given MetricsService")
		{
			t.Log("When called on fresh service")
			{
				sut := Create()
				report := sut.GetReport()
				expected := "Valid inputs: 0\nInvalid inputs: 0\n"
				if report != expected {
					t.Errorf("It should return valid report:\n%q\ngot:\n%q\n", expected, report)
				} else {
					t.Log("It should return valid report:", expected)
				}
			}
			t.Log("When called on service with non-zero values")
			{
				sut := Create()
				sut.valid = 7
				sut.invalid = 4
				report := sut.GetReport()
				expected := "Valid inputs: 7\nInvalid inputs: 4\n"
				if report != expected {
					t.Errorf("It should return valid report:\n%q\ngot:\n%q\n", expected, report)
				} else {
					t.Log("It should return valid report:", expected)
				}
			}
		}
	}
}
