package apis
import (
	"encoding/json"
    "net/http"
	"myapi/models"
	"myapi/dals"
	"myapi/utils"
)

func GetProducts(response http.ResponseWriter, request *http.Request) {
	products := dals.GetProducts()
	utils.ResponseWithJSON(response, http.StatusOK, products)
}

func CreateProduct(response http.ResponseWriter, request *http.Request) {
    var product models.Product
    err := json.NewDecoder(request.Body).Decode(&product)
    if err != nil {
        utils.ResponseWithError(response, http.StatusBadRequest, err.Error())
    } else {
        result := dals.CreateProduct(&product)
        if !result {
            utils.ResponseWithError(response, http.StatusBadRequest, "Could not create product")
            return
        }
        utils.ResponseWithJSON(response, http.StatusOK, product)
    }
}

func UpdateProduct(response http.ResponseWriter, request *http.Request) {
    var product models.Product
    err := json.NewDecoder(request.Body).Decode(&product)
    if err != nil {
        utils.ResponseWithError(response, http.StatusBadRequest, err.Error())
    } else {
        result := dals.UpdateProduct(&product)
        if !result {
            utils.ResponseWithError(response, http.StatusBadRequest, "Could not update product")
            return
        }
        utils.ResponseWithJSON(response, http.StatusOK, "Update product successfully")
    }
}
 
func Delete(response http.ResponseWriter, request *http.Request) {
    ids, ok := request.URL.Query()["id"]
    if !ok || len(ids) < 1 {
        utils.ResponseWithError(response, http.StatusBadRequest, "Url Param id is missing")
        return
    }
    result := dals.DeleteProduct(ids[0])
    if !result {
        utils.ResponseWithError(response, http.StatusBadRequest, "Could not delete product")
        return
    }
    utils.ResponseWithJSON(response, http.StatusOK, "Delete product successfully")
}