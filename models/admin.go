package models

import (
	DB "application-web/db"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	//ID        uuid.UUID      `json:"id,omitempty" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	ID   string `json:"id,omitempty" gorm:"type:string;primary_key;"`
	Type string `json:"type_user,omitempty" binding:"required"`
}

type Driver struct {
	UserID           string          `json:"id,omitempty" binding:"-" gorm:"type:string;primary_key;" gorm:"foreignKey:UserID"`
	LinkImageProfile string          `json:"link_image_profile" binding:"-"`
	Name             string          `json:"name,omitempty" binding:"-"`
	Plate            string          `json:"plate,omitempty" binding:"required" gorm:"unique"`
	User             User            `json:"user,omitempty" binding:"-" gorm:"foreignKey:UserID"`
	AddressDriverID  uuid.UUID       `json:"address_id,omitempty" binding:"-"`
	AddressDriver    Address         `json:"address_driver,omitempty" binding:"-" gorm:"foreignKey:AddressDriverID"`
	AddressOwnerID   uuid.UUID       `json:"address_owner_id,omitempty" binding:"-"`
	AddressOwner     Address         `json:"address_owner,omitempty" binding:"-" gorm:"foreignKey:AddressOwnerID"`
	Contact          string          `json:"contact,omitempty" binding:"-"`
	FieldsDocuments  FieldsDocuments `json:"fields_documents" binding:"-" gorm:"foreignKey:DriverID"`
	Status           string          `json:"status" binding:"-"`
	Score            int             `json:"score" binding:"-"`
	CnhDocComplete   string          `json:"cnh_doc_complete" binding:"required" `
	RgDocComplete    string          `json:"rg_doc_complete" binding:"required" `
	CrlvDocComplete  string          `json:"crlv_doc_complete" binding:"required" `
}

type FieldsDocuments struct {
	ID        uuid.UUID      `json:"id,omitempty" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// STRUCT CADASTRO DO CARRO KRONA

	DriverID string `json:"driver_id" binding:"-"`

	CarPlate        string `json:"car_plate"`
	Renavam         string `json:"renavam"`
	Chassi          string `json:"car_chassi"`
	CarBrand        string `json:"car_brand"`
	CarModel        string `json:"car_model"`
	CarColor        string `json:"car_color"`
	CarYear         string `json:"car_year"`
	CarType         string `json:"car_type"`
	CarCapacity     string `json:"car_capacity"`
	AnttNumber      string `json:"antt_number"`
	AnttValidity    string `json:"antt_validity"`
	CarOwner        string `json:"car_owner"`
	CarOwnerCnpjCpf string `json:"car_owner_cnpj_cpf"`
	CarTelephone    string `json:"car_telephone"`
	CarCellPhone    string `json:"car_cellphone"`
	CarTypePeople   string `json:"car_type_people"`
	//CarStreet           string `json:"car_street"`
	//CarHouseNumber      string `json:"car_housenumber"`
	//CarComplement       string `json:"car_complement"`
	//CarNeighborhood     string `json:"car_neighborhood"`
	//CarCity             string `json:"car_city"`
	//CarUf               string `json:"car_uf"`
	//CarCep              string `json:"car_cep"`
	CarTechnology       string `json:"car_technology"`
	CarTrackerID        string `json:"car_tracker_id"`
	CarCommunication    string `json:"car_communication"`
	CarSecTechnology    string `json:"car_sec_technology"`
	CarIDTrackerSec     string `json:"car_id_rastreador_sec"`
	CarCommunicationSec string `json:"car_communication_sec"`
	CarFixed            string `json:"car_fixed"`
	CarBondType         string `json:"car_bond_type"`

	// STRUCT CADASTRO DO MOTORISTA KRONA

	DriverName          string `json:"driver_name"`
	DriverCpf           string `json:"driver_cpf"`
	DriverRg            string `json:"driver_rg"`
	DriverUfRg          string `json:"driver_uf_rg"`
	DriverIssuingAgency string `json:"driver_issuing_agency"`
	DriverCityBirth     string `json:"driver_city_birth"`
	DriverBirthDate     string `json:"driver_birth_date"`
	DriverMotherName    string `json:"driver_mother_name"`
	DriverMaritalStatus string `json:"driver_marital_status"`
	DriverSchooling     string `json:"driver_schooling"`
	DriverCnhNumber     string `json:"driver_cnh_number"`
	DriverCnhCategory   string `json:"driver_cnh_category"`
	DriverBallotCnh     string `json:"driver_ballot_cnh"`
	DriverCnhExpiration string `json:"driver_cnh_vencimento"`
	DriverUfCnh         string `json:"driver_uf_cnh"`
	//DriverStreet        string `json:"driver_street"`
	//DriverHouseNumber   string `json:"driver_housenumber"`
	//DriverComplement    string `json:"driver_complement"`
	//DriverNeighborhood  string `json:"driver_neighborhood"`
	//DriverCity          string `json:"driver_city"`
	//DriverUf            string `json:"driver_uf"`
	//DriverCep           string `json:"driver_cep"`
	DriverTelephone string `json:"driver_telephone"`
	DriverCellPhone string `json:"driver_cellphone"`
	DriverNextel    string `json:"driver_nextel"`
	DriverMopp      string `json:"driver_mopp"`
	DriverAso       string `json:"driver_aso"`
	DriverCdd       string `json:"driver_cdd"`
	DriverCapacity  string `json:"driver_capacity"`
	DriverBond      string `json:"driver_bond"`
}

func (f *FieldsDocuments) BeforeCreate(tx *gorm.DB) (err error) {

	f.ID = uuid.New()

	return

}

type Address struct {
	ID        uuid.UUID      `json:"id,omitempty" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	Cep          string `json:"cep,omitempty" binding:"required"`
	Street       string `json:"street,omitempty" binding:"required"`
	Complement   string `json:"complement,omitempty" binding:""`
	Neighborhood string `json:"neighborhood,omitempty" binding:"required"`
	City         string `json:"city,omitempty" binding:"required"`
	Uf           string `json:"uf,omitempty" binding:"required"`
	HouseNumber  string `json:"house_number,omitempty" binding:"required"`
}

