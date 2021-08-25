package repository

import (
    "testing"
    "github.com/amrchnk/advertisement_service/pkg/models"
    "github.com/zhashkevych/go-sqlxmock"
    "github.com/stretchr/testify/assert"
)

func Test_CreatePhoto(t *testing.T){
    db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r:=NewPhotoPostgres(db)

	type args struct {
        photo   models.Photo
        advId int
    }

    tests := []struct {
		name    string
		mock    func()
		input   args
		want    int
		wantErr bool
	}{
        {
            name:"OK",
            input: args{
                photo:models.Photo{
                    Link:"link",
                    First:true,
                },
                advId:1,
            },
            want:2,
            mock:func() {
                mock.ExpectBegin()

                rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
                mock.ExpectQuery("INSERT INTO photos").
                WithArgs("link","first","advert_id").WillReturnRows(rows)

                mock.ExpectCommit()
            },
        },
	}

    for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.CreatePhoto(tt.input.photo, tt.input.advId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}