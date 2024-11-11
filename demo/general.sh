#!/bin/bash

# 将下划线转成驼峰
function underscore_to_camelcase() {
    input_string="$1"
    output_string=""

    # Split the string by underscore
    IFS='_' read -ra words <<< "$input_string"

    for word in "${words[@]}"; do
        # Capitalize the first letter of each word
        capitalized_word=$(echo "${word:0:1}" | tr '[:lower:]' '[:upper:]')$(echo "${word:1}" | tr '[:upper:]' '[:lower:]')
        output_string+="$capitalized_word"
    done

    echo "$output_string"
}

if [ $# -eq 0 ]; then
    echo "请输入 table name"
    exit 1
fi

tableName=$1
echo "table name："$tableName
controllerName=$(underscore_to_camelcase "$tableName")
echo "controller name: "$controllerName

cd /Users/yiweilin/GolandProjects/demo

# Controller
echo 'package cache

import (
  "github.com/gookit/goutil"
  "github.com/gookit/goutil/jsonutil"
  "gorm.io/gen"

  "mecordserver/demo/admin/dao/model"
  "mecordserver/demo/admin/dao/query"
  "mecordserver/demo/admin/request"
  "mecordserver/demo/admin/response"
  "mecordserver/demo/admin/utils/validator"
  "mecordserver/util/errors"
)

type '$controllerName' struct {
}

func (s *'$controllerName') List(param map[string]interface{}) (rsp *response.ResData) {
  // 解析请求体
  req := request.'$controllerName'List{}
  request.ParseParam(param, &req)

  var (
    rows  []*model.'$controllerName'
    count int64
    err   error
  )

  c := query.'$controllerName'

  // 查询条件
  filters := make([]gen.Condition, 0)
  if req.ID > 0 {
    filters = append(filters, c.ID.Eq(goutil.Int64(req.ID)))
  }
  do := c.Where(filters...)
  do = do.Order(c.ID.Desc())

  if req.IsExport {
    // 导出处理
    rows, err = do.Limit(10000).Find()
  } else {
    rows, count, err = do.FindByPage(req.Offset, req.Limit)
  }

  if err != nil {
    rsp = response.Error(err)
    return
  }

  items := make([]*response.'$controllerName', 0)
  for _, row := range rows {
      item := &response.'$controllerName'{}
  		jsonutil.Mapping(row, item)
      items = append(items, item)
  }

  // 返回结果
  rsp = response.SuccessWithList(items, count)
  return
}

func (s *'$controllerName') Create(param map[string]interface{}) (rsp *response.ResData) {
  // 解析请求体
  req := &request.'$controllerName'Create{}
  request.ParseParam(param, &req)

  // 参数校验
  err := validator.Validate(req)
  if err != nil {
    rsp = response.Error(err)
    return
  }

  m := &model.'$controllerName'{}
  err = jsonutil.Mapping(req, m)
  if err != nil {
    rsp = response.Error(err)
    return
  }

  qModel := query.'$controllerName'
  err = qModel.Save(m)
  if err != nil {
    rsp = response.Error(err)
    return
  }

  // 返回结果
  rsp = response.Success()
  return
}

func (s *'$controllerName') Update(param map[string]interface{}) (rsp *response.ResData) {
  // 解析请求体
  req := &request.'$controllerName'Update{}
  request.ParseParam(param, &req)

  // 参数校验
  err := validator.Validate(req)
  if err != nil {
    rsp = response.Error(err)
    return
  }

  qModel := query.'$controllerName'
  m, err := qModel.Where(qModel.ID.Eq(req.ID)).Take()
  if err != nil {
    rsp = response.Error(err)
    return
  }

  err = jsonutil.Mapping(req, m)
  if err != nil {
    rsp = response.Error(err)
    return
  }

  err = qModel.Save(m)
  if err != nil {
    rsp = response.Error(err)
    return
  }

  // 返回结果
  rsp = response.Success()
  return
}

func (s *'$controllerName') Delete(param map[string]interface{}) (rsp *response.ResData) {

  // 解析请求体
  req := &request.'$controllerName'Delete{}
  request.ParseParam(param, &req)

  // 参数校验
  err := validator.Validate(req)
  if err != nil {
    rsp = response.Error(err)
    return
  }
  if req.ID <= 0 {
    err = errors.NewParams()
    rsp = response.Error(err)
    return
  }

  qModel := query.'$controllerName'
  row, err := qModel.Where(qModel.ID.Eq(req.ID)).Take()
  if err != nil {
    rsp = response.Error(err)
    return
  }

  _, err = qModel.Delete(row)
  if err != nil {
    rsp = response.Error(err)
    return
  }

  return response.Success()
}

' > $tableName.go


# Request
echo 'package request

import (
  "mecordserver/demo/admin/dao/model"
)

type '$controllerName'List struct {
  Admin
  Paginate
  model.'$controllerName'
}

type '$controllerName'Create struct {
  Admin
  model.'$controllerName'
}

type '$controllerName'Update struct {
  Admin
  model.'$controllerName'
}

type '$controllerName'Delete struct {
  Admin
  model.'$controllerName'
}
' > $tableName.go

# Test
echo 'package tests

import (
	"testing"

	"github.com/gookit/goutil/jsonutil"

	"mecordserver/demo/admin/controller"
	"mecordserver/demo/admin/request"
)

func Test_'$controllerName'_List(t *testing.T) {
  req := request.'$controllerName'List{
    Paginate: request.Paginate{
      Page:     1,
      PageSize: 20,
    },
  }

  param := map[string]interface{}{}
  jsonutil.Mapping(req, &param)

  rsp := new(controller.'$controllerName').List(param)
  JsonPrint(param, rsp)
}

func Test_'$controllerName'_Create(t *testing.T) {

  req := request.'$controllerName'Create{}

  param := map[string]interface{}{}
  jsonutil.Mapping(req, &param)

  rsp := new(controller.'$controllerName').Create(param)
  JsonPrint(param, rsp)
}

func Test_'$controllerName'_Update(t *testing.T) {

  req := request.'$controllerName'Update{}

  param := map[string]interface{}{}
  jsonutil.Mapping(req, &param)

  rsp := new(controller.'$controllerName').Update(param)
  JsonPrint(param, rsp)
}

func Test_'$controllerName'_Delete(t *testing.T) {

  req := request.'$controllerName'Delete{}

  param := map[string]interface{}{}
  jsonutil.Mapping(req, &param)

  rsp := new(controller.'$controllerName').Delete(param)
  JsonPrint(param, rsp)
}

' > $tableName'_test.go'


# Response
echo 'package response

import (
  "mecordserver/demo/admin/dao/model"
)

type '$controllerName' struct {
  model.'$controllerName'
}

' > $tableName.go
