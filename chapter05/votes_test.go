package chapter05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckVoted(t *testing.T) {
	m := New()
	got := m.CheckVoted("me")
	assert.Equal(t, "thank you for your vote", got)
	got = m.CheckVoted("me")
	assert.Equal(t, "already voted", got)
}
