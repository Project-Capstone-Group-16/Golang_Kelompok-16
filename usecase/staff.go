package usecase

import (
	"Capstone/models"
	"Capstone/models/payload"
	"Capstone/repository/database"
	"errors"
	"time"
)

func CreateStaff(req *payload.CreateStaffRequest) (resp payload.ManageStaffResponse, err error) {
	BirthDate, err := time.Parse("02/01/2006", req.BirthDate)
	if err != nil {
		return
	}
	newStaff := &models.Staff{
		FullName:    req.FullName,
		Occupation:  req.Occupation,
		WarehouseID: req.WarehouseID,
		BirthDate:   &BirthDate,
		PhoneNumber: "0" + req.PhoneNumber,
		Address:     req.Address,
	}

	err = database.CreateStaff(newStaff)
	if err != nil {
		return
	}

	resp = payload.ManageStaffResponse{
		FullName:    newStaff.FullName,
		Occupation:  newStaff.Occupation,
		WarehouseID: newStaff.WarehouseID,
		BirthDate:   newStaff.BirthDate,
		PhoneNumber: newStaff.PhoneNumber,
		Address:     newStaff.Address,
	}

	return
}

func UpdateStaff(staff *models.Staff, req *payload.UpdateStaffRequest) (resp payload.ManageStaffResponse, err error) {
	birthDate, err := time.Parse("02/01/2006", req.BirthDate)
	if err != nil {
		return
	}
	staff.BirthDate = &birthDate
	staff.PhoneNumber = req.PhoneNumber
	staff.Address = req.Address

	err = database.UpdateStaff(staff)
	if err != nil {
		return resp, errors.New("Can't update staff")
	}

	updatedStaff, _ := database.GetStaffByID(uint64(staff.ID))

	resp = payload.ManageStaffResponse{
		FullName:    updatedStaff.FullName,
		Occupation:  updatedStaff.Occupation,
		WarehouseID: updatedStaff.WarehouseID,
		BirthDate:   updatedStaff.BirthDate,
		PhoneNumber: updatedStaff.PhoneNumber,
		Address:     updatedStaff.Address,
	}

	return resp, nil
}

func GetStaffByID(id uint64) (staff *models.Staff, err error) {
	staff, err = database.GetStaffByID(id)
	if err != nil {
		return staff, errors.New("Staff not found")
	}

	return staff, nil
}

func GetAllStaffs() (resp []payload.GetAllStaffsResponse, err error) {
	Staffs, err := database.GetAllStaffs()
	if err != nil {
		return
	}

	resp = []payload.GetAllStaffsResponse{}
	for _, staff := range Staffs {
		resp = append(resp, payload.GetAllStaffsResponse{
			ID:          staff.ID,
			WarehouseID: staff.WarehouseID,
			FullName:    staff.FullName,
			Occupation:  staff.Occupation,
			BirthDate:   staff.BirthDate,
			PhoneNumber: staff.PhoneNumber,
			Address:     staff.Address,
		})
	}
	return
}

func DeleteStaff(staff *models.Staff) error {
	err := database.DeleteStaff(staff)
	if err != nil {
		return err
	}
	return nil
}