func GetAllDrivers() []Driver {
	db := DB.Conn

	var drivers []Driver

	err := db.Joins("FieldsDocuments").Find(&drivers).Error
	if err != nil {
		panic(err.Error())
	}

	return drivers
}
func GetDriversNoDoc() []Driver {
	db := DB.Conn

	var driversNoDoc []Driver

	err := db.Raw("SELECT * FROM drivers A WHERE A.user_id NOT IN(SELECT driver_id FROM fields_documents);").Find(&driversNoDoc).Error
	if err != nil {
		panic(err.Error())
	}

	return driversNoDoc
}

func GetDriversWithDoc() []Driver {
	db := DB.Conn

	var driversDoc []Driver

	err := db.Raw("SELECT * FROM drivers A WHERE A.user_id IN(SELECT driver_id FROM fields_documents);").Find(&driversDoc).Error
	if err != nil {
		panic(err.Error())
	}

	return driversDoc
}

func EditDriver(user_id string) Driver {
	db := DB.Conn

	var driver Driver

	err := db.Joins("FieldsDocuments").Where("user_id = ?", user_id).Find(&driver).Error
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(driver)

	return driver
}

func CreateDocument(slice []byte) {
	db := DB.Conn

	var fields_documents FieldsDocuments

	err := json.Unmarshal(slice, &fields_documents)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}

	query := db.Create(&fields_documents).Error
	if query != nil {
		panic(query.Error)
		return
	}

	return
}
func AtualizaDriver(slice []byte, user_id string) {
	db := DB.Conn

	var atualizadriver Driver

	err := json.Unmarshal(slice, &atualizadriver)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	type score_str struct {
		score string `json:"score2"`
	}
	var scorestr score_str

	err = json.Unmarshal(slice, &scorestr)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}

	scoreparaint, _ := strconv.Atoi(scorestr.score)

	atualizadriver.Score = scoreparaint

	query := db.Where("user_id= ?", user_id).Updates(&atualizadriver).Error
	if query != nil {
		panic(query.Error)
		return
	}

	return
}

func AtualizaFieldsDocuments(slice []byte, driver_id string) {
	db := DB.Conn

	var fields_documents FieldsDocuments

	err := json.Unmarshal(slice, &fields_documents)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}

	query := db.Where("driver_id= ?", driver_id).Updates(&fields_documents).Error
	if query != nil {
		panic(query.Error)
		return
	}

	return
}
