package menusrv

import (
	"reflect"
	"testing"

	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/models"
	"github.com/shopspring/decimal"
)

func Test_getFlavorsMenu(t *testing.T) {
	type args struct {
		menuList []models.Menu
	}
	tests := []struct {
		name string
		args args
		want []models.Menu
	}{
		{
			name: "test_getFlavorsMenu_success",
			args: args{
				menuList: []models.Menu{
					{
						MenuID: "FLV01",
						Name:   "test1",
						Price:  decimal.NewFromFloat(5.99),
					},
					{
						MenuID: "FLV02",
						Name:   "test2",
						Price:  decimal.NewFromFloat(3.95),
					},
					{
						MenuID: "DRK01",
						Name:   "Drink",
						Price:  decimal.NewFromFloat(4.99),
					},
					{
						MenuID: "SPC02",
						Name:   "Special",
						Price:  decimal.NewFromFloat(5.00),
					},
					{
						MenuID: "FLV05",
						Name:   "test5",
						Price:  decimal.NewFromFloat(3.95),
					},
				},
			},
			want: []models.Menu{
				{
					MenuID: "FLV01",
					Name:   "test1",
					Price:  decimal.NewFromFloat(5.99),
				},
				{
					MenuID: "FLV02",
					Name:   "test2",
					Price:  decimal.NewFromFloat(3.95),
				},
				{
					MenuID: "FLV05",
					Name:   "test5",
					Price:  decimal.NewFromFloat(3.95),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFlavorsMenu(tt.args.menuList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFlavorsMenu() = %v, want %v", got, tt.want)
			}
		})
	}
}
