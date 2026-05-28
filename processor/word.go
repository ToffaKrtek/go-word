package processor

import (
	"github.com/ToffaKrtek/go-word/vars"
	"github.com/gomutex/godocx"
	"github.com/gomutex/godocx/docx"
)

type AppendData struct {
	styles map[string]string
	rows   []string
}

type ReplaceData struct {
	data map[string]any
	vars []vars.Variable
}

type File struct {
	pathSource  string
	pathTarget  string
	appendData  []AppendData
	replaceData []ReplaceData
}

func NewFile(pathSource string, pathTarget string) File {
	return File{
		pathSource:  pathSource,
		pathTarget:  pathTarget,
		appendData:  []AppendData{},
		replaceData: []ReplaceData{},
	}
}

func (f File) Proccess() error {
	doc, err := f.open()
	if err != nil {
		return err
	}
	err = f.replace(doc)
	if err != nil {
		return err
	}
	err = f.append(doc)
	return f.save(doc)
}

func (f File) save(doc *docx.RootDoc) error {
	return doc.SaveTo(f.pathTarget)
}

func (f File) append(doc *docx.RootDoc) error {
	// for _, ad := range f.appendData {
	//
	// }
	return nil
}

func (f File) replace(doc *docx.RootDoc) error {
	// for _, rd := range f.appendData {
	// 	//replace.go
	// }
	return nil
}

func (f File) open() (*docx.RootDoc, error) {
	doc, err := godocx.OpenDocument(f.pathSource)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func newAppendData(styles map[string]string, rows []string) AppendData {
	return AppendData{styles, rows}
}

func (f *File) AddAppendData(styles map[string]string, rows []string) {
	(*f).appendData = append(f.appendData, newAppendData(styles, rows))
}

func (f *File) AddReplaceData(data map[string]any, vars []vars.Variable) {
	(*f).replaceData = append(f.replaceData, ReplaceData{data, vars})
}
