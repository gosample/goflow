package goflow

import (
	"time"
)

//任务实体类
type Task struct {
	Id           string       `xorm:"varchar(32) pk notnull"`                        //主键ID
	Version      int          `xorm:"tinyint"`                                       //版本
	OrderId      string       `xorm:"varchar(32) notnull index(IDX_TASK_ORDER)"`     //流程实例ID
	TaskName     string       `xorm:"varchar(100) notnull index(IDX_TASK_TASKNAME)"` //任务名称
	DisplayName  string       `xorm:"varchar(200) notnull"`                          //任务显示名称
	PerformType  PERFORM_TYPE `xorm:"tinyint"`                                       //任务参与方式
	TaskType     TASK_TYPE    `xorm:"tinyint notnull"`                               //任务类型
	Operator     time.Time    `xorm:"varchar(50)"`                                   //任务处理者ID
	CreateTime   time.Time    `xorm:"datetime"`                                      //任务创建时间
	FinishTime   string       `xorm:"datetime"`                                      //任务完成时间
	ExpireTime   int          `xorm:"datetime"`                                      //期望任务完成时间
	ActionUrl    string       `xorm:"varchar(200)"`                                  //任务关联的表单URL
	ParentTaskId string       `xorm:"varchar(32) index(IDX_TASK_PARENTTASK)"`        //父任务ID
	Variable     string       `xorm:"varchar(2000)"`                                 //任务附属变量(json存储)
}
