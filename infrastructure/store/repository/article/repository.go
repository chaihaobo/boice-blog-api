package article

import (
	"context"

	"github.com/chaihaobo/boice-blog-api/infrastructure/store/client"
	"github.com/chaihaobo/boice-blog-api/model/entity"
	"github.com/chaihaobo/boice-blog-api/model/generic"
)

type (
	Repository interface {
		ListArticles(ctx context.Context, pagination generic.Pagination) ([]*entity.Article, int64, error)
		GetArticleTags(ctx context.Context, id uint64) ([]*entity.Tag, error)
		Save(ctx context.Context, data *entity.Article) error
		Get(ctx context.Context, id uint64) (*entity.Article, error)
	}
	repository struct {
		client client.Client
	}
)

func (r repository) Get(ctx context.Context, id uint64) (*entity.Article, error) {
	var data entity.Article
	if result := r.client.DB(ctx).Where("id = ?", id).Find(&data); result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}
	tags, err := r.GetArticleTags(ctx, data.ID)
	if err != nil {
		return nil, err
	}
	data.Tags = tags
	return &data, nil
}

func (r repository) Save(ctx context.Context, data *entity.Article) error {
	db := r.client.DB(ctx)
	if err := db.Save(data).Error; err != nil {
		return err
	}
	// 删除所有tag
	if err := db.Exec("delete from article_tag where article_id=?", data.ID).Error; err != nil {
		return err
	}
	// 保存所有tag
	for _, tag := range data.Tags {
		if err := db.Save(tag).Error; err != nil {
			return err
		}
		if err := db.Exec("insert into article_tag(article_id,tag_id) values(?,?)", data.ID, tag.ID).Error; err != nil {
			return err
		}
	}
	return nil

}

func (r repository) GetArticleTags(ctx context.Context, id uint64) ([]*entity.Tag, error) {
	var tags = make([]*entity.Tag, 0)
	if err := r.client.DB(ctx).
		Raw("select * from article_tag t1 inner join tag t2 on t1.tag_id=t2.id where t1.article_id=?", id).
		Scan(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r repository) ListArticles(ctx context.Context, pagination generic.Pagination) ([]*entity.Article, int64, error) {
	// 查询文章列表
	var articles []*entity.Article
	err := r.client.DB(ctx).Offset((pagination.Page - 1) * pagination.Size).Order("id desc").Limit(pagination.Size).Find(&articles).Error
	if err != nil {
		return nil, 0, err
	}

	// 查询总记录数
	var total int64
	err = r.client.DB(ctx).Model(&entity.Article{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	for _, article := range articles {
		// 查询文章标签
		tags, err := r.GetArticleTags(ctx, article.ID)
		if err != nil {
			return nil, 0, err
		}
		article.Tags = tags
	}

	return articles, total, nil
}

func NewRepository(client client.Client) Repository {
	return &repository{
		client: client,
	}
}
