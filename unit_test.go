package main

import (
	"testing"
)

func TestRetailerPoint(t *testing.T) {
	result := retailerPoint("Target")
	if result != 6 {
		t.Fail()
	}

}

func TestRetailerPoint2(t *testing.T) {
	result := retailerPoint("M&M Corner Market")
	if result != 14 {
		t.Fail()
	}
}

func TestRoundDollarPoint(t *testing.T) {
	result := roundDollarPoint("35.35")
	if result != 0 {
		t.Fail()
	}
}

func TestRoundDollarPoint2(t *testing.T) {
	result := roundDollarPoint("9.00")
	if result != 50 {
		t.Fail()
	}
}

func TestQuarterMultiplePoint(t *testing.T) {
	result := quarterMultiplePoint("35.35")
	if result != 0 {
		t.Fail()
	}
}

func TestQuarterMultiplePoint2(t *testing.T) {
	result := quarterMultiplePoint("9.00")
	if result != 25 {
		t.Fail()
	}
}

func TestTwoItemPoint(t *testing.T) {
	result := twoItemPoint([]Item{{"Mountain Dew 12PK", "6.49"}, {"Emils Cheese Pizza", "12.25"},
		{"Knorr Creamy Chicken", "1.26"}, {"Doritos Nacho Cheese", "3.35"}, {"   Klarbrunn 12-PK 12 FL OZ  ", "12.00"}})
	if result != 10 {
		t.Fail()
	}
}

func TestTwoItemPoint2(t *testing.T) {
	result := twoItemPoint([]Item{{"Gatorade", "2.25"}, {"Gatorade", "2.25"},
		{"Gatorade", "2.25"}, {"Gatorade", "2.25"}})
	if result != 10 {
		t.Fail()
	}
}

func TestTrimItemPoint(t *testing.T) {
	result := trimItemPoint([]Item{{"Mountain Dew 12PK", "6.49"}, {"Emils Cheese Pizza", "12.25"},
		{"Knorr Creamy Chicken", "1.26"}, {"Doritos Nacho Cheese", "3.35"}, {"   Klarbrunn 12-PK 12 FL OZ  ", "12.00"}})
	if result != 6 {
		t.Fail()
	}
}

func TestTrimItemPoint2(t *testing.T) {
	result := trimItemPoint([]Item{{"Gatorade", "2.25"}, {"Gatorade", "2.25"},
		{"Gatorade", "2.25"}, {"Gatorade", "2.25"}})
	if result != 0 {
		t.Fail()
	}
}

func TestOddDayPoint(t *testing.T) {
	result := oddDayPoint("2022-01-01")
	if result != 6 {
		t.Fail()
	}
}

func TestOddDayPoint2(t *testing.T) {
	result := oddDayPoint("2022-03-20")
	if result != 0 {
		t.Fail()
	}
}

func TestBetweenTwoFourPoint(t *testing.T) {
	result := betweenTwoFourPoint("13:01")
	if result != 0 {
		t.Fail()
	}
}

func TestBetweenTwoFourPoint2(t *testing.T) {
	result := betweenTwoFourPoint("14:33")
	if result != 10 {
		t.Fail()
	}
}
