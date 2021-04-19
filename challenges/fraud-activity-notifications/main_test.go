package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestActivityNotifications_Case1(t *testing.T) {
	ret := activityNotifications([]int32{10, 20, 30, 40, 50}, 3)
	assert.Equal(t, int32(1), ret)
}

func TestActivityNotifications_Case2(t *testing.T) {
	ret := activityNotifications([]int32{2, 3, 4, 2, 3, 6, 8, 4, 5}, 5)
	assert.Equal(t, int32(2), ret)
}

func TestActivityNotifications_Case3(t *testing.T) {
	ret := activityNotifications([]int32{1, 2, 3, 4, 4}, 4)
	assert.Equal(t, int32(0), ret)
}
