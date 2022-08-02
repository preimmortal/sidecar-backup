package main

// Need Mocking to test this
/*
func Test_worker(t *testing.T) {
	type args struct {
		jobs <-chan Job
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Worker",
			args: args{
				jobs: make(<-chan Job),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			worker(tt.args.jobs)

		})
	}
}
*/