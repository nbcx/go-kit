package tz

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func TestSetTimezone(t *testing.T) {
	err := SetTimeZone("Asia/Ho_Chi_Minh")
	checkError(err)

	loc := GetTimezone()
	fmt.Printf("%+v\n", loc)

	date := time.Now().In(loc).Format(FullFormat)
	fmt.Println(date)

	date2 := LocalNow()
	fmt.Println(date2)
}

func TestGetTimezone(t *testing.T) {
	loc := GetTimezone()
	fmt.Printf("%+v\n", loc)

	date2 := LocalNow()
	fmt.Println(date2)
}

func TestSetSystemTime(t *testing.T) {
	tests := []struct {
		name string
		dt   string
		want string
	}{
		{"test", "2002-12-12 19:07:55", "2002-12-12 19:07:55"},
		{"test", "", GetFormatStr("2006-01-02 15:04:05")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetSystemTime(tt.dt); err != nil {
				t.Errorf("SetSystemTime() error = %v", err)
			}
			if w := GetFormatStr(FullFormat); w != tt.want {
				t.Errorf("SetSystemTime() w  = %v, want = %v", w, tt.want)
			}
		})
	}
}

func TestGetDays(t *testing.T) {
	tests := []struct {
		name string
		to   string
		from string
		want int64
	}{
		{"test1", "2002-12-12 19:07:55", "2002-12-12 19:07:55", 0},
		{"test1", "2002-12-12 19:07:55", "2002-12-13 19:07:55", 1},
		{"test1", "2002-12-12 19:07:55", "2002-11-13 19:07:55", 29},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t1, _ := Parse("2006-01-02 15:04:05", tt.to)
			t2, _ := Parse("2006-01-02 15:04:05", tt.from)
			if day := GetDays(t1, t2); day != tt.want {
				t.Errorf("GetDays() day  = %v, want = %v", day, tt.want)
			}
		})
	}

}

func TestGetWeek(t *testing.T) {
	SetSystemTime("2024-05-07 10:10:10")
	tests := []struct {
		name     string
		n        int
		wantYear int
		wantWeek int
	}{
		{"test1", 0, 2024, 19},
		{"test1", -1, 2024, 18},
		{"test1", -3, 2024, 16},
		{"test1", -20, 2023, 51},
		{"test1", 1, 2024, 20},
		{"test1", 3, 2024, 22},
		{"test1", 20, 2024, 39},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotYear, gotWeek := GetWeek(LocalNow(), tt.n)
			if gotYear != tt.wantYear {
				t.Errorf("GetLastWeek() gotYear = %v, want %v", gotYear, tt.wantYear)
			}
			if gotWeek != tt.wantWeek {
				t.Errorf("GetLastWeek() gotWeek = %v, want %v", gotWeek, tt.wantWeek)
			}
		})
	}
}

func TestGetWeekStartTime(t *testing.T) {
	SetSystemTime("2024-05-20 10:10:10")
	tests := []struct {
		name string
		n    int
		want string
	}{
		{"test1", -1, "2024-05-13 00:00:00"},
		{"test2", 0, "2024-05-20 00:00:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want, _ := Parse("2006-01-02 15:04:05", tt.want)
			if got := GetWeekStartTime(LocalNow(), tt.n); !reflect.DeepEqual(got, want) {
				t.Errorf("GetWeekStartTime() = %v, want %v", got, want)
			}
		})
	}
}
