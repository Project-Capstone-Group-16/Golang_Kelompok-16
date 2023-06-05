package usecase

import (
	"Capstone/constants"
	"Capstone/models"
	"Capstone/models/payload"
	"Capstone/repository/database"
	"context"
	"errors"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// logic create new warehouse
func CreateWarehouse(fileHeader *multipart.FileHeader, req *payload.CreateWarehouseRequest) (resp payload.CreateWarehouseResponse, err error) {

	ctx := context.Background()

	file, _ := fileHeader.Open()

	cldService, _ := cloudinary.NewFromURL(os.Getenv("CLD_URL"))
	responseImage, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: "Inventron" + "/" + fileHeader.Filename,
	})

	newWarehouse := models.Warehouse{
		Name:     req.Name,
		Location: req.Location,
		Status:   constants.Available,
		ImageURL: responseImage.SecureURL,
	}

	err = database.CreateWarehouse(&newWarehouse)
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
	ctx := context.Background()

	cldService, _ := cloudinary.NewFromURL(os.Getenv("CLD_URL"))

	_, err := cldService.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: warehouses.ImageURL,
	})
	if err != nil {
		return errors.New("Failed Delete warehouse")
	}

	err = database.DeleteWarehouse(warehouses)
	if err != nil {
		return err
	}
	return nil
}

// Logic Get All Warehouse

// logic by status warehouse
func GetWarehouses(warehouse *models.Warehouse) (resp []payload.GetAllWarehouseResponse, err error) {
	warehouses, err := database.GetWarehouses(warehouse)
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
			Favorite: totalFavorite[i],
			Status:   warehouse.Status,
			Capacity: warehouse.Capacity,
			ImageURL: warehouse.ImageURL,
		})
	}

	return
}

func GetRecomendedWarehouse(warehouse *models.Warehouse) (resp []payload.GetAllWarehouseResponse, err error) {
	warehouses, err := database.GetRecomendedWarehouses(warehouse)
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
			Favorite: totalFavorite[i],
			Status:   warehouse.Status,
			Capacity: warehouse.Capacity,
			ImageURL: warehouse.ImageURL,
		})
	}

	return
}

// logic update warehouse
func UpdateWarehouse(file multipart.File, warehouse *models.Warehouse) (resp payload.UpdateWarehouseResponse, err error) {
	ctx := context.Background()

	cldService, _ := cloudinary.NewFromURL(os.Getenv("CLD_URL"))
	path, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{
		UniqueFilename: api.Bool(true),
		Overwrite:      api.Bool(true),
	})

	warehouse.ImageURL = path.SecureURL

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
