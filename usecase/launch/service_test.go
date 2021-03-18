package launch

import (
	"fmt"
	"github.com/soyantonio-w/academy-go-q12021/entity"
	"github.com/soyantonio-w/academy-go-q12021/entity/mocks"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	type args struct {
		r entity.LaunchRepo
	}

	mockRepo := new(mocks.LaunchRepo)

	tests := []struct {
		name string
		args args
		want *Service
	}{
		{name: "Should create a service with the given repo", args: args{
			r: mockRepo,
		}, want: &Service{
			repo: mockRepo,
		}},
		{name: "Should create a service with nil repo", args: args{}, want: &Service{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetLaunch(t *testing.T) {
	type fields struct {
		repo entity.LaunchRepo
	}
	type args struct {
		launchId string
	}

	mockRepo := new(mocks.LaunchRepo)
	mockRepo.On("Get", entity.LaunchId(24)).Return(entity.Launch{}, fmt.Errorf("non found id"))
	mockRepo.On("Get", entity.LaunchId(2)).Return(entity.Launch{LaunchId: 2}, nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Launch
		wantErr bool
	}{
		{
			name: "Should get an error when the id is not numeric", fields: fields{
				repo: mockRepo,
			}, args: args{
				launchId: "NonNumericId",
			},
			want: entity.Launch{}, wantErr: true,
		},
		{
			name: "Should get an error when the id is no present in the repo", fields: fields{
				repo: mockRepo,
			}, args: args{
				launchId: "24",
			},
			want: entity.Launch{}, wantErr: true,
		},
		{
			name: "Should return the launch with the matched id", fields: fields{
				repo: mockRepo,
			}, args: args{
				launchId: "2",
			},
			want: entity.Launch{
				LaunchId: 2,
			}, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: tt.fields.repo,
			}
			got, err := s.GetLaunch(tt.args.launchId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLaunch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLaunch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ListLaunches(t *testing.T) {
	type fields struct {
		repo entity.LaunchRepo
	}

	emptyRepo := new(mocks.LaunchRepo)
	emptyRepo.On("GetLaunches").Return([]entity.Launch{}, nil)

	mockedLaunches := []entity.Launch{{LaunchId: 4}, {LaunchId: 1}}
	errorRepo := new(mocks.LaunchRepo)
	errorRepo.On("GetLaunches").Return(mockedLaunches, fmt.Errorf("unexpected error"))

	mockRepo := new(mocks.LaunchRepo)
	mockRepo.On("GetLaunches").Return(mockedLaunches, nil)

	tests := []struct {
		name    string
		fields  fields
		want    []entity.Launch
		wantErr bool
	}{
		{
			name:    "Should return empty launches without any error",
			fields:  fields{repo: emptyRepo},
			want:    []entity.Launch{},
			wantErr: false,
		},
		{
			name:    "Should return an empty list when there is an error in the repo",
			fields:  fields{repo: errorRepo},
			want:    []entity.Launch{},
			wantErr: true,
		},
		{
			name:    "Should return all the launches",
			fields:  fields{repo: mockRepo},
			want:    mockedLaunches,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: tt.fields.repo,
			}
			got, err := s.ListLaunches()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListLaunches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListLaunches() got = %v, want %v", got, tt.want)
			}
		})
	}
}
