package usecase

import (
	"Capstone/constants"
	"Capstone/models"
	"Capstone/models/payload"
	"Capstone/repository/database"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/gosimple/slug"
)

// Login Upload Image Warehouse
func UploadImage(file *multipart.FileHeader, warehouseName string) (string, error) {
	slugWarehouse := slug.Make(warehouseName)
	slugFileName := slug.Make(file.Filename)
	path := fmt.Sprintf("images/warehouse/%s-%s.png", slugWarehouse, slugFileName)

	//upload the avatar
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	// Create a new file on disk
	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	// Copy the uploaded file to the destination file
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}
	return path, nil
}

// logic create new warehouse
func CreateWarehouse(file *multipart.FileHeader, req *payload.CreateWarehouseRequest) (resp payload.CreateWarehouseResponse, err error) {
	req.WarehouseImage, err = UploadImage(file, req.Name)
	if err != nil {
		return
	}

	path := fmt.Sprintf("%s/%s", constants.Base_Url, req.WarehouseImage)

	newWarehouse := &models.Warehouse{
		Name:     req.Name,
		Location: req.Location,
		Status:   constants.Available,
		ImageURL: path,
	}

	err = database.CreateWarehouse(newWarehouse)
	if err != nil {
		return
	}

	resp = payload.CreateWarehouseResponse{
		Name:     newWarehouse.Name,
		Location: newWarehouse.Location,
		Status:   newWarehouse.Status,
		ImageURL: newWarehouse.ImageURL,
	}

	return
}

// Logic Delete Warahouse
func DeleteWarehouse(warehouses *models.Warehouse) error {
	err := database.DeleteWarehouse(warehouses)
	if err != nil {
		return err
	}
	return nil
}

// Logic Get All Warehouse

// logic by status warehouse
func GetAllByStatusWarehouse(warehouse *models.Warehouse) (resp []payload.GetAllWarehouseResponse, err error) {
	warehouses, err := database.GetAllWarehouses(warehouse)
	if err != nil {
		return resp, err
	}

	var totalFavorite []int
	for _, v := range warehouses {
		warehouse_id := v.ID
		totalCount := database.CountFavoriteByWarehouseId(warehouse_id)
		totalFavorite = append(totalFavorite, int(totalCount))
	}

	resp = []payload.GetAllWarehouseResponse{}
	for i, warehouse := range warehouses {
		resp = append(resp, payload.GetAllWarehouseResponse{
			ID:       warehouse.ID,
			Name:     warehouse.Name,
			Location: warehouse.Location,
			Favorite: uint(totalFavorite[i]),
			Status:   warehouse.Status,
			Capacity: warehouse.Capacity,
			ImageURL: warehouse.ImageURL,
		})
	}

	return
}

// logic update warehouse
func UpdateWarehouse(warehouse *models.Warehouse) (resp payload.UpdateWarehouseResponse, err error) {

	err = database.UpdateWarehouse(warehouse)
	if err != nil {
		return resp, errors.New("Can't update warehouse")
	}

	resp = payload.UpdateWarehouseResponse{
		Name:     warehouse.Name,
		Location: warehouse.Location,
		Status:   warehouse.Status,
		ImageURL: warehouse.ImageURL,
	}

	return resp, nil
}

// Get Warehouse By Id
func GetWarehouseByID(id uint64) (warehouse *models.Warehouse, err error) {
	warehouse, err = database.GetWarehouseByID(id)
	if err != nil {
		return warehouse, errors.New("Warehouse not found")
	}

	return warehouse, nil
}
