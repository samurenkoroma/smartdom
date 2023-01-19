package parser

import (
	"log"
	"os"
	"path/filepath"
	"smartdom/services/libparser/internal/domain/book"
	"smartdom/services/libparser/internal/useCase"
	"smartdom/services/libparser/internal/useCase/parser/adapters"
)

type UseCase struct {
	store adapters.Storage
}

func (u UseCase) Parse(dir string) (books []book.Book, err error) {
	err = filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				ext := filepath.Ext(path)
				books = append(books, book.New(info.Name()[:len(info.Name())-len(ext)], path, uint64(info.Size()), nil, nil))
			}
			return err
		})
	if err != nil {
		log.Println(err)
	}
	return books, err
}

func (u UseCase) Save(books []book.Book) error {
	return u.store.Create(books)
}

func New(store adapters.Storage) useCase.LibraryParser {
	return &UseCase{
		store: store,
	}
}

//
//import (
//	"context"
//	"log"
//	"os"
//	"path/filepath"
//	"smarthome/library/internal/config"
//	"smarthome/library/internal/domain/book"
//	"smarthome/library/internal/domain/book/db"
//	"smarthome/library/pkg/client/mongodb"
//)
//
//type file struct {
//	title string
//	path  string
//	ext   string
//	size  int64
//}
//
//type Parser struct {
//	files    []file
//	dirParse string
//}
//
//func NewParser(dirParse string) *Parser {
//	return &Parser{dirParse: dirParse}
//}
//
//func (p *Parser) Run() {
//	cfg := config.GetConfig()
//	client, err := mongodb.NewClient(context.Background(), mongodb.)
//	if err != nil {
//		log.Fatal(err)
//	}
//	saver := book.NewService(db.NewStorage(client, "Books"))
//	p.parse()
//	if err != nil {
//		return
//	}
//	for _, file := range p.files {
//		saver.Create(context.Background(), book.Book{
//			Title: file.title,
//			Path:  file.path,
//			Ext:   file.ext,
//		})
//	}
//}
//
//func (p *Parser) parse() error {
//	err := filepath.Walk(p.dirParse,
//		func(path string, info os.FileInfo, err error) error {
//			if err != nil {
//				return err
//			}
//			if !info.IsDir() {
//				ext := filepath.Ext(path)
//				p.files = append(p.files, file{title: info.Name()[:len(info.Name())-len(ext)], path: path, ext: ext, size: info.Size()})
//			}
//			return nil
//		})
//	if err != nil {
//		log.Println(err)
//	}
//	return nil
//}
