package service

import (
	"errors"
	"go-study/global"
	"go-study/internal/model"
	"go-study/utils"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type SolutionService struct{}

var addSolutionLock sync.Mutex

// 获取解决方案列表
func (solutionService *SolutionService) GetSolutionList(c *gin.Context) ([]model.StoreSolution, error) {
	if name := c.Query("name"); name != "" {
		// 模糊查询
		global.DB = global.DB.Where("name like ?", "%"+name+"%")
	}
	if solutionType := c.Query("type"); solutionType != "" {
		global.DB = global.DB.Where("type = ?", solutionType)
	}
	if status := c.Query("status"); status != "" {
		global.DB = global.DB.Where("status = ?", status)
	}
	var solutionList []model.StoreSolution
	result := global.DB.Find(&solutionList)
	if result.Error != nil {
		return solutionList, result.Error
	}
	return solutionList, nil
}

// 创建解决方案
func (solutionService *SolutionService) AddSolution(solution model.AddSolutionReq) (model.StoreSolution, error) {

	addSolutionLock.Lock()
	defer addSolutionLock.Unlock()

	var newSolution model.StoreSolution
	//name，type， picUrl， status不能为空
	if solution.Name == "" {
		return newSolution, errors.New("名称不能为空")
	}
	if strconv.Itoa(solution.Type) == "" {
		return newSolution, errors.New("类型不能为空")
	}
	if solution.PicUrl == "" {
		return newSolution, errors.New("图片不能为空")
	}
	// 名称只能由中文字母、数字、下划线组成
	if !utils.IsLetterOrDigit(solution.Name) {
		return newSolution, errors.New("名称只能由中文、字母、数字、下划线组成")
	}

	//若name和version相同，则不允许添加
	findDuplicate := global.DB.First(&newSolution, "name = ? and version = ?", newSolution.Name, newSolution.Version)
	if findDuplicate.Error == nil {
		return newSolution, errors.New("已存在该解决方案")
	}

	result := global.DB.Create(&newSolution)
	if result.Error != nil {
		return newSolution, result.Error
	}
	return newSolution, nil
}

func (solutionService *SolutionService) UpdateSolution(newSolution model.StoreSolution) (model.StoreSolution, error) {
	//首先根据id查找
	var solutionEntity model.StoreSolution
	findStoreSolution, err := solutionService.GetSolutionById(newSolution.Id)
	if err != nil {
		return solutionEntity, err
	}
	solutionEntity.Name = findStoreSolution.Name
	solutionEntity.Version = findStoreSolution.Version
	solutionEntity.Type = findStoreSolution.Type
	solutionEntity.Introduction = findStoreSolution.Introduction
	solutionEntity.PicUrl = findStoreSolution.PicUrl
	solutionEntity.Status = findStoreSolution.Status
	result := global.DB.Updates(&solutionEntity)
	if result.Error != nil {
		return solutionEntity, result.Error
	}
	return solutionEntity, nil
}

// 删除解决方案
func (solutionService *SolutionService) DeleteSolution(id int) error {
	var solution model.StoreSolution
	result := global.DB.First(&solution, "id = ?", id)
	if result.Error != nil {
		return errors.New("未找到该解决方案")
	}
	result = global.DB.Delete(&solution)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 根据id获取解决方案
func (solutionService *SolutionService) GetSolutionById(id int) (model.StoreSolution, error) {
	var solution model.StoreSolution
	result := global.DB.First(&solution, "id = ?", id)
	if result.Error != nil {
		return solution, errors.New("未找到该解决方案")
	}
	return solution, nil
}
