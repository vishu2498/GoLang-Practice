//Testify utility is a tool for performing tests more effectively with various built-in functions.
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIt(t *testing.T) {
	//'assert' package is mainly used for getting failure messages and checking/comparing purposes.
	//All the functions of 'assert' package return a boolean value showing that test was successful or not.
	t.Log(assert.Equal(t, 125, 125))
	t.Log(assert.NotEqual(t, 125, 237))
	var slice1 []int
	t.Log(assert.Empty(t, slice1))
	var1 := 62
	t.Log(assert.Nil(t, var1))

	//If 'assert' package is utilized many times, it is better to create an object for it.
	eg := assert.New(t)
	t.Log(eg.NotEqual(534, 534, "this is not equal"))

	//'require' implements the same assertions as the `assert` package but stops test execution when a test fails.
	var2 := 62
	var3 := 20
	require.Equal(t, var2, var3, "they are not equal")
}

