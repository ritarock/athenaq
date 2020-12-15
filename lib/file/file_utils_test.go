package file

import "testing"

func TestRead(t *testing.T) {
	filepath := "../../test_data/test.sql"
	result := Read(filepath)
	expect := "SELECT id FROM test_table WHERE id > 5;"

	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}
