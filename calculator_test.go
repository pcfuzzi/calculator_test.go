package main

import "testing"


func TestAdd (t *testing.T) {
c := Calculator{}
result := c.Add(2, 3)
expected := 5
if result != expected {
	t.Errorf("Add(2, 3)= %d; want %de", result, expected)
}

}
func TestSubtract (t *testing.T) {
	c := Calculator{}
	result := c.Subtract(5, 3)
	expected := 2
	if result != expected {
		t.Errorf("Subtract(2, 3)= %d; want %de", result, expected)
	}
	
	} 
	func TestMultiply (t *testing.T) {
		c := Calculator{}
		result := c.Multiply(5, 3)
		expected := 15
		if result != expected {
			t.Errorf("Multiply(2, 3)= %d; want %de", result, expected)
		}
		
		} 
		func TestDivide (t *testing.T) {
			c := Calculator{}
			result, err := c.Divide(15, 5)
			expected := 3
			if err != nil {
				t.Errorf("Divide(15, 5) returned an error: %s", err.Error())
			}
			if result != expected {
				t.Errorf("Divide(6, 3)= %d; want %de", result, expected)
				
			}
			
			} 
			type Calculator struct{}
			
			func(c Calculator) Add(a, b int) int{
				return a + b
			}
			func(c Calculator) Subtract(a, b int) int{
				return a - b
			}
			func(c Calculator) Multiply(a, b int) int{
				return a * b
			}
			func(c Calculator) Divide(a, b int) (int, error) {
				if b == 0 {
					return 0, &DividebyZeroError{}
				}
				return a / b, nil
			}
type DividebyZeroError struct{}

func (e *DividebyZeroError) Error() string{
	return "Division durch Null"
}