package usecase

import (
	"Capstone/models"
	"Capstone/models/payload"
	"Capstone/repository/database"
	"errors"
)

func CreateStaff(req *payload.CreateStaffRequest) (resp payload.ManageStaffResponse, err error) {
	newStaff := &models.Staff{
		FullName:    req.FullName,
		BirthDate:   req.BirthDate,
		PhoneNumber: req.PhoneNumber,
	}

	err = database.CreateStaff(newStaff)
	if err != nil {
		return
	}

	resp = payload.ManageStaffResponse{
		FullName:    newStaff.FullName,
		BirthDate:   newStaff.BirthDate,
		PhoneNumber: newStaff.PhoneNumber,
	}

	return
}

func UpdateStaff(staff *models.Staff) (resp payload.ManageStaffResponse, err error) {

	err = database.UpdateStaff(staff)
	if err != nil {
		return resp, errors.New("Can't update staff")
	}

	resp = payload.ManageStaffResponse{
		FullName:    staff.FullName,
		BirthDate:   staff.BirthDate,
		PhoneNumber: staff.PhoneNumber,
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
			BirthDate:   staff.BirthDate,
			PhoneNumber: staff.PhoneNumber,
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
