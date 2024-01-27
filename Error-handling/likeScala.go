package main

func Match[T any](v T, validateFunc ...func(v T)) {
	for _, fn := range validateFunc {
		fn(v)
	}
}
