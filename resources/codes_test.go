package resources

import (
	"perx-go-test/models"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

func getTmpDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	db = db.Set("gorm:auto_preload", true)

	if err := models.InitAllModels(db); err != nil {
		panic(err)
	}

	return db
}

func TestNewCodesResource(t *testing.T) {
	db := &gorm.DB{}
	alphabet := []rune("abc")
	res := NewCodesResource(db, alphabet)

	assert.Equal(t, res.db, db)
	assert.Equal(t, res.alphabet, alphabet)
}

func TestCodesResource_Create(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name      string
		prepCodes []models.Code
		alphabet  []rune

		args    args
		wantErr bool
	}{
		{
			name:     "simple",
			alphabet: []rune("a"),
			args: args{
				size: 1,
			},

			wantErr: false,
		},

		{
			name:      "no free values",
			prepCodes: []models.Code{{Code: "a"}},
			alphabet:  []rune("a"),
			args: args{
				size: 1,
			},

			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := getTmpDB()
			defer func() { assert.NoError(t, db.Close()) }()

			// fill db
			for _, prepCode := range tt.prepCodes {
				if !assert.NoError(t, db.Create(&prepCode).Error) {
					t.FailNow()
				}
			}

			c := &CodesResource{
				db:       db,
				alphabet: tt.alphabet,
			}

			got, err := c.Create(tt.args.size)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				if assert.NoError(t, err) {
					// test code
					assert.Len(t, got.Code, tt.args.size)
					for _, rune := range got.Code {
						if !assert.Contains(t, tt.alphabet, rune,
							"Code contains rune '%s' not in the alphabet", rune) {
							t.FailNow()
						}
					}

					// test other fields
					assert.Nil(t, got.UsedAt)
					assert.Nil(t, got.DeletedAt)
					assert.NotEqual(t, got.CreatedAt, time.Time{})

				}
			}
		})
	}
}

func TestCodesResource_generateString(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name     string
		alphabet []rune
		args     args
	}{
		{
			name:     "one symbol",
			alphabet: []rune("a"),
			args: args{
				size: 1,
			},
		},
		{
			name:     "multiple symbols",
			alphabet: []rune("abcdef"),
			args: args{
				size: 1000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CodesResource{
				alphabet: tt.alphabet,
			}

			got := c.generateString(tt.args.size)
			for _, rune := range got {
				if !assert.Contains(t, tt.alphabet, rune,
					"CodesResource.generateString() contains rune '%s' not in the alphabet", rune) {
					t.FailNow()
				}
			}
		})
	}
}

func TestCodesResource_Check(t *testing.T) {
	type args struct {
		codeStr string
	}
	tests := []struct {
		name      string
		prepCodes []models.Code
		alphabet  []rune

		args    args
		want    bool
		wantErr bool
		errObj  error
	}{
		{
			name: "simple",
			prepCodes: []models.Code{
				{Code: "test"},
			},
			args: args{
				codeStr: "test",
			},
			want: true,
		},
		{
			name: "code already used",
			prepCodes: func() []models.Code {
				now := time.Now()
				return []models.Code{
					{Code: "test", UsedAt: &now},
				}
			}(),
			args: args{
				codeStr: "test",
			},
			wantErr: true,
			errObj:  ErrCodeAlreadyUsed,
		},
		{
			name: "code not found",
			args: args{
				codeStr: "test",
			},
			wantErr: true,
			errObj:  ErrCodeNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := getTmpDB()
			defer func() { assert.NoError(t, db.Close()) }()

			// fill db
			for _, prepCode := range tt.prepCodes {
				if !assert.NoError(t, db.Create(&prepCode).Error) {
					t.FailNow()
				}
			}

			c := &CodesResource{
				db:       db,
				alphabet: tt.alphabet,
			}

			got, err := c.Check(tt.args.codeStr)

			if tt.wantErr {
				if tt.errObj != nil {
					assert.EqualError(t, err, tt.errObj.Error())
				} else {
					assert.Error(t, err)
				}
			} else {
				if !assert.NoError(t, err) {
					t.FailNow()
				}
				assert.Equal(t, tt.want, got)

				var code models.Code
				assert.NoError(t, db.First(&code).Error)
				assert.NotNil(t, code.UsedAt)
				assert.NotEqual(t, &time.Time{}, code.UsedAt)
			}

		})
	}
}
