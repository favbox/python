// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q             = new(Query)
	Author        *author
	Category      *category
	Note          *note
	NoteCategory  *noteCategory
	NoteContent   *noteContent
	NoteImage     *noteImage
	NoteInteract  *noteInteract
	NotePipeline  *notePipeline
	NoteTag       *noteTag
	NoteVideo     *noteVideo
	OriginalURL   *originalURL
	Tag           *tag
	WeixinMp      *weixinMp
	WeixinRequest *weixinRequest
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Author = &Q.Author
	Category = &Q.Category
	Note = &Q.Note
	NoteCategory = &Q.NoteCategory
	NoteContent = &Q.NoteContent
	NoteImage = &Q.NoteImage
	NoteInteract = &Q.NoteInteract
	NotePipeline = &Q.NotePipeline
	NoteTag = &Q.NoteTag
	NoteVideo = &Q.NoteVideo
	OriginalURL = &Q.OriginalURL
	Tag = &Q.Tag
	WeixinMp = &Q.WeixinMp
	WeixinRequest = &Q.WeixinRequest
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:            db,
		Author:        newAuthor(db, opts...),
		Category:      newCategory(db, opts...),
		Note:          newNote(db, opts...),
		NoteCategory:  newNoteCategory(db, opts...),
		NoteContent:   newNoteContent(db, opts...),
		NoteImage:     newNoteImage(db, opts...),
		NoteInteract:  newNoteInteract(db, opts...),
		NotePipeline:  newNotePipeline(db, opts...),
		NoteTag:       newNoteTag(db, opts...),
		NoteVideo:     newNoteVideo(db, opts...),
		OriginalURL:   newOriginalURL(db, opts...),
		Tag:           newTag(db, opts...),
		WeixinMp:      newWeixinMp(db, opts...),
		WeixinRequest: newWeixinRequest(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Author        author
	Category      category
	Note          note
	NoteCategory  noteCategory
	NoteContent   noteContent
	NoteImage     noteImage
	NoteInteract  noteInteract
	NotePipeline  notePipeline
	NoteTag       noteTag
	NoteVideo     noteVideo
	OriginalURL   originalURL
	Tag           tag
	WeixinMp      weixinMp
	WeixinRequest weixinRequest
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:            db,
		Author:        q.Author.clone(db),
		Category:      q.Category.clone(db),
		Note:          q.Note.clone(db),
		NoteCategory:  q.NoteCategory.clone(db),
		NoteContent:   q.NoteContent.clone(db),
		NoteImage:     q.NoteImage.clone(db),
		NoteInteract:  q.NoteInteract.clone(db),
		NotePipeline:  q.NotePipeline.clone(db),
		NoteTag:       q.NoteTag.clone(db),
		NoteVideo:     q.NoteVideo.clone(db),
		OriginalURL:   q.OriginalURL.clone(db),
		Tag:           q.Tag.clone(db),
		WeixinMp:      q.WeixinMp.clone(db),
		WeixinRequest: q.WeixinRequest.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:            db,
		Author:        q.Author.replaceDB(db),
		Category:      q.Category.replaceDB(db),
		Note:          q.Note.replaceDB(db),
		NoteCategory:  q.NoteCategory.replaceDB(db),
		NoteContent:   q.NoteContent.replaceDB(db),
		NoteImage:     q.NoteImage.replaceDB(db),
		NoteInteract:  q.NoteInteract.replaceDB(db),
		NotePipeline:  q.NotePipeline.replaceDB(db),
		NoteTag:       q.NoteTag.replaceDB(db),
		NoteVideo:     q.NoteVideo.replaceDB(db),
		OriginalURL:   q.OriginalURL.replaceDB(db),
		Tag:           q.Tag.replaceDB(db),
		WeixinMp:      q.WeixinMp.replaceDB(db),
		WeixinRequest: q.WeixinRequest.replaceDB(db),
	}
}

type queryCtx struct {
	Author        IAuthorDo
	Category      ICategoryDo
	Note          INoteDo
	NoteCategory  INoteCategoryDo
	NoteContent   INoteContentDo
	NoteImage     INoteImageDo
	NoteInteract  INoteInteractDo
	NotePipeline  INotePipelineDo
	NoteTag       INoteTagDo
	NoteVideo     INoteVideoDo
	OriginalURL   IOriginalURLDo
	Tag           ITagDo
	WeixinMp      IWeixinMpDo
	WeixinRequest IWeixinRequestDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Author:        q.Author.WithContext(ctx),
		Category:      q.Category.WithContext(ctx),
		Note:          q.Note.WithContext(ctx),
		NoteCategory:  q.NoteCategory.WithContext(ctx),
		NoteContent:   q.NoteContent.WithContext(ctx),
		NoteImage:     q.NoteImage.WithContext(ctx),
		NoteInteract:  q.NoteInteract.WithContext(ctx),
		NotePipeline:  q.NotePipeline.WithContext(ctx),
		NoteTag:       q.NoteTag.WithContext(ctx),
		NoteVideo:     q.NoteVideo.WithContext(ctx),
		OriginalURL:   q.OriginalURL.WithContext(ctx),
		Tag:           q.Tag.WithContext(ctx),
		WeixinMp:      q.WeixinMp.WithContext(ctx),
		WeixinRequest: q.WeixinRequest.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
