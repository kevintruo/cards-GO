package main

import "testing"

// Introduction to unit testing
// Because I have a repository, using the command go test will throw me an error
// To solve this error, check this link: 'https://stackoverflow.com/questions/67306638/go-test-results-in-go-cannot-find-main-module-but-found-git-config-in-users'
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got length of %v", len(d))
	}
}
