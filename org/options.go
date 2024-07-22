package org

import "math/rand/v2"

// Option is a function that configures a Faker instance
type Option func(*Faker)

// WithRandSource sets the custome random source for the Faker
func WithRandSource(src rand.Source) Option {
	return func(f *Faker) {
		f.rnd = rand.New(src)
	}
}
