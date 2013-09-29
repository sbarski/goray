package geometry

import (
	"math"
	"testing"
)

const (
	x  = 4.5
	y  = 2.4
	z  = 3.4
	x2 = 8.7
	y2 = 1.4
	z2 = 3.1
)

func CreateVector() *Vector {
	v := new(Vector)

	v.x = x
	v.y = y
	v.z = z

	return v
}

func CreateVector2() *Vector {
	v := new(Vector)

	v.x = x2
	v.y = y2
	v.z = z2

	return v
}

func TestCreateVector(t *testing.T) {
	v := CreateVector()

	if v.x != x || v.y != y || v.z != z {
		t.Fatal("Vector values not assigned correctly")
	}
}

func TestCloneVector(t *testing.T) {
	v := CreateVector()

	v2 := v.Clone()

	if v2.x != x || v2.y != y || v2.z != z {
		t.Fatalf("Vector values not copied correctly %f - %f, %f - %f, %f - %f", v2.x, x, v2.y, y, v2.z, z)
	}
}

func TestCloneVectorAndReAssign(t *testing.T) {
	v := CreateVector()

	v2 := v.Clone()

	v.x = x2
	v.y = y2
	v.z = z2

	if v == v2 {
		t.Fatalf("Vector addresses are the same: %v - %v", v, v2)
	}

	if v.x != x2 || v2.x != x || v.y != y2 || v2.y != y || v.z != z2 || v2.z != z {
		t.Fatalf("Vector values not copied correctly and assigned %f - %f, %f - %f, %f - %f", v.x, v2.x, v.y, v2.y, v.z, v2.z)
	}
}

func TestVectorMagnitude(t *testing.T) {
	v := CreateVector()

	m := math.Sqrt(x*x + y*y + z*z)
	magnitude := v.Magnitude()

	if magnitude != m {
		t.Fatalf("Magnitude calculated incorrectly %f - %f", magnitude, m)
	}
}

func TestVectorNormalisation(t *testing.T) {
	v := CreateVector()

	v2, _ := v.Normalize()

	if v2.x != 0.7341620244157445 || v2.y != 0.3915530796883971 || v2.z != 0.5547001962252291 {
		t.Fatal("Could not normalize the Vector correctly", v2.x, v2.y, v2.z)
	}

	if v.x != x || v.y != y || v.z != z {
		t.Fatal("Original Vector shouldn't change when performing a normalisation")
	}
}

func TestVectorDotProduct(t *testing.T) {
	v := CreateVector()
	v2 := CreateVector2()

	dotProduct := v.DotProduct(v2)
	dotProductReverse := v2.DotProduct(v)

	if dotProduct != dotProductReverse || dotProduct != 53.05 {
		t.Fatalf("Dot product calculated incorrectly", dotProduct)
	}
}

func TestVectorCrossProduct(t *testing.T) {
	v := CreateVector()
	v2 := CreateVector2()

	v3 := v.CrossProduct(v2)

	if v3.x != 2.6799999999999997 || v3.y != 15.629999999999997 || v3.z != -14.579999999999998 {
		t.Fatalf("Cross product calculated incorrectly", v3)
	}
}

func TestVectorAdd(t *testing.T) {
	v := CreateVector()
	v2 := CreateVector2()

	v3 := v.Add(v2)

	if v3.x != 13.2 || v3.y != 3.8 || v3.z != 6.5 {
		t.Fatalf("Could not add two Vectors together", v3)
	}
}
