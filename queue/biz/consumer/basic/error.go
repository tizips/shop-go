package basic

import (
	"encoding/json"
	"github.com/golang-module/carbon/v2"
	"github.com/gookit/color"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/queue"
	"gorm.io/gorm"
	constant "project.io/shop/constants/queue"
	"project.io/shop/model"
)

func Error() {

	q := queue.NewQueue()

	err := q.Consumer(doError, constant.BasicError)

	if err != nil {
		color.Errorf("[queue] basic error: %v", err)
		return
	}
}

func doError(data []byte) {

	var body constant.BasicErrorMessage

	var err error

	if err = json.Unmarshal(data, &body); err != nil {
		return
	}

	queue := model.SysQueue{
		Queue:     body.Queue,
		Message:   body.Message,
		Error:     body.Error,
		IsOmitted: util.No,
		IsTried:   util.No,
		IsHandled: util.No,
		CreatedAt: carbon.Carbon{},
		UpdatedAt: carbon.Carbon{},
		DeletedAt: gorm.DeletedAt{},
	}

	facades.Gorm.Create(&queue)
}
