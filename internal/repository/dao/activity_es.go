package dao

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	dtov1 "github.com/zhuguangfeng/go-chat/dto"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/utils"
	"strconv"
)

type ActivityEsDao interface {
	InputActivity(ctx context.Context, activity model.ActivityEs) error
	SearchActivity(ctx context.Context, req dtov1.ActivitySearchReq) ([]model.ActivityEs, error)
}

type OlivereActivityEsDao struct {
	esCli *elastic.Client
}

func NewActivityEsDao(esCli *elastic.Client) ActivityEsDao {
	return &OlivereActivityEsDao{
		esCli: esCli,
	}
}

func (dao *OlivereActivityEsDao) InputActivity(ctx context.Context, activity model.ActivityEs) error {
	_, err := dao.esCli.Index().Index(model.ActivityIndexName).Id(strconv.Itoa(int(activity.ID))).BodyJson(activity).Do(ctx)
	return err
}

func (dao *OlivereActivityEsDao) SearchActivity(ctx context.Context, req dtov1.ActivitySearchReq) ([]model.ActivityEs, error) {
	quer := dao.buildQuery(req)
	resp, err := dao.esCli.Search(model.ActivityIndexName).Query(quer).Do(ctx)
	if err != nil {
		return nil, err
	}
	var res []model.ActivityEs

	for _, hit := range resp.Hits.Hits {
		var art model.ActivityEs
		err := json.Unmarshal(hit.Source, &art)
		if err != nil {
			return nil, err
		}
		res = append(res, art)
	}
	return res, nil
}

func (dao *OlivereActivityEsDao) buildQuery(req dtov1.ActivitySearchReq) elastic.Query {
	titleQuery := elastic.NewMatchQuery("title", req.SearchKey)
	descQuery := elastic.NewMatchQuery("desc", req.SearchKey)
	or := elastic.NewBoolQuery().Should(titleQuery, descQuery)

	query := elastic.NewBoolQuery().Must(or)
	//or 查询

	if req.Address != "" {
		query = query.Must(elastic.NewMatchQuery("address", req.Address))
	}
	if req.AgeRestrict > 0 {
		query = query.Must(elastic.NewTermQuery("ageRestrict", req.AgeRestrict))
	}
	if req.GenderRestrict > 0 {
		query = query.Must(elastic.NewTermQuery("genderRestrict", req.GenderRestrict))
	}
	if req.Visibility > 0 {
		query = query.Must(elastic.NewMatchQuery("visibility", req.Visibility))
	}
	if req.Category > 0 {
		query = query.Must(elastic.NewTermQuery("category", req.Category))
	}
	startTime := utils.StringToTimeUnix(req.StartTime)
	if startTime > 0 {
		query = query.Must(elastic.NewRangeQuery("startTime").Lte(startTime))
	}
	endTime := utils.StringToTimeUnix(req.StartTime)
	if endTime > 0 {
		query = query.Must(elastic.NewRangeQuery("startTime").Gte(endTime))
	}
	return query
}
