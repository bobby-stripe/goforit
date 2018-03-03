package goforit

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Tset basic CSV parsing
func TestParseCSV(t *testing.T) {
	t.Parallel()

	path := filepath.Join("fixtures", "flags_example.csv")
	expected := []SampleFlag{
		{
			"go.sun.money",
			0,
		},
		{
			"go.moon.mercury",
			1,
		},
		{
			"go.stars.money",
			0.5,
		},
	}

	// Parse it
	file, err := os.Open(path)
	require.NoError(t, err)
	defer file.Close()
	flags, lastMod, err := CsvFileFormat{}.Read(file)
	assert.NoError(t, err)
	assert.True(t, lastMod.IsZero())

	// Check that the results are correct
	for i, f := range flags {
		sf, ok := f.(SampleFlag)
		require.True(t, ok)
		assert.Equal(t, expected[i].Name(), f.Name())
		assert.Equal(t, expected[i].Rate, sf.Rate)
	}
}

// Test that we get errors from broken CSVs
func TestParseCSVBroken(t *testing.T) {
	t.Parallel()

	// Bad number of fields
	path := filepath.Join("fixtures", "flags_broken1.csv")
	file, err := os.Open(path)
	defer file.Close()
	_, _, err = CsvFileFormat{}.Read(file)
	assert.Error(t, err)

	// Rate not parseable as a number
	path = filepath.Join("fixtures", "flags_broken2.csv")
	file2, err := os.Open(path)
	defer file2.Close()
	_, _, err = CsvFileFormat{}.Read(file2)
	assert.Error(t, err)
}

func TestNewCsvBackend(t *testing.T) {
	t.Parallel()

	path := filepath.Join("fixtures", "flags_example.csv")
	backend := NewCsvBackend(path, DefaultRefreshInterval)
	defer backend.Close()
}
