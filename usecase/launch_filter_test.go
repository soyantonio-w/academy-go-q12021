package usecase

import (
	"reflect"
	"testing"

	"github.com/soyantonio-w/academy-go-q12021/entity"
)

func Test_removeOddLaunches(t *testing.T) {
	type args struct {
		launches []entity.Launch
	}

	evenLaunches := []entity.Launch{
		{LaunchId: 0},
		{LaunchId: 2},
		{LaunchId: 4},
	}

	oddLaunches := []entity.Launch{
		{LaunchId: 1},
		{LaunchId: 7},
		{LaunchId: 11},
	}

	tests := []struct {
		name string
		args args
		want []entity.Launch
	}{
		{
			name: "Should not remove any element from a even list",
			args: args{
				launches: evenLaunches,
			},
			want: evenLaunches,
		},
		{
			name: "Should remove all the launches with odd id",
			args: args{
				launches: oddLaunches,
			},
			want: []entity.Launch{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeOddLaunches(tt.args.launches); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeOddLaunches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeEvenLaunches(t *testing.T) {
	type args struct {
		launches []entity.Launch
	}

	evenLaunches := []entity.Launch{
		{LaunchId: 0},
		{LaunchId: 2},
		{LaunchId: 4},
	}

	oddLaunches := []entity.Launch{
		{LaunchId: 1},
		{LaunchId: 7},
		{LaunchId: 11},
	}

	tests := []struct {
		name string
		args args
		want []entity.Launch
	}{
		{
			name: "Should not remove any element from a odd list",
			args: args{
				launches: oddLaunches,
			},
			want: oddLaunches,
		},
		{
			name: "Should remove all the launches with even id",
			args: args{
				launches: evenLaunches,
			},
			want: []entity.Launch{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeEvenLaunches(tt.args.launches); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeEvenLaunches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterLaunchesByType(t *testing.T) {
	type args struct {
		launches   []entity.Launch
		filterType LaunchFilterType
	}

	evenLaunches := []entity.Launch{
		{LaunchId: 0},
		{LaunchId: 2},
		{LaunchId: 4},
	}

	oddLaunches := []entity.Launch{
		{LaunchId: 1},
		{LaunchId: 7},
		{LaunchId: 11},
	}

	tests := []struct {
		name string
		args args
		want []entity.Launch
	}{
		{
			name: "Should not remove odd launches when 'odd' is provided",
			args: args{
				launches:   oddLaunches,
				filterType: Odd,
			},
			want: oddLaunches,
		},
		{
			name: "Should remove even launches when 'odd' is provided",
			args: args{
				launches:   evenLaunches,
				filterType: Odd,
			},
			want: []entity.Launch{},
		},
		{
			name: "Should not remove even launches when 'even' is provided",
			args: args{
				launches:   evenLaunches,
				filterType: Even,
			},
			want: evenLaunches,
		},
		{
			name: "Should remove odd launches when 'even' is provided",
			args: args{
				launches:   oddLaunches,
				filterType: Even,
			},
			want: []entity.Launch{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterLaunchesByType(tt.args.launches, tt.args.filterType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterLaunchesByType() = %v, want %v", got, tt.want)
			}
		})
	}
}
