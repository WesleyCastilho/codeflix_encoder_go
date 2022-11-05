package domain

import "time"
import "github.com/asaskevich/govalidator"

type Video struct {
	ID         string    `json:"encoded_video_id" valid:"uuid"`                        //ID do microsservice converter
	ResourceID string    `json:"resource_id" valid:"notnull" gorm:"type:varchar(255)"` //ID do recurso no servico que esta solicitando o converter
	FilePath   string    `json:"file_path" valid:"notnull" gorm:"type:varchar(255)"`   //Caminho do arquivo no bucket
	CreatedAt  time.Time `json:"-" valid:"-"`                                          //Data de criacao do video convertido
	Jobs       []*Job    `json:"-" valid:"-" gorm:"ForeignKey:VideoID"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewVideo() *Video {
	return &Video{}
}

func (video *Video) Validate() error {
	_, err := govalidator.ValidateStruct(video)
	if err != nil {
		return err
	}

	return nil

}
