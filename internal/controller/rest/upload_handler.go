package rest

import (
	"net/http"
	"sppg-backend/internal/model"
	"sppg-backend/internal/usecase"
	"sppg-backend/pkg/upload"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadRoutes(r *gin.RouterGroup) {
	u := r.Group("/upload")
	{
		u.PATCH("/supplier/:id/document", uploadSupplierDocument)
		u.PATCH("/product/:id/image", uploadProductImage)
	}
}

func uploadSupplierDocument(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.SupplierFail("ID tidak valid"))
		return
	}
	docType := c.PostForm("document_type")
	if docType != "nib" && docType != "halal" && docType != "other" {
		c.JSON(http.StatusBadRequest, model.SupplierFail("document_type harus: nib, halal, atau other"))
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, model.SupplierFail("File tidak ditemukan dalam request"))
		return
	}
	savedPath, err := upload.SaveDocument(file, "supplier")
	if err != nil {
		c.JSON(http.StatusBadRequest, model.SupplierFail(err.Error()))
		return
	}
	if err := usecase.UpdateSupplierDocument(id, docType, savedPath); err != nil {
		_ = upload.DeleteFile(savedPath)
		c.JSON(http.StatusInternalServerError, model.SupplierFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.SupplierOK("Dokumen berhasil diupload", gin.H{
		"document_type": docType,
		"file_url":      savedPath,
	}))
}

func uploadProductImage(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ProductFail("ID tidak valid"))
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ProductFail("File tidak ditemukan dalam request"))
		return
	}
	savedPath, err := upload.SaveImage(file, "products")
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ProductFail(err.Error()))
		return
	}
	if err := usecase.UpdateProductImage(id, savedPath); err != nil {
		_ = upload.DeleteFile(savedPath)
		c.JSON(http.StatusInternalServerError, model.ProductFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.ProductOK("Foto produk berhasil diupload", gin.H{
		"file_url": savedPath,
	}))
}