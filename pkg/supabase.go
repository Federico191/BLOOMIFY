package pkg

import (
	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"mime/multipart"
)

type SupabaseStorageItf interface {
	Upload(file *multipart.FileHeader) (string, error)
	Delete(link string) error
}

type SupabaseStorage struct {
	client *supabasestorageuploader.Client
}

func NewSupabaseStorage(client *supabasestorageuploader.Client) SupabaseStorageItf {
	return &SupabaseStorage{client: client}
}

func (s SupabaseStorage) Upload(file *multipart.FileHeader) (string, error) {
	link, err := s.client.Upload(file)
	if err != nil {
		return "", err
	}

	return link, nil
}

func (s SupabaseStorage) Delete(link string) error {
	err := s.client.Delete(link)
	if err != nil {
		return err
	}

	return nil
}
